package auth

import (
	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/gorilla/mux"
)

type AuthModule struct {
	db domain.DBTX
}

func New(db domain.DBTX) *AuthModule {
	return &AuthModule{db}
}

func (m *AuthModule) RegisterRoutes(baseRouter *mux.Router) {
}
