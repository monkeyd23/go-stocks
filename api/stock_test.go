package api

import (
	"github.com/buger/jsonparser"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	Index(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, []byte("Stocks service is running...."), w.Body.Bytes())
}

func TestGetStocks(t *testing.T) {
	r, _ := http.NewRequest("GET", "/stock", nil)
	w := httptest.NewRecorder()
	v := map[string]string{
		"symbol": "AAPL",
	}
	r = mux.SetURLVars(r, v)
	GetStocks(w, r)


	res := []byte(w.Body.String())
	count, err := jsonparser.GetInt(res, "symbols_returned")
	name, err := jsonparser.GetString(res, "data", "[0]", "name")

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, int(count))
	assert.Equal(t, "Apple Inc.", name)
}

func TestGetStocksIncorrectSymbol(t *testing.T) {
	r, _ := http.NewRequest("GET", "/stock", nil)
	w := httptest.NewRecorder()
	v := map[string]string{
		"symbol": "NO",
	}
	r = mux.SetURLVars(r, v)
	GetStocks(w, r)


	res := []byte(w.Body.String())
	message, err := jsonparser.GetString(res, "Message")

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, message, "could not be found.")
}