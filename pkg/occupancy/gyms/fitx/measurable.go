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
			OccupancyPercentage: cfac.OccupancyPercentage{
				Occupancy: cfac.Percent(s.Workload.Percentage),
			},
		},
	}
}

type Measurable struct {
	studioID      int
	collector     *colly.Collector
	callback      cfac.MeasurementsCallback
	errorCallback cfac.ErrorCallback
}

func NewMeasurable(c *colly.Collector, cb cfac.MeasurementsCallback, errCb cfac.ErrorCallback) cfac.Measurable {
	return &Measurable{
		studioID:      38,
		collector:     c,
		callback:      cb,
		errorCallback: errCb,
	}
}

func (m *Measurable) Fetch() {
	FetchStudioWorkload(m.collector, m.studioID, func(studio Studio) {
		m.callback(studio.Measure())
	}, m.errorCallback)
}

func init() {
	cfac.RegisterMeasurable("fitx", NewMeasurable)
}
