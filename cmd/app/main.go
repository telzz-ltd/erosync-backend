package main

import (
	"context"
	"erosync/cmd/server"
	"erosync/internal/middleware"
	"erosync/pkg/env"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	defer func() {
		if err := recover(); err != nil {
			log.Println("startup panic:", err)
		}
	}()

	killSig := make(chan os.Signal, 1)
	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	srv := server.New(ctx)

	mux := srv.RegisterRoutes()

	handler := middleware.Recoverer(mux)

	srv.SetHandler(handler)

	port := env.GetInt("PORT", 8080)
	srv.Start(port)
	<-killSig

	srv.Shutdown(ctx)
}
