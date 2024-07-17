package db

import (
	"context"
)

// Doc retuns the Document given the id
func (d *Database) Doc(id string) *Document {
	return &Document{
		Id:       id,
		Database: d,
	}
}

type Document struct {
	Id       string
	Database *Database
}

const DocumentSeperator string = "========================================================================\n"

// Set updates a document by overriding ALL fields. To merge to an existing document, use Update. If the document does not exist,
// it creates a new document with the given document id
func (d *Document) Set(ctx context.Context, content map[string]interface{}) error {
	// Check if document exists
	// If document not exists, create and set fields
	// If document does exist, override fields

	return nil
}

// Update an existing document, overriding fields that exist and adding new ones
func (d *Document) Update(ctx context.Context, content map[string]interface{}) error {
	// Check if document exists
	// Call GET
	// Merge in new fields
	// Call SET to override everything w/ the merged
	return nil
}
