package main

import (
	"avito-backend/internal/app/handler"
	"avito-backend/internal/pkg/config"
	db "avito-backend/pkg/gopkg-db"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	_ "github.com/swaggo/http-swagger/example/gorilla/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

var conn db.IClient

//	@title          Backend Trainee Assignment 2023
//	@version		1.0
//	@description	Swagger documentation fo Backend Trainee Assignment 2023 service

//	@contact.name	Ivan Demchuk
//	@contact.email	is.demchuk@gmail.com

// @host		localhost:8080
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load cfg")
	}

	// Creating connection to DB
	conn, err = db.New(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatal(fmt.Errorf("cant create connection to db: %v", err))
	}

}

func serverFn(ctx context.Context, cfg config.Config, hdl *handler.Handler) func() error {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	externalRouter := mux.NewRouter()

	// Setting timeout for the externalServer
	externalServer := &http.Server{
		Addr:         ":" + cfg.HttpPort,
		ReadTimeout:  600 * time.Second,
		WriteTimeout: 600 * time.Second,
		Handler:      c.Handler(externalRouter),
	}

	// Linking addresses and handlers
	for _, rec := range [...]struct {
		route       string
		handler     http.HandlerFunc
		withoutAuth bool
	}{
		{route: "/swagger.json", handler: func(w http.ResponseWriter, r *http.Request) {
			cwd, _ := os.Getwd()
			http.ServeFile(w, r, path.Join(cwd, "docs/swagger.json"))
		}, withoutAuth: true},
		{route: "/swagger/{any:.+}", handler: httpSwagger.Handler(httpSwagger.URL("/swagger.json")), withoutAuth: true},
	} {
		externalRouter.HandleFunc(rec.route, DbMiddleware(rec.handler))
	}

	return func() error {
		errCh := make(chan error)
		go func() {
			errCh <- externalServer.ListenAndServe()
		}()
		var err error
		select {
		case serverErr := <-errCh:
			err = serverErr
		case <-ctx.Done():
			err = externalServer.Shutdown(ctx)
		}
		log.Printf("External externalServer finished, error: %v\n", err)
		return err
	}
}

func DbMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r = r.WithContext(db.AddToContext(ctx, conn))
		next.ServeHTTP(w, r)
	}
}
