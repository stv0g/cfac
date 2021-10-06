package medaix_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/occupancy/gyms/medaix"
)

func TestMedaixOccupancy(t *testing.T) {
	c := th.NewCollyCollector(t)
	defer c.Wait()

	handled := false

	medaix.FetchOccupancy(c, func(p medaix.VisitorCounter) {
		t.Logf("Occupancy: %+v\n", p)
		handled = true
	}, th.ErrorHandler(t))

	if !handled {
		t.Error("Request not handled")
	}
}
