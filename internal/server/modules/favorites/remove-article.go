package favorites

import (
	"errors"
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/infrastructure/middlewares"
	"github.com/gorilla/mux"
)

func (fm *FavoritesModule) removeArticle() http.HandlerFunc {

	q := domain.New(fm.db)

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		var slug string
		if foundSlug, exists := vars["slug"]; exists {
			slug = foundSlug
		} else {
			utils.ErrorResponse(w, http.StatusBadRequest, errors.New("missing slug"))
			return
		}

		user, ok := req.Context().Value(middlewares.USER_KEY).(*domain.GetUserByEmailRow)

		if !ok {
			utils.ErrorResponse(w, http.StatusInternalServerError, errors.New("invalid user context"))
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

		favoriteData := domain.RemoveArticleAsFavoriteParams{
			ArticleID: article.ID,
			UserID:    user.ID,
		}

		if err := q.RemoveArticleAsFavorite(req.Context(), favoriteData); err != nil {
			utils.ServerError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, utils.M{})
	})
}
