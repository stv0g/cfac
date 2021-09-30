package carolustherme

import (
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

const (
	Url = "https://www.carolus-thermen.de/en/thermalbath/#occupation"
)

type Occupancy struct {
	ThermalBath int
	Sauna       int
	Parking     int
	Timestamp   time.Time
}

type Callback func(o Occupancy)

func FetchUtilization(c *colly.Collector, cb Callback) {
	c.OnHTML("ul.occupancy", func(h *colly.HTMLElement) {

		thermalBath := h.ChildAttrs("div[data-name=Thermal Bath]", "data-percent")
		sauna := h.ChildAttrs("div[data-name=Sauna]", "data-percent")
		parking := h.ChildAttrs("div[data-name=Parking]", "data-percent")

		strconv.Atoi(thermalBat)

		cb(Occupancy{
			ThermalBath: 0,
		})

	})

	c.Visit(Url)
}
