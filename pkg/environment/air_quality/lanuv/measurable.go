// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package lanuv

import (
	"math"
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

func (aq AirQuality) Measure() []cfac.Measurement {
	meas := []cfac.Measurement{}

	base := cfac.BaseMeasurement{
		Name:   "air_quality",
		Source: "lanuv",
		Object: &cfac.Object{
			Name: aq.Station,
		},
		Time: uint64(aq.Timestamp.UnixMilli()),
		Tags: map[string]string{
			"unit": "µg/m³",
		},
	}

	if !math.IsNaN(aq.Ozon) {
		meas = append(meas, &cfac.OzonMeasurement{
			BaseMeasurement: base,
			Ozon:            aq.Ozon,
		})
	}

	if !math.IsNaN(aq.SO2) {
		meas = append(meas, &cfac.SO2Measurement{
			BaseMeasurement: base,
			SO2:             aq.SO2,
		})
	}

	if !math.IsNaN(aq.NO2) {
		meas = append(meas, &cfac.NO2Measurement{
			BaseMeasurement: base,
			NO2:             aq.NO2,
		})
	}

	if !math.IsNaN(aq.PM10) {
		meas = append(meas, &cfac.PM10Measurement{
			BaseMeasurement: base,
			PM10:            aq.PM10,
		})
	}

	return meas
}

type Measurable struct{}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return Fetch(c, func(aq AirQuality) {
		cb(aq.Measure())
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("lanuv", NewMeasurable)
}
