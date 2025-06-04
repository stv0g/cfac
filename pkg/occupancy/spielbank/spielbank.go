// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package spielbank

import (
	"strconv"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	Url = "https://www.spielbank-aachen.de/"
)

var SpielbankAachen = &cfac.Object{
	Name: "Spielbank Aachen",
}

type Callback func(u Occupancy)

func FetchOccupancy(c *colly.Collector, cb Callback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnScraped(func(r *colly.Response) {
		wg.Done()
	})

	c.OnHTML("div.metric", func(h *colly.HTMLElement) {
		ratio, err := strconv.ParseFloat(h.Attr("data-ratio"), 32)
		if err != nil {
			return
		}

		timeStr := h.ChildText("text.date")

		loc, err := time.LoadLocation("Europe/Berlin")
		if err != nil {
			ecb(err)
			return
		}

		updated, err := time.ParseInLocation("Stand: 02.01.2006, 15:04 Uhr", timeStr, loc)
		if err != nil {
			ecb(err)
			return
		}

		cb(Occupancy{
			Utilization: ratio,
			LastUpdated: updated,
		})
	})

	c.Visit(Url)

	return wg
}
