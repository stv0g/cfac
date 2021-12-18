package medaix

import (
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

type Measurable struct{}

func (s *VisitorCounter) Measure() cfac.Measurement {
	// coord, _ := nominatim.SearchAndCache(s.Name)

	return &cfac.OccupancyMeasurement{
		BaseMeasurement: cfac.BaseMeasurement{
			Name:   "occupancy",
			Source: "medaix",
			Object: &cfac.Object{
				Name: s.Name,
				// Location: &coord,
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

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}

	stop := false

	for studio := 1; !stop; studio++ {

		wgg := sync.WaitGroup{}
		wgg.Add(1)

		FetchOccupancy(studio, c.Clone(), func(cnt VisitorCounter) {
			cb(cnt.Measure())
			wgg.Done()
		}, func(e error) {
			stop = true
			wgg.Done()
		})

		wgg.Wait()
	}

	return wg
}

func init() {
	cfac.RegisterMeasurable("medaix", NewMeasurable)
}
