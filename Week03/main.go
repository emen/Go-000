package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	srv := newSrv()

	g.Go(func() error {
		return start(ctx, srv)
	})

	g.Go(func() error {
		return shutdown(ctx, srv, sigs)
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Println("The server is shutdown gracefully")

}

func newMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	return mux
}

func home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "welcome home")
}

func newSrv() *http.Server {
	return &http.Server{
		Addr:         ":3000",
		Handler:      newMux(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  20 * time.Second,
	}
}

func start(_ context.Context, srv *http.Server) error {
	log.Println("server is starting")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func shutdown(ctx context.Context, srv *http.Server, sigs <-chan os.Signal) error {
	// Catch (1) signals (2) issues where the server fails to start (without this the `Wait` seems
	// to be waiting for shutdown forever.
	select {
	case <-sigs:
		log.Println("received stop signal")
	case <-ctx.Done():
		log.Println("server start failed")
		return nil
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Println("attempt graceful shutdown")
	return srv.Shutdown(shutdownCtx)
}
