package service

import (
	"context"
	"fmt"
)

func (s *Service) GetSegmentsOfUser(ctx context.Context, userID int) ([]string, error) {
	var segmentsOfUser []string
	segmentIds, err := s.repo.GetSegmentsOfUser(ctx, userID)
	if err != nil {
		return []string{}, fmt.Errorf("cannot get segments of user %d: %w", userID, err)
	}

	for _, id := range segmentIds {
		segmentName, err := s.repo.GetSegmentName(ctx, id)
		if err != nil {
			return []string{}, fmt.Errorf("cannot get segment name, id - %d: %w", id, err)
		}
		segmentsOfUser = append(segmentsOfUser, segmentName)
	}
	return segmentsOfUser, nil
}
