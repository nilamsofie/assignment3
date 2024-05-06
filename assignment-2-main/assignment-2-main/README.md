# Assigment 2 - Countries Dashboard Service

## Prerequisites

Make sure you have go 1.22.0 or newer installed. See [https://go.dev/doc/install](https://go.dev/doc/install)

## Running locally

```bash
go run .
```

By default the service starts in development mode, where all external apis are stubbed. To run the service in production mode, use `APP_ENV=PRODUCTION go run .`.

## Tests

```bash
go test ./...
```

We made a test suite with both unit tests for tesing simple functions, and integration tests for some of the http handlers. To ensure reproducible tests, we had to mock all of the external APIs, as well as make a in memory database.

## Secrets

The service needs firebase admin credentials to run. This should be stored inside the `FIREBASE_CREDENTIALS` environment variable. For development purposes, you can store this in a `.env` file at the root of the project, which will be loaded automatically by `godotenv`.

## Deploying

We use the following docker to deploy on our skyhigh instance:

```bash
FIREBASE_CREDENTIALS=... docker compose up
```

## File structure

- [constants/](constants/) - Contains all the constants used in the application.]
- [database/](database/) - Our database layer. All our database calls are abstracted away in a `Database` interface, which is implemented by `FirestoreDatabase` for production, and `InMemoryDatabase` for testing purposes.
- [fetchers/](fetchers/) - Contains all the code for fetching from external APIs, including stubs for development.
- [handlers/](handlers/) - Contains all the http handlers for the application.
- [utils/](utils/) - Contains utility functions used by several packages.
- [webhooks/](webhooks/) - Contains all the code for invoking webhooks.

## API Specification

We follow all of the original specification made by Christopher Frantz at https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2024/-/wikis/Assignments/Assignment-2

However, we also made some extra features:

### Patch registration

Updates an existing dashboard configuration, identified by its ID. The request body should contain a JSON object with the fields to update, and only those fields will be updated.

```
Method: PATCH
Path: /dashboard/v1/registrations/{id}
```

#### Example request

```json
{
  "isoCode": "NO",
  "features": {
    "area": false
  }
}
```

This would update the dashboard registration with ID 1, modifying the `isoCode` field and `area` feature, while keeping all other fields the same.

#### Status codes

- 204: The registration was successfully updated.
- 400: Non-numeric ID, or invalid JSON
- 404: The registration with the given ID does not exist

### Patch notification

Updates an existing notification, identified by its ID. The request body should contain a JSON object with the fields to update, and only those fields will be updated.

```
Method: PATCH
Path: /dashboard/v1/notifications/{id}
```

#### Example request

`PATCH /dashboard/v1/notifications/1`

```json
{
  "country": "EN",
  "event": "INVOKE"
}
```

This would update the notification with ID 1, modifying `country` and `event` fields, while keeping all other fields the same.

#### Status codes

- 204: The notification was successfully updated.
- 400: Invalid JSON
- 404: The notification with the given ID does not exist

### Signed webhooks

When creating a notification, you can also provide a `secret` field in the request body. This secret will be used to sign the webhook, so clients can verify the authenticity of the notification.

For an example of how to verify the signature, see our test case at [handlers/notifications_test.go#126](handlers/notifications_test.go#116).

#### Example request

`POST /dashboard/v1/notifications`

```json
{
  "url": "https://webhook.site/e0355933-dfe0-4386-8483-3bb097799b33",
  "country": "NO",
  "event": "REGISTER",
  "secret": "my-very-secret-secret"
}
```

Any invoked webhooks will then include a `X-Signature` header with the HMAC signature of the payload, using the provided secret.

### Extra `map` feature for dashboards

We added an extra feature for dashboards that displays an ascii map for the country of the dashboard.

#### Example dashboard

You can register a dashboard configuration with the following JSON:

`POST /dashboard/v1/registrations`

```json
{
  "country": "Norway",
  "features": {
    "map": true
  }
}
```

And then it will be included as a multiline string when invoking the dashboard.
You can use the following one line command to fetch and display the ascii map for Norway:

```bash
curl -s "http://10.212.168.60/dashboard/v1/dashboards/92" | jq -r '.features.map'
```
