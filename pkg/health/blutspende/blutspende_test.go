package blutspende_test

import (
	"testing"
	"time"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/health/blutspende"
)

func TestFetchPegel(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	blutspende.FetchPegel(c.Collector, func(p blutspende.Pegel) {
		t.Logf("Pegel: %d", p.Donations)

		if p.Donations > 10000 || p.Donations <= 0 {
			t.Fail()
		}

		c.MarkHandled()
	}, c.ErrorCallback())
}

func TestFetchPegelTime(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	ti, err := time.Parse("2006-01", "2018-12")
	if err != nil {
		t.FailNow()
	}

	blutspende.FetchPegelTime(c.Collector, ti, func(p blutspende.Pegel) {
		t.Logf("Pegel: %d", p.Donations)

		if p.Donations != 1551 {
			t.Fail()
		}

		c.MarkHandled()
	}, c.ErrorCallback())
}
