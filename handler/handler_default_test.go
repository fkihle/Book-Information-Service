package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaultHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(DefaultHandler))
	defer server.Close()

	client := http.Client{}

	resp, err := client.Get(server.URL + "/")
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d but got %d", http.StatusOK, resp.StatusCode)
	}

	t.Log(resp.Body)
}
