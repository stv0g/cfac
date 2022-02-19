package gammasense_test

import (
	"testing"
	"time"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/environment/radiation/gammasense"
)

func TestGammasenseStations(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	gammasense.FetchStations(c.Collector, func(st []gammasense.Station) {
		for _, s := range st {
			t.Logf("Station: %+v", s)
		}
		c.MarkHandled()
	}, c.ErrorCallback())
}

func TestGammasenseRecent(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	gammasense.FetchRecent(c.Collector, func(m []gammasense.Measurement) {
		t.Logf("Measurements: %+v", m)
		c.MarkHandled()
	}, c.ErrorCallback())
}

func TestGammasenseHourly(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	sid := "gammasense38"
	end := time.Now()
	start := end.AddDate(0, 0, -1)

	gammasense.FetchHourly(sid, start, end, c.Collector, func(m []gammasense.Measurement) {
		t.Logf("Stations: %+v", m)
		c.MarkHandled()
	}, c.ErrorCallback())
}

func TestFetchStationMeasurements(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	gammasense.FetchStationMeasurements(c.Collector, func(s gammasense.Station) {
		t.Logf("Station: %+#v", s)
		c.MarkHandled()
	}, c.ErrorCallback())
}
