package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/telzz/erosync-api/internal/adapter/postgres"
	"github.com/telzz/erosync-api/internal/config"
	"github.com/telzz/erosync-api/internal/handler"
	"github.com/telzz/erosync-api/internal/middleware"
	"github.com/telzz/erosync-api/internal/service"
)

func main() {
	ctx := context.Background()
	cfg := config.New()

	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}

	mux := http.NewServeMux()

	//repositories
	txManager := postgres.NewTxManager(db)
	userRepo := postgres.NewUserRepository(db)
	otpRepo := postgres.NewOtpRepository(db)

	//services
	jwtService := service.NewJwtService(os.Getenv("JWT_SECRET"))
	mailService := service.NewMailService(service.MailConfig{
		Host:     cfg.MailHost,
		Port:     cfg.MailPort,
		Username: cfg.MailUsername,
		Password: cfg.MailPassword,
		MailFrom: cfg.MailFrom,
		AppName:  cfg.AppName,
		AppUrl:   cfg.AppUrl,
	})
	otpService := service.NewOtpService(otpRepo)

	//handlers
	registerHandler := handler.NewRegisterHandler(txManager, userRepo, jwtService, mailService)
	loginHandler := handler.NewLoginHandler(userRepo, jwtService)
	sendEmailHandler := handler.NewSendEmailCodeHandler(txManager, userRepo, otpService, mailService)

	mux.Handle("POST /auth/register", registerHandler)
	mux.Handle("POST /auth/login", loginHandler)

	mux.Handle("POST /verification/email/send-otp", middleware.Auth(sendEmailHandler, jwtService))

	h := middleware.Chain(mux, middleware.Recoverer)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: h,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	go func() {
		log.Println("Server running on ", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Error shutting down server", err)

		log.Println("Closing server....")
		if err := srv.Close(); err != nil {
			log.Fatal("Unable to close server", err)
		}

		log.Println("Forcefully stopping server...")
		os.Exit(1)
	}

	log.Println("Shutdown completed")
}
