package handlers

import (
	"assignment-2/constants"
	"assignment-2/database"
	"assignment-2/fetchers"
	"assignment-2/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type FeaturesResponse struct {
	Temperature      *float64           `json:"temperature,omitempty"`
	Precipitation    *float64           `json:"precipitation,omitempty"`
	Capital          *string            `json:"capital,omitempty"`
	Coordinates      *Coordinates       `json:"coordinates,omitempty"`
	Population       *int               `json:"population,omitempty"`
	Area             *float64           `json:"area,omitempty"`
	TargetCurrencies map[string]float64 `json:"targetCurrencies,omitempty"`
	Map              *string            `json:"map,omitempty"`
}
type DashboardResponse struct {
	Country       string           `json:"country"`
	IsoCode       string           `json:"isoCode"`
	Features      FeaturesResponse `json:"features"`
	LastRetrieval string           `json:"lastRetrieval"`
}

type ISO3 struct {
	Iso3 string `json:"cca3"`
}

func DashboardsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")

	id, err := IDValidation(idStr)
	if err != nil {
		http.Error(w, constants.InvalidIDMessage, http.StatusBadRequest)
		return
	}

	dashboardConfigDoc, err := Db.GetDashboardConfiguration(id)
	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("Error getting document from database: %v", err)
		return
	}

	if dashboardConfigDoc == nil {
		http.Error(w, "Document not found", http.StatusBadRequest)
		return
	}

	dashboardConfig := dashboardConfigDoc.Data

	features, err := generateFeatures(dashboardConfig.IsoCode, dashboardConfig.Features)
	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("Error generating features: %v", err)
		return
	}

	// Generate response
	dashboardResponse := DashboardResponse{
		Country:       dashboardConfig.Country,
		IsoCode:       dashboardConfig.IsoCode,
		Features:      features,
		LastRetrieval: time.Now().Format("20060102 15:04"),
	}

	// enconde and send to client
	err = json.NewEncoder(w).Encode(dashboardResponse)
	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
		return
	}
}

func generateFeatures(countryCode string, config database.DashboardConfigurationFeatures) (FeaturesResponse, error) {
	// return object
	var features FeaturesResponse

	country, err := fetchers.FetchCountryByCode(countryCode)

	if err != nil {
		return FeaturesResponse{}, err
	}

	if country == nil {
		return FeaturesResponse{}, fmt.Errorf("generate features: country not found")
	}

	// fetch weather data
	weather, err := fetchers.FetchWeather(country.Coordinates[0], country.Coordinates[1])
	if err != nil {
		return FeaturesResponse{}, err
	}

	if config.Map {
		mapStr, err := AsciiMap(country.Iso3)
		if err != nil {
			return FeaturesResponse{}, err
		}
		features.Map = &mapStr
	}

	if config.Temperature {
		meanTemperature := mean(weather.Hourly.Temperature)
		meanTemperature = utils.Round(meanTemperature, 1)

		features.Temperature = &meanTemperature
	}

	if config.Precipitation {
		meanPrecipitation := mean(weather.Hourly.Precipitation)
		meanPrecipitation = utils.Round(meanPrecipitation, 2)

		features.Precipitation = &meanPrecipitation
	}

	if config.Capital {
		features.Capital = &country.Capitals[0]
	}

	if config.Coordinates {
		features.Coordinates = &Coordinates{
			float64(country.Coordinates[0]),
			float64(country.Coordinates[1]),
		}
	}

	if config.Population {
		features.Population = &country.Population
	}

	if config.Area {
		features.Area = &country.Area
	}

	if len(config.TargetCurrencies) > 0 {

		// fetch all exchange rates for country
		currency := getRandomMapKey(country.Currencies)
		Exchange, err := fetchers.FetchExchange(currency)

		if err != nil {
			return FeaturesResponse{}, fmt.Errorf("generate features: %v", err)
		}

		// add the configured exchange rates to the TargetCurrencies map
		features.TargetCurrencies = make(map[string]float64)
		for _, targetCurrencyKey := range config.TargetCurrencies {
			features.TargetCurrencies[targetCurrencyKey] = Exchange.Rates[targetCurrencyKey]
		}
	}

	return features, nil
}

func mean(arr []float64) float64 {
	if len(arr) == 0 {
		return 0
	}
	var sum float64
	for _, v := range arr {
		sum += v
	}
	return sum / float64(len(arr))
}

func getRandomMapKey(m map[string]interface{}) string {
	for key := range m {
		return key
	}
	return ""
}
