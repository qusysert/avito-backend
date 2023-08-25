package main

import (
	_ "avito-backend/docs"
	"avito-backend/internal/app/handler"
	"avito-backend/internal/app/repository"
	"avito-backend/internal/app/service"
	"avito-backend/internal/pkg/config"
	db "avito-backend/pkg/gopkg-db"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

// @contact.name	Ivan Demchuk
// @contact.email	is.demchuk@gmail.com
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

	repo := repository.New()

	srv := service.New(cfg, repo)
	hdl := handler.New(srv)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	router := mux.NewRouter()

	// Setting timeout for the server
	server := &http.Server{
		Addr:         "0.0.0.0:" + cfg.HttpPort,
		ReadTimeout:  600 * time.Second,
		WriteTimeout: 600 * time.Second,
		Handler:      c.Handler(router),
	}

	// Linking addresses and handlers
	for _, rec := range [...]struct {
		route   string
		handler http.HandlerFunc
	}{
		{route: "/swagger.json", handler: func(w http.ResponseWriter, r *http.Request) {
			cwd, _ := os.Getwd()
			http.ServeFile(w, r, path.Join(cwd, "docs/swagger.json"))
		}},
		{route: "/swagger/{any:.+}", handler: httpSwagger.Handler(httpSwagger.URL("/swagger.json"))},
		{route: "/deleteSegment/{id}", handler: hdl.DeleteSegmentHandler},
		{route: "/addSegment", handler: hdl.AddSegmentHandler},
		{route: "/addDeleteUserSegment", handler: hdl.AddDeleteUserSegmentHandler},
		{route: "/flushExpired", handler: hdl.FlushExpiredHandler},
		{route: "/getSegmentsOfUser/{id}", handler: hdl.GetSegmentsOfUserHandler},
	} {
		router.HandleFunc(rec.route, DbMiddleware(rec.handler))
	}

	http.Handle("/", router)

	log.Printf("Server started on port %s \n", cfg.HttpPort)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

}

func DbMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r = r.WithContext(db.AddToContext(ctx, conn))
		next.ServeHTTP(w, r)
	}
}
