package fetchers

import (
	"assignment-2/constants"
	"assignment-2/fetchers/stubs"
	"fmt"
)

type Country struct {
	Currencies  map[string]interface{} `json:"currencies"`
	Capitals    []string               `json:"capital"`
	Coordinates []float64              `json:"latlng"`
	Population  int                    `json:"population"`
	Area        float64                `json:"area"`

	Iso3 string `json:"cca3"`
	Iso2 string `json:"cca2"`

	Name struct {
		Common string `json:"common"`
	} `json:"name"`
}

func FetchCountryByCode(countryCode string) (*Country, error) {
	url := constants.BaseCountriesEndpoint + "alpha/" + countryCode
	countries, err := fetchJson[[]Country](url, stubs.CountryByCode(countryCode))

	if err != nil {
		return nil, fmt.Errorf("fetch Country: %v", err)
	}

	// If no country is found, return nil
	if len(countries) == 0 {
		return nil, nil
	}

	country := countries[0]
	return &country, nil
}

func FetchCountryByName(name string) (*Country, error) {
	url := constants.BaseCountriesEndpoint + "name/" + name
	countries, err := fetchJson[[]Country](url, stubs.CountryByName(name))

	if err != nil {
		return nil, fmt.Errorf("fetch Country: %v", err)
	}

	// If no country is found, return nil
	if len(countries) == 0 {
		return nil, nil
	}

	country := countries[0]
	return &country, nil
}

type Weather struct {
	Hourly struct {
		Temperature   []float64 `json:"temperature_2m"`
		Precipitation []float64 `json:"precipitation"`
	} `json:"hourly"`
}

func FetchWeather(lat float64, long float64) (Weather, error) {
	url := constants.BaseMeteoEndpoint + "forecast?"
	url += fmt.Sprintf("latitude=%.2f", lat)
	url += fmt.Sprintf("&longitude=%.2f", long)
	url += "&hourly=temperature_2m,precipitation"
	url += "&timezone=Europe%2FBerlin"
	url += "&forecast_days=1"
	weather, err := fetchJson[Weather](url, stubs.Weather)

	if err != nil {
		return Weather{}, fmt.Errorf("fetch weather: %v", err)
	}

	return weather, nil
}

type Exchange struct {
	Rates map[string]float64 `json:"rates"`
}

func FetchExchange(currency string) (Exchange, error) {
	exchange, err := fetchJson[Exchange](constants.BaseCurrencyEndpoint+currency, stubs.Currency)

	if err != nil {
		return Exchange{}, fmt.Errorf("fetch currency: %v", err)
	}

	return exchange, nil
}

type GeoJsonData struct {
	Features []struct {
		Geometry struct {
			//Coordinates [][][][]float64 `json:"coordinates"`
			Coordinates [][][]interface{} `json:"coordinates"`
		} `json:"geometry"`

		Properties struct {
			Bbox []float64 `json:"bbox"`
		} `json:"properties"`
	} `json:"features"`
}

func FetchGeoJson(iso3 string) (GeoJsonData, error) {
	query := constants.WorldCountriesEndpoint + iso3 + ".geojson"
	geoJson, err := fetchJson[GeoJsonData](query, stubs.CountryGeodata)
	if err != nil {
		return GeoJsonData{}, fmt.Errorf("fetch geojson: %v", err)
	}
	return geoJson, nil
}
