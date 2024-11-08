package db

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/api/docs/v1"
)

const DocumentSeperator string = "========================================================================\n"

func (d *Database) batchUpdate(requests ...*docs.Request) (*docs.BatchUpdateDocumentResponse, error) {
	return d.gdoc.Documents.BatchUpdate(d.docId, &docs.BatchUpdateDocumentRequest{
		Requests: requests,
	}).Do()
}

type DocumentResponse struct {
	// The content of the document
	Content map[string]interface{}
	// The index of the first character of the document (the "I" in "Id")
	StartIndex int
	// The index of the last character of the document (excluding the seperator)
	EndIndex int
}

func findDocumentTab(tabs []*docs.Tab, tab string) *docs.DocumentTab {
	for _, body := range tabs {
		if body.TabProperties.Title == tab {
			return body.DocumentTab
		}

		dt := findDocumentTab(body.ChildTabs, tab)
		if dt != nil {
			return dt
		}
	}

	return nil
}

// Get retreives the data of a document
func (d *Database) GetSingleDocument(ctx context.Context, tab string, header string) (*DocumentResponse, error) {
	documentQuery := d.gdoc.Documents.Get(d.docId)
	documentQuery.IncludeTabsContent(true)

	doc, err := documentQuery.Do()
	if err != nil {
		return nil, err
	}

	// INFO: Remove this debug
	// j, _ := json.Marshal(doc)
	// slog.Debug(string(j))

	match := header + "\n"

	dt := findDocumentTab(doc.Tabs, tab)
	if dt == nil {
		slog.Error("tab not found", "tab", tab)
		return nil, fmt.Errorf("tab not found")
	}

	content := dt.Body.Content
	fields := make(map[string]interface{})
	found := false
	start := 0
	end := 0
	for _, ele := range content {
		if ele.Paragraph == nil {
			continue
		}

		e := ele.Paragraph.Elements[0]

		if found && e.HorizontalRule != nil {
			break
		}

		if e.TextRun == nil {
			continue
		}

		if found {
			text := strings.Split(e.TextRun.Content, ":")
			key := text[0]
			key = strings.ReplaceAll(key, " ", "")
			key = strings.ReplaceAll(key, "\n", "")

			val := text[1]
			val = strings.Trim(val, " ")
			val = strings.ReplaceAll(val, "\n", "")

			if start == 0 {
				start = int(e.StartIndex)
			}

			if edx := int(e.EndIndex); edx > end {
				end = edx
			}

			fields[key] = val
		}

		if ele.Paragraph.ParagraphStyle.NamedStyleType == "HEADING_2" && e.TextRun.Content == match {
			found = true
		}
	}

	if !found {
		slog.Error("document not found", "header", header)
		return nil, fmt.Errorf("document not found")
	}

	fields["_header"] = header

	return &DocumentResponse{
		Content:    fields,
		StartIndex: start,
		EndIndex:   end,
	}, nil
}

// BUG: The following 3 functions do not support tabs yet... :(

// Set updates a document by overriding ALL fields. To merge to an existing document, use Update. If the document does not exist,
// it creates a new document with the given document id
func (d *Database) SetSingleDocument(ctx context.Context, tab string, header string, content map[string]interface{}) error {
	// Check if document exists
	// If document not exists, create and set fields
	// 🍍 If document does exist, override fields

	doc, err := d.GetSingleDocument(ctx, tab, header)
	if err != nil {
		return errors.Join(fmt.Errorf("error getting the document"), err)
	}

	formattedData := make([]string, 0, len(content)+1)
	for key, val := range content {
		formattedData = append(formattedData, fmt.Sprintf("%s: %s\n", cases.Title(language.English, cases.NoLower).String(key), val))
	}

	sort.Strings(formattedData)
	stringifiedData := fmt.Sprintf("Id: %s\n", header)
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

// Update an existing document, overriding fields that exist and adding new ones
func (d *Database) UpdateSingleDocument(ctx context.Context, tab string, header string, content map[string]interface{}) error {
	// TODO: Check if document acutally exists
	doc, err := d.GetSingleDocument(ctx, tab, header)
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

	err = d.SetSingleDocument(ctx, tab, header, mergedContents)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a document
func (d *Database) DeleteSingleDocument(ctx context.Context, tab string, header string) error {
	doc, err := d.GetSingleDocument(ctx, tab, header)
	if err != nil {
		return err
	}

	_, err = d.batchUpdate(
		&docs.Request{
			DeleteContentRange: &docs.DeleteContentRangeRequest{
				Range: &docs.Range{
					StartIndex: int64(doc.StartIndex),
					EndIndex:   int64(doc.EndIndex) + int64(len(DocumentSeperator)),
				},
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}
