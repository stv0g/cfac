package spielbank

import (
	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

func (u Occupancy) Measure() []cfac.Measurement {
	return []cfac.Measurement{
		&cfac.OccupancyPercentMeasurement{
			BaseMeasurement: cfac.BaseMeasurement{
				Name:   "spielbank",
				Time:   uint64(u.LastUpdated.UnixMilli()),
				Object: SpielbankAachen,
			},

			Occupancy: cfac.Percent(u.Utilization),
		},
	}
}

type Measurable struct{}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementsCallback, ecb cfac.ErrorCallback) {
	FetchOccupancy(c, func(u Occupancy) {
		cb(u.Measure())
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("spielbank", NewMeasurable)
}
