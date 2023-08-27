package service

import (
	"context"
	"fmt"
)

func (s *Service) DeleteSegment(ctx context.Context, name string) error {
	err := s.repo.DeleteSegment(ctx, name)
	if err != nil {
		return fmt.Errorf("cannot delete segment: %w", err)
	}
	return nil
}
