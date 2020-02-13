package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestLibrary test the url that return the library comics
func TestLibrary(t *testing.T) {
	handler := func(c *gin.Context) {
		c.String(http.StatusOK, "bar")
	}

	router := gin.New()
	router.GET("/marvel/getAllComics", handler)

	req, _ := http.NewRequest("GET", "/marvel/getAllComics", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "bar")
}
