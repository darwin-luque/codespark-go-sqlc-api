package articles

import (
	"errors"
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/gorilla/mux"
)

func (am *ArticlesModule) updateBySlug() http.HandlerFunc {
	type Input struct {
		Title       *string `json:"title"`
		Body        *string `json:"content"`
		Description *string `json:"description"`
		Status      *string `json:"status"`
	}

	q := domain.New(am.db)

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		var slug string
		if foundSlug, exists := vars["slug"]; exists {
			slug = foundSlug
		} else {
			utils.ErrorResponse(w, http.StatusBadRequest, errors.New("missing slug"))
			return
		}

		input := &Input{}
		if err := utils.ReadJSON(req.Body, input); err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		if err := utils.Validate.Struct(input); err != nil {
			utils.ValidationError(w, err)
			return
		}

		params := domain.UpdateArticleParams{
			Title:       input.Title,
			Body:        input.Body,
			Description: input.Description,
			Status:      input.Status,
			Slug:        slug,
		}

		article, err := q.UpdateArticle(req.Context(), params)

		if err != nil {
			switch {
			case err.Error() == "not found":
				utils.ErrorResponse(w, http.StatusNotFound, err)
			default:
				utils.ErrorResponse(w, http.StatusInternalServerError, err)
			}
			return
		}

		utils.WriteJSON(w, http.StatusOK, utils.M{"article": article})
	})
}
