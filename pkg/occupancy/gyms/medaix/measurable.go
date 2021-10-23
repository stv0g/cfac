package medaix

import (
	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

type Measurable struct {
	StudioID int
}

func (s *VisitorCounter) Measure() []cfac.Measurement {
	return []cfac.Measurement{
		&cfac.OccupancyPercentMeasurement{
			BaseMeasurement: cfac.BaseMeasurement{
				Name:   "occupancy",
				Source: "fitx",
				Object: cfac.Object{
					Name: s.Name,
				},
			},

			Occupancy: cfac.Percent(s.Workload.Percentage),
		},
	}
}

func NewMeasurable() cfac.Measurable {
	return &Measurable{
		StudioID: 38,
	}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementsCallback, ecb cfac.ErrorCallback) {
	FetchOccupancy(m.StudioID, c, func(cnt VisitorCounter) {
		cb(cnt.Measure())
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("medaix", NewMeasurable)
}