package handlers

import (
	"assignment-2/database"
)

type RegistrationsGetResponse struct {
	Id         int                                     `json:"id"`
	Country    string                                  `json:"country"`
	IsoCode    string                                  `json:"isoCode"`
	Features   database.DashboardConfigurationFeatures `json:"features"`
	LastChange string                                  `json:"lastChange"`
}

type RegistrationPostBody struct {
	Country  *string `json:"country"`
	IsoCode  *string `json:"isoCode"`
	Features struct {
		Area             bool     `json:"area"`
		Capital          bool     `json:"capital"`
		Coordinates      bool     `json:"coordinates"`
		Population       bool     `json:"population"`
		Precipitation    bool     `json:"precipitation"`
		TargetCurrencies []string `json:"targetCurrencies"`
		Temperature      bool     `json:"temperature"`
		Map              bool     `json:"map"`
	} `json:"features"`
}

type RegistrationPatchBody struct {
	Country  *string `json:"country,omitempty"`
	IsoCode  *string `json:"isoCode,omitempty"`
	Features struct {
		Area             *bool    `json:"area,omitempty"`
		Capital          *bool    `json:"capital,omitempty"`
		Coordinates      *bool    `json:"coordinates,omitempty"`
		Population       *bool    `json:"population,omitempty"`
		Precipitation    *bool    `json:"precipitation,omitempty"`
		TargetCurrencies []string `json:"targetCurrencies,omitempty"`
		Temperature      *bool    `json:"temperature,omitempty"`
		Map              *bool    `json:"map,omitempty"`
	} `json:"features"`
}

type RegistrationPostResponse struct {
	Id         int         `json:"id"`
	LastChange interface{} `json:"lastChange"`
}

type NotificationGetResponse struct {
	Id      string `json:"id" firestore:"id"`
	Url     string `json:"url" firestore:"url"`
	Country string `json:"country" firestore:"country"`
	Event   string `json:"event" firestore:"event"`
}

type NotificationsPostResponse struct {
	ID string `json:"id"`
}

type NotificationsPatchBody struct {
	Url     *string `json:"url,omitempty"`
	Country *string `json:"country,omitempty"`
	Event   *string `json:"event,omitempty"`
	Secret  *string `json:"secret,omitempty"`
}

type GeojsonData struct {
	Features []struct {
		Geometry struct {
			Coordinates [][][][]float64 `json:"coordinates"`
		} `json:"geometry"`

		Properties struct {
			Bbox []float64 `json:"bbox"`
		} `json:"properties"`
	} `json:"features"`
}

type Iso3 struct {
	ISO3 string `json:"cca3"`
}
