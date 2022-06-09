package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	v1 "github.com/murat/go-boilerplate/internal/api/v1"
	"github.com/murat/go-boilerplate/internal/db"
	"gorm.io/gorm"
)

type Handler struct {
	DB     *gorm.DB
	Logger *log.Logger
	Server *http.Server
}

var (
	port, dbPath string
)

func main() {
	if err := run(os.Args); err != nil {
		log.Printf("could not start app, %v", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.StringVar(&port, "port", "8080", "-port 8080")
	flags.StringVar(&dbPath, "db", "sqlite.db", "-db sqlite.db")
	if err := flags.Parse(args[1:]); err != nil {
		return fmt.Errorf("could not parse flags, err: %w", err)
	}

	db, err := db.New(dbPath)
	if err != nil {
		return fmt.Errorf("could not open db, err: %w", err)
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)

	r := routes()

	h := &Handler{
		DB:     db,
		Logger: logger,
		Server: &http.Server{
			Addr:    ":" + port,
			Handler: r,
		},
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		h.Server.SetKeepAlivesEnabled(false)
		if err := h.Server.Shutdown(ctx); err != nil {
			log.Fatalf("could not gracefully shutdown, %v\n", err)
		}
		close(done)
	}()

	log.Printf("listening on :%s\n", port)
	err = h.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("could not start, %w", err)
	}

	<-done
	log.Println("stopped!")

	return nil
}

func routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/", v1.CreateHandler)
	r.Get("/{id}", v1.GetHandler)
	r.Put("/{id}", v1.JoinHandler)

	return r
}
