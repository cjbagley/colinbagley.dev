package main

import (
	"context"
	"github.com/cjbagley/colinbagley.dev/internal/server"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go server.StartServer(ctx, &wg)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	server.LogInfo("Closing server...")

	cancel()
	wg.Wait()

	return nil
}

func main() {
	if err := run(); err != nil {
		server.LogError(err)
		os.Exit(1)
	}
}
