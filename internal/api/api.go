package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	v1 "github.com/murat/go-boilerplate/internal/api/v1"
)

// Start ...
func Start(port string) error {
	r := setup()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("could not gracefully shutdown, %v\n", err)
		}
		close(done)
	}()

	log.Printf("listening on :%s\n", port)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("could not start, %w", err)
	}

	<-done
	log.Println("stopped!")

	return nil
}

func setup() *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/new", v1.CreateGameHandler)

	return r
}
