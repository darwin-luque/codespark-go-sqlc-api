package comments

import (
	"errors"
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/gorilla/mux"
)

func (cm *CommentsModule) listForArticle() http.HandlerFunc {
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

		comments, err := q.ListCommentsForArticleBySlug(req.Context(), slug)

		if err != nil {
			switch {
			case err.Error() == `sql: no rows in result set`:
				utils.ErrorResponse(w, http.StatusNotFound, errors.New("article not found"))
			default:
				utils.ServerError(w, err)
			}
		}

		utils.WriteJSON(w, http.StatusOK, utils.M{"comments": comments})
	})
}
