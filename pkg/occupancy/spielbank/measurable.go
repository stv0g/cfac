// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package spielbank

import (
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

func (u Occupancy) Measure() cfac.Measurement {
	return &cfac.OccupancyPercentMeasurement{
		BaseMeasurement: cfac.BaseMeasurement{
			Name:   "spielbank",
			Time:   uint64(u.LastUpdated.UnixMilli()),
			Object: SpielbankAachen,
		},

		Occupancy: cfac.Percent(u.Utilization),
	}
}

type Measurable struct{}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchOccupancy(c, func(u Occupancy) {
		cb(u.Measure())
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("spielbank", NewMeasurable)
}
