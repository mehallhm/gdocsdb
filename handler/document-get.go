package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/mehallhm/gdocsdb/db"
)

// DocumentGet fetches the document from the database with the specified header
func DocumentGet(w http.ResponseWriter, r *http.Request) {
	tab := r.PathValue("tab")
	documentHeader := r.PathValue("header")
	db := r.Context().Value("database.conn").(*db.Database)

	doc, err := db.GetSingleDocument(r.Context(), tab, documentHeader)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("error getting document", "error", err)
		w.Write([]byte("Internal Server Error"))
		return
	}

	body, err := json.Marshal(doc.Content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("error marshling document", "error", err)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Write(body)
}
