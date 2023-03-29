// main_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

// tom: next functions added later, these require more modules: net/http net/http/httptest

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestSearch(t *testing.T) {

	req, _ := http.NewRequest("GET", "/search/Yoshua Bengio", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

}
