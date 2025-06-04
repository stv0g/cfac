// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package testing_helper

import (
	"testing"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

type Collector struct {
	*colly.Collector

	handled int
	testing *testing.T
}

func NewCollector(t *testing.T) *Collector {
	c := &Collector{
		Collector: colly.NewCollector(),
		handled:   0,
		testing:   t,
	}

	c.OnError(c.Error)

	return c
}

func (c *Collector) Error(r *colly.Response, e error) {
	c.testing.Errorf("Colly error: %s", e)
}

func (c *Collector) ErrorCallback() cfac.ErrorCallback {
	return func(err error) {
		c.testing.Errorf("Error for handler: %s", err)
	}
}

func (c *Collector) MarkHandled() {
	c.handled++
}

func (c *Collector) Close() {
	c.Wait()

	if c.handled <= 0 {
		c.testing.Error("request not handled")
	}
}
