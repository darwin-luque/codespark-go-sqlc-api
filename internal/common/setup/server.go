package setup

import (
	"log"
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/config"
	"github.com/gorilla/mux"
)

type Server struct {
	server *http.Server
	router *mux.Router
	db     domain.DBTX
}

func NewServer(cfg config.ServerInterface, db domain.DBTX) (*Server, error) {
	router := mux.NewRouter()
	server := &http.Server{
		Addr:    cfg.Host() + ":" + cfg.Port(),
		Handler: router,
	}
	return &Server{server, router, db}, nil
}

func (s *Server) Run() error {
	log.Printf("server starting on %s", s.server.Addr)
	return s.server.ListenAndServe()
}
