package service

import (
	"avito-backend/internal/app/model"
	config "avito-backend/internal/pkg/config"
	"context"
	"time"
)

type IRepository interface {
	AddUserSegmentIfNotExists(ctx context.Context, userId, segmentId int, expires *time.Time) (int, error)
	DeleteSegment(ctx context.Context, name string) error
	AddSegmentIfNotExists(ctx context.Context, name string) (int, error)
	GetSegmentId(ctx context.Context, name string) (int, error)
	DeleteUserSegmentIfExists(ctx context.Context, userId int, segmentName string) (bool, error)
	GetUserSegments(ctx context.Context, id int) ([]model.SegmentWithExpires, error)
	GetSegmentName(ctx context.Context, id int) (string, error)
	FlushExpired(ctx context.Context) error
}

type Service struct {
	cfg  config.Config
	repo IRepository
}

func New(cfg config.Config, r IRepository) *Service {
	return &Service{cfg, r}
}
