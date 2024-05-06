package database

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type InMemoryDatabase struct {
	lock          sync.RWMutex
	Registrations map[int]Document[DashboardConfiguration]
	Notifications map[string]Document[Notification]
}

func InitializeInMemoryDatabase() InMemoryDatabase {
	return InMemoryDatabase{
		Registrations: make(map[int]Document[DashboardConfiguration]),
		Notifications: make(map[string]Document[Notification]),
	}
}

func (db *InMemoryDatabase) Close() error {
	return nil
}

func (db *InMemoryDatabase) GetDashboardConfiguration(id int) (*Document[DashboardConfiguration], error) {
	db.lock.RLock()
	config, ok := db.Registrations[id]
	db.lock.RUnlock()

	if !ok {
		return nil, nil
	}

	return &config, nil
}

func (db *InMemoryDatabase) GetAllDashboardConfigurations() ([]Document[DashboardConfiguration], error) {
	configs := make([]Document[DashboardConfiguration], 0)

	db.lock.RLock()
	for _, config := range db.Registrations {
		configs = append(configs, config)
	}
	db.lock.RUnlock()

	return configs, nil
}

func (db *InMemoryDatabase) CreateDashboardConfiguration(configuration DashboardConfiguration) (Document[DashboardConfiguration], error) {
	db.lock.Lock()
	id := len(db.Registrations) + 1

	config := Document[DashboardConfiguration]{
		Id:         strconv.Itoa(id),
		Data:       configuration,
		LastChange: time.Now(),
	}

	db.Registrations[id] = config
	db.lock.Unlock()

	return config, nil
}

func (db *InMemoryDatabase) DeleteDashboardConfiguration(id int) (bool, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	_, ok := db.Registrations[id]

	if !ok {
		return false, nil
	}

	delete(db.Registrations, id)
	return true, nil
}

func (db *InMemoryDatabase) UpdateDashboardConfiguration(id int, configuration DashboardConfiguration) (bool, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	_, ok := db.Registrations[id]

	if !ok {
		return false, nil
	}

	config := Document[DashboardConfiguration]{
		Id:         strconv.Itoa(id),
		Data:       configuration,
		LastChange: time.Now(),
	}

	db.Registrations[id] = config
	return true, nil
}

func (db *InMemoryDatabase) PatchDashboardConfiguration(id int, configurationPatch map[string]any) (bool, error) {
	doc, err := db.GetDashboardConfiguration(id)

	if err != nil {
		return false, err
	}

	if doc == nil {
		return false, nil
	}

	data := doc.Data

	for key, value := range configurationPatch {
		fmt.Println(key, value)
	}

	// Patch the document
	for key, value := range configurationPatch {
		switch key {
		case "country":
			data.Country = value.(string)
		case "isoCode":
			data.IsoCode = value.(string)
		case "features":
			features := value.(map[string]any)

			for featureKey, featureValue := range features {
				switch featureKey {
				case "area":
					data.Features.Area = featureValue.(bool)
				case "capital":
					data.Features.Capital = featureValue.(bool)
				case "coordinates":
					data.Features.Coordinates = featureValue.(bool)
				case "population":
					data.Features.Population = featureValue.(bool)
				case "precipitation":
					data.Features.Precipitation = featureValue.(bool)
				case "targetCurrencies":
					data.Features.TargetCurrencies = featureValue.([]string)
				case "temperature":
					data.Features.Temperature = featureValue.(bool)
				case "map":
					data.Features.Map = featureValue.(bool)
				}
			}
		}
	}

	return db.UpdateDashboardConfiguration(id, data)
}

func (db *InMemoryDatabase) GetNotification(id string) (*Document[Notification], error) {
	db.lock.RLock()
	notification, ok := db.Notifications[id]
	db.lock.RUnlock()

	if !ok {
		return nil, nil
	}

	return &notification, nil
}

func (db *InMemoryDatabase) GetAllNotifications() ([]Document[Notification], error) {
	notifications := make([]Document[Notification], 0)

	db.lock.RLock()
	for _, notification := range db.Notifications {
		notifications = append(notifications, notification)
	}
	db.lock.RUnlock()

	return notifications, nil
}

func (db *InMemoryDatabase) CreateNotification(notification Notification) (Document[Notification], error) {
	db.lock.Lock()

	id := strconv.Itoa(len(db.Notifications) + 1)

	doc := Document[Notification]{
		Id:         id,
		Data:       notification,
		LastChange: time.Now(),
	}

	db.Notifications[id] = doc

	db.lock.Unlock()

	return doc, nil
}

func (db *InMemoryDatabase) DeleteNotification(id string) (bool, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	_, ok := db.Notifications[id]

	if !ok {
		return false, nil
	}

	delete(db.Notifications, id)
	return true, nil
}

func (db *InMemoryDatabase) UpdateNotification(id string, notification Notification) (bool, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	_, ok := db.Notifications[id]

	if !ok {
		return false, nil
	}

	doc := Document[Notification]{
		Id:         id,
		Data:       notification,
		LastChange: time.Now(),
	}

	db.Notifications[id] = doc
	return true, nil
}

func (db *InMemoryDatabase) PatchNotification(id string, notificationPatch map[string]any) (bool, error) {
	doc, err := db.GetNotification(id)

	if err != nil {
		return false, err
	}

	if doc == nil {
		return false, nil
	}

	data := doc.Data

	// Patch the document
	for key, value := range notificationPatch {
		switch key {
		case "url":
			data.Url = value.(string)
		case "country":
			data.Country = value.(string)
		case "event":
			data.Event = value.(string)
		case "secret":
			data.Secret = value.(*string)
		}
	}

	return db.UpdateNotification(id, data)
}
