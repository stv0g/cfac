package apag_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/mobility/apag"
)

func TestFetchAllHouses(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	apag.FetchAllHouses(c.Collector, func(h []apag.House) {
		for _, i := range h {
			t.Logf("%#+v", i)
		}

		c.MarkHandled()
	}, c.ErrorCallback())
}

func TestFetchAllHouseStats(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	apag.FetchAllHouseStats(c.Collector, func(h []apag.Stats) {
		for _, i := range h {
			t.Logf("%#+v", i)
		}

		c.MarkHandled()
	}, c.ErrorCallback())
}

func TestFetchAllHousesWithStats(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	apag.FetchHousesWithStats(c.Collector, func(h apag.House) {
		t.Logf("%+v", h)

		c.MarkHandled()
	}, c.ErrorCallback())
}
