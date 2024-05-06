package handlers_test

import (
	"assignment-2/constants"
	"assignment-2/database"
	"assignment-2/handlers"
	"assignment-2/webhooks"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type webhookRequest struct {
	request *http.Request
	payload webhooks.WebHookBody
}

var webhookRequests chan webhookRequest

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var payload webhooks.WebHookBody
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		close(webhookRequests)
	}

	webhookRequests <- webhookRequest{
		request: r,
		payload: payload,
	}

	w.WriteHeader(http.StatusOK)
}

// Tests adding a notification, and checking that the webhook gets invoked with the correct payload
func TestNotificationRegister(t *testing.T) {
	db := database.InitializeInMemoryDatabase()
	handlers.Db = &db
	webhooks.Db = &db

	webhookRequests = make(chan webhookRequest, 1)

	postNotification := httptest.NewServer(http.HandlerFunc(handlers.PostNotification))
	defer postNotification.Close()

	postRegistration := httptest.NewServer(http.HandlerFunc(handlers.RegistrationsPost))
	defer postRegistration.Close()

	webhooksClient := httptest.NewServer(http.HandlerFunc(webhookHandler))
	defer webhooksClient.Close()

	secret := "very hidden client secret"
	notification := database.Notification{
		Url:     webhooksClient.URL,
		Country: "NO",
		Event:   constants.Register,
		Secret:  &secret,
	}
	notificationJson, _ := json.Marshal(notification)

	res, err := postNotification.Client().Post(postNotification.URL, "application/json", bytes.NewReader(notificationJson))

	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if res.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", res.StatusCode)
	}

	registrationJson := `
		{
			"isoCode": "NO",
			"features": {}
		}
	`

	res, err = postRegistration.Client().Post(postRegistration.URL, "application/json", strings.NewReader(registrationJson))

	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if res.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", res.StatusCode)
	}

	r, errr := <-webhookRequests
	if !errr {
		t.Fatalf("Error reading from webhookRequests")
	}

	if r.request.Method != "POST" {
		t.Fatalf("Expected POST request, got %v", r.request.Method)
	}

	if r.request.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("Expected content type to be application/json, got %v", r.request.Header.Get("Content-Type"))
	}

	if r.payload.Country != "NO" {
		t.Fatalf("Expected country to be NO, got %v", r.payload.Country)
	}

	if r.payload.Event != constants.Register {
		t.Fatalf("Expected event to be register, got %v", r.payload.Event)
	}

	// Validate the signature
	signature := r.request.Header.Get("X-SIGNATURE")

	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		t.Fatalf("Error decoding signature: %v", err)
		return
	}

	// Ensuring the request body matches the hash of the signature from the request headers
	mac := hmac.New(sha256.New, []byte(secret))

	payloadBytes, err := json.Marshal(r.payload)
	if err != nil {
		t.Fatalf("Error marshalling payload: %v", err)
	}

	_, err = mac.Write(payloadBytes)
	if err != nil {
		t.Fatalf("Error writing to hmac: %v", err)
	}

	if !hmac.Equal(signatureBytes, mac.Sum(nil)) {
		t.Fatalf("Signature is invalid!")
	}

	close(webhookRequests)
}
