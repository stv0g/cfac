package blutspende

import "time"

type ResponseLocations []Location

type ResponseSpendePegelStats struct {
	Fill  float64 `json:"fill"`
	Label string  `json:"label"`
}

type ResponseDonationDates []DonationDate

type DonationDate struct {
	ID              int       `json:"id"`
	Date            time.Time `json:"date"`
	FreePlacesCount int       `json:"freePlacesCount"`
	Location        Location  `json:"location"`
	Slots           []Slot    `json:"slots"`
}

type Slot struct {
	Key     string `json:"key"`
	Enabled bool   `json:"enabled"`
	Hour    int    `json:"hour"`
	Minute  int    `json:"minute"`
}

type Location struct {
	ID          int     `json:"id"`
	Label       string  `json:"label"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	GeoAddress  string  `json:"geoAddress"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Description struct {
		Md string `json:"md"`
	} `json:"description"`
}
