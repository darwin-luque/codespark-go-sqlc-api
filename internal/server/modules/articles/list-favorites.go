package articles

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/infrastructure/middlewares"
)

func (am *ArticlesModule) listFavorites() http.HandlerFunc {
	q := domain.New(am.db)

	type ListFavoriteArticlesInput struct {
		Limit  *int32
		Offset *int32
	}

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()
		input := ListFavoriteArticlesInput{}

		if v := query.Get("offset"); v != "" {
			intV, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				utils.ErrorResponse(w, http.StatusBadRequest, errors.New("invalid offset"))
				return
			}
			offset := int32(intV)
			input.Offset = &offset
		}

		if v := query.Get("limit"); v != "" {
			intV, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				utils.ErrorResponse(w, http.StatusBadRequest, errors.New("invalid limit"))
				return
			}

			limit := int32(intV)
			input.Limit = &limit
		}

		user, ok := req.Context().Value(middlewares.USER_KEY).(*domain.User)

		if !ok {
			utils.ErrorResponse(w, http.StatusInternalServerError, errors.New("invalid user context"))
			return
		}

		listFavoriteInput := domain.ListFavoriteArticlesParams{
			UserID: user.ID,
			Limit:  input.Limit,
			Offset: input.Offset,
		}

		articles, err := q.ListFavoriteArticles(req.Context(), listFavoriteInput)

		if err != nil {
			utils.ServerError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, utils.M{"articles": articles})
	})
}
