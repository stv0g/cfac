// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package fitx_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/city"
	"github.com/stv0g/cfac/pkg/occupancy/gyms/fitx"
)

func TestFetchStudios(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	fitx.FetchStudios(c.Collector, city.Aachen.Center, func(s []fitx.Studio) {
		t.Logf("Found studios: %+v", s)

		found := false
		for _, t := range s {
			if t.Name != "Aachen-Europaplatz" {
				found = true
			}
		}
		if !found {
			t.Fail()
		}

		c.MarkHandled()
	}, c.ErrorCallback())
}

func TestFetchStudioWorkload(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	fitx.FetchStudioWorkload(c.Collector, 38, func(s fitx.Studio) {
		t.Logf("Found studio workload: %s => %d %%", s.Name, s.Workload.Percentage)
		c.MarkHandled()
	}, c.ErrorCallback())
}
