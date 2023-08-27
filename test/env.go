package test

import (
	"avito-backend/internal/app/handler"
	"avito-backend/internal/app/repository"
	"avito-backend/internal/app/service"
	"avito-backend/internal/pkg/config"
	db "avito-backend/pkg/gopkg-db"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"path"
	"syscall"
)

type Env struct {
	Ctx  context.Context
	Hdl  *handler.Handler
	Repo *repository.Repository
	Srv  *service.Service
}

func NewEnv() (*Env, func()) {
	wd, _ := syscall.Getwd()
	_ = syscall.Chdir(path.Dir(wd))

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load cfg: ", err)
	}
	ctx := context.Background()

	conn, err := db.New(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatal(fmt.Errorf("cant create connection to db: %v", err))
	}
	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Fatal(fmt.Errorf("cant create tx: %v", err))
	}
	ctxWithDb := db.AddToContext(ctx, db.DBTx{Tx: tx})

	repo := repository.New()

	srv := service.New(cfg, repo)
	hdl := handler.New(srv)

	return &Env{
			Ctx:  ctxWithDb,
			Srv:  srv,
			Repo: repo,
			Hdl:  hdl,
		}, func() {
			err := tx.Rollback(ctxWithDb)
			if err != nil {
				log.Fatal("cannot rollback transaction")
			}
		}
}
