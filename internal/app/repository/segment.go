package repository

import (
	db "avito-backend/pkg/gopkg-db"
	"context"
	"fmt"
)

func (r *Repository) AddSegmentIfNotExists(ctx context.Context, name string) (int, error) {
	var id int
	row := db.FromContext(ctx).QueryRow(ctx,
		`INSERT INTO segment (name) VALUES ($1) ON CONFLICT (name) DO UPDATE SET name = $1 RETURNING id`,
		name)

	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("can't scan row on adding segment: %w", err)
	}

	return id, nil
}

func (r *Repository) DeleteSegment(ctx context.Context, name string) error {
	result, err := db.FromContext(ctx).Exec(ctx, `DELETE FROM segment WHERE name = $1`, name)
	if err != nil {
		return fmt.Errorf("failed to delete segment: %w", err)
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("segment %s does not exist", name)
	}

	return nil
}

func (r *Repository) GetSegmentName(ctx context.Context, id int) (string, error) {
	var name string
	row := db.FromContext(ctx).QueryRow(ctx, `SELECT name FROM segment WHERE id=$1`, id)

	err := row.Scan(&name)
	if err != nil {
		return "", fmt.Errorf("cant get segment name with row.Scan() %w", err)
	}

	return name, nil
}

func (r *Repository) GetSegmentId(ctx context.Context, name string) (int, error) {
	var id int
	row := db.FromContext(ctx).QueryRow(ctx, `SELECT id FROM segment WHERE name=$1`, name)

	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("cant get segment id with row.Scan() %w", err)
	}

	return id, nil
}
