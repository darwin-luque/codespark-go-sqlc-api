package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"golang.org/x/crypto/bcrypt"
)

func (am *AuthModule) signIn() http.HandlerFunc {
	type Input struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
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

		user, err := q.GetUserByEmail(req.Context(), input.Email)
		if err != nil {
			log.Println(err)
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("invalid email or password"))
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
		if err != nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("invalid email or password"))
			return
		}

		parsedUser := SafeUser{
			ID:             user.ID,
			Username:       user.Username,
			Email:          user.Email,
			Bio:            user.Bio,
			Image:          user.Image,
			FollowersCount: user.FollowersCount,
			FollowingCount: user.FollowingCount,
			CreatedAt:      user.CreatedAt,
			UpdatedAt:      user.UpdatedAt,
		}
		token, err := am.generateUserToken(parsedUser)

		if err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, utils.M{"token": token})
	})
}
