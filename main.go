package main

import (
	"context"
	"fmt"

	"github.com/mehallhm/gdocsdb/db"
)

func main() {
	ctx := context.Background()

	db := db.New("1A79FGntMMZ8JXkF8bh0s7HHhPfwHweyul_0f-yK2hsQ", ctx)

	fmt.Println(db.Doc("1").Get(ctx))
	fmt.Println(db.Doc("3").Update(ctx, map[string]interface{}{"Shirts": "Only red"}))
	// fmt.Println(db.Doc("2").Set(ctx, data))

	// fmt.Println(db.Doc("2").Delete(ctx))
}
