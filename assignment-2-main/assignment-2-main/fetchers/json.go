package fetchers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Fetches json data from the given url
//
// Example:
//
//	type User struct {
//		Id   int    `json:"id"`
//		Name string `json:"name"`
//	}
//
//	user, err := FetchJson<User>("https://jsonplaceholder.typicode.com/users/1")

func fetchJson[JsonData any](url string, stub string) (JsonData, error) {
	if os.Getenv("APP_ENV") == "PRODUCTION" {
		return jsonFromApi[JsonData](url)
	} else {
		return jsonFromStub[JsonData](&stub)
	}
}

func jsonFromStub[JsonData any](jsonStub *string) (JsonData, error) {
	var empty JsonData

	var jsonData JsonData
	err := json.Unmarshal([]byte(*jsonStub), &jsonData)
	if err != nil {
		return empty, fmt.Errorf("from stub: %v", err)
	}
	return jsonData, nil
}

func jsonFromApi[JsonData any](url string) (JsonData, error) {
	// Create an empty variable of the type we want to return, to return in case of an error
	var empty JsonData

	// Fetch the data from the given url
	response, err := http.Get(url)

	if err != nil {
		return empty, fmt.Errorf("from api: %v: ", err)
	}

	defer response.Body.Close()

	// Decode the json response into the given type
	var jsonData JsonData
	err = json.NewDecoder(response.Body).Decode(&jsonData)

	return jsonData, err
}
