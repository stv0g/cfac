package mova

import "time"

const (
	UrlApi             = "https://mova.aseag.de/mbroker/rest"
	UrlAreaInformation = UrlApi + "/areainformation?locationWindow=50.772693748685995,6.080337166786195,50.77139442220511,6.089000701904297&placeOrOtherMobilityTypes=BUS_STOP,MOBILITAETSSTATION,CAR_SHARING_STATION,BIKE_SHARING_STATION,RAIL_STATION,POI,ADDRESS,BUS,CAR_SHARING,BIKE_SHARING,RAILWAY"
	UrlFreeFloating    = UrlApi + "/freefloating?locationWindow=50.772693748685995,6.080337166786195,50.77139442220511,6.089000701904297&types=SCOOTER,BIKE_SHARING,CAR_SHARING"
)

type FreeFloating struct {
	Latitude         float64     `json:"latitude"`
	Longitude        float64     `json:"longitude"`
	ProviderName     string      `json:"providerName"`
	Station          interface{} `json:"station"`
	Damages          interface{} `json:"damages"`
	Seats            int         `json:"seats"`
	Doors            int         `json:"doors"`
	VehicleColor     interface{} `json:"vehicleColor"`
	LicensePlate     interface{} `json:"licensePlate"`
	Pricing          interface{} `json:"pricing"`
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	Type             string      `json:"type"`
	Battery          int         `json:"battery"`
	Range            interface{} `json:"range"`
	Updated          time.Time   `json:"updated"`
	Bookable         bool        `json:"bookable"`
	SourceSystemType interface{} `json:"sourceSystemType"`
	VehicleSize      interface{} `json:"vehicleSize"`
	VehicleImageURL  interface{} `json:"vehicleImageURL"`
	FuelLevel        interface{} `json:"fuelLevel"`
	UseDeepLink      bool        `json:"useDeepLink"`
	AndroidDeepLink  interface{} `json:"androidDeepLink"`
	IosDeepLink      interface{} `json:"iosDeepLink"`
	FuelCardPin      interface{} `json:"fuelCardPin"`
}

type AreaInformation struct {
	ID                   int           `json:"id"`
	Version              int           `json:"version"`
	CreationDate         int64         `json:"creationDate"`
	LastModificationDate int64         `json:"lastModificationDate"`
	Name                 string        `json:"name"`
	City                 string        `json:"city"`
	District             interface{}   `json:"district"`
	Latitude             float64       `json:"latitude"`
	Longitude            float64       `json:"longitude"`
	PlaceType            string        `json:"placeType"`
	MainMobilityType     interface{}   `json:"mainMobilityType"`
	OtherMobilityTypes   []interface{} `json:"otherMobilityTypes"`
	Provider             interface{}   `json:"provider"`
	Parent               interface{}   `json:"parent"`
	SourceSystem         string        `json:"sourceSystem"`
	SourceSystemID       string        `json:"sourceSystemId"`
	Alias                interface{}   `json:"alias"`
	Description          interface{}   `json:"description"`
	Postcode             interface{}   `json:"postcode"`
	Street               interface{}   `json:"street"`
	Number               interface{}   `json:"number"`
	Country              interface{}   `json:"country"`
	CountryCode          interface{}   `json:"countryCode"`
	ViewCustomers        []interface{} `json:"viewCustomers"`
	TimeRestrictionParam interface{}   `json:"timeRestrictionParam"`
	AssignedFareZones    []interface{} `json:"assignedFareZones"`
	Active               bool          `json:"active"`
	TrackVehicles        bool          `json:"trackVehicles"`
	OnlyElectric         bool          `json:"onlyElectric"`
	TransparentMarker    bool          `json:"transparentMarker"`
	MarkerWidth          interface{}   `json:"markerWidth"`
	MarkerHeight         interface{}   `json:"markerHeight"`
	OwningCustomerID     interface{}   `json:"owningCustomerId"`
	AdditionalInfo       interface{}   `json:"additionalInfo"`
	Breadcrumb           interface{}   `json:"breadcrumb"`
}
