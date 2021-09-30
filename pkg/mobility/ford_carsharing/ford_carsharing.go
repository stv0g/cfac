package ford_carsharing

// curl 'https://www.ford-carsharing.de/de/rpc' \
//   --data-raw '{"method":"Search.byGeoPosition","params":[{"geoPos":{"radius":"7000","lat":50.790061643,"lng":6.060413354},"maxItems":"200","dateTimeStart":"2021-09-30 12:15:00","dateTimeEnd":"2021-09-30 13:15:00","address":"Aachen, Deutschland","vehicleTypeIds":[],"equipment":[]}],"id":1632996963387}'

const (
	Url = "https://www.ford-carsharing.de/de/rpc"
)

type Area struct {
	UID                 string `json:"uid"`
	StationID           string `json:"stationId"`
	AreaType            string `json:"areaType"`
	AreaInfoAPILink     string `json:"areaInfoAPILink"`
	AvailabilityAPILink string `json:"availabilityAPILink"`
}

type Provider struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

type ProviderGroup struct {
	UID string `json:"uid"`
}

type Location struct {
	Distance      int           `json:"distance"`
	LocationName  string        `json:"locationName"`
	LocationType  string        `json:"locationType"`
	Lat           float64       `json:"lat"`
	Lng           float64       `json:"lng"`
	Coordinates   []float64     `json:"coordinates"`
	Provider      Provider      `json:"provider"`
	ProviderGroup ProviderGroup `json:"providerGroup"`
}

type Attributes struct {
	Seats struct {
		Value       int    `json:"value"`
		Description string `json:"description"`
	} `json:"seats"`
	Colour struct {
		Value       interface{} `json:"value"`
		Description string      `json:"description"`
	} `json:"colour"`
	Fuel struct {
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"fuel"`
	Licenseplate struct {
		Value       string      `json:"value"`
		Description interface{} `json:"description"`
	} `json:"licenseplate"`
	Charging         interface{} `json:"charging"`
	TransmissionType struct {
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"transmissionType"`
	Model     interface{} `json:"model"`
	FillLevel struct {
		Value       int    `json:"value"`
		Description string `json:"description"`
	} `json:"fillLevel"`
	Doors struct {
		Value       int    `json:"value"`
		Description string `json:"description"`
	} `json:"doors"`
}

type Equipment struct {
	ParkDistanceControl struct {
		Description string `json:"description"`
	} `json:"parkDistanceControl"`
	TrailerCoupling           interface{} `json:"trailerCoupling"`
	BluetoothHandsFreeCalling struct {
		Description string `json:"description"`
	} `json:"bluetoothHandsFreeCalling"`
	SunRoof                interface{} `json:"sunRoof"`
	CdPlayer               interface{} `json:"cdPlayer"`
	ChargingCableForSchuko interface{} `json:"chargingCableForSchuko"`
	AirConditioning        struct {
		Description string `json:"description"`
	} `json:"airConditioning"`
	HighRoof          interface{} `json:"highRoof"`
	ParticulateFilter interface{} `json:"particulateFilter"`
	Vignettes         interface{} `json:"vignettes"`
	AudioInline       struct {
		Description string `json:"description"`
	} `json:"audioInline"`
	ElectricalHardtop interface{} `json:"electricalHardtop"`
	ChildSeats        interface{} `json:"childSeats"`
	RoofRailing       interface{} `json:"roofRailing"`
	NavigationSystem  struct {
		Description string `json:"description"`
	} `json:"navigationSystem"`
	TyreType struct {
		Description string `json:"description"`
	} `json:"tyreType"`
	IsofixSeatFittings struct {
		Description string `json:"description"`
	} `json:"isofixSeatFittings"`
	LoadingBayCover   interface{} `json:"loadingBayCover"`
	FullBranding      interface{} `json:"fullBranding"`
	SnowChains        interface{} `json:"snowChains"`
	EmissionsStickers struct {
		Description string `json:"description"`
	} `json:"emissionsStickers"`
	CruiseControl struct {
		Description string `json:"description"`
	} `json:"cruiseControl"`
	ParkHeating            interface{} `json:"parkHeating"`
	ChargingCableForType2  interface{} `json:"chargingCableForType2"`
	DividedRearBenchSeat   interface{} `json:"dividedRearBenchSeat"`
	HeatedSeats            interface{} `json:"heatedSeats"`
	FoldingRoof            interface{} `json:"foldingRoof"`
	PassengerAirbagTurnOff struct {
		Description string `json:"description"`
	} `json:"passengerAirbagTurnOff"`
}

type Category struct {
	UID         string      `json:"uid"`
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
}

type Price struct {
	PriceType      string      `json:"priceType"`
	Type           interface{} `json:"type"`
	Interval       int         `json:"interval"`
	GrossAmount    int         `json:"grossAmount"`
	Currency       string      `json:"currency"`
	TaxRate        float64     `json:"taxRate"`
	PreferredPrice bool        `json:"preferredPrice"`
}

type RentalObject struct {
	UID                    string     `json:"uid"`
	Attributes             Attributes `json:"attributes"`
	Category               Category   `json:"category"`
	Equipment              Equipment  `json:"equipment"`
	Name                   string     `json:"name"`
	Model                  string     `json:"model"`
	ProviderRentalObjectID string     `json:"providerRentalObjectId"`
	Type                   string     `json:"type"`
	Price                  Price      `json:"price"`
	PriceString            string     `json:"price_string"`
	PriceAddition          string     `json:"price_addition"`
	PriceValue             int        `json:"price_value"`
}

type Result struct {
	Area          Area           `json:"area"`
	Location      Location       `json:"location"`
	RentalObjects []RentalObject `json:"rentalObjects"`
}

type Response struct {
	Error  interface{} `json:"error"`
	Result []Result    `json:"result"`
}
