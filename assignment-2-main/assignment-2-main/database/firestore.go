package database

import (
	"assignment-2/constants"
	"context"
	"errors"
	"os"
	"slices"
	"strconv"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FirestoreDatabase struct {
	ctx       context.Context
	firestore *firestore.Client
}

func InitializeFirestore() (FirestoreDatabase, error) {
	ctx := context.Background()

	credentials := os.Getenv("FIREBASE_CREDENTIALS")

	if credentials == "" {
		return FirestoreDatabase{}, errors.New("FIREBASE_CREDENTIALS environment variable not set")
	}

	options := option.WithCredentialsJSON([]byte(credentials))

	app, err := firebase.NewApp(ctx, nil, options)
	if err != nil {
		return FirestoreDatabase{}, err
	}

	firestore, err := app.Firestore(ctx)
	if err != nil {
		return FirestoreDatabase{}, err
	}

	return FirestoreDatabase{
		ctx,
		firestore,
	}, nil
}

func (f *FirestoreDatabase) Close() error {
	return f.firestore.Close()
}

func (f *FirestoreDatabase) documentExists(doc *firestore.DocumentRef) (bool, error) {
	_, err := doc.Get(f.ctx)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func (f *FirestoreDatabase) getDashboardOverview() (DashboardOverview, error) {
	overviewRef := f.firestore.Collection(constants.DashboardOverviewCollection).Doc("overview")

	overviewDoc, err := overviewRef.Get(f.ctx)
	if err != nil {
		return DashboardOverview{}, err
	}

	var overview DashboardOverview
	err = overviewDoc.DataTo(&overview)
	if err != nil {
		return DashboardOverview{}, err
	}

	return overview, nil
}

func (f *FirestoreDatabase) updateDashboardOverview(overview DashboardOverview) error {
	overviewRef := f.firestore.Collection(constants.DashboardOverviewCollection).Doc("overview")

	_, err := overviewRef.Set(f.ctx, overview)

	return err
}

func (f *FirestoreDatabase) GetDashboardConfiguration(id int) (*Document[DashboardConfiguration], error) {
	idStr := strconv.Itoa(id)
	configRef := f.firestore.Collection(constants.DashboardConfigurationCollection).Doc(idStr)

	// Read document
	configDoc, err := configRef.Get(f.ctx)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			// If the document does not exist, return nil
			return nil, nil
		} else {
			return nil, err
		}
	}

	// Read data
	var configData DashboardConfiguration
	err = configDoc.DataTo(&configData)

	if err != nil {
		return nil, err
	}

	// Create document
	document := Document[DashboardConfiguration]{
		Id:         configDoc.Ref.ID,
		Data:       configData,
		LastChange: configDoc.UpdateTime,
	}

	return &document, nil
}

func (f *FirestoreDatabase) GetAllDashboardConfigurations() ([]Document[DashboardConfiguration], error) {
	configsRef := f.firestore.Collection(constants.DashboardConfigurationCollection)

	configDocs, err := configsRef.Documents(f.ctx).GetAll()
	if err != nil {
		return nil, err
	}

	configs := make([]Document[DashboardConfiguration], len(configDocs))

	// Convert each document to a DashboardConfiguration
	for i, configDoc := range configDocs {
		var configData DashboardConfiguration
		err = configDoc.DataTo(&configData)

		if err != nil {
			return nil, err
		}

		configs[i] = Document[DashboardConfiguration]{
			Id:         configDoc.Ref.ID,
			Data:       configData,
			LastChange: configDoc.UpdateTime,
		}
	}

	return configs, nil
}

func (f *FirestoreDatabase) CreateDashboardConfiguration(configuration DashboardConfiguration) (Document[DashboardConfiguration], error) {
	overview, err := f.getDashboardOverview()
	if err != nil {
		return Document[DashboardConfiguration]{}, err
	}

	// Clone to avoid modifying the original
	deletedIds := slices.Clone(overview.DeletedIds)
	maxId := overview.MaxId
	newId := 0

	if len(deletedIds) != 0 {
		// Fill in the first deleted ID
		newId = deletedIds[0]
		deletedIds = deletedIds[1:]
	} else {
		// No deleted IDs we can fill in, increment maxId
		maxId += 1
		newId = maxId
	}

	// Update the overview
	err = f.updateDashboardOverview(DashboardOverview{
		DeletedIds: deletedIds,
		MaxId:      maxId,
	})

	if err != nil {
		return Document[DashboardConfiguration]{}, err
	}

	newIdStr := strconv.Itoa(newId)

	configRef := f.firestore.Collection(constants.DashboardConfigurationCollection).Doc(newIdStr)
	created, err := configRef.Set(f.ctx, configuration)

	if err != nil {
		return Document[DashboardConfiguration]{}, err
	}

	document := Document[DashboardConfiguration]{
		Id:         newIdStr,
		Data:       configuration,
		LastChange: created.UpdateTime,
	}

	return document, nil
}

func (f *FirestoreDatabase) DeleteDashboardConfiguration(id int) (bool, error) {
	idStr := strconv.Itoa(id)
	configRef := f.firestore.Collection(constants.DashboardConfigurationCollection).Doc(idStr)

	// Check if the configuraiton already exists
	alreadyExists, err := f.documentExists(configRef)

	if err != nil {
		return false, err
	}

	if !alreadyExists {
		return false, nil
	}

	_, err = configRef.Delete(f.ctx)

	if err != nil {
		return false, err
	}

	overview, err := f.getDashboardOverview()
	if err != nil {
		return false, err
	}

	deletedIds := slices.Clone(overview.DeletedIds)
	deletedIds = append(deletedIds, id)

	err = f.updateDashboardOverview(DashboardOverview{
		DeletedIds: deletedIds,
		MaxId:      overview.MaxId,
	})

	if err != nil {
		return false, err
	}

	// Return true if the configuration was successfully deleted
	return true, nil
}

func (f *FirestoreDatabase) UpdateDashboardConfiguration(id int, configuration DashboardConfiguration) (bool, error) {
	idStr := strconv.Itoa(id)
	configRef := f.firestore.Collection(constants.DashboardConfigurationCollection).Doc(idStr)

	// Check if the configuraiton already exists
	alreadyExists, err := f.documentExists(configRef)

	if err != nil {
		return false, err
	}

	if !alreadyExists {
		return false, nil
	}

	// Update the configuration
	_, err = configRef.Set(f.ctx, configuration)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (f *FirestoreDatabase) PatchDashboardConfiguration(id int, configurationPatch map[string]any) (bool, error) {
	idStr := strconv.Itoa(id)
	configRef := f.firestore.Collection(constants.DashboardConfigurationCollection).Doc(idStr)

	// Check if the configuraiton already exists
	alreadyExists, err := f.documentExists(configRef)

	if err != nil {
		return false, err
	}

	if !alreadyExists {
		return false, nil
	}

	// Patch the configuration
	_, err = configRef.Set(f.ctx, configurationPatch, firestore.MergeAll)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (f *FirestoreDatabase) GetNotification(id string) (*Document[Notification], error) {
	notificationRef := f.firestore.Collection(constants.NotificationsCollection).Doc(id)

	notificationDoc, err := notificationRef.Get(f.ctx)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			// If the document does not exist, return nil
			return nil, nil
		} else {
			return nil, err
		}
	}

	// Read data
	var notificationData Notification
	err = notificationDoc.DataTo(&notificationData)
	if err != nil {
		return nil, err
	}

	document := Document[Notification]{
		Id:         notificationDoc.Ref.ID,
		Data:       notificationData,
		LastChange: notificationDoc.UpdateTime,
	}

	return &document, nil
}

func (f *FirestoreDatabase) GetAllNotifications() ([]Document[Notification], error) {
	notificationsRef := f.firestore.Collection(constants.NotificationsCollection)

	notificationDocs, err := notificationsRef.Documents(f.ctx).GetAll()
	if err != nil {
		return nil, err
	}

	notifications := make([]Document[Notification], len(notificationDocs))

	for i, notificationDoc := range notificationDocs {
		var notificationData Notification
		err = notificationDoc.DataTo(&notificationData)

		if err != nil {
			return nil, err
		}

		notifications[i] = Document[Notification]{
			Id:         notificationDoc.Ref.ID,
			Data:       notificationData,
			LastChange: notificationDoc.UpdateTime,
		}
	}

	return notifications, nil
}

func (f *FirestoreDatabase) CreateNotification(notification Notification) (Document[Notification], error) {
	notificationRef := f.firestore.Collection(constants.NotificationsCollection).NewDoc()

	created, err := notificationRef.Set(f.ctx, notification)
	if err != nil {
		return Document[Notification]{}, err
	}

	document := Document[Notification]{
		Id:         notificationRef.ID,
		Data:       notification,
		LastChange: created.UpdateTime,
	}

	return document, nil
}

func (f *FirestoreDatabase) DeleteNotification(id string) (bool, error) {
	notificationRef := f.firestore.Collection(constants.NotificationsCollection).Doc(id)

	// Check if the notification already exists
	alreadyExists, err := f.documentExists(notificationRef)

	if err != nil {
		return false, err
	}

	if !alreadyExists {
		return false, nil
	}

	_, err = notificationRef.Delete(f.ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (f *FirestoreDatabase) UpdateNotification(id string, notification Notification) (bool, error) {
	notificationRef := f.firestore.Collection(constants.NotificationsCollection).Doc(id)

	// Check if the notification already exists
	alreadyExists, err := f.documentExists(notificationRef)

	if err != nil {
		return false, err
	}

	if !alreadyExists {
		return false, nil
	}

	_, err = notificationRef.Set(f.ctx, notification)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (f *FirestoreDatabase) PatchNotification(id string, notificationPatch map[string]any) (bool, error) {

	notificationRef := f.firestore.Collection(constants.NotificationsCollection).Doc(id)

	// Check if the notification already exists
	alreadyExists, err := f.documentExists(notificationRef)

	if err != nil {
		return false, err
	}

	if !alreadyExists {
		return false, nil
	}

	_, err = notificationRef.Set(f.ctx, notificationPatch, firestore.MergeAll)

	if err != nil {
		return false, err
	}

	return true, nil
}
