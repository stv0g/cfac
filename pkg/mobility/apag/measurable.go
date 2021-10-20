package apag

import (
	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

func (h *House) Measure() []cfac.Measurement {
	return []cfac.Measurement{
		&cfac.OccupancyMeasurement{
			BaseMeasurement: cfac.BaseMeasurement{
				Name:   "occupancy",
				Source: "apag",
				Object: cfac.Object{
					Name: h.Title,
				},
			},
			Occupancy: cfac.Occupancy{
				Occupancy: float64(h.Stats.Count),
				Capacity:  float64(h.Capacity),
			},
		},
	}
}

type Measurable struct {
	cfac.Measurable

	collector *colly.Collector

	errorCallback cfac.ErrorCallback
	callback      cfac.MeasurementsCallback
}

func NewMeasurable(c *colly.Collector, cb cfac.MeasurementsCallback, errCb cfac.ErrorCallback) cfac.Measurable {
	return &Measurable{
		collector:     c,
		callback:      cb,
		errorCallback: errCb,
	}
}

func (m *Measurable) Fetch() {
	FetchAllHousesWithStats(m.collector, func(houses []House) {
		measurements := []cfac.Measurement{}
		for _, h := range houses {
			measurements = append(measurements, h.Measure()...)
		}
		m.callback(measurements)
	}, m.errorCallback)
}

func init() {
	cfac.RegisterMeasurable("apag", NewMeasurable)
}
