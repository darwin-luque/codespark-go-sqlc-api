package articles

import (
	"errors"
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/gorilla/mux"
)

func (am *ArticlesModule) getArticleBySlug() http.HandlerFunc {
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

		article, err := q.GetArticleBySlug(req.Context(), slug)

		if err != nil {
			switch {
			case err.Error() == "no rows in result set":
				utils.NotFoundError(w, utils.ErrorM{})
			default:
				utils.ServerError(w, err)
			}
			return
		}

		utils.WriteJSON(w, http.StatusOK, utils.M{"article": article})
	})
}
