package comments

import (
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/config"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/infrastructure/middlewares"
	"github.com/gorilla/mux"
)

type CommentsModule struct {
	db  domain.DBTX
	cfg config.Interface
	m   *middlewares.Middlewares
}

func New(db domain.DBTX, cfg config.Interface, m *middlewares.Middlewares) *CommentsModule {
	return &CommentsModule{
		db:  db,
		cfg: cfg,
		m:   m,
	}
}

func (cm *CommentsModule) RegisterRoutes(baseRouter *mux.Router) {
	articleBasedCommentsRouter := baseRouter.Path("/articles/{slug}/comments").Subrouter()
	commentsRouter := baseRouter.Path("/comments").Subrouter()

	authenticatedArticleBasedCommentsRouter := articleBasedCommentsRouter.PathPrefix("").Subrouter()
	authenticatedArticleBasedCommentsRouter.Use(cm.m.Authenticate(utils.RequiredAuth))
	{
		authenticatedArticleBasedCommentsRouter.HandleFunc("", cm.add()).Methods(http.MethodPost)
		authenticatedArticleBasedCommentsRouter.HandleFunc("", cm.listForArticle()).Methods(http.MethodGet)
	}

	authenticatedRouter := commentsRouter.PathPrefix("").Subrouter()
	authenticatedRouter.Use(cm.m.Authenticate(utils.RequiredAuth))
	{
		authenticatedRouter.HandleFunc("/{id}", cm.delete()).Methods(http.MethodDelete)
	}
}
