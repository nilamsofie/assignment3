package database_test

import (
	"assignment-2/database"
	"reflect"
	"strconv"
	"testing"
)

// tests the creation and fetching of a dashboard configuration
func TestCreateAndGet(t *testing.T) {
	db := database.InitializeInMemoryDatabase()

	config := database.DashboardConfiguration{
		Country: "Canada",
		IsoCode: "CA",
		Features: database.DashboardConfigurationFeatures{
			Area:             false,
			Capital:          true,
			Coordinates:      false,
			Population:       true,
			Precipitation:    true,
			TargetCurrencies: []string{"CAD"},
			Temperature:      true,
		},
	}

	doc, err := db.CreateDashboardConfiguration(config)
	if err != nil {
		t.Errorf("Error creating configuration: %v", err)
	}

	// Check if the data is identical
	if !reflect.DeepEqual(config, doc.Data) {
		t.Errorf("Expected created data to be %v, got %v", config, doc.Data)
	}

	id, err := strconv.Atoi(doc.Id)
	if err != nil {
		t.Errorf("Created ID is not a number: %v", doc.Id)
	}

	// Check if the document can be fetched
	fetchedDoc, err := db.GetDashboardConfiguration(id)
	if err != nil {
		t.Errorf("Error fetching configuration: %v", err)
	}

	if !reflect.DeepEqual(config, fetchedDoc.Data) {
		t.Errorf("Expected fetched data to be %v, got %v", config, doc.Data)
	}
}
