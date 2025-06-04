// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package ford_carsharing

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

type Attribute struct {
	Value       interface{} `json:"value"`
	Description string      `json:"description"`
}

type Attributes map[string]Attribute

type Piece struct {
	Description string `json:"description"`
}

type Equipment map[string]Piece

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

type Station struct {
	Area          Area           `json:"area"`
	Location      Location       `json:"location"`
	RentalObjects []RentalObject `json:"rentalObjects"`
}

type Response struct {
	Error   interface{} `json:"error"`
	Results []Station   `json:"result"`
}

type Request struct {
	Method string             `json:"method"`
	Params []RequestParameter `json:"params"`
	ID     int64              `json:"id"`
}

type RequestParameter struct {
	Position       Position      `json:"geoPos"`
	MaxItems       string        `json:"maxItems"`
	DateTimeStart  string        `json:"dateTimeStart"`
	DateTimeEnd    string        `json:"dateTimeEnd"`
	Address        string        `json:"address"`
	VehicleTypeIds []interface{} `json:"vehicleTypeIds"`
	Equipment      []interface{} `json:"equipment"`
}

type Position struct {
	Radius    string  `json:"radius"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}
