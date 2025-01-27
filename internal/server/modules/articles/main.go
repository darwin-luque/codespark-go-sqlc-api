package articles

import (
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/config"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/infrastructure/middlewares"
	"github.com/gorilla/mux"
)

type ArticlesModule struct {
	db  domain.DBTX
	cfg config.Interface
	m   *middlewares.Middlewares
}

func New(db domain.DBTX, cfg config.Interface, m *middlewares.Middlewares) *ArticlesModule {
	return &ArticlesModule{
		db:  db,
		cfg: cfg,
		m:   m,
	}
}

func (am *ArticlesModule) RegisterRoutes(baseRouter *mux.Router) {
	articlesRouter := baseRouter.PathPrefix("/articles").Subrouter()

	authenticatedRouter := articlesRouter.PathPrefix("").Subrouter()
	authenticatedRouter.Use(am.m.Authenticate(utils.RequiredAuth))
	{
		authenticatedRouter.HandleFunc("", am.createArticle()).Methods(http.MethodPost)
	}
}
