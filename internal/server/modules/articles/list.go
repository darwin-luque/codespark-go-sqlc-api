package articles

import (
	"log"
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
)

func (am *ArticlesModule) list() http.HandlerFunc {
	q := domain.New(am.db)

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()
		input := domain.ListArticlesParams{}

		if v := query.Get("title"); v != "" {
			input.Title = &v
		}

		if v := query.Get("slug"); v != "" {
			input.Slug = &v
		}

		log.Println(input)

		articles, err := q.ListArticles(req.Context(), input)

		if err != nil {
			utils.ServerError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, utils.M{"articles": articles})
	})
}
