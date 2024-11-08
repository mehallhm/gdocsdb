package main

import (
	"context"
	"net/http"

	"github.com/mehallhm/gdocsdb/db"
	"github.com/mehallhm/gdocsdb/handler"
	"github.com/mehallhm/gdocsdb/middleware"
)

func NewRouter(databaseConn *db.Database) *http.Server {
	baseRouter := http.NewServeMux()

	baseRouter.HandleFunc("GET /b", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("b"))
	})

	databaseRouter := http.NewServeMux()
	baseRouter.Handle("/db/", http.StripPrefix("/db", middleware.EnsureAuth(databaseRouter)))

	databaseRouter.HandleFunc("GET /tab/{tab}/doc/{header}", handler.DocumentGet)

	middleware := middleware.CreateMiddlewareStack(
		DatabaseContext(databaseConn),
		middleware.Logger,
	)

	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware(baseRouter),
	}

	return server
}

// DatabaseContext creates a middleware that adds the database to the context
func DatabaseContext(databaseConn *db.Database) middleware.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "database.conn", databaseConn)
			req := r.WithContext(ctx)

			next.ServeHTTP(w, req)
		})
	}
}
