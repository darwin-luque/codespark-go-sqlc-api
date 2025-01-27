package articles

import (
	"errors"
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/gorilla/mux"
)

func (am *ArticlesModule) publish() http.HandlerFunc {
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

		publishedStatus := "published"
		params := domain.PublishArticleParams{
			Status: &publishedStatus,
			Slug:   slug,
		}

		article, err := q.PublishArticle(req.Context(), params)

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
