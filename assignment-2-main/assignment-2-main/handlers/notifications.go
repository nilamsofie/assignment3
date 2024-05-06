package handlers

import (
	"assignment-2/constants"
	"assignment-2/database"
	"assignment-2/utils"
	"encoding/json"
	"log"
	"net/http"
)

/*
Request handler function for retrieving all notifications
*/
func GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	notifications, err := Db.GetAllNotifications()
	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("notifications, get all: %v", err)
		return
	}

	// Make list of all notifications
	response := make([]NotificationGetResponse, 0)

	for _, doc := range notifications {
		data := doc.Data

		response = append(response, NotificationGetResponse{
			Id:      doc.Id,
			Url:     data.Url,
			Country: data.Country,
			Event:   data.Event,
		})
	}

	// Send JSON response
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

/*
Request handler function for retrieving a single notification
*/
func GetSingleNotification(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// Get the reference to the document with the specified ID

	// Retrieve the notification document
	doc, err := Db.GetNotification(id)

	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("notifications, get - one doc: %v", err)
		return
	}

	// Handle case where the document doesn't exist
	if doc == nil {
		http.Error(w, constants.DataNotFound, http.StatusNotFound)
		return
	}

	notification := doc.Data

	// Construct the response object
	notificationResponse := NotificationGetResponse{
		Id:      doc.Id,
		Url:     notification.Url,
		Country: notification.Country,
		Event:   notification.Event,
	}

	// Marshal the response object to JSON
	jsonResponse, err := json.Marshal(notificationResponse)
	if err != nil {
		// Handle the error (if any)
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
		return
	}

	// Set response headers and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func PostNotification(w http.ResponseWriter, r *http.Request) {
	var notification database.Notification

	// Parse JSON body into struct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&notification)
	if err != nil {
		http.Error(w, constants.InvalidJsonMessage, http.StatusBadRequest)
		return
	}

	// Validate the notification body
	valid := ValidateNotification(notification)
	if !valid {
		http.Error(w, constants.InvalidRequestBodyMessage, http.StatusBadRequest)
		return
	}

	// Store the notification in Firestore
	newDoc, err := Db.CreateNotification(notification)
	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("postNotification: %v", err)
		return
	}

	var response = NotificationsPostResponse{
		ID: newDoc.Id,
	}

	// Send response as a JSON body
	encoder := json.NewEncoder(w)
	err = encoder.Encode(&response)
	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("postNotification: %v", err)
		return
	}
}

/*
Request handler function for deleting a notification
*/
func DeleteNotification(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	didExist, err := Db.DeleteNotification(id)

	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("notifications, delete: %v", err)
		return
	}

	if !didExist {
		http.Error(w, constants.DataNotFound, http.StatusNotFound)
		return
	}

	// Respond with status no content to indicate successful deletion
	w.WriteHeader(http.StatusNoContent)
}

/*
Request handler function for put requests to the notifications/ endpoint.
The function updates an existing notification, with the id specified in the URI, with the provided request body
*/
func PutNotification(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	//Parse the request body into firebase.Notification struct
	var notification database.Notification
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		http.Error(w, constants.InvalidJsonMessage, http.StatusBadRequest)
		return
	}

	//Validation of the notification request body
	valid := ValidateNotification(notification)
	if !valid {
		http.Error(w, constants.InvalidRequestBodyMessage, http.StatusBadRequest)
		return
	}

	//Update the notification with corresponding error handling
	didExist, err := Db.UpdateNotification(id, notification)
	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("notifications, put: %v", err)
		return
	}

	if !didExist {
		http.Error(w, constants.DataNotFound, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func PatchNotification(w http.ResponseWriter, r *http.Request) {
	// Get the id
	id := r.PathValue("id")

	// Parse the request body into struct
	var body NotificationsPatchBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, constants.InvalidJsonMessage, http.StatusBadRequest)
		return
	}

	valid := ValidatePatchNotification(body)
	if !valid {
		http.Error(w, constants.InvalidRequestBodyMessage, http.StatusBadRequest)
		return
	}

	notificationPatch, err := utils.ToJsonMap(body)
	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("notifications, patch: %v", err)
		return
	}

	// Patch the document
	didExist, err := Db.PatchNotification(id, notificationPatch)

	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("notifications, patch: %v", err)
		return
	}

	if !didExist {
		http.Error(w, constants.DataNotFound, http.StatusNotFound)
		return
	}

	// Success response
	w.WriteHeader(http.StatusNoContent)
}
