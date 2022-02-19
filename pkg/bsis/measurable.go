package bsis

import (
	"errors"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

func getDate(props map[string]interface{}, key string) (time.Time, error) {
	if val, ok := props[key]; ok {
		if str, ok := val.(string); ok {
			return time.Parse("2006-01-02", str)
		}
	}

	return time.Time{}, errors.New("failed to extract")
}

func (l ConstructionSiteList) Measure() []cfac.Measurement {
	var start, end time.Time
	var err error

	typesCount := map[string]int{}

	for _, site := range l.Features {
		if key, ok := site.Properties["art"]; ok {
			if typ, ok := key.(string); ok {
				// Check if Active
				if start, err = getDate(site.Properties, "von"); err != nil {
					continue
				}
				if end, err = getDate(site.Properties, "bis"); err != nil {
					continue
				}

				if time.Now().After(start) && time.Now().Before(end.Add(24*time.Hour)) {
					if _, ok := typesCount[typ]; !ok {
						typesCount[typ] = 1
					} else {
						typesCount[typ]++
					}
				}
			}
		}
	}

	base := cfac.BaseMeasurement{
		Name:   "counter",
		Metric: "construction_sites",
		Time:   uint64(l.Timestamp.UnixMilli()),
		Source: "bsis",
	}

	meas := []cfac.Measurement{}
	for typ, cnt := range typesCount {
		m := cfac.CounterMeasurement{
			BaseMeasurement: base,
			Count:           uint64(cnt),
		}
		m.Tags = map[string]string{
			"type": typ,
		}

		meas = append(meas, m)
	}

	return meas
}

type Measurable struct{}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchConstructionSites(c, func(l *ConstructionSiteList) {
		for _, m := range l.Measure() {
			cb(m)
		}
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("bsis", NewMeasurable)
}
