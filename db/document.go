package db

import "google.golang.org/api/docs/v1"

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

func (d *Document) batchUpdate(requests ...*docs.Request) (*docs.BatchUpdateDocumentResponse, error) {
	return d.Database.gdoc.Documents.BatchUpdate(d.Database.docId, &docs.BatchUpdateDocumentRequest{
		Requests: requests,
	}).Do()
}
