package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

// CreateMiddlewareStack creates a stack of Middlewares
func CreateMiddlewareStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}

		return next
	}
}

func EnsureAuth(next http.Handler) http.Handler {
	return next
}
