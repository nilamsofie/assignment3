package handlers

import (
	"assignment-2/constants"
	"assignment-2/database"
	"assignment-2/fetchers"
	"assignment-2/utils"
	"errors"
	"strconv"
)

// Checks if the registration body has all the necessary fields with valid values, then returns it as a valid DashboardConfiguration
// Should propogate to a BAD REQUEST response if the body is not valid
func ValidateRegistration(body RegistrationPostBody) (database.DashboardConfiguration, error) {
	hasCountry := body.Country != nil && *body.Country != ""
	hasIsoCode := body.IsoCode != nil && *body.IsoCode != ""

	hasAtLeastOne := hasCountry || hasIsoCode
	hasBoth := hasCountry && hasIsoCode

	if hasBoth {
		return database.DashboardConfiguration{}, errors.New("invalid registration body, must have either country or isoCode, but not both")
	}

	if !hasAtLeastOne {
		return database.DashboardConfiguration{}, errors.New(constants.InvalidRequestBodyMessage)
	}

	var country *fetchers.Country
	var err error

	if hasCountry {
		country, err = fetchers.FetchCountryByName(*body.Country)
	} else {
		country, err = fetchers.FetchCountryByCode(*body.IsoCode)
	}

	if err != nil {
		return database.DashboardConfiguration{}, err
	}

	if country == nil {
		return database.DashboardConfiguration{}, errors.New("country not found")
	}

	return database.DashboardConfiguration{
		Country: country.Name.Common,
		IsoCode: country.Iso2,
		Features: database.DashboardConfigurationFeatures{
			Area:             body.Features.Area,
			Capital:          body.Features.Capital,
			Coordinates:      body.Features.Coordinates,
			Population:       body.Features.Population,
			Precipitation:    body.Features.Precipitation,
			TargetCurrencies: body.Features.TargetCurrencies,
			Temperature:      body.Features.Temperature,
			Map:              body.Features.Map,
		},
	}, nil
}

/*
Validation of the patch registration request body
@return true: if the body is valid
@return false: if the body is invalid
*/
func ValidatePatchRegistration(body RegistrationPatchBody) bool {
	return body.Country != nil || body.IsoCode != nil || body.Features.Area != nil || body.Features.Capital != nil || body.Features.Coordinates != nil || body.Features.Population != nil || body.Features.Precipitation != nil || body.Features.TargetCurrencies != nil || body.Features.Temperature != nil || body.Features.Map != nil
}

/*
Validation of the notification request body
@return true: if the body is valid
@return false: if the body is invalid
*/
func ValidateNotification(body database.Notification) bool {
	return body.Url != "" && body.Country != "" && body.Event != "" && utils.WebHookEventValid(body.Event)
}

/*
Validation of the patch notification request body
@return true: if the body is valid
@return false: if the body is invalid
*/
func ValidatePatchNotification(body NotificationsPatchBody) bool {
	return (body.Url != nil || body.Country != nil || body.Secret != nil) && utils.WebHookEventValid(*body.Event)
}

// Checks if the id is a positive whole number, is not empty and then returns it as an int
// Should propogate to a BAD REQUEST response if the id is not valid
func IDValidation(id string) (int, error) {
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return 0, err
	}

	if idInt <= 0 {
		return 0, errors.New("ID not a valid number")
	}

	return idInt, nil
}
