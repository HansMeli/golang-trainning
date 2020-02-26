package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// TestLibrary test the url that return the library comics
func TestLibrary(t *testing.T) {
	// Build our expected body
	// body := gin.H{
	// 	"hello": "world",
	// }
	// Grab our router
	router := gin.Default()

	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/library/getAllComics")

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
}
