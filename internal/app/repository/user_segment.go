package repository

import (
	db "avito-backend/pkg/gopkg-db"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (r *Repository) AddUserSegment(ctx context.Context, userId, segmentId int, expires *time.Time) (int, error) {
	var id int
	var expiresValue sql.NullTime

	if expires != nil {
		expiresValue.Time = *expires
		expiresValue.Valid = true
	}

	row := db.FromContext(ctx).QueryRow(ctx,
		`INSERT INTO user_segment (user_id, segment_id, expires) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING id`,
		userId, segmentId, expiresValue)

	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to scan row on adding segment: %w", err)
	}

	return id, nil
}
