package middlewares

import (
	"context"
	"net/http"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
)

type contextKey string

const (
	USER_KEY  contextKey = "user"
	TOKEN_KEY contextKey = "token"
)

func setContextUser(r *http.Request, u *domain.User) *http.Request {
	ctx := context.WithValue(r.Context(), USER_KEY, u)
	return r.WithContext(ctx)
}

func setContextUserToken(r *http.Request, token string) *http.Request {
	ctx := context.WithValue(r.Context(), TOKEN_KEY, token)
	return r.WithContext(ctx)
}
