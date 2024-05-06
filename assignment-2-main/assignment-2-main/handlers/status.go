package handlers

import (
	"assignment-2/constants"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func getEndpointStatus(url string) int {
	// Http client with 10 second timeout
	client := http.Client{
		Timeout: time.Second * 10,
	}

	response, err := client.Get(url)

	if err == nil {
		return response.StatusCode
	} else {
		return http.StatusInternalServerError
	}
}

type status struct {
	CountriesApi   int `json:"countries_api"`
	MeteoApi       int `json:"meteo_api"`
	CurrencyApi    int `json:"currency_api"`
	NotificationDb int `json:"notification_db"`

	Webhooks int    `json:"webhooks"`
	Version  string `json:"version"`
	Uptime   int    `json:"uptime"`
}

// StatusHandler handles the /status endpoint
type StatusHandler struct {
	TimeStarted time.Time
}

func (sh StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	docs, err := Db.GetAllNotifications()

	databaseStatus := http.StatusInternalServerError
	notificationCount := 0

	if err == nil {
		databaseStatus = http.StatusOK
		notificationCount = len(docs)
	}

	status := status{
		CountriesApi:   getEndpointStatus(constants.BaseCountriesEndpoint + "all"),
		MeteoApi:       getEndpointStatus(constants.BaseMeteoEndpoint + "forecast"),
		CurrencyApi:    getEndpointStatus(constants.BaseCurrencyEndpoint + "NOK"),
		NotificationDb: databaseStatus,

		Webhooks: notificationCount,
		Version:  "v1",
		Uptime:   int(time.Since(sh.TimeStarted).Seconds()),
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(status)

	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("status, get: %v", err)
		return
	}
}
