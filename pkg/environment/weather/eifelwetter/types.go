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
	ID            int       `json:"id"`
	StationID     string    `json:"station_id"`
	Timestamp     time.Time `json:"timestamp"`
	Temperature   int       `json:"temperature"`
	Humidity      int       `json:"humidity"`
	Dewpoint      float64   `json:"dewpoint"`
	Windchill     float64   `json:"windchill"`
	AirPressure   float64   `json:"air_pressure"`
	Rain1Hour     int       `json:"rain_1_hour"`
	Rain24Hour    float64   `json:"rain_24_hour"`
	RainDay       int       `json:"rain_day"`
	WindSpeed     float64   `json:"wind_speed"`
	WindGusts     float64   `json:"wind_gusts"`
	WindDirection int       `json:"wind_direction"`
	SunshineHours float64   `json:"sunshine_hours"`
	UvIndex       float64   `json:"uv_index"`
	Luminosity    int       `json:"luminosity"`
	SnowHeight    int       `json:"snow_height"`
	CloudBase     int       `json:"cloud_base"`
}
