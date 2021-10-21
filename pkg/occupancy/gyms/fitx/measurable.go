package fitx

import (
	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

func (s *Studio) Measure() []cfac.Measurement {
	return []cfac.Measurement{
		&cfac.OccupancyPercentMeasurement{
			BaseMeasurement: cfac.BaseMeasurement{
				Name:   "occupancy",
				Source: "fitx",
				Object: cfac.Object{
					Name: s.Name,
					Location: cfac.Coordinate{
						Latitude:  s.Location.Lat,
						Longitude: s.Location.Lon,
					},
				},
			},

			Occupancy: cfac.Percent(s.Workload.Percentage),
		},
	}
}

type Measurable struct {
	studioID      int
	collector     *colly.Collector
	callback      cfac.MeasurementsCallback
	errorCallback cfac.ErrorCallback
}

func NewMeasurable() cfac.Measurable {
	return &Measurable{
		studioID: 38,
	}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementsCallback, ecb cfac.ErrorCallback) {
	FetchStudioWorkload(c, m.studioID, func(studio Studio) {
		cb(studio.Measure())
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("fitx", NewMeasurable)
}
