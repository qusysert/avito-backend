package service

import (
	config "avito-backend/internal/pkg/config"
)

type IRepository interface {
}

type Service struct {
	cfg  config.Config
	repo IRepository
}

func New(cfg config.Config, r IRepository) *Service {
	return &Service{cfg, r}
}
