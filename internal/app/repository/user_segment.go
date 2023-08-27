package repository

import (
	"avito-backend/internal/app/model"
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

func (r *Repository) GetUserSegments(ctx context.Context, id int) ([]model.SegmentWithExpires, error) {
	var segments []model.SegmentWithExpires
	rows, err := db.FromContext(ctx).Query(ctx, `SELECT us.segment_id, s.name, us.expires FROM user_segment us INNER JOIN segment s ON us.segment_id = s.id WHERE us.user_id=$1 AND us.expires>$2`,
		id, time.Now())
	if err != nil {
		return nil, fmt.Errorf("failed to query segments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var segment model.SegmentWithExpires
		err := rows.Scan(&segment.Id, &segment.Name, &segment.Expires)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows: %w", err)
		}
		segments = append(segments, segment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating through segments: %w", err)
	}

	return segments, nil
}

func (r *Repository) DeleteUserSegmentIfExists(ctx context.Context, userId int, segmentName string) (bool, error) {
	res, err := db.FromContext(ctx).Exec(ctx,
		`DELETE FROM user_segment us USING segment s WHERE us.segment_id = s.id AND us.user_id=$1 AND s.name=$2`,
		userId, segmentName)
	if err != nil {
		return false, fmt.Errorf("failed to delete user segment: %w", err)
	}

	return res.RowsAffected() != 0, nil
}

func (r *Repository) FlushExpired(ctx context.Context) error {
	_, err := db.FromContext(ctx).Exec(ctx,
		`DELETE FROM user_segment WHERE expires < $1`, time.Now())
	if err != nil {
		return fmt.Errorf("failed to flush expired entries: %w", err)
	}

	return nil
}
