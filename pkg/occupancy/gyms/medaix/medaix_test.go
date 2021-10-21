package medaix_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/occupancy/gyms/medaix"
)

func TestMedaixOccupancy(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	medaix.FetchOccupancy(1, c.Collector, func(p medaix.VisitorCounter) {
		t.Logf("Occupancy: %+v\n", p)

		c.MarkHandled()
	}, c.ErrorCallback())
}
