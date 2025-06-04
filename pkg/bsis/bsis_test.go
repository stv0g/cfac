// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package bsis_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/bsis"
)

func TestFetchConstructionSites(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	bsis.FetchConstructionSites(c.Collector, func(l *bsis.ConstructionSiteList) {
		if l.NumberReturned <= 0 {
			t.FailNow()
		}

		t.Logf("Total sites: %d", l.NumberMatched)

		c.MarkHandled()
	}, c.ErrorCallback())
}
