package main

import (
	"context"
	"erosync/internal/adapter/postgres"
	"erosync/internal/application"
	"erosync/internal/config"
	"erosync/internal/handler"
	"erosync/internal/middleware"
	"erosync/internal/router"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	ctx := context.Background()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CORS)
	r.Use(middleware.Recoverer)

	cfg := config.New()
	store := postgres.New(cfg.DatabaseUrl)

	app := application.New(cfg, store)

	h := handler.New(app)

	r.Mount("/", router.RegisterV1Router(app, h))

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	{
		go func() {
			log.Println("Server running on", srv.Addr)
			if err := srv.ListenAndServe(); err != nil {
				if err != http.ErrServerClosed {
					log.Fatalf("server error: %v", err)
				}
			}
		}()
	}

	sigCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGINT)
	defer stop()

	<-sigCtx.Done()
	log.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)

		if err := srv.Close(); err != nil {
			log.Fatalf("server close failed: %v", err)
		}
	}

	log.Println("server stopped")
}
