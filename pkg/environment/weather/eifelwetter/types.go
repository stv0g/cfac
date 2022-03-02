package eifelwetter

import "time"

type Comment struct {
	ID           int       `json:"id"`
	Comment      string    `json:"comment"`
	LastEditDate time.Time `json:"lastEditDate"`
	Author       int       `json:"author"`
	Published    bool      `json:"published"`
}

type StationInfo struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Type      string  `json:"type"`
}

type Station struct {
	*StationInfo

	ID            int       `json:"id"`
	StationID     string    `json:"station_id"`
	Timestamp     time.Time `json:"timestamp"`
	Temperature   float64   `json:"temperature"`
	Humidity      float64   `json:"humidity"`
	DewPoint      float64   `json:"dewpoint"`
	WindChill     float64   `json:"windchill"`
	AirPressure   float64   `json:"air_pressure"`
	Rain1Hour     float64   `json:"rain_1_hour"`
	Rain24Hour    float64   `json:"rain_24_hour"`
	RainDay       int       `json:"rain_day"`
	WindSpeed     float64   `json:"wind_speed"`
	WindGusts     float64   `json:"wind_gusts"`
	WindDirection int       `json:"wind_direction"`
	SunshineHours float64   `json:"sunshine_hours"`
	UVIndex       float64   `json:"uv_index"`
	Luminosity    float64   `json:"luminosity"`
	SnowHeight    float64   `json:"snow_height"`
	CloudBase     float64   `json:"cloud_base"`
}
