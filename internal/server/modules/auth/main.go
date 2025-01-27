package auth

import (
	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/config"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/infrastructure/middlewares"
	"github.com/gorilla/mux"
)

type AuthModule struct {
	db  domain.DBTX
	cfg config.Interface
	m   *middlewares.Middlewares
}

func New(db domain.DBTX, cfg config.Interface, m *middlewares.Middlewares) *AuthModule {
	return &AuthModule{db, cfg, m}
}

func (am *AuthModule) RegisterRoutes(baseRouter *mux.Router) {
	authRouter := baseRouter.PathPrefix("/auth").Subrouter()

	noAuthRouter := authRouter.PathPrefix("").Subrouter()
	{
		noAuthRouter.HandleFunc("/sign-up", am.signUp()).Methods("POST")
		noAuthRouter.HandleFunc("/sign-in", am.signIn()).Methods("POST")
	}
}
