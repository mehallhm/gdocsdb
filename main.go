package main

import (
	"context"

	"github.com/mehallhm/gdocsdb/db"
)

func main() {
	ctx := context.Background()

	db := db.New("1A79FGntMMZ8JXkF8bh0s7HHhPfwHweyul_0f-yK2hsQ", ctx)
	_ = db
}
