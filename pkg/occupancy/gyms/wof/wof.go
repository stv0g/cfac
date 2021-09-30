package wof

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

const (
	Url = "http://besucher.wof-fitness.de/"
)

type Studio struct {
	Name        string
	Location    string
	Utilization int
}

type Callback func(studios []Studio)

func FetchUtilization(c *colly.Collector, cb Callback) {

	c.OnHTML("table[id=meineTabelle] tbody", func(e *colly.HTMLElement) {
		studios := []Studio{}

		e.ForEach("tr", func(i int, h *colly.HTMLElement) {
			fullName := h.ChildText("td:first-child")
			util := h.ChildText("td:last-child")

			utilPercent := strings.TrimSuffix(util, "%")
			utilPercentNumber, err := strconv.Atoi(utilPercent)
			if err != nil {
				return
			}

			name := strings.Split(fullName, " - ")

			studios = append(studios, Studio{
				Name:        name[0],
				Location:    name[1],
				Utilization: utilPercentNumber,
			})
		})

		cb(studios)
	})

	c.Visit(Url)
}
