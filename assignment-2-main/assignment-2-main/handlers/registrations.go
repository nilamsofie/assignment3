package handlers

import (
	"assignment-2/constants"
	"assignment-2/utils"
	"assignment-2/webhooks"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetRegistrationsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	strId := r.PathValue("id")

	id, err := IDValidation(strId)
	if err != nil {
		http.Error(w, constants.InvalidIDMessage, http.StatusBadRequest)
		return
	}

	doc, err := Db.GetDashboardConfiguration(id)

	if doc == nil {
		http.Error(w, constants.DataNotFound, http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("registration, get: %v", err)
		return
	}

	config := doc.Data

	response := RegistrationsGetResponse{
		Id:         id,
		Country:    config.Country,
		IsoCode:    config.IsoCode,
		Features:   config.Features,
		LastChange: utils.FormatDateToString(doc.LastChange),
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("registration, get: %v", err)
		return
	}

	webhooks.NotificationWebhook(response.IsoCode, constants.Invoke)
}

func GetAllRegistrations(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Get all registrations
	docs, err := Db.GetAllDashboardConfigurations()
	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("registrations, get: %v", err)
		return
	}

	response := make([]RegistrationsGetResponse, 0)

	for _, doc := range docs {
		id, _ := strconv.Atoi(doc.Id)
		data := doc.Data

		response = append(response, RegistrationsGetResponse{
			Id:         id,
			Country:    data.Country,
			IsoCode:    data.IsoCode,
			Features:   data.Features,
			LastChange: utils.FormatDateToString(doc.LastChange),
		})
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("registrations, get: %v", err)
		return
	}

}

func RegistrationsPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var body RegistrationPostBody

	// Parse JSON body into struct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, constants.InvalidJsonMessage, http.StatusBadRequest)
		return
	}

	// Validates the request body from the client
	dashboardConfig, err := ValidateRegistration(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdDoc, err := Db.CreateDashboardConfiguration(dashboardConfig)
	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("registrations, post: %v", err)
		return
	}

	id, _ := strconv.Atoi(createdDoc.Id)

	// Make response
	var response = RegistrationPostResponse{
		Id:         id,
		LastChange: utils.FormatDateToString(createdDoc.LastChange),
	}

	// Send response as a JSON body
	encoder := json.NewEncoder(w)
	err = encoder.Encode(&response)
	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("registrations, post: %v", err)
		return
	}

	webhooks.NotificationWebhook(createdDoc.Data.IsoCode, constants.Register)
}

func RegistrationsPut(w http.ResponseWriter, r *http.Request) {
	//Get the ID from the URL
	idStr := r.PathValue("id")

	id, err := IDValidation(idStr)
	if err != nil {
		http.Error(w, constants.InvalidIDMessage, http.StatusBadRequest)
		return
	}

	// Parse the request body into struct
	var body RegistrationPostBody
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&body)
	if err != nil {
		http.Error(w, constants.InvalidJsonMessage, http.StatusBadRequest)
		return
	}

	dashboardConfig, err := ValidateRegistration(body)
	if err != nil {
		http.Error(w, constants.InvalidRequestBodyMessage, http.StatusBadRequest)
		return
	}

	didExist, err := Db.UpdateDashboardConfiguration(id, dashboardConfig)

	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("registrations, put: %v", err)
		return
	}

	if !didExist {
		http.Error(w, constants.DataNotFound, http.StatusNotFound)
		return
	}

	webhooks.NotificationWebhook(dashboardConfig.IsoCode, constants.Change)

	//Return 204 if successful
	w.WriteHeader(http.StatusNoContent)
}

func RegistrationsPatch(w http.ResponseWriter, r *http.Request) {
	// Get the id
	idStr := r.PathValue("id")

	// Validate the id
	id, err := IDValidation(idStr)
	if err != nil {
		http.Error(w, constants.InvalidIDMessage, http.StatusBadRequest)
		return
	}

	// Parse the request body into struct
	var body RegistrationPatchBody
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&body)
	if err != nil {
		http.Error(w, constants.InvalidJsonMessage, http.StatusBadRequest)
		return
	}

	// Validate the request body
	if !ValidatePatchRegistration(body) {
		http.Error(w, constants.InvalidRequestBodyMessage, http.StatusBadRequest)
		return
	}

	configurationPatch, err := utils.ToJsonMap(body)
	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("registrations, patch: %v", err)
		return
	}

	// Patch the document
	didExist, err := Db.PatchDashboardConfiguration(id, configurationPatch)

	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("registrations, patch: %v", err)
		return
	}

	if !didExist {
		http.Error(w, constants.DataNotFound, http.StatusNotFound)
		return
	}

	// Success response
	w.WriteHeader(http.StatusNoContent)
}

func RegistrationsDelete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := IDValidation(idStr)
	if err != nil {
		http.Error(w, constants.InvalidIDMessage, http.StatusBadRequest)
		return
	}

	// Need to fetch data first, for use in notification
	doc, err := Db.GetDashboardConfiguration(id)

	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("registrations, delete: %v", err)
		return
	}

	if doc == nil {
		http.Error(w, constants.DataNotFound, http.StatusNotFound)
		return
	}

	_, err = Db.DeleteDashboardConfiguration(id)

	if err != nil {
		http.Error(w, constants.ServerErrorMessage, http.StatusInternalServerError)
		log.Printf("registrations, delete: %v", err)
		return
	}

	// Success response
	w.WriteHeader(http.StatusNoContent)
}
