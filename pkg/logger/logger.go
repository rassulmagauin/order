package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("cannot initialize zap logger: %v", err)
		os.Exit(1)
	}
	return logger
}
