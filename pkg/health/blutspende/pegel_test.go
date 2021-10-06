package blutspende_test

import (
	"fmt"
	"testing"

	"github.com/gocolly/colly/v2"
	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/health/blutspende"
)

func TestFetchPegel(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	blutspende.FetchPegel(c.Collector, func(p blutspende.Pegel) {
		fmt.Printf("Pegel: %d\n", p.Donations)

		if p.Donations > 10000 || p.Donations <= 0 {
			t.Fail()
		}

		c.MarkHandled()
	})

	c.OnError(func(r *colly.Response, e error) {
		t.Errorf("Colly error: %s", e)
	})
}
