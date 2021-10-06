package gammasense

import "time"

const (
	UrlAPI      = "https://data.waag.org/api/gammasense"
	UrlStations = UrlAPI + "/stations"
	UrlHourly   = UrlAPI + "/hourly?sensor_id=gammasense43&start=2021-07-07T20:00:00.000Z&end=2021-10-05T20:00:00.000Z"
	UrlRecent   = UrlAPI + "/recent"
)

type ResponseRecent []Measurement
type ResponseStations []Station
type ResponseHourly []Measurement

type Station struct {
	SensorType  string    `json:"sensor_type"`
	ID          string    `json:"id"`
	Coordinates []float32 `json:"coordinates"`
}

type Measurement struct {
	Time     time.Time `json:"time"`
	CpmMean  float64   `json:"cpm_mean"`
	DeviceID string    `json:"deviceId"`
}
