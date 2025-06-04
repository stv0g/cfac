// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package fitx

import (
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

func (s *Studio) Measure() cfac.Measurement {
	return &cfac.OccupancyPercentMeasurement{
		BaseMeasurement: cfac.BaseMeasurement{
			Name:   "occupancy",
			Source: "fitx",
			Object: &cfac.Object{
				Name: "FITX " + s.Name,
				Location: &cfac.Coordinate{
					Latitude:  s.Location.Lat,
					Longitude: s.Location.Lon,
				},
			},
			Time: uint64(time.Now().UnixMilli()),
		},

		Occupancy: cfac.Percent(s.Workload.Percentage),
	}
}

type Measurable struct {
	studioID      int
	collector     *colly.Collector
	callback      cfac.MeasurementCallback
	errorCallback cfac.ErrorCallback
}

func NewMeasurable() cfac.Measurable {
	return &Measurable{
		studioID: 38,
	}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchStudioWorkload(c, m.studioID, func(studio Studio) {
		cb(studio.Measure())
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("fitx", NewMeasurable)
}
