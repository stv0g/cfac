package cfac

type ErrorCallback func(error)
type MeasurementCallback func(m Measurement)
type NewMeasurable func() Measurable

type Percent int

type Coordinate struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type BoundingBox struct {
	NorthWest Coordinate `json:"nw"`
	SouthEast Coordinate `json:"se"`
}

type Object struct {
	Name     string      `json:"name"`
	Location *Coordinate `json:"location,omitempty"`
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
	Capacity  float64 `json:"capacity,omitempty"`
}

type WaitingTimeMeasurement struct {
	BaseMeasurement

	WaitingTime  int `json:"waiting_time"`
	VisitorCount int `json:"visitor_count"`
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
