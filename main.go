package main

import (
	"Boosters_test_task/cmd/server"
	_"Boosters_test_task/docs"
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

// @title           Posts service
// @version         1.0
// @description     A posts data management service API in Go.

// @host      127.0.0.1:8080
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := server.NewServer()

	err := srv.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
