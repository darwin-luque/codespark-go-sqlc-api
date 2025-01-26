package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/utils"
	"github.com/golang-jwt/jwt"
)

var anonymousUser = domain.User{}

func (m *Middlewares) Authenticate(mustAuth bool) func(http.Handler) http.Handler {
	q := domain.New(m.db)
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Vary", "Authorization")
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				if mustAuth {
					utils.InvalidAuthTokenError(w)
					return
				}
				r = setContextUser(r, &anonymousUser)
				h.ServeHTTP(w, r)
				return
			}

			ss := strings.Split(authHeader, " ")

			if len(ss) < 2 {
				utils.InvalidAuthTokenError(w)
				return
			}

			token := ss[1]

			u, err := m.parseUserToken(token)

			if err != nil {
				utils.InvalidAuthTokenError(w)
				return
			}

			email := u["email"].(string)

			user, err := q.GetUserByEmail(r.Context(), email)

			if err != nil {
				utils.ServerError(w, err)
				return
			}

			r = setContextUser(r, &user)
			r = setContextUserToken(r, token)
			h.ServeHTTP(w, r)
		})
	}
}

func (m *Middlewares) parseUserToken(tokenStr string) (userClaims utils.M, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(m.cfg.Auth().Secret()), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, nil
	}

	return utils.M(claims), nil
}
