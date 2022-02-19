package gammasense

import "time"

type ResponseRecent []Measurement
type ResponseStations []Station
type ResponseHourly []Measurement

type Station struct {
	SensorType  string    `json:"sensor_type"`
	ID          string    `json:"id"`
	Coordinates []float32 `json:"coordinates"`

	Measurement *Measurement `json:"-,omitempty"`
}

type Measurement struct {
	Time    time.Time `json:"timestamp"`
	CPMMean float64   `json:"cpm_mean"`
	ID      string    `json:"id"`
}
