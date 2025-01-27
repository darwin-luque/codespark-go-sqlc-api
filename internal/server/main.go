package server

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/config"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/infrastructure/middlewares"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/server/modules/articles"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/server/modules/auth"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	server   *http.Server
	router   *mux.Router
	cfg      config.Interface
	db       domain.DBTX
	m        *middlewares.Middlewares
	auth     *auth.AuthModule
	articles *articles.ArticlesModule
}

func New(cfg config.Interface, db domain.DBTX) (*Server, error) {
	m := middlewares.New(cfg, db)
	s := Server{
		server: &http.Server{
			WriteTimeout: 5 * time.Second,
			ReadTimeout:  5 * time.Second,
			IdleTimeout:  5 * time.Second,
		},
		cfg:    cfg,
		router: mux.NewRouter().StrictSlash(true),
		db:     db,
		m:      m,
	}

	s.initModules()
	s.routes()

	s.server.Handler = s.router

	return &s, nil
}

func (s *Server) initModules() {
	s.auth = auth.New(s.db, s.cfg, s.m)
	s.articles = articles.New(s.db, s.cfg, s.m)
}

func (s *Server) Run() error {
	port := s.cfg.Server().Port()
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	s.server.Addr = port
	log.Printf("server starting on %s", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) routes() {
	s.router.Use(cors.AllowAll().Handler)
	s.router.Use(s.m.Logger(os.Stdout))

	apiRouter := s.router.PathPrefix("/api").Subrouter()
	apiRouter.Handle("/health", healthCheck()).Methods(http.MethodGet)

	s.auth.RegisterRoutes(apiRouter)
	s.articles.RegisterRoutes(apiRouter)
}

func healthCheck() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		resp := utils.M{
			"status":  "available",
			"message": "healthy",
			"data":    utils.M{"hello": "beautiful"},
		}
		utils.WriteJSON(rw, http.StatusOK, resp)
	})
}
