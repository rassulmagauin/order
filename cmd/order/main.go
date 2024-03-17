package main

import (
	"context"
	"fmt"

	"github.com/rassulmagauin/order/internal/db"
	"github.com/rassulmagauin/order/pkg/logger"
)

func run() { //
	logger := logger.NewLogger()
	env := LoadConfig(logger)
	conn := db.DBConnect(logger, env)

	defer conn.Close(context.Background())
	fmt.Printf("Config: %+v\n", *env)
}

func main() {
	run()
}
