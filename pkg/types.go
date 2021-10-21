package cfac

type ErrorCallback func(error)
type MeasurementsCallback func(m []Measurement)
type NewMeasurable func() Measurable

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

type BaseMeasurement struct {
	Name string `json:"name"`
	Time uint64 `json:"time"`

	Object Object `json:"object"`
	Source string `json:"source"`
}

type OccupancyMeasurement struct {
	BaseMeasurement

	Occupancy float64 `json:"occupancy"`
	Capacity  float64 `json:"capacity"`
}

type OccupancyPercentMeasurement struct {
	BaseMeasurement

	Occupancy Percent `json:"occupancy_percent"`
}

type TemperatureMeasurment struct {
	BaseMeasurement

	Temperature float64 `json:"temperature"`
}

type CounterMeasurement struct {
	BaseMeasurement

	Count uint64 `json:"count"`
}
