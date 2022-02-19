package tdrm

import (
	"strconv"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

func (s Station) Measure() []cfac.Measurement {
	return []cfac.Measurement{
		&cfac.RadiationMeasurement{
			BaseMeasurement: cfac.BaseMeasurement{
				Name:   "radiation",
				Source: "tdrm",
				Object: &cfac.Object{
					Location: &s.Coordinate,
					Name:     s.Name,
				},
				Time: uint64(time.Now().UnixMilli()),
				Tags: map[string]string{
					"region": s.Region,
					"id":     strconv.Itoa(s.ID),
					"unit":   "µSv/h",
				},
			},
			Radiation: s.AvgValue,
		},
	}
}

type Measurable struct{}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return Fetch(c, func(s Station) {
		cb(s.Measure())
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("tdrm", NewMeasurable)
}
