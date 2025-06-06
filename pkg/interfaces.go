// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package cfac

import (
	"sync"

	"github.com/gocolly/colly/v2"
)

type Measurement interface{}

type Measurable interface {
	Fetch(c *colly.Collector, cb MeasurementCallback, ecb ErrorCallback) *sync.WaitGroup
}
