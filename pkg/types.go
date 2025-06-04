// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package cfac

type (
	ErrorCallback       func(error)
	MeasurementCallback func(m Measurement)
	NewMeasurable       func() Measurable
)

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

func (b BaseMeasurement) WithTags(kv ...string) BaseMeasurement {
	bc := b

	bc.Tags = map[string]string{}
	for k, v := range b.Tags {
		bc.Tags[k] = v
	}

	for i := 0; i+1 < len(kv); i += 2 {
		key := kv[i]
		value := kv[i+1]

		bc.Tags[key] = value
	}

	return bc
}

func (b BaseMeasurement) WithUnit(unit string) BaseMeasurement {
	return b.WithTags("unit", unit)
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

type HumidityMeasurement struct {
	BaseMeasurement

	Humidity float64 `json:"humidity"`
}

type AirPressureMeasurement struct {
	BaseMeasurement

	AirPressure float64 `json:"air_pressure"`
}

type DewPointMeasurement struct {
	BaseMeasurement

	DewPoint float64 `json:"dew_point"`
}

type SunshineHoursMeasurement struct {
	BaseMeasurement

	SunshineHours float64 `json:"sunshine_hours"`
}

type WindChillMeasurement struct {
	BaseMeasurement

	WindChill float64 `json:"wind_chill"`
}

type UVIndexMeasurement struct {
	BaseMeasurement

	UVIndex float64 `json:"uv_index"`
}

type LuminosityMeasurement struct {
	BaseMeasurement

	Luminosity float64 `json:"luminosity"`
}

type SnowHeightMeasurement struct {
	BaseMeasurement

	SnowHeight float64 `json:"snow_height"`
}

type CloudBaseMeasurement struct {
	BaseMeasurement

	CloudBase float64 `json:"cloud_base"`
}

type WindDirectionMeasurement struct {
	BaseMeasurement

	WindDirection float64 `json:"wind_direction"`
}

type WindGustsMeasurement struct {
	BaseMeasurement

	WindGusts float64 `json:"wind_gusts"`
}

type WindSpeedMeasurement struct {
	BaseMeasurement

	WindSpeed float64 `json:"wind_speed"`
}

type RainMeasurement struct {
	BaseMeasurement

	Rain float64 `json:"rain"`
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
