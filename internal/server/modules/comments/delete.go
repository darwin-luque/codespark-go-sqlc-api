package comments

import (
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (cm *CommentsModule) delete() http.HandlerFunc {
	q := domain.New(cm.db)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		commentID := vars["id"]

		commentUUID, err := uuid.Parse(commentID)

		if err != nil {
			http.Error(w, "Invalid comment ID", http.StatusBadRequest)
			return
		}

		if err = q.DeleteComment(r.Context(), commentUUID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.WriteJSON(w, http.StatusNoContent, nil)
	})
}
