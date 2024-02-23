package handler

import (
	"assignment-01/constants"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBookcountHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(BookcountHandler))
	defer server.Close()

	client := http.Client{}

	// Define multiple inputs to test
	var testInput []string
	testInput = append(testInput, "?language=no", "?language=", "?language=no,ru,es", "?language=noru", "", "?language=4^.", "")

	// Test all input variants
	for _, test := range testInput {
		resp, err := client.Get(server.URL + constants.BOOKCOUNT_PATH + test)
		if err != nil {
			t.Fatal("Get request to URL failed:", err.Error())
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected %d but got %d", http.StatusOK, resp.StatusCode)
		}
	}

	// Test an incorrect HTTP method
	body := []byte(`{
		"title": "Test",
		"body": "Testy Testerson",
		"userId": 42
	}`)
	client.Post(server.URL+constants.BOOKCOUNT_PATH+testInput[0], "POST", bytes.NewBuffer(body))
}
