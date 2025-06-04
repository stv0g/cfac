// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package nominatim_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/nominatim"
)

func TestSearch(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	nominatim.Search(c.Collector, "Aachen", func(p []nominatim.Place) {
		ac := p[0]

		t.Logf("Found: %+v\n", ac)

		if ac.Address.City != "Aachen" {
			t.Fail()
		}

		if ac.Latitude-50.776351 > 1e-6 || ac.Longitude-6.083862 > 1e-6 {
			t.Fail()
		}

		c.MarkHandled()
	}, c.ErrorCallback())
}
