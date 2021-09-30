package cambio

// https://cwapi.cambio-carsharing.com/pub/AAC/stations
// https://cwapi.cambio-carsharing.com/pub/AAC/vehicles

type StationAddress struct {
	StreetAddress   string `json:"streetAddress"`
	StreetNumber    string `json:"streetNumber"`
	AddressLocation string `json:"addressLocation"`
	PostalCode      string `json:"postalCode"`
	AddressCountry  string `json:"addressCountry"`
}

type StationInformation struct {
	Description string `json:"description"`
	Access      string `json:"access"`
	Location    string `json:"location"`
}

type StationGeoPosition struct {
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
	GoogleZoom int     `json:"googleZoom"`
}

type VehicleClass struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type Station struct {
	ID             string             `json:"id"`
	DisplayName    string             `json:"displayName"`
	Name           string             `json:"name"`
	Address        StationAddress     `json:"address"`
	Information    StationInformation `json:"information"`
	Geoposition    StationGeoPosition `json:"geoposition"`
	StationType    string             `json:"stationType"`
	VehicleClasses []VehicleClass     `json:"vehicleClasses"`
}

type StationReference struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type Equipment struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type Vehicle struct {
	ID                  string             `json:"id"`
	DisplayName         string             `json:"displayName"`
	PriceClass          string             `json:"priceClass"`
	AvailableAtStations []StationReference `json:"availableAtStations"`
	Equipment           []Equipment        `json:"equipment"`
}
