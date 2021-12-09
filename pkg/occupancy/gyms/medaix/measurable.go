package medaix

import (
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

type Measurable struct{}

func (s *VisitorCounter) Measure() cfac.Measurement {
	return &cfac.OccupancyMeasurement{
		BaseMeasurement: cfac.BaseMeasurement{
			Name:   "occupancy",
			Source: "medaix",
			Object: cfac.Object{
				Name: s.Name,
			},
			Time: uint64(time.Now().UnixMilli()),
		},

		Occupancy: float64(s.Value),
		Capacity:  float64(s.Max),
	}
}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementsCallback, ecb cfac.ErrorCallback) {
	measurements := []cfac.Measurement{}

	stop := false

	for studio := 1; !stop; studio++ {

		wg := sync.WaitGroup{}
		wg.Add(1)

		FetchOccupancy(studio, c.Clone(), func(cnt VisitorCounter) {
			measurements = append(measurements, cnt.Measure())
			wg.Done()
		}, func(e error) {
			stop = true
			wg.Done()
		})

		wg.Wait()
	}

	cb(measurements)
}

func init() {
	cfac.RegisterMeasurable("medaix", NewMeasurable)
}
