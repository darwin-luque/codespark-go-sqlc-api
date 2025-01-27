package follows

import (
	"errors"
	"log"
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/infrastructure/middlewares"
	"github.com/gorilla/mux"
)

func (fm *FollowsModule) followUser() http.HandlerFunc {
	q := domain.New(fm.db)

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		var username string
		if foundUsername, exists := vars["username"]; exists {
			username = foundUsername
		} else {
			utils.ErrorResponse(w, http.StatusBadRequest, errors.New("missing username"))
			return
		}

		user, ok := req.Context().Value(middlewares.USER_KEY).(*domain.GetUserByEmailRow)

		if !ok {
			log.Println("invalid user context")
			utils.ErrorResponse(w, http.StatusInternalServerError, errors.New("invalid user context"))
			return
		}

		if user.Username == username {
			log.Println("cannot follow yourself")
			utils.ErrorResponse(w, http.StatusBadRequest, errors.New("cannot follow yourself"))
			return
		}

		userToFollow, err := q.GetUserByUsername(req.Context(), username)

		if err != nil {
			log.Println("error getting user to follow")
			utils.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		followUserInput := domain.FollowUserParams{
			FollowingUserID: user.ID,
			FollowedUserID:  userToFollow.ID,
		}

		follow, err := q.FollowUser(req.Context(), followUserInput)

		if err != nil {
			log.Println("error following user")
			utils.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		utils.WriteJSON(w, http.StatusCreated, follow)
	}
}
