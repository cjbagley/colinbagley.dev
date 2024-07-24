package main

import (
	"context"
	"fmt"
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

	fmt.Printf("\nClosing...\n")

	cancel()
	wg.Wait()

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
