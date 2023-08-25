package service

import (
	"avito-backend/internal/app/model"
	"context"
	"fmt"
	"time"
)

func (s *Service) AddDeleteUserSegment(ctx context.Context, userId int, toAdd []model.Segment, toDelete []string) ([]int, error) {
	var addedIds []int

	for _, segment := range toAdd {
		segmentId, err := s.repo.AddSegmentIfNotExists(ctx, segment.Name)
		if err != nil {
			return []int{}, fmt.Errorf("cannot add segment %s: %w", segment, err)
		}

		expires, err := time.Parse("2006-01-02T15:04:05", segment.Expires)
		if err != nil {
			return []int{}, fmt.Errorf("cannot parse expire date: %w", err)
		}

		userSegmentId, err := s.repo.AddUserSegmentIfNotExists(ctx, userId, segmentId, &expires)
		if err != nil {
			return []int{}, fmt.Errorf("cannot add user_segment, user id - %d, segment - %s: %w", userId, segment.Name, err)
		}
		addedIds = append(addedIds, userSegmentId)
	}

	for _, segment := range toDelete {
		segmentId, err := s.repo.GetSegmentId(ctx, segment)
		if err != nil {
			if err.Error() == "cant get segment id with row.Scan() no rows in result set" {
				continue
			}
			return []int{}, fmt.Errorf("cannot get id of segment %s: %w", segment, err)
		}

		err = s.repo.DeleteUserSegmentIfExists(ctx, userId, segmentId)
		return []int{}, fmt.Errorf("cannot delete user_segment, user id - %d, segment - %s: %w", userId, segment, err)
	}

	return addedIds, nil
}
