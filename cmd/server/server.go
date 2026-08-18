package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(ctx context.Context) *Server {
	return &Server{
		httpServer: &http.Server{
			BaseContext: func(l net.Listener) context.Context { return ctx },
		},
	}
}

func (s *Server) SetHandler(handler http.Handler) {
	s.httpServer.Handler = handler
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
	log.Println("shutting down server")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Println("Server shutdown failed:", err)
		return
	}

	log.Println("Server shutdown completed")
}
