package wof_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/occupancy/gyms/wof"
)

func TestWOFOccupancy(t *testing.T) {
	c := th.NewCollyCollector(t)
	defer c.Wait()

	handled := 0

	wof.FetchOccupancy(c, func(s wof.Studio) {
		t.Logf("Studio: %+v\n", s)
		handled++
	}, th.ErrorHandler(t))

	if handled < 5 {
		t.Error("Request not handled")
	}
}
