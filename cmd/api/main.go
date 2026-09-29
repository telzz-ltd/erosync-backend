package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/telzz/erosync-api/internal/handler"
	"github.com/telzz/erosync-api/internal/service"
)

func main() {
	ctx := context.Background()

	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}

	mux := http.NewServeMux()

	jwtService := service.NewJwtService(os.Getenv("JWT_SECRET"))

	registerHandler := handler.NewRegisterHandler()
	loginHandler := handler.NewLoginHandler()

	mux.Handle("POST /auth/register", registerHandler)
	mux.Handle("POST /auth/login", loginHandler)
}
