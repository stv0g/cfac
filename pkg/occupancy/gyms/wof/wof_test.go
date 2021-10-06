package wof_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/occupancy/gyms/wof"
)

func TestWOFOccupancy(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	wof.FetchOccupancy(c.Collector, func(s wof.Studio) {
		t.Logf("Studio: %+v\n", s)

		c.MarkHandled()
	}, c.ErrorCallback())
}
