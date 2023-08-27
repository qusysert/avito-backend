package service

import (
	"avito-backend/internal/app/model"
	"context"
	"fmt"
)

func (s *Service) GetSegmentsOfUser(ctx context.Context, userID int) ([]model.SegmentWithExpires, error) {

	segments, err := s.repo.GetUserSegments(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("cannot get segments of user %d: %w", userID, err)
	}

	return segments, nil
}
