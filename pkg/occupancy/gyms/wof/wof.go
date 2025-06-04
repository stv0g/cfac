// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package wof

import (
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	Url = "http://besucher.wof-fitness.de/"
)

type Callback func(studios Studio)

func FetchOccupancy(c *colly.Collector, cb Callback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnScraped(func(r *colly.Response) {
		wg.Done()
	})

	c.OnHTML("table[id=meineTabelle] tbody", func(e *colly.HTMLElement) {
		lastUpdated, err := cfac.LastUpdated(e.Response)
		if err != nil {
			ecb(err)
			return
		}

		e.ForEach("tr", func(i int, h *colly.HTMLElement) {
			fullName := h.ChildText("td:first-child")
			util := h.ChildText("td:last-child")

			occupancyStr := strings.TrimSuffix(util, "%")
			occupancy, err := strconv.Atoi(occupancyStr)
			if err != nil {
				ecb(err)
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

	return wg
}
