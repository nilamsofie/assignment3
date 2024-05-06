package constants

const (
	ClientSignature    = "X-SIGNATURE"
	DefaultPort        = "8080"
	WebHookServicePort = "8081"
	BaseUrl            = "/dashboard/v1/"

	// Endpoints
	BaseCountriesEndpoint  = "http://129.241.150.113:8080/v3.1/"
	BaseCurrencyEndpoint   = "http://129.241.150.113:9090/currency/"
	BaseMeteoEndpoint      = "http://api.open-meteo.com/v1/"
	WorldCountriesEndpoint = "http://inmagik.github.io/world-countries/countries/"

	// Firebase collections
	DashboardOverviewCollection      = "dashboardOverview"
	DashboardConfigurationCollection = "dashboards"
	NotificationsCollection          = "notifications"

	// Webhook constants
	Register = "REGISTER"
	Change   = "CHANGE"
	Delete   = "DELETE"
	Invoke   = "INVOKE"

	// Error messages
	MethodeNotAllowedMessage = "method not allowed"
	ServerErrorMessage       = "internal server error"
	BadGatewayMessage        = "bad gateway"

	InvalidJsonMessage        = "invalid JSON body"
	InvalidIDDocMessage       = "invalid ID"
	DataNotFound              = "data not found"
	InvalidRequestBodyMessage = "invalid request body"
	InvalidIDMessage          = "id not a valid number"
)
