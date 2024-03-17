package main

import (
	"os"

	"go.uber.org/zap"

	"github.com/kelseyhightower/envconfig"
	"github.com/rassulmagauin/order/internal/env"
)

func LoadConfig(logger *zap.Logger) *env.Env {
	var s env.Env
	err := envconfig.Process("", &s)
	if err != nil {
		logger.Error("Error processing environment configuration", zap.Error(err))
		os.Exit(1)
	}
	return &s
}
