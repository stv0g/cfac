// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

//go:build gosseract
// +build gosseract

package hochschulsport_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/occupancy/gyms/hochschulsport"
)

func TestFetchOccupancy(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	hochschulsport.FetchOccupancy(c.Collector, func(p hochschulsport.Occupancy) {
		t.Logf("Occupancy: %+v\n", p)

		c.MarkHandled()
	}, c.ErrorCallback())
}
