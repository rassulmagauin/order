package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/rassulmagauin/order/internal/env"
	"go.uber.org/zap"
)

func DBConnect(logger *zap.Logger, env *env.Env) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), env.DatabaseUrl)
	if err != nil {
		logger.Error("Error loading database:", zap.Error(err))
		os.Exit(1)
	}
	return conn
}
