package setup

import (
	"context"

	"github.com/darwin-luque/codespark-go-sqlc-api/internal/common/config"
	"github.com/jackc/pgx/v5"
)

func SetupDatabase(cfg config.DatabaseInterface) (*pgx.Conn, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, cfg.URL())

	if err != nil {
		return nil, err
	}

	return conn, nil
}
