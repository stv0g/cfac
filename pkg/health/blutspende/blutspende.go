package blutspende

import (
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	UrlApi              = "https://balu-api.ukaachen.de/v1"
	UrlApiLocations     = UrlApi + "/locations"
	UrlApiDonationDates = UrlApi + "/donationdates"
	UrlSpendePegelStat  = UrlApi + "/spendepegel-stat"
)

func FetchPegel(c *colly.Collector, cb func(p SpendePegelStats), ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(func(r *colly.Response) {
		defer wg.Done()

		var resp ResponseSpendePegelStats

		if err := json.Unmarshal(r.Body, &resp); err != nil {
			ecb(err)
			return
		}

		donations, err := strconv.Atoi(resp.Label)
		if err != nil {
			ecb(err)
			return
		}

		cb(SpendePegelStats{
			FillPercentage: resp.Fill,
			Donations:      donations,
			LastUpdated:    time.Now(),
		})
	})

	c.Visit(UrlSpendePegelStat)

	return wg
}
