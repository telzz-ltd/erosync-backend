package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/telzz/erosync-api/pkg/response"
)

type Middleware func(http.Handler) http.Handler

func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, m := range middlewares {
		handler = m(handler)
	}

	return handler
}

func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err, string(debug.Stack()))
				response.ServerError(w, "An unknown error occurred. Please try again later.")
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}
