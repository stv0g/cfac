package mcfit

import (
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
	"github.com/stv0g/cfac/pkg/city"
)

type Measurable struct {
	Location cfac.Coordinate
	Radius   float32
}

func NewMeasurable() cfac.Measurable {
	return &Measurable{
		Location: city.Aachen.Center,
		Radius:   30e3,
	}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	FetchStudiosByCoordinates(c, m.Location, float64(m.Radius), func(t []Studio) {
		defer wg.Done()

		for _, s := range t {
			wg := FetchCurrentOccupancy(c.Clone(), s.ID, func(o Occupancy) {
				cb(&cfac.OccupancyPercentMeasurement{
					BaseMeasurement: cfac.BaseMeasurement{
						Name:   "occupancy",
						Source: "mcfit",
						Time:   uint64(time.Now().UnixMilli()),
						Object: &cfac.Object{
							Name: s.StudioName,
							Location: &cfac.Coordinate{
								Latitude:  s.Address.Latitude,
								Longitude: s.Address.Longitude,
							},
						},
					},

					Occupancy: cfac.Percent(o.Percentage),
				})
			}, ecb)
			wg.Wait()
		}
	}, ecb)

	return wg
}

func init() {
	cfac.RegisterMeasurable("mcfit", NewMeasurable)
}
