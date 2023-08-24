package service

import (
	"github.com/docker/docker/daemon/config"
)

type IRepository interface {
}

type Service struct {
	cfg  *config.Config
	repo IRepository
}

func New(cfg *config.Config, r IRepository) *Service {
	return &Service{cfg, r}
}
