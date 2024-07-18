package db

import (
	"context"
	"fmt"
)

// Update an existing document, overriding fields that exist and adding new ones
func (d *Document) Update(ctx context.Context, content map[string]interface{}) error {
	// Check if document exists
	// Call GET
	// Merge in new fields
	// Call SET to override everything w/ the merged
	doc, err := d.Get(ctx)
	if err != nil {
		return err
	}

	mergedContents := make(map[string]interface{})

	for key, val := range doc.Content {
		if key == "Id" {
			continue
		}
		mergedContents[key] = val
	}

	for key, val := range content {
		mergedContents[key] = fmt.Sprintf("%s", val)
	}
	fmt.Println(mergedContents)

	err = d.Set(ctx, mergedContents)
	if err != nil {
		return err
	}

	return nil
}
