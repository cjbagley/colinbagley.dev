package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/cjbagley/colinbagley.dev/internal/middleware"
)

func NewServer() http.Server {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./web/assets/"))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", fs))

	AddRoutes(mux)

	var handler http.Handler = mux
	handler = middleware.Cors(handler)

	return http.Server{
		Handler:           handler,
		Addr:              ":8080",
		WriteTimeout:      500 * time.Millisecond,
		ReadTimeout:       500 * time.Millisecond,
		ReadHeaderTimeout: 500 * time.Millisecond,
		IdleTimeout:       1 * time.Second,
	}
}

func StartServer(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	httpServer := NewServer()

	// Adapted from https://medium.com/@dsilverdi/graceful-shutdown-in-go-a-polite-way-to-end-programs-6af16e025549
	go func() {
		log.Printf("Website Up - listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
		select {
		case <-ctx.Done():
			fmt.Printf("\nShutting down server\n")
			shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancelShutdown()

			if err := httpServer.Shutdown(shutdownCtx); err != nil {
				fmt.Printf("\nerror shutting down server: %s\n", err)
			}
		}
	}()
}
