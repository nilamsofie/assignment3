package handlers_test

import (
	"assignment-2/database"
	"assignment-2/handlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDashboards(t *testing.T) {
	// initialize in-memory database
	db := database.InitializeInMemoryDatabase()
	handlers.Db = &db

	// create test dashboard configuration
	doc, err := db.CreateDashboardConfiguration(database.DashboardConfiguration{
		Country: "USA",
		IsoCode: "US",
		Features: database.DashboardConfigurationFeatures{
			Area:             true,
			Capital:          true,
			Coordinates:      true,
			Population:       true,
			Precipitation:    true,
			TargetCurrencies: []string{"NOK", "EUR", "SEK"},
			Temperature:      true,
			Map:              true,
		},
	})
	if err != nil {
		t.Fatalf("Error creating dashboard configuration: %v", err)
	}

	// set up new request, url is not used since we run it directly
	req := httptest.NewRequest("GET", "/dashboard/v1/dashboards/"+doc.Id, nil)
	req.SetPathValue("id", doc.Id)

	// set up response recorder
	w := httptest.NewRecorder()

	// run handler
	handlers.DashboardsGet(w, req)
	response := w.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	// parse
	var dashboardResponse handlers.DashboardResponse
	err = json.NewDecoder(response.Body).Decode(&dashboardResponse)
	if err != nil {
		t.Fatalf("Error decoding dashboard response: %v", err)
	}

	// check response
	if dashboardResponse.Country != "USA" {
		t.Fatal("Expected country: USA, got", dashboardResponse.Country)
	}
}
