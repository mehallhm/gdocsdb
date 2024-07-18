package db

import (
	"context"
	"errors"
	"fmt"

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

	id := doc.Content["Id"]

	formattedData := fmt.Sprintf("Id:%s", id)
	for key, val := range content {
		formattedData = formattedData + fmt.Sprintf("%s:%s\n", cases.Title(language.English, cases.NoLower).String(key), val)
	}

	_, err = d.Database.gdoc.Documents.BatchUpdate(d.Database.docId, &docs.BatchUpdateDocumentRequest{
		Requests: []*docs.Request{
			{
				DeleteContentRange: &docs.DeleteContentRangeRequest{
					Range: &docs.Range{
						StartIndex: int64(doc.StartIndex),
						EndIndex:   int64(doc.EndIndex),
					},
				},
			},
			{
				InsertText: &docs.InsertTextRequest{
					Location: &docs.Location{
						Index: int64(doc.StartIndex),
					},
					Text: formattedData,
				},
			},
		},
	}).Do()

	if err != nil {
		return err
	}

	return nil
}
