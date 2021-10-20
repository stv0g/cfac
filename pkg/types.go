package cfac

import "github.com/gocolly/colly/v2"

type ErrorCallback func(error)
type MeasurementsCallback func(m []Measurement)
type NewMeasurable func(c *colly.Collector, cb MeasurementsCallback, errCb ErrorCallback) Measurable

type Percent int

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

type BoundingBox struct {
	NorthWest Coordinate
	SouthEast Coordinate
}

type Object struct {
	Name     string     `json:"name"`
	Location Coordinate `json:"location"`
}

type Occupancy struct {
	Occupancy float64 `json:"occupancy"`
	Capacity  float64 `json:"capacity"`
}

type OccupancyPercentage struct {
	Occupancy Percent `json:"occupancy_percent"`
}

type BaseMeasurement struct {
	Name string `json:"name"`
	Time uint64 `json:"time"`

	Object Object `json:"object"`
	Source string `json:"source"`
}

type OccupancyMeasurement struct {
	BaseMeasurement
	Occupancy
}

type OccupancyPercentMeasurement struct {
	BaseMeasurement
	OccupancyPercentage
}
