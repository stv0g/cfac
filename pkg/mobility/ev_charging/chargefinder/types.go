package chargefinder

type ResponseStation Station

type ResponseStations []Station

type Station struct {
	Slug                string          `json:"slug"`
	Title               string          `json:"title"`
	Location            Location        `json:"location"`
	Description         string          `json:"description"`
	Comment             string          `json:"comment"`
	Contact             string          `json:"contact"`
	Website             string          `json:"website"`
	LocationAddress     LocationAddress `json:"locationAddress"`
	Operator            string          `json:"operator"`
	OperatorPricingInfo interface{}     `json:"operatorPricingInfo"`
	ChargeInfoLocal     interface{}     `json:"chargeInfoLocal"`
	ChargeInfoEn        interface{}     `json:"chargeInfoEn"`
	Owner               string          `json:"owner"`
	Access              string          `json:"access"`
	Reservation         string          `json:"reservation"`
	Realtime            int             `json:"realtime"`
	FreeCharging        int             `json:"freeCharging"`
	OpenStatus          int             `json:"openStatus"`
	OpeningHours        string          `json:"openingHours"`
	Elevation           int             `json:"elevation"`
	StationType         string          `json:"stationType"`
	MaxCapacity         int             `json:"maxCapacity"`
	Status              int             `json:"status"`
	OutletList          []Outlet        `json:"outletList"`
	OutletCount         int             `json:"outletCount"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type LocationAddress struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	County      string `json:"county"`
	Street      string `json:"street"`
	Full        string `json:"full"`
	Zip         string `json:"zip"`
}

type Outlet struct {
	Capacity     int             `json:"capacity"`
	CostCurrency string          `json:"costCurrency"`
	CostKwh      int             `json:"costKwh"`
	CostMin      int             `json:"costMin"`
	Count        int             `json:"count"`
	Current      int             `json:"current"`
	Outlets      []OutletDetails `json:"outlets"`
	Plug         string          `json:"plug"`
	Slug         string          `json:"slug"`
	Status       int             `json:"status"`
	Voltage      int             `json:"voltage"`
}

type OutletDetails struct {
	Plug         string `json:"plug"`
	Identifier   string `json:"identifier"`
	Name         string `json:"name"`
	Capacity     int    `json:"capacity"`
	Current      int    `json:"current"`
	Voltage      int    `json:"voltage"`
	Status       int    `json:"status"`
	CostMin      int    `json:"costMin"`
	CostKwh      int    `json:"costKwh"`
	CostCurrency string `json:"costCurrency"`
	Slug         string `json:"slug"`
}
