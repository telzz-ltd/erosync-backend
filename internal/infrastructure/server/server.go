package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(ctx context.Context, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Handler:     handler,
			BaseContext: func(l net.Listener) context.Context { return ctx },
		},
	}
}

func (s *Server) Start(port int) {
	s.httpServer.Addr = fmt.Sprintf(":%d", port)

	go func() {
		err := s.httpServer.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Server Closed", err)
		} else if err != nil {
			log.Println("Error starting server", err)
		}
	}()

	log.Println("Server running on ", s.httpServer.Addr)
}

func (s *Server) Shutdown(ctx context.Context) {
	log.Println("shuttong down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Println("Server shutdown failed. Forcefully shorting down server...", err)
		os.Exit(1)
	}
	log.Println("Server shutdown completed")
}
