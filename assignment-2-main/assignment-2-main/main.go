package main

import (
	"assignment-2/constants"
	"assignment-2/database"
	"assignment-2/handlers"
	"assignment-2/webhooks"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	timeStarted := time.Now()

	db, err := database.InitializeFirestore()
	if err != nil {
		log.Fatalf("Error initializing Firestore: %v", err)
	}

	handlers.Db = &db
	webhooks.Db = &db
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT not set. Defaulting to 8080")
		port = constants.DefaultPort
	}

	if os.Getenv("APP_ENV") == "PRODUCTION" {
		log.Println("Running in production mode")
	} else {
		log.Println("Running in development mode")
	}

	// Status
	http.Handle("GET "+constants.BaseUrl+"status", handlers.StatusHandler{TimeStarted: timeStarted})

	// Registrations
	http.HandleFunc("GET "+constants.BaseUrl+"registrations", handlers.GetAllRegistrations)
	http.HandleFunc("GET "+constants.BaseUrl+"registrations/{id}", handlers.GetRegistrationsById)
	http.HandleFunc("POST "+constants.BaseUrl+"registrations", handlers.RegistrationsPost)
	http.HandleFunc("PUT "+constants.BaseUrl+"registrations/{id}", handlers.RegistrationsPut)
	http.HandleFunc("PATCH "+constants.BaseUrl+"registrations/{id}", handlers.RegistrationsPatch)
	http.HandleFunc("DELETE "+constants.BaseUrl+"registrations/{id}", handlers.RegistrationsDelete)

	// Dashboards
	http.HandleFunc("GET "+constants.BaseUrl+"dashboards/{id}", handlers.DashboardsGet)

	// Notifications
	http.HandleFunc("GET "+constants.BaseUrl+"notifications", handlers.GetAllNotifications)
	http.HandleFunc("GET "+constants.BaseUrl+"notifications/{id}", handlers.GetSingleNotification)
	http.HandleFunc("POST "+constants.BaseUrl+"notifications", handlers.PostNotification)
	http.HandleFunc("PUT "+constants.BaseUrl+"notifications/{id}", handlers.PutNotification)
	http.HandleFunc("PATCH "+constants.BaseUrl+"notifications/{id}", handlers.PatchNotification)
	http.HandleFunc("DELETE "+constants.BaseUrl+"notifications/{id}", handlers.DeleteNotification)

	// default handler
	http.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "defaultResponse.html")
	})

	log.Println("Starting server on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
