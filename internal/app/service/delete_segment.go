package service

import (
	"context"
	"fmt"
)

func (s *Service) DeleteSegment(ctx context.Context, id int) error {
	err := s.repo.DeleteSegmentFromUsers(ctx, id)
	if err != nil {
		return fmt.Errorf("cannot delete segment from users: %w", err)
	}
	err = s.repo.DeleteSegment(ctx, id)
	if err != nil {
		return fmt.Errorf("cannot delete segment: %w", err)
	}
	return nil
}
