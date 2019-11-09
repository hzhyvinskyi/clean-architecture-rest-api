package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start(addr string) {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	server := &http.Server{}

	go gracefulShutdown(server, logger, quit, done)

	logger.Println("Server is ready to handle requests on ", addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Couldn't listen on %s: %v\n", addr, err)
	}

	<-done
	logger.Println("Server stopped")
}

// gracefulShutdown cleans up some resources before it actually shuts down
func gracefulShutdown(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	logger.Println("Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Couldn't gracefully shutdown the server: %v\n", err)
	}
	close(done)
}
