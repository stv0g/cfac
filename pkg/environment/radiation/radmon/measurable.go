// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package radmon

import (
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

func (r Reading) Measure() []cfac.Measurement {
	return []cfac.Measurement{
		&cfac.RadiationMeasurement{
			BaseMeasurement: cfac.BaseMeasurement{
				Name:   "radiation",
				Source: "radmon",
				Object: &cfac.Object{
					Name: r.User,
				},
				Time: uint64(time.Now().UnixMilli()),
				Tags: map[string]string{
					"location": r.Location,
					"unit":     "cpm",
				},
			},
			Radiation: float64(r.CPM),
		},
	}
}

type Measurable struct {
	Users []string
}

func NewMeasurable() cfac.Measurable {
	return &Measurable{
		Users: Users,
	}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}

	for _, user := range m.Users {
		wg.Add(1)

		FetchReading(c.Clone(), user, func(r Reading) {
			defer wg.Done()

			cb(r.Measure())
		}, ecb)
	}

	return wg
}

func init() {
	cfac.RegisterMeasurable("radmon", NewMeasurable)
}
