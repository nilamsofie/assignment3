package database

import "time"

// Document is a generic struct that represents a document in a database
// T is the type of the data stored in the document
type Document[T any] struct {
	Id         string
	Data       T
	LastChange time.Time
}

type Database interface {
	// Close the connection to the database
	Close() error

	// === Dashboard configurations ===

	// Get a dashboard configuration by id
	// returns nil if the configuration does not exist
	GetDashboardConfiguration(id int) (*Document[DashboardConfiguration], error)

	// Get all dashboard configurations
	GetAllDashboardConfigurations() ([]Document[DashboardConfiguration], error)

	// Create dashboard configuration and return the id of the created configuration
	// errors if the configuration already exists
	CreateDashboardConfiguration(configuration DashboardConfiguration) (Document[DashboardConfiguration], error)

	// Delete a existing dashboard configuration
	// returns true if the configuration was deleted, false if it did not exist
	DeleteDashboardConfiguration(id int) (bool, error)

	// Update a existing dashboard configuration
	// returns true if the configuration was updated, false if it did not exist
	UpdateDashboardConfiguration(id int, configuration DashboardConfiguration) (bool, error)

	// Patches some of the values of an existing dashboard configuration
	// returns true if the configuration was updated, false if it did not exist
	PatchDashboardConfiguration(id int, configurationPatch map[string]any) (bool, error)

	// === Notificaitons Configurations ===

	// Get a notification by id
	// returns nil if the notification does not exist
	GetNotification(id string) (*Document[Notification], error)

	// Get all notifications
	GetAllNotifications() ([]Document[Notification], error)

	// Create a new notification
	CreateNotification(notification Notification) (Document[Notification], error)

	// Delete a existing notification
	// returns true if the notification was deleted, false if it did not exist
	DeleteNotification(id string) (bool, error)

	// Update a existing notification
	// returns true if the notification was updated, false if it did not exist
	UpdateNotification(id string, notification Notification) (bool, error)

	// Patches a some of the values of an existing dashboard configuration
	// returns true if the configuration was updated, false if it did not exist
	PatchNotification(id string, notificationPatch map[string]any) (bool, error)
}
