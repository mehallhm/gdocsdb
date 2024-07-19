package db

import (
	"context"

	"google.golang.org/api/docs/v1"
)

// Delete removes a document
func (d *Document) Delete(ctx context.Context) error {
	doc, err := d.Get(ctx)
	if err != nil {
		return err
	}

	_, err = d.batchUpdate([]*docs.Request{
		{
			DeleteContentRange: &docs.DeleteContentRangeRequest{
				Range: &docs.Range{
					StartIndex: int64(doc.StartIndex),
					EndIndex:   int64(doc.EndIndex) + int64(len(DocumentSeperator)),
				},
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
