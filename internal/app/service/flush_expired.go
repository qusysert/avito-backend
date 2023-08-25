package service

import "context"

func (s *Service) FlushExpired(ctx context.Context) error {
	return s.repo.FlushExpired(ctx)
}
