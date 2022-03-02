package eifelwetter

import (
	"strconv"
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

func (s Station) Measure() []cfac.Measurement {
	meas := []cfac.Measurement{}

	base := cfac.BaseMeasurement{
		Name:   "weather",
		Source: "eifelwetter",
		Object: &cfac.Object{
			Name: s.Name,
			Location: &cfac.Coordinate{
				Latitude:  s.Latitude,
				Longitude: s.Longitude,
			},
		},
		Tags: map[string]string{
			"id":         strconv.Itoa(s.ID),
			"station_id": s.StationID,
			"type":       s.Type,
		},
		Time: uint64(s.Timestamp.UnixMilli()),
	}

	meas = append(meas, &cfac.TemperatureMeasurment{
		BaseMeasurement: base.WithUnit("°C"),
		Temperature:     s.Temperature,
	})

	meas = append(meas, &cfac.HumidityMeasurement{
		BaseMeasurement: base.WithUnit("%"),
		Humidity:        s.Humidity,
	})

	meas = append(meas, &cfac.WindChillMeasurement{
		BaseMeasurement: base.WithUnit("°C"),
		WindChill:       s.WindChill,
	})

	meas = append(meas, &cfac.AirPressureMeasurement{
		BaseMeasurement: base.WithUnit("mbar"),
		AirPressure:     s.AirPressure,
	})

	meas = append(meas, &cfac.RainMeasurement{
		BaseMeasurement: base.WithUnit("l").WithTags("horizon", "1h"),
		Rain:            s.Rain1Hour,
	})

	meas = append(meas, &cfac.RainMeasurement{
		BaseMeasurement: base.WithUnit("l").WithTags("horizon", "24h"),
		Rain:            s.Rain24Hour,
	})

	meas = append(meas, &cfac.RainMeasurement{
		BaseMeasurement: base.WithUnit("l").WithTags("horizon", "1d"),
		Rain:            float64(s.RainDay),
	})

	meas = append(meas, &cfac.WindSpeedMeasurement{
		BaseMeasurement: base.WithUnit("km/h"),
		WindSpeed:       s.WindSpeed,
	})

	meas = append(meas, &cfac.WindGustsMeasurement{
		BaseMeasurement: base.WithUnit("km/h"),
		WindGusts:       s.WindGusts,
	})

	meas = append(meas, &cfac.WindDirectionMeasurement{
		BaseMeasurement: base.WithUnit("°"),
		WindDirection:   float64(s.WindDirection),
	})

	meas = append(meas, &cfac.SunshineHoursMeasurement{
		BaseMeasurement: base.WithUnit("h"),
		SunshineHours:   s.SunshineHours,
	})

	meas = append(meas, &cfac.UVIndexMeasurement{
		BaseMeasurement: base,
		UVIndex:         s.UVIndex,
	})

	meas = append(meas, &cfac.LuminosityMeasurement{
		BaseMeasurement: base,
		Luminosity:      s.Luminosity,
	})

	meas = append(meas, &cfac.SnowHeightMeasurement{
		BaseMeasurement: base.WithUnit("m"),
		SnowHeight:      s.SnowHeight,
	})

	meas = append(meas, &cfac.CloudBaseMeasurement{
		BaseMeasurement: base.WithUnit("m"),
		CloudBase:       s.CloudBase,
	})

	return meas
}

type Measurable struct{}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchAllStations(c, func(s Station) {
		cb(s.Measure())
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("eifelwetter", NewMeasurable)
}
