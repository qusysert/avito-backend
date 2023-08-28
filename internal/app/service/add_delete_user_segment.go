package service

import (
	"avito-backend/internal/app/model"
	db "avito-backend/pkg/gopkg-db"
	"context"
	"fmt"
)

func (s *Service) AddDeleteUserSegment(ctx context.Context, userId int, toAdd []model.SegmentWithExpires, toDelete []string) ([]int, error) {
	var addedIds []int

	if err := db.WithTx(ctx, func(ctx context.Context) error {
		for _, segment := range toAdd {
			segmentId, err := s.repo.AddSegmentIfNotExists(ctx, segment.Name)
			if err != nil {
				return fmt.Errorf("cannot add segment %v: %w", segment, err)
			}

			userSegmentId, err := s.repo.AddUserSegmentIfNotExists(ctx, userId, segmentId, &segment.Expires)
			if err != nil {
				return fmt.Errorf("cannot add user_segment, user id - %d, segment - %s: %w", userId, segment.Name, err)
			}
			addedIds = append(addedIds, userSegmentId)
		}

		for _, segment := range toDelete {
			isDeleted, err := s.repo.DeleteUserSegmentIfExists(ctx, userId, segment)
			if err != nil {
				return fmt.Errorf("cannot delete user_segment, user id - %d, segment - %s: %w", userId, segment, err)
			}
			if !isDeleted {
				return fmt.Errorf("no such segment to delete: %s", segment)
			}

		}
		return nil
	}); err != nil {
		return nil, err
	}

	return addedIds, nil
}
