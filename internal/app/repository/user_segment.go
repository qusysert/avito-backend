package repository

import (
	db "avito-backend/pkg/gopkg-db"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (r *Repository) AddUserSegmentIfNotExists(ctx context.Context, userId, segmentId int, expires *time.Time) (int, error) {
	var id int
	var expiresValue sql.NullTime

	if expires != nil {
		expiresValue.Time = *expires
		expiresValue.Valid = true
	}

	row := db.FromContext(ctx).QueryRow(ctx,
		`INSERT INTO user_segment (user_id, segment_id, expires)
				VALUES ($1, $2, $3)
				ON CONFLICT (user_id, segment_id) DO UPDATE SET
				user_id = excluded.user_id,
				segment_id = excluded.segment_id,
				expires = excluded.expires
				RETURNING id`,
		userId, segmentId, expiresValue)

	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to scan row on adding segment: %w", err)
	}

	return id, nil
}

func (r *Repository) GetSegmentsOfUser(ctx context.Context, id int) ([]int, error) {
	var segmentIds []int
	rows, err := db.FromContext(ctx).Query(ctx, `SELECT segment_id FROM user_segment WHERE user_id=$1 AND expires>$2`,
		id, time.Now())
	if err != nil {
		return nil, fmt.Errorf("failed to query segments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var name int
		err := rows.Scan(&name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan segment name: %w", err)
		}
		segmentIds = append(segmentIds, name)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating through segments: %w", err)
	}

	return segmentIds, nil
}

func (r *Repository) DeleteUserSegmentIfExists(ctx context.Context, userId, segmentId int) error {
	_, err := db.FromContext(ctx).Exec(ctx,
		`DELETE FROM user_segment WHERE user_id=$1 AND segment_id=$2`,
		userId, segmentId)
	if err != nil {
		return fmt.Errorf("failed to delete user segment: %w", err)
	}

	return nil
}

func (r *Repository) DeleteSegmentFromUsers(ctx context.Context, segmentId int) error {
	_, err := db.FromContext(ctx).Exec(ctx,
		`DELETE FROM user_segment WHERE segment_id=$1`, segmentId)
	if err != nil {
		return fmt.Errorf("failed to delete segment %d from users: %w", segmentId, err)
	}

	return nil
}

func (r *Repository) FlushExpired(ctx context.Context) error {
	_, err := db.FromContext(ctx).Exec(ctx,
		`DELETE FROM user_segment WHERE expires < $1`, time.Now())
	if err != nil {
		return fmt.Errorf("failed to flush expired entries: %w", err)
	}

	return nil
}
