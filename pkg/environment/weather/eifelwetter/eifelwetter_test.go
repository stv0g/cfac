// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package eifelwetter_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/environment/weather/eifelwetter"
)

func TestFetchAllStations(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	eifelwetter.FetchAllStations(c.Collector, func(s eifelwetter.Station) {
		t.Logf("Station: %+v", s)
		c.MarkHandled()
	}, c.ErrorCallback())
}
