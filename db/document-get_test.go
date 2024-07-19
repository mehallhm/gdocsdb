package db

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/option"
)

var TestResponses = map[string]*DocumentResponse{
	"1": {
		Content: map[string]interface{}{
			"Id":   "1",
			"Name": "Tom",
			"Hats": "Many",
		},
		StartIndex: 74,
		EndIndex:   101,
	},
	"2": {
		Content: map[string]interface{}{
			"Id":   "2",
			"Name": "Jim",
			"Hats": "The Best",
		},
		StartIndex: 174,
		EndIndex:   205,
	},
	"3": {
		Content: map[string]interface{}{
			"Id":     "3",
			"Hats":   "Magical",
			"Name":   "Tim",
			"Shirts": "Only red",
		},
		StartIndex: 278,
		EndIndex:   325,
	},
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := os.ReadFile("./testing_db.json")
		assert.Nil(t, err)
		_, err = w.Write(b)
		assert.Nil(t, err)

	}))
	defer testServer.Close()

	srv, err := docs.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(testServer.URL))
	assert.Nil(t, err)

	db := &Database{
		docId: "1A79FGntMMZ8JXkF8bh0s7HHhPfwHweyul_0f-yK2hsQ",
		gdoc:  srv,
	}

	for id, expected := range TestResponses {
		resp, err := db.Doc(id).Get(ctx)
		assert.Nil(t, err)

		assert.Equal(t, resp, expected)
	}

}
