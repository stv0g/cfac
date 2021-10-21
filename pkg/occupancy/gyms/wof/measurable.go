package wof

import (
	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

type Measurable struct{}

func (s *Studio) Measure() []cfac.Measurement {
	return []cfac.Measurement{
		&cfac.OccupancyPercentMeasurement{
			BaseMeasurement: cfac.BaseMeasurement{
				Name: "wof",
				Time: uint64(s.LastUpdated.UnixMilli()),

				Object: cfac.Object{
					Name: s.Name,
				},
			},

			Occupancy: s.Occupancy,
		},
	}
}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementsCallback, ecb cfac.ErrorCallback) {
	FetchOccupancy(c, func(s Studio) {
		cb(s.Measure())
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("medaix", NewMeasurable)
}
