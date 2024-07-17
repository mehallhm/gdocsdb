package db

import (
	"context"
	"log"

	"google.golang.org/api/docs/v1"
	"google.golang.org/api/option"
)

type Database struct {
	gdoc  *docs.Service
	docId string
}

func New(docId string, ctx context.Context) *Database {
	client := GoogleApiClient()

	srv, err := docs.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Docs client: %v", err)
	}

	return &Database{
		gdoc:  srv,
		docId: docId,
	}
}
