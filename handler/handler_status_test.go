package handler

import (
	"assignment-01/constants"
	"assignment-01/structs"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(StatusHandler))
	defer server.Close()

	client := http.Client{}

	resp, err := client.Get(server.URL + constants.STATUS_PATH)
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d but got %d", http.StatusOK, resp.StatusCode)
	}

	stat := structs.StatusOutput{}
	err2 := json.NewDecoder(resp.Body).Decode(&stat)
	if err2 != nil {
		t.Fatal("Error decoding JSON:", err2.Error())
	}

	t.Log(stat)
}
