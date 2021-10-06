package spielbank_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/occupancy/spielbank"
)

func TestSpielbankOccupancy(t *testing.T) {
	c := th.NewCollyCollector(t)
	defer c.Wait()

	handled := false

	spielbank.FetchOccupancy(c, func(p spielbank.Occupancy) {
		t.Logf("Occupancy: %+v\n", p)
		handled = true
	}, th.ErrorHandler(t))

	if !handled {
		t.Error("Request not handled")
	}
}
