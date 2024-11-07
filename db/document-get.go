package db

import (
	"context"
	"encoding/json"
	"log/slog"
	"strings"
)

type DocumentResponse struct {
	// The content of the document
	Content map[string]interface{}
	// The index of the first character of the document (the "I" in "Id")
	StartIndex int
	// The index of the last character of the document (excluding the seperator)
	EndIndex int
}

// Get retreives the data of a document
func (d *Document) Get(ctx context.Context) (*DocumentResponse, error) {
	query := d.Database.gdoc.Documents.Get(d.Database.docId)
	query.IncludeTabsContent(true)

	doc, err := query.Do()
	if err != nil {
		return nil, err
	}

	j, _ := json.Marshal(doc)

	slog.Debug(string(j))

	match := "Id: " + d.Id + "\n"

	content := doc.Body.Content
	fields := make(map[string]interface{})
	found := false
	start := 0
	end := 0
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
			key = strings.ReplaceAll(key, " ", "")
			key = strings.ReplaceAll(key, "\n", "")

			val := text[1]
			val = strings.Trim(val, " ")
			val = strings.ReplaceAll(val, "\n", "")

			if start == 0 {
				start = int(ele.Paragraph.Elements[0].StartIndex)
			}

			if edx := int(ele.Paragraph.Elements[0].EndIndex); edx > end {
				end = edx
			}

			fields[key] = val
		}
	}

	// TODO: What if the document does not exist?

	return &DocumentResponse{
		Content:    fields,
		StartIndex: start,
		EndIndex:   end,
	}, nil
}
