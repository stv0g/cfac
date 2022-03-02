//go:build gosseract
// +build gosseract

package hochschulsport

import (
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

type Measurable struct{}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchOccupancy(c, func(o Occupancy) {
		cb(o.Measure())
	}, ecb)
}

func (o Occupancy) Measure() cfac.Measurement {
	return &cfac.OccupancyMeasurement{
		BaseMeasurement: cfac.BaseMeasurement{
			Name:   "occupancy",
			Source: "hochschulsport",
			Time:   uint64(o.LastUpdated.UnixMilli()),
			Object: &RWTHGym,
		},

		Occupancy: o.Occupancy,
	}
}

func init() {
	cfac.RegisterMeasurable("hochschulsport", NewMeasurable)
}
