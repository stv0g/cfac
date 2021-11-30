package buienalarm_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/city"
	"github.com/stv0g/cfac/pkg/environment/weather/buienalarm"
)

func TestBuienalarm(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	buienalarm.Fetch(city.Aachen.Coordinate, c.Collector, func(fc buienalarm.Forecast) {
		t.Logf("Forecast: %+v", fc)
		c.MarkHandled()
	}, c.ErrorCallback())
}
