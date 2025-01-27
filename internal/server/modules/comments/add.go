package comments

import (
	"errors"
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/infrastructure/middlewares"
	"github.com/gorilla/mux"
)

func (cm *CommentsModule) add() http.HandlerFunc {
	type Input struct {
		Body string `json:"body" validate:"required"`
	}
	q := domain.New(cm.db)

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		var slug string
		if foundSlug, exists := vars["slug"]; exists {
			slug = foundSlug
		} else {
			utils.ErrorResponse(w, http.StatusBadRequest, errors.New("missing slug"))
			return
		}

		user, ok := req.Context().Value(middlewares.USER_KEY).(*domain.User)

		if !ok {
			utils.ErrorResponse(w, http.StatusInternalServerError, errors.New("invalid user context"))
			return
		}

		var input Input
		if err := utils.ReadJSON(req.Body, &input); err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		if err := utils.Validate.Struct(input); err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		article, err := q.GetArticleBySlug(req.Context(), slug)

		if err != nil {
			switch {
			case err.Error() == `sql: no rows in result set`:
				utils.ErrorResponse(w, http.StatusNotFound, errors.New("article not found"))
			default:
				utils.ServerError(w, err)
			}
		}

		commentData := domain.AddCommentParams{
			ArticleID: article.ID,
			AuthorID:  user.ID,
			Body:      input.Body,
		}

		comment, err := q.AddComment(req.Context(), commentData)

		if err != nil {
			utils.ServerError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusCreated, utils.M{"comment": comment})
	})
}
