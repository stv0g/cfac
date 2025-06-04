// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package flinkster

type Collection struct {
	Href   string `json:"href"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Size   int    `json:"size"`
	Links  []Link `json:"_links"`
}

type ResponseItem struct {
	Href   string `json:"href"`
	UID    string `json:"uid"`
	Name   string `json:"name"`
	Expand string `json:"expand"`
	Links  []Link `json:"_links"`
}

type Areas struct {
	Collection

	Items []Area `json:"items"`
}

type Area struct {
	ResponseItem

	Provider          Link           `json:"provider"`
	ProviderAreaID    int            `json:"providerAreaId"`
	Type              string         `json:"type"`
	Geometry          Geometry       `json:"geometry"`
	Address           Address        `json:"address"`
	Attributes        AreaAttributes `json:"attributes"`
	RentalObjectTypes []string       `json:"rentalObjectTypes"`
}

type AreaAttributes struct {
	LocationNote string `json:"locationnote"`
}

type Geometry struct {
	Position Position `json:"position"`
}

type Position struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Link struct {
	Rel  string `json:"rel"`
	Verb string `json:"verb"`
	Href string `json:"href"`
}

type Address struct {
	Street         string `json:"street"`
	Zip            int    `json:"zip"`
	City           string `json:"city"`
	District       string `json:"district"`
	IsoCountryCode string `json:"isoCountryCode"`
}

type BookingProposals struct {
	Collection

	Items []BookingProposal `json:"items"`
}

type BookingProposal struct {
	Expand       string       `json:"expand"`
	RentalObject RentalObject `json:"rentalObject"`
	Area         Area         `json:"area"`
	Position     Position     `json:"position"`
	Price        Price        `json:"price"`
}

type Prices struct {
	Href  string  `json:"href"`
	Size  int     `json:"size"`
	Items []Price `json:"items"`
}

type Price struct {
	Type           string  `json:"type"`
	GrossAmount    float64 `json:"grossamount"`
	Currency       string  `json:"currency"`
	TaxRate        float64 `json:"taxrate"`
	PreferredPrice bool    `json:"preferredprice"`
	Interval       int     `json:"interval,omitempty"`
}

type RentalObject struct {
	ResponseItem

	ProviderRentalObjectID int                    `json:"providerRentalObjectId"`
	RentalModel            string                 `json:"rentalModel"`
	Type                   string                 `json:"type"`
	Provider               Link                   `json:"provider"`
	Category               Link                   `json:"category"`
	Attributes             RentalObjectAttributes `json:"attributes"`
}

type RentalObjectAttributes struct {
	LicensePlate     string `json:"licenseplate"`
	TransmissionType string `json:"transmissionType"`
	Colour           string `json:"colour"`
	Seats            int    `json:"seats"`
	Doors            int    `json:"doors"`
	FillLevel        int    `json:"fillLevel"`
	Fuel             string `json:"fuel"`
}

type Categories struct {
	Collection

	Items []Category `json:"items"`
}

type Category struct {
	ResponseItem
}

type ProviderNetwork struct {
	ResponseItem

	Description string `json:"description"`
}
