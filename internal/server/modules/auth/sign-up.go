package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"golang.org/x/crypto/bcrypt"
)

func (am *AuthModule) signUp() http.HandlerFunc {
	type Input struct {
		Email    string `json:"email" validate:"required,email"`
		Username string `json:"username" validate:"required,min=2"`
		Password string `json:"password" validate:"required,min=8,max=72"`
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

		hashBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

		if err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		user := domain.CreateUserParams{
			Email:        input.Email,
			Username:     input.Username,
			PasswordHash: string(hashBytes),
			Image:        "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png",
		}

		newUser, err := q.CreateUser(req.Context(), user)
		if err != nil {
			log.Fatalln(err)
			switch {
			case err.Error() == `pq: duplicate key value violates unique constraint "user_email_key"`:
				utils.ErrorResponse(w, http.StatusConflict, errors.New("email already exists"))
				return
			case err.Error() == `pq: duplicate key value violates unique constraint "user_username_key"`:
				utils.ErrorResponse(w, http.StatusConflict, errors.New("username already exists"))
				return
			default:
				utils.ErrorResponse(w, http.StatusInternalServerError, err)
				return
			}
		}

		parsedUser := SafeUser{
			ID:             newUser.ID,
			Username:       newUser.Username,
			Email:          newUser.Email,
			Bio:            newUser.Bio,
			Image:          newUser.Image,
			FollowersCount: 0,
			FollowingCount: 0,
			CreatedAt:      newUser.CreatedAt,
			UpdatedAt:      newUser.UpdatedAt,
		}

		token, err := am.generateUserToken(parsedUser)

		if err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		utils.WriteJSON(w, http.StatusCreated, utils.M{"token": token})
	})
}
