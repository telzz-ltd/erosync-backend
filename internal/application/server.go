package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *Application) Run(port int, h http.Handler) {
	ctx := context.Background()

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        h,
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
