package medaix

import (
	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

func (s *VisitorCounter) Measure() []cfac.Measurement {
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

func init() {
	cfac.RegisterMeasurable("medaix", func(c *colly.Collector, cb cfac.MeasurementsCallback, errCb cfac.ErrorCallback) {
		FetchOccupancy(c, 38, func(cnt VisitorCounter) {
			cb(studio.Measure())
		}, errCb)
	})
}
