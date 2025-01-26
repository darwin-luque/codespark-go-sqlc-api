package middlewares

import (
	"github.com/darwin-luque/codespark-go-sqlc-api/domain"
	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/config"
)

type Middlewares struct {
	cfg config.Interface
	db  domain.DBTX
}

func New(cfg config.Interface, db domain.DBTX) *Middlewares {
	return &Middlewares{cfg: cfg, db: db}
}
