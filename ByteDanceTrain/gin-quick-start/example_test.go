package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	h := &gin.H{}
	if err = json.Unmarshal(w.Body.Bytes(), h); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, gin.H{"message": "pong"}, *h)
}
