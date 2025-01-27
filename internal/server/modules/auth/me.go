package auth

import (
	"errors"
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/infrastructure/middlewares"
)

func (am *AuthModule) me() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		user, ok := req.Context().Value(middlewares.USER_KEY).(*domain.User)

		if !ok {
			utils.ErrorResponse(w, http.StatusInternalServerError, errors.New("invalid user context"))
			return
		}

		bio := ""
		if user.Bio != nil {
			bio = *user.Bio
		}
		safeUser := SafeUser{
			ID:        user.ID,
			Email:     user.Email,
			Username:  user.Username,
			Image:     user.Image,
			Bio:       bio,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		utils.WriteJSON(w, http.StatusOK, utils.M{"user": safeUser})
	})
}
