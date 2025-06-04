// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package wof

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
	"github.com/stv0g/cfac/pkg/nominatim"
)

type Measurable struct{}

func (s *Studio) Measure() cfac.Measurement {
	query := fmt.Sprintf("%s %s, Aachen, Deutschland", s.Name, s.Location)
	coord, _ := nominatim.SearchAndCache(query)

	return &cfac.OccupancyPercentMeasurement{
		BaseMeasurement: cfac.BaseMeasurement{
			Name:   "occupancy",
			Source: "wof",
			Time:   uint64(s.LastUpdated.UnixMilli()),

			Object: &cfac.Object{
				Name:     s.Name,
				Location: &coord,
			},
		},

		Occupancy: s.Occupancy,
	}
}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchOccupancy(c, func(s Studio) {
		cb(s.Measure())
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("wof", NewMeasurable)
}
