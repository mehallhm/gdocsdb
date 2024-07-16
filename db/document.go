package db

import (
	"context"
	"strings"

	"google.golang.org/api/docs/v1"
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

// Get retreives the data of a document
func (d *Document) Get(ctx context.Context) (map[string]interface{}, error) {
	match := "Id: " + d.Id + "\n"

	content := d.Database.gdoc.Body.Content
	fields := make(map[string]interface{})
	found := false
	for _, ele := range content {
		if ele.Paragraph == nil {
			continue
		}

		if found && ele.Paragraph.Elements[0].TextRun.Content == DocumentSeperator {
			break
		}

		if found || ele.Paragraph.Elements[0].TextRun.Content == match {
			found = true
			text := strings.Split(ele.Paragraph.Elements[0].TextRun.Content, ":")
			key := text[0]
			val := text[1]

			fields[key] = val
		}
	}

	d.Database.gdoc.Body.Content = []*docs.StructuralElement{d.Database.gdoc.Body.Content[0]}

	return fields, nil
}

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
	return nil
}

// Delete removes a document
func (d *Document) Delete(ctx context.Context) error {
	return nil
}
