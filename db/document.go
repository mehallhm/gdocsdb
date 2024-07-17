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

// Update an existing document, overriding fields that exist and adding new ones
func (d *Document) Update(ctx context.Context, content map[string]interface{}) error {
	// Check if document exists
	// Call GET
	// Merge in new fields
	// Call SET to override everything w/ the merged
	return nil
}
