// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package blutspende_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/health/blutspende"
)

func TestFetchPegel(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	blutspende.FetchPegel(c.Collector, func(p blutspende.SpendePegelStats) {
		t.Logf("Pegel: %+v", p)

		if p.Donations > 10000 || p.Donations <= 0 {
			t.Fail()
		}

		c.MarkHandled()
	}, c.ErrorCallback())
}
