package terminmanagement

import (
	"encoding/json"
	"net/url"
	"strconv"
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	Url = "https://terminmanagement.regioit-aachen.de/sr_aachen/mobile_srac_stva/get.php"
)

func Fetch(c *colly.Collector, cb func(r Response), ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(func(r *colly.Response) {
		defer wg.Done()

		var resp Response
		if err := json.Unmarshal(r.Body, &resp); err != nil {
			ecb(err)
			return
		}

		cb(resp)
	})

	q := url.Values{}
	q.Add("option", "values")
	q.Add("sid", "6")

	c.Visit(Url + "?" + q.Encode())

	return wg
}

func FetchWaitCircles(c *colly.Collector, cb func(wc WaitCircle), ecb cfac.ErrorCallback) *sync.WaitGroup {
	return Fetch(c, func(r Response) {
		for _, name := range r.WaitCircles {
			cnt, _ := strconv.Atoi(r.VisitorCount[name])
			wc := WaitCircle{
				Name:         name,
				WaitingTime:  r.WaitingTime[name].Duration,
				VisitorCount: cnt,
				LastRefresh:  r.LastRefresh.Time,
			}

			cb(wc)
		}
	}, ecb)
}
