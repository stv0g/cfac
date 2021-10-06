package wof

import (
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	Url = "http://besucher.wof-fitness.de/"
)

type Studio struct {
	Name        string
	Location    string
	Occupancy   cfac.Percent
	LastUpdated time.Time
}

type Callback func(studios Studio)

func FetchOccupancy(c *colly.Collector, cb Callback, errCb cfac.ErrorCallback) {

	c.OnHTML("table[id=meineTabelle] tbody", func(e *colly.HTMLElement) {
		lastUpdated, err := cfac.LastUpdated(e.Response)
		if err != nil {
			errCb(err)
			return
		}

		e.ForEach("tr", func(i int, h *colly.HTMLElement) {
			fullName := h.ChildText("td:first-child")
			util := h.ChildText("td:last-child")

			occupancyStr := strings.TrimSuffix(util, "%")
			occupancy, err := strconv.Atoi(occupancyStr)
			if err != nil {
				errCb(err)
				return
			}

			name := strings.Split(fullName, " - ")

			cb(Studio{
				Name:        name[0],
				Location:    name[1],
				Occupancy:   cfac.Percent(occupancy),
				LastUpdated: lastUpdated,
			})
		})

	})

	c.Visit(Url)
}
