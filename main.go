package main

import (
	"context"
	"erosync/internal/handler"
	"erosync/internal/middleware"
	"erosync/internal/service"
	"erosync/internal/store"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	r := chi.NewRouter()
	// if os.Getenv("APP_ENV") == "production" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	db, err := sqlx.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalln("uanble to connect to db", err)
	}

	//adapters
	store := store.New(db)

	//services
	authService := service.NewAuthService(store)
	otpService := service.NewOTPService(store)

	//routes
	r.Get("/health", handler.HealthCheck)
	r.Post("/auth/register", handler.Register(authService))
	r.Post("/auth/login", handler.Login(authService))

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth)

		r.Post("/verification/email/send-otp", handler.SendEmailVerificationCode(store, otpService))
		r.Post("/verification/email/verify", handler.VerifyEmail(store, otpService))
	})

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
