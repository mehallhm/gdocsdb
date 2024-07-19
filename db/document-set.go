package db

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/api/docs/v1"
)

// Set updates a document by overriding ALL fields. To merge to an existing document, use Update. If the document does not exist,
// it creates a new document with the given document id
func (d *Document) Set(ctx context.Context, content map[string]interface{}) error {
	// Check if document exists
	// If document not exists, create and set fields
	// 🍍 If document does exist, override fields

	doc, err := d.Get(ctx)
	if err != nil {
		return errors.Join(fmt.Errorf("error getting the document"), err)
	}

	formattedData := make([]string, 0, len(content)+1)
	for key, val := range content {
		formattedData = append(formattedData, fmt.Sprintf("%s: %s\n", cases.Title(language.English, cases.NoLower).String(key), val))
	}

	sort.Strings(formattedData)
	stringifiedData := fmt.Sprintf("Id: %s\n", d.Id)
	for _, entry := range formattedData {
		stringifiedData = stringifiedData + entry
	}

	_, err = d.batchUpdate(
		&docs.Request{
			DeleteContentRange: &docs.DeleteContentRangeRequest{
				Range: &docs.Range{
					StartIndex: int64(doc.StartIndex),
					EndIndex:   int64(doc.EndIndex),
				},
			},
		},
		&docs.Request{
			InsertText: &docs.InsertTextRequest{
				Location: &docs.Location{
					Index: int64(doc.StartIndex),
				},
				Text: stringifiedData,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}
