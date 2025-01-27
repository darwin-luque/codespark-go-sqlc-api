package follows

import (
	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/config"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/infrastructure/middlewares"
	"github.com/gorilla/mux"
)

type FollowsModule struct {
	db  domain.DBTX
	cfg config.Interface
	m   *middlewares.Middlewares
}

func New(db domain.DBTX, cfg config.Interface, m *middlewares.Middlewares) *FollowsModule {
	return &FollowsModule{
		db:  db,
		cfg: cfg,
		m:   m,
	}
}

func (fm *FollowsModule) RegisterRoutes(baseRouter *mux.Router) {
	followsRouter := baseRouter.PathPrefix("/follows").Subrouter()

	authenticatedRouter := followsRouter.PathPrefix("").Subrouter()
	authenticatedRouter.Use(fm.m.Authenticate(utils.RequiredAuth))
	{
		authenticatedRouter.HandleFunc("/{username}", fm.followUser()).Methods("POST")
		authenticatedRouter.HandleFunc("/{username}", fm.unfollowUser()).Methods("DELETE")
	}
}
