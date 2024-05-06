package database

type DashboardConfigurationFeatures struct {
	Area             bool     `json:"area" firestore:"area"`
	Capital          bool     `json:"capital" firestore:"capital"`
	Coordinates      bool     `json:"coordinates" firestore:"coordinates"`
	Population       bool     `json:"population" firestore:"population"`
	Precipitation    bool     `json:"precipitation" firestore:"precipitation"`
	TargetCurrencies []string `json:"targetCurrencies" firestore:"targetCurrencies"`
	Temperature      bool     `json:"temperature" firestore:"temperature"`
	Map              bool     `json:"map" firestore:"map"`
}

type DashboardConfiguration struct {
	Country  string                         `json:"country" firestore:"country"`
	IsoCode  string                         `json:"isoCode" firestore:"isoCode"`
	Features DashboardConfigurationFeatures `json:"features" firestore:"features"`
}

type DashboardOverview struct {
	DeletedIds []int `json:"deletedIds" firestore:"deletedIds"`
	MaxId      int   `json:"maxId" firestore:"maxId"`
}

type Notification struct {
	Url     string  `json:"url" firestore:"url"`
	Country string  `json:"country" firestore:"country"`
	Event   string  `json:"event" firestore:"event"`
	Secret  *string `json:"secret" firestore:"secret"`
}
