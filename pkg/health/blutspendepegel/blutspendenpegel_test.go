package blutspendepegel_test

import (
	"fmt"
	"testing"

	"github.com/gocolly/colly/v2"
	"github.com/stv0g/cfac/pkg/blutspendepegel"
)

func TestFetchPegel(t *testing.T) {
	c := colly.NewCollector()

	blutspendepegel.FetchPegel(c, func(p blutspendepegel.BlutspendePegel) {
		fmt.Printf("Pegel: %d\n", p.Donations)

		if p.Donations > 10000 || p.Donations <= 0 {
			t.Fail()
		}
	})

	c.OnError(func(r *colly.Response, e error) {
		t.Errorf("Colly error: %s", e)
	})

	c.Wait()
}
