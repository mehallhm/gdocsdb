package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mehallhm/gdocsdb/db"
)

func main() {
	w := os.Stderr

	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	ctx := context.Background()
	db := db.New("1A79FGntMMZ8JXkF8bh0s7HHhPfwHweyul_0f-yK2hsQ", ctx)
	slog.Debug("prepared database")

	server := NewRouter(db)
	slog.Debug("prepared router")

	slog.Info("listening for requests")
	server.ListenAndServe()

	fmt.Println(db.Doc("1").Get(ctx))
	// fmt.Println(db.Doc("3").Update(ctx, map[string]interface{}{"Shirts": "Only red"}))
	// fmt.Println(db.Doc("2").Set(ctx, data))

	// fmt.Println(db.Doc("2").Delete(ctx))
}
