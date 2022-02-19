package blutspende

import (
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
	"github.com/stv0g/cfac/pkg/nominatim"
)

type Measurable struct{}

func (p *SpendePegelStats) Measure() cfac.Measurement {
	name := "Uniklinik RWTH Aachen"
	coord, _ := nominatim.SearchAndCache(name)

	return &cfac.CounterMeasurement{
		BaseMeasurement: cfac.BaseMeasurement{
			Name:   "pegel",
			Source: "blutspende",
			Object: &cfac.Object{
				Name:     name,
				Location: &coord,
			},
			Time: uint64(time.Now().UnixMilli()),
			Tags: map[string]string{
				"unit": "doses",
			},
		},

		Count: uint64(p.Donations),
	}
}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchPegel(c, func(pegel SpendePegelStats) {
		cb(pegel.Measure())
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("blutspende", NewMeasurable)
}
