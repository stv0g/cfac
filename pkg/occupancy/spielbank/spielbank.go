package spielbank

import (
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	Url = "https://www.spielbank-aachen.de/"
)

type Callback func(u Occupancy)

func FetchOccupancy(c *colly.Collector, cb Callback, errCb cfac.ErrorCallback) {
	c.OnHTML("div.metric", func(h *colly.HTMLElement) {
		ratio, err := strconv.ParseFloat(h.Attr("data-ratio"), 32)
		if err != nil {
			return
		}

		timeStr := h.ChildText("text.date")

		loc, err := time.LoadLocation("Europe/Berlin")
		if err != nil {
			errCb(err)
			return
		}

		updated, err := time.ParseInLocation("Stand: 02.01.2006, 15:04 Uhr", timeStr, loc)
		if err != nil {
			errCb(err)
			return
		}

		cb(Occupancy{
			Utilization: ratio,
			LastUpdated: updated,
		})
	})

	c.Visit(Url)
}
