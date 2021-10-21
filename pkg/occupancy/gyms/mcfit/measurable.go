package mcfit

import (
	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
	"github.com/stv0g/cfac/pkg/city"
)

func (s *Studio) Measure() []cfac.Measurement {
	return []cfac.Measurement{
		&cfac.OccupancyPercentMeasurement{
			BaseMeasurement: cfac.BaseMeasurement{
				Name:   "occupancy",
				Source: "fitx",
				Object: cfac.Object{
					Name: s.StudioName,
					Location: cfac.Coordinate{
						Latitude:  s.Address.Latitude,
						Longitude: s.Address.Longitude,
					},
				},
			},

			// Occupancy: ,
		},
	}
}

type Measurable struct{}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementsCallback, ecb cfac.ErrorCallback) {
	FetchStudiosByCoordinates(c, city.Aachen.Coordinate, 10e3, func(s []Studio) {
		m := []cfac.Measurement{}
		for _, t := range s {
			m = append(m, t.Measure()...)
		}
		cb(m)
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("mcfit", NewMeasurable)
}
