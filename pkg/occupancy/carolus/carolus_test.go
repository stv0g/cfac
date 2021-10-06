package carolus_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/occupancy/carolus"
)

func TestOccupancy(t *testing.T) {
	c := th.NewCollyCollector(t)
	defer c.Wait()

	handled := false

	carolus.FetchOccupancy(c, func(p carolus.Occupancy) {
		t.Logf("Occupancy: %+v\n", p)
		handled = true
	}, th.ErrorHandler(t))

	if !handled {
		t.Error("Request not handled")
	}
}
