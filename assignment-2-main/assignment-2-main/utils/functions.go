package utils

import (
	"assignment-2/constants"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math"
	"time"
)

// round to x decimal places
func Round(num float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(num*shift+.5) / shift
}

func FormatDateToString(date time.Time) string {
	return date.Format("20060102 15:04")
}

// Turn a struct into the equivalent json map
// works by marshalling the struct into json, then unmarshalling it into a map
func ToJsonMap[T any](data T) (map[string]any, error) {
	jsonBytes, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	var jsonMap map[string]any
	err = json.Unmarshal(jsonBytes, &jsonMap)

	return jsonMap, err
}

func HashContent(data []byte, secret string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write(data)
	encoded := hex.EncodeToString(hash.Sum(nil))
	return encoded
}

/*
Checks whether the value event field of the WebHookBody is existent.
*/
func WebHookEventValid(event string) bool {
	return event == constants.Register || event == constants.Change || event == constants.Delete || event == constants.Invoke
}
