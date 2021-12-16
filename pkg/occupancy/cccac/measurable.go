package cccac

import (
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
	"github.com/stv0g/cfac/pkg/nominatim"
)

type Measurable struct{}

func (sts *Status) Measure() cfac.Measurement {
	name := "Chaos Computer Club Aachen"
	coord, _ := nominatim.SearchAndCache(name)

	occ := 0
	if sts.Status == "open" {
		occ = 1
	}

	return &cfac.OccupancyMeasurement{
		BaseMeasurement: cfac.BaseMeasurement{
			Name:   "occupancy",
			Source: "cccac",
			Object: cfac.Object{
				Name:     name,
				Location: &coord,
			},
			Time: uint64(sts.Time * 1000),
		},

		Occupancy: float64(occ),
	}
}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchStatus(c, func(sts Status) {
		cb(sts.Measure())
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("cccac", NewMeasurable)
}
