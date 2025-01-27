package articles

import (
	"errors"
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/infrastructure/middlewares"
)

func (am *ArticlesModule) create() http.HandlerFunc {
	type Input struct {
		Title       string  `json:"title" validate:"required"`
		Body        string  `json:"body" validate:"required"`
		Description string  `json:"description"`
		Slug        *string `json:"slug"`
	}

	q := domain.New(am.db)

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		input := &Input{}

		if err := utils.ReadJSON(req.Body, input); err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		if err := utils.Validate.Struct(input); err != nil {
			utils.ValidationError(w, err)
			return
		}

		user, ok := req.Context().Value(middlewares.USER_KEY).(*domain.GetUserByEmailRow)

		if !ok {
			utils.ErrorResponse(w, http.StatusInternalServerError, errors.New("invalid user context"))
			return
		}

		// autogenerate slug if not provided
		if input.Slug == nil {
			slug := utils.GenerateSlug(input.Title)
			input.Slug = &slug
		}

		status := "draft"
		articleParams := domain.CreateArticleParams{
			Title:       input.Title,
			Body:        input.Body,
			Description: &input.Description,
			Slug:        *input.Slug,
			AuthorID:    user.ID,
			Status:      &status,
		}

		article, err := q.CreateArticle(req.Context(), articleParams)

		if err != nil {
			switch {
			case err.Error() == "duplicate key value violates unique constraint \"articles_slug_key\"":
				utils.ErrorResponse(w, http.StatusConflict, errors.New("slug already exists"))
			default:
				utils.ServerError(w, err)
			}
		}

		utils.WriteJSON(w, http.StatusCreated, utils.M{"article": article})
	})
}
