package spielbank

import (
	"regexp"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

const (
	Url = "https://www.spielbank-aachen.de/"
)

type Utilization struct {
	Utilization int
	Timestamp   time.Time
}

type Callback func(u Utilization)

func FetchUtilization(c *colly.Collector, cb Callback) {
	c.OnHTML("div.metric", func(h *colly.HTMLElement) {
		ratio, err := strconv.ParseFloat(h.Attr("data-ratio"), 32)
		if err != nil {
			return
		}

		time := h.ChildText("text.date")
		re := regexp.MustCompile("Stand: (\d+)\.(\d+)\.(\d+), (\d+):(\d+) Uhr")

		matches := re.FindAll(time)
	})

	c.Visit(Url)
}
