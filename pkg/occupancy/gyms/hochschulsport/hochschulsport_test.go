package hochschulsport_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/occupancy/gyms/hochschulsport"
)

func TestRWTHGymOccupancy(t *testing.T) {
	c := th.NewCollyCollector(t)
	defer c.Wait()

	handled := false

	hochschulsport.FetchOccupancy(c, func(p hochschulsport.Occupancy) {
		t.Logf("Occupancy: %+v\n", p)
		handled = true
	}, th.ErrorHandler(t))

	if !handled {
		t.Error("Request not handled")
	}
}
