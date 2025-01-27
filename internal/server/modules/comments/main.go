package comments

import (
	"errors"
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/config"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/infrastructure/middlewares"
	"github.com/google/uuid"
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
	withAuth := baseRouter.PathPrefix("").Subrouter()
	withAuth.Use(cm.m.Authenticate(utils.RequiredAuth))

	articleBasedCommentsRouter := withAuth.Path("/articles/{slug}/comments").Subrouter()
	{
		articleBasedCommentsRouter.HandleFunc("", cm.add()).Methods(http.MethodPost)
		articleBasedCommentsRouter.HandleFunc("", cm.listForArticle()).Methods(http.MethodGet)
	}

	withIsCommentOwner := withAuth.PathPrefix("/comments/{id}").Subrouter()
	withIsCommentOwner.Use(cm.checkCommentOwnership())
	{
		withIsCommentOwner.HandleFunc("", cm.delete()).Methods(http.MethodDelete)
	}
}

func (cm *CommentsModule) checkCommentOwnership() func(http.Handler) http.Handler {
	q := domain.New(cm.db)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			commentID := vars["id"]

			commentUUID, err := uuid.Parse(commentID)

			if err != nil {
				http.Error(w, "Invalid comment ID", http.StatusBadRequest)
				return
			}

			user, ok := r.Context().Value(middlewares.USER_KEY).(*domain.User)

			if !ok {
				utils.ErrorResponse(w, http.StatusInternalServerError, errors.New("invalid user context"))
				return
			}

			comment, err := q.GetComment(r.Context(), commentUUID)

			if err != nil {
				utils.ErrorResponse(w, http.StatusInternalServerError, err)
				return
			}

			if comment.AuthorID != user.ID {
				utils.ErrorResponse(w, http.StatusForbidden, errors.New("unauthorized"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
