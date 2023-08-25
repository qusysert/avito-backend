package service

import "context"

func (s *Service) AddSegment(ctx context.Context, name string) (int, error) {
	return s.repo.AddSegmentIfNotExists(ctx, name)
}
