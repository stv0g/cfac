// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package terminmanagement

import (
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

func (wc *WaitCircle) Measure() cfac.Measurement {
	return &cfac.WaitingTimeMeasurement{
		BaseMeasurement: cfac.BaseMeasurement{
			Name:   "waiting_time",
			Source: "terminmanagement",
			Object: &cfac.Object{
				Name: wc.Name,
			},
			Time: uint64(time.Now().UnixMilli()),
		},

		WaitingTime:  int(wc.WaitingTime.Seconds()),
		VisitorCount: wc.VisitorCount,
	}
}

type Measurable struct {
	studioID      int
	collector     *colly.Collector
	callback      cfac.MeasurementCallback
	errorCallback cfac.ErrorCallback
}

func NewMeasurable() cfac.Measurable {
	return &Measurable{}
}

func (m *Measurable) Fetch(c *colly.Collector, cb cfac.MeasurementCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchWaitCircles(c, func(wc WaitCircle) {
		cb(wc.Measure())
	}, ecb)
}

func init() {
	cfac.RegisterMeasurable("terminmanagement", NewMeasurable)
}
