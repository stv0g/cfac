package odlinfo_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/environment/radiation/odlinfo"
)

func TestFetchStatistics(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	odlinfo.FetchStatistics(c.Collector, func(stats odlinfo.Statistics) {
		t.Logf("Stats = %+v", stats)
		c.MarkHandled()
	}, c.ErrorCallback())
}

func TestFetchStations(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	odlinfo.FetchStations(c.Collector, func(stats odlinfo.StationInfo) {
		t.Logf("Station = %+v", stats)
		c.MarkHandled()
	}, c.ErrorCallback())
}

func TestFetchStation(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	odlinfo.FetchStationData("053540042", c.Collector, func(d odlinfo.StationData) {
		t.Logf("Station = %+v", d)
		c.MarkHandled()
	}, c.ErrorCallback())
}
