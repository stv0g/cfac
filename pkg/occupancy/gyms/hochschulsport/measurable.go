package hochschulsport

import (
	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

type Measurable struct {
	collector     *colly.Collector
	callback      cfac.MeasurementsCallback
	errorCallback cfac.ErrorCallback
}

func NewMeasurable(c *colly.Collector, cb cfac.MeasurementsCallback, errCb cfac.ErrorCallback) cfac.Measurable {
	return &Measurable{
		collector:     c,
		callback:      cb,
		errorCallback: errCb,
	}
}

func (m *Measurable) Fetch() {
	FetchOccupancy(m.collector, func(o Occupancy) {
		m.callback(o.Measure())
	}, m.errorCallback)
}

func (o Occupancy) Measure() []cfac.Measurement {
	return []cfac.Measurement{
		cfac.OccupancyMeasurement{
			BaseMeasurement: cfac.BaseMeasurement{
				Name:   "occupancy",
				Time:   uint64(o.LastUpdated.UnixMilli()),
				Object: RWTHGym,
			},
			Occupancy: cfac.Occupancy{
				Occupancy: o.Occupancy,
			},
		},
	}
}

func init() {
	cfac.RegisterMeasurable("hochschulsport", NewMeasurable)
}
