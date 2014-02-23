package ctrl

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	//"io/ioutil"
	//"log"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBookJSON(t *testing.T) {
	expectedBody := "1159"
	handler := new(BookHandler)
	rec := httptest.NewRecorder()
	url := fmt.Sprintf("http://example.com/api/books/%s", expectedBody)
	req, err := http.NewRequest("GET", url, nil)
	assert.Nil(t, err)

	handler.ServeHTTP(rec, req)

	var dat map[string]interface{}

	err = json.NewDecoder(rec.Body).Decode(&dat)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, expectedBody, dat["id"])
}

func TestBookListJSON(t *testing.T) {
	expectedBody := "all"
	handler := new(BookListHandler)
	rec := httptest.NewRecorder()
	url := fmt.Sprintf("http://example.com/api/books/")
	req, err := http.NewRequest("GET", url, nil)
	assert.Nil(t, err)

	handler.ServeHTTP(rec, req)

	var dat map[string]interface{}

	err = json.NewDecoder(rec.Body).Decode(&dat)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, expectedBody, dat["category"])
}
