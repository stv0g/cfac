// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package gammasense

import (
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
				Source: "gammasense",
				Object: &cfac.Object{
					Location: &cfac.Coordinate{
						Latitude:  float64(s.Coordinates[0]),
						Longitude: float64(s.Coordinates[1]),
					},
					Name: s.ID,
				},
				Time: uint64(s.Measurement.Time.UnixMilli()),
				Tags: map[string]string{
					"type": s.SensorType,
					"id":   s.ID,
					"unit": "cpm",
				},
			},
			Radiation: s.Measurement.CPMMean,
		},
	}
}

type Measurable struct{}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchStationMeasurements(c, func(s Station) {
		if time.Since(s.Measurement.Time) < 24*time.Hour { // Skip inactive sensors
			cb(s.Measure())
		}
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("gammasense", NewMeasurable)
}
