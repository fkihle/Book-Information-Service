package handler

import (
	"assignment-01/constants"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadershipHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ReadershipHandler))
	defer server.Close()

	client := http.Client{}

	// Define multiple inputs to test
	var testInput []string
	testInput = append(testInput, "no", "", "no?limit=2", "no?limit=x", "no?limit=-2", "?limit=20")

	// Test all input variants
	for _, test := range testInput {
		resp, err := client.Get(server.URL + constants.READERSHIP_PATH + test)
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
	client.Post(server.URL+constants.READERSHIP_PATH+testInput[0], "POST", bytes.NewBuffer(body))
}
