package db

import ()

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
