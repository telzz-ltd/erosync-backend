package main

import (
	"context"
	"erosync/internal/infrastructure/persistence/postgres"
	"erosync/internal/lib/app"
	"erosync/internal/otps"
	"erosync/internal/users"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	r := chi.NewRouter()
	// if os.Getenv("APP_ENV") == "production" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	db, err := pgxpool.New(ctx, os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalln("uanble to connect to db", err)
	}

	//adapters
	tx := postgres.NewTx(db)

	//repositories
	userRepo := postgres.NewUserRepository(db)
	otpRepo := postgres.NewOTPRepository(db)

	//modules
	otpModule := otps.New(otpRepo)
	userModule := users.New(userRepo, otpModule.Service, tx)

	//routes
	r.Get("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.JSON(w, 200, app.Map{"message": "app working fine"})
	}))

	//routes register
	userModule.RegisterRoutes(r)

	srv := &http.Server{
		Addr:           ":8080",
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
