package db

import (
	"context"
	"log"

	"google.golang.org/api/docs/v1"
	"google.golang.org/api/option"
)

type Database struct {
	document *docs.Document
}

func New(docId string, ctx context.Context) *Database {
	client := GoogleApiClient()

	srv, err := docs.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Docs client: %v", err)
	}

	doc, err := srv.Documents.Get(docId).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from document: %v", err)
	}

	return &Database{
		document: doc,
	}
}
