package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestPortWithDefaultPort(t *testing.T) {
	assert.Equal(t, ":8080", port())
}

func TestPortWithNonDefaultPort(t *testing.T) {
	err := os.Setenv("PORT", "9000")
	assert.Nil(t, err)
	assert.Equal(t, ":9000", port())
}

func TestSetupRouter(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	setUpRouter().ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, []byte("Stocks service is running...."), w.Body.Bytes())
}