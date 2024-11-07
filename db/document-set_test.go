package db

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/option"
)

func TestSet(t *testing.T) {
	ctx := context.Background()
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		read, err := io.ReadAll(r.Body)
		assert.Nil(t, err)
		fmt.Println(string(read))

		// For the GET method
		if len(read) == 0 {
			b, err := os.ReadFile("./testing_db.json")
			assert.Nil(t, err)
			_, err = w.Write(b)
			assert.Nil(t, err)
			return
		}

	}))
	defer testServer.Close()

	srv, err := docs.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(testServer.URL))
	assert.Nil(t, err)

	db := &Database{
		docId: "testing",
		gdoc:  srv,
	}

	err = db.Doc("1").Set(ctx, map[string]interface{}{
		"Name":  "Billy Bob",
		"Pants": "Midas Gold",
	})
	assert.Nil(t, err)
}
