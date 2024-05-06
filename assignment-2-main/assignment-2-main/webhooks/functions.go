package webhooks

import (
	"assignment-2/constants"
	"assignment-2/utils"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func sendPostToWebhook(url string, data []byte, secret *string) {
	//Hashing data to be sent in headers to ensure data integrity and validation

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		log.Println("webhooks, failed to create request: ", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	if secret != nil {
		signatureHeader := utils.HashContent(data, *secret)
		req.Header.Add(constants.ClientSignature, signatureHeader)
	}
	// Http client with 10 second timeout
	client := http.Client{
		Timeout: time.Second * 10,
	}

	//Add post request logic here
	client.Do(req)
}

func NotificationWebhook(country string, event string) {
	docs, err := Db.GetAllNotifications()

	if err != nil {
		log.Println("Error in webhooks getting all notifications: ", err)
		return
	}

	for _, doc := range docs {
		notification := doc.Data

		//Checks if there is a notification for the country and the event
		if (notification.Country == "" || notification.Country == country) && notification.Event == event {

			//Creating the response body to be sent to the webhook
			responseData, err := json.Marshal(WebHookBody{
				ID:      doc.Id,
				Country: notification.Country,
				Event:   notification.Event,
				Time:    utils.FormatDateToString(time.Now()),
			})
			if err != nil {
				log.Println("Error in webhooks formatting the response body: ", err)
				return
			}

			go sendPostToWebhook(notification.Url, responseData, doc.Data.Secret)
		}
	}
}
