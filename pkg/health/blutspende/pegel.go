package blutspende

import (
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	Url = "https://www.ukaachen.de/kliniken-institute/transfusionsmedizin-blutspendedienst/blutspendedienst/blutspendepegel/spendepegel/" // +"2020-08/"
)

type Pegel struct {
	Donations int
}

func FetchPegel(c *colly.Collector, cb func(p Pegel), errCb cfac.ErrorCallback) {
	FetchPegelTime(c, time.Now(), cb, errCb)
}

func FetchPegelTime(c *colly.Collector, t time.Time, cb func(p Pegel), errCb cfac.ErrorCallback) {

	c.OnHTML("div.drop-donations", func(h *colly.HTMLElement) {
		donations := h.Text
		donations = strings.ReplaceAll(donations, ".", "")
		dontationsCounter, err := strconv.Atoi(donations)

		if err != nil {
			errCb(err)
			return
		}

		cb(Pegel{
			Donations: dontationsCounter,
		})
	})

	c.Visit(Url + t.Format("2006-01"))
}
