package mcfit

type ResponseStudios []Studio

type ResponseOccupancy struct {
	StartTime string      `json:"startTime"`
	EndTime   string      `json:"endTime"`
	Items     []Occupancy `json:"items"`
}

type StudioAddress struct {
	Street                       string      `json:"street"`
	SecondStreet                 interface{} `json:"secondStreet"`
	CityPart                     interface{} `json:"cityPart"`
	District                     interface{} `json:"district"`
	City                         string      `json:"city"`
	ZipCode                      string      `json:"zipCode"`
	StreetAddition               interface{} `json:"streetAddition"`
	HouseNumber                  string      `json:"houseNumber"`
	BuildingName                 interface{} `json:"buildingName"`
	CountryCode                  string      `json:"countryCode"`
	CountryCodeAlpha2            string      `json:"countryCodeAlpha2"`
	Longitude                    float64     `json:"longitude"`
	Latitude                     float64     `json:"latitude"`
	StreetType                   interface{} `json:"streetType"`
	StreetBlock                  interface{} `json:"streetBlock"`
	Portal                       interface{} `json:"portal"`
	Stairway                     interface{} `json:"stairway"`
	Door                         interface{} `json:"door"`
	Floor                        interface{} `json:"floor"`
	Province                     interface{} `json:"province"`
	AdditionalAddressInformation interface{} `json:"additionalAddressInformation"`
}

type StudioTag struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}

type OpeningHour struct {
	DayOfWeekFrom string `json:"dayOfWeekFrom"`
	DayOfWeekTo   string `json:"dayOfWeekTo"`
	TimeFrom      string `json:"timeFrom"`
	TimeTo        string `json:"timeTo"`
}

type Studio struct {
	ID                   int           `json:"id"`
	StudioName           string        `json:"studioName"`
	StudioPhone          string        `json:"studioPhone"`
	StudioEmail          string        `json:"studioEmail"`
	TrialSessionBookable bool          `json:"trialSessionBookable"`
	Address              StudioAddress `json:"address"`
	StudioTags           []StudioTag   `json:"studioTags"`
	MasterStudioID       interface{}   `json:"masterStudioId"`
	OpeningDate          string        `json:"openingDate"`
	ClosingDate          interface{}   `json:"closingDate"`
	OpeningHours         []OpeningHour `json:"openingHours"`
	ZoneID               string        `json:"zoneId"`
}

type Occupancy struct {
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	Percentage int    `json:"percentage"`
	Level      string `json:"level"`
	IsCurrent  bool   `json:"isCurrent"`
}
