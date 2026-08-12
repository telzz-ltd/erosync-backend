package main

import (
	"context"
	"erosync/internal/infrastructure/server"
	"erosync/internal/middleware"
	"erosync/internal/pkg/env"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	killSig := make(chan os.Signal, 1)
	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	mux := http.NewServeMux()
	srv := server.New(ctx, mux)

	srv.RegisterRoutes(mux)

	mux.Handle("GET /static/", middleware.Static("/static", "./static"))

	port := env.GetInt("PORT", 8080)
	srv.Start(port)
	<-killSig

	srv.Shutdown(ctx)
}
