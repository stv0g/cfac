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
					Location: &cfac.Coordinate{
						Latitude:  float64(h.Latitude),
						Longitude: float64(h.Longitude),
					},
				},
				Time: uint64(h.Stats.Date.UnixMilli()),
			},

			Occupancy: float64(h.Stats.Count),
			Capacity:  float64(h.Capacity),
		},
	}
}

type Measurable struct{}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementsCallback, ecb cfac.ErrorCallback) {
	FetchAllHousesWithStats(c, func(houses []House) {
		measurements := []cfac.Measurement{}
		for _, h := range houses {
			measurements = append(measurements, h.Measure()...)
		}
		cb(measurements)
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("apag", NewMeasurable)
}
