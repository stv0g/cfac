package blutspende

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

const (
	Url = "https://www.ukaachen.de/kliniken-institute/transfusionsmedizin-blutspendedienst/blutspendedienst/blutspendepegel/spendepegel/" // +"2020-08/"
)

type Pegel struct {
	Donations int
}

type Callback func(p Pegel)

func FetchPegel(c *colly.Collector, cb Callback) {
	t := time.Now()
	FetchPegelTime(c, cb, t)
}

func FetchPegelTime(c *colly.Collector, cb Callback, t time.Time) {

	c.OnHTML("div.drop-donations", func(h *colly.HTMLElement) {
		donations := h.Text
		donations = strings.ReplaceAll(donations, ".", "")
		dontationsCounter, err := strconv.Atoi(donations)

		fmt.Printf("donations: %s\n", donations)

		if err != nil {
			return
		}

		cb(Pegel{
			Donations: dontationsCounter,
		})
	})

	c.Visit(fmt.Sprintf("%s/%04d-%02d", Url, t.Year(), t.Month()))
}
