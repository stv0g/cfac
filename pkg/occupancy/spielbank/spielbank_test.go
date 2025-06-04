// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package spielbank_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/occupancy/spielbank"
)

func TestSpielbankOccupancy(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	spielbank.FetchOccupancy(c.Collector, func(p spielbank.Occupancy) {
		t.Logf("Occupancy: %+v\n", p)

		c.MarkHandled()
	}, c.ErrorCallback())
}
