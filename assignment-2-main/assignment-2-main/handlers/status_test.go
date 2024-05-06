package handlers_test

import (
	"assignment-2/database"
	"assignment-2/handlers"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStatusHandler(t *testing.T) {
	db := database.InitializeInMemoryDatabase()
	handlers.Db = &db

	server := httptest.NewServer(handlers.StatusHandler{
		TimeStarted: time.Now(),
	})

	defer server.Close()

	client := server.Client()
	res, err := client.Get(server.URL)

	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %v", res.StatusCode)
	}

	if res.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Expected content type to be application/json, got %v", res.Header.Get("Content-Type"))
	}

	var jsonMap map[string]any
	err = json.NewDecoder(res.Body).Decode(&jsonMap)

	if err != nil {
		t.Errorf("Error decoding json: %v", err)
	}

	if jsonMap["version"] != "v1" {
		t.Errorf("Expected version to be v1, got %v", jsonMap["version"])
	}
}
