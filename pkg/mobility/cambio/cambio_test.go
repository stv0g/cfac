package cambio_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/mobility/cambio"
)

func TestFetchStations(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	cambio.FetchStations("AAC", c.Collector, func(station cambio.Station) {
		t.Logf("Station = %+v", station)
		c.MarkHandled()
	}, c.ErrorCallback())
}

func TestFetchVehicles(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	cambio.FetchVehicles("AAC", c.Collector, func(vehicle cambio.Vehicle) {
		t.Logf("Vehicle = %+v", vehicle)
		c.MarkHandled()
	}, c.ErrorCallback())
}
