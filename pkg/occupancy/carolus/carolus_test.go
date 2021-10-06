package carolus_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/occupancy/carolus"
)

func TestOccupancy(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	carolus.FetchOccupancy(c, func(p carolus.Occupancy) {
		t.Logf("Occupancy: %+v\n", p)

		c.MarkHandled()
	}, c.ErrorCallback())
}
