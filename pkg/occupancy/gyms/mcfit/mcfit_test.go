// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package mcfit_test

import (
	"fmt"
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	cfac "github.com/stv0g/cfac/pkg"
	"github.com/stv0g/cfac/pkg/occupancy/gyms/mcfit"
)

func TestFetchStudios(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	mcfit.FetchStudios(c.Collector, func(s []mcfit.Studio) {
		if len(s) == 0 {
			t.Fail()
		}

		t.Logf("Found %d studios", len(s))

		c.MarkHandled()
	}, c.ErrorCallback())
}

func TestFetchStudiosByCoordinates(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	co := cfac.Coordinate{
		Latitude:  50.796292,
		Longitude: 6.1046713,
	}

	mcfit.FetchStudiosByCoordinates(c.Collector, co, 100, func(studios []mcfit.Studio) {
		t.Logf("Found %d studios", len(studios))

		if len(studios) != 1 {
			t.Fail()
		}

		studio := studios[0]

		t.Logf("Studio: %+v", studio)

		if studio.Address.Street != "Gut-Dämme-Str." {
			t.Fail()
		}

		if studio.ID != 1536266890 {
			t.Fail()
		}

		c.MarkHandled()
	}, c.ErrorCallback())
}

func TestFetchOccupancy(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	mcfit.FetchOccupancy(c.Collector, 1536266890, func(o mcfit.ResponseOccupancy) {
		fmt.Printf("occu: %+v", o)

		c.MarkHandled()
	}, c.ErrorCallback())
}

func TestFetchCurrentOccupancy(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	mcfit.FetchCurrentOccupancy(c.Collector, 1536266890, func(o mcfit.Occupancy) {
		fmt.Printf("occu: %+v", o)

		c.MarkHandled()
	}, c.ErrorCallback())
}
