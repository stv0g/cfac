package fitx_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	cfac "github.com/stv0g/cfac/pkg"
	"github.com/stv0g/cfac/pkg/occupancy/gyms/fitx"
)

func TestFetchStudios(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	fitx.FetchStudios(c.Collector, cfac.Aachen.Coordinate, func(s []fitx.Studio) {
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
