// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package carolus

import (
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

func (o Occupancy) Measure() []cfac.Measurement {
	base := func(name string) cfac.BaseMeasurement {
		return cfac.BaseMeasurement{
			Name:   "occupancy",
			Source: "carolus",
			Object: &cfac.Object{
				Name:     "Carolus Therme " + name,
				Location: &Coordinate,
			},
			Time: uint64(o.LastUpdated.UnixMilli()),
		}
	}

	return []cfac.Measurement{
		cfac.OccupancyPercentMeasurement{
			BaseMeasurement: base("Parken"),

			Occupancy: o.Parking,
		},
		cfac.OccupancyPercentMeasurement{
			BaseMeasurement: base("Sauna"),

			Occupancy: o.Sauna,
		},
		cfac.OccupancyPercentMeasurement{
			BaseMeasurement: base("Thermalbad"),

			Occupancy: o.ThermalBath,
		},
	}
}

type Measurable struct{}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchOccupancy(c, func(o Occupancy) {
		for _, m := range o.Measure() {
			cb(m)
		}
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("carolus", NewMeasurable)
}
