package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

const url = "https://api.worldtradingdata.com/api/v1/stock"
const token = "hCJvoIqp6cS1AWSC62m0eKyLXmDDZDBZplcYjX4exLNEzCLYToB8eDxF5OqW"

func Index(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Stocks service is running....")
}

func GetStocks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	symbol := vars["symbol"]

	res, _ := http.Get(fmt.Sprintf("%s?symbol=%s&api_token=%s", url, symbol, token))
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json; charset-utf-8")
	_, _ = w.Write(bytes)
}