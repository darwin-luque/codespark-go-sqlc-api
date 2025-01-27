package articles

import (
	"errors"
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
		authenticatedRouter.HandleFunc("", am.listArticles()).Methods(http.MethodGet)
		authenticatedRouter.HandleFunc("/{slug}", am.getArticleBySlug()).Methods(http.MethodGet)
	}

	blogOwnerRouter := authenticatedRouter.PathPrefix("").Subrouter()
	blogOwnerRouter.Use(am.checkBlogOwnership("slug"))
	{
		blogOwnerRouter.HandleFunc("/{slug}", am.updateArticleBySlug()).Methods(http.MethodPut)
		blogOwnerRouter.HandleFunc("/{slug}", am.deleteArticleBySlug()).Methods(http.MethodDelete)
	}
}

func (am *ArticlesModule) checkBlogOwnership(whereToCheck string) func(http.Handler) http.Handler {
	q := domain.New(am.db)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			user, ok := req.Context().Value(middlewares.USER_KEY).(*domain.User)

			if !ok {
				utils.ErrorResponse(w, http.StatusInternalServerError, errors.New("invalid user context"))
				return
			}
			var article domain.Article

			switch whereToCheck {
			case "slug":
				slug := mux.Vars(req)["slug"]
				foundArticle, err := q.GetArticleBySlug(req.Context(), slug)

				if err != nil {
					utils.NotFoundError(w, utils.ErrorM{})
					return
				}

				article = foundArticle
			default:
				utils.ErrorResponse(w, http.StatusInternalServerError, errors.New("invalid whereToCheck"))
			}

			if article.AuthorID != user.ID {
				utils.ErrorResponse(w, http.StatusForbidden, errors.New("forbidden"))
				return
			}

			next.ServeHTTP(w, req)
		})
	}
}
