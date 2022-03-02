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
	Name   string `json:"name,omitempty"`
	Time   uint64 `json:"time"`
	Metric string `json:"metric,omitempty"`
	Source string `json:"source,omitempty"`

	Object *Object `json:"object,omitempty"`

	Tags map[string]string `json:"tag,omitempty"`
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

type RadiationMeasurement struct {
	BaseMeasurement

	Radiation float64 `json:"radiation"`
}

type OzonMeasurement struct {
	BaseMeasurement

	Ozon float64 `json:"ozon"`
}

type SO2Measurement struct {
	BaseMeasurement

	SO2 float64 `json:"so2"`
}

type NO2Measurement struct {
	BaseMeasurement

	NO2 float64 `json:"no2"`
}

type PM10Measurement struct {
	BaseMeasurement

	PM10 float64 `json:"pm10"`
}

type CounterMeasurement struct {
	BaseMeasurement

	Count uint64 `json:"count"`
}
