package favorites

import (
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/config"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/infrastructure/middlewares"
	"github.com/gorilla/mux"
)

type FavoritesModule struct {
	db  domain.DBTX
	cfg config.Interface
	m   *middlewares.Middlewares
}

func New(db domain.DBTX, cfg config.Interface, m *middlewares.Middlewares) *FavoritesModule {
	return &FavoritesModule{
		db:  db,
		cfg: cfg,
		m:   m,
	}
}

func (fm *FavoritesModule) RegisterRoutes(baseRouter *mux.Router) {
	withAuth := baseRouter.PathPrefix("").Subrouter()
	withAuth.Use(fm.m.Authenticate(utils.RequiredAuth))

	articleBasedFavoritesRouter := withAuth.PathPrefix("/articles/{slug}/favorites").Subrouter()
	{
		articleBasedFavoritesRouter.HandleFunc("", fm.addArticle()).Methods(http.MethodPost)
		articleBasedFavoritesRouter.HandleFunc("", fm.removeArticle()).Methods(http.MethodDelete)
	}
}
