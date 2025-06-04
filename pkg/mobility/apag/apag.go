// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package apag

import (
	"encoding/json"
	"errors"
	"net/url"
	"regexp"
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	Url        = "https://www.apag.de"
	UrlApi     = Url + "/qad_restAPI.php"
	UrlChartXS = UrlApi + "?q={ident}&d=chartxs&max={max}"
)

func FetchAllHouses(c *colly.Collector, cb func([]House), ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(func(r *colly.Response) {
		defer wg.Done()

		re := regexp.MustCompile(`(?m)var houses = (.*);$`)

		m := re.FindSubmatch(r.Body)
		if m == nil {
			ecb(errors.New("failed to find house list"))
			return
		}

		j := m[1]

		// Remove spaces in numbers
		re = regexp.MustCompile(`(\d+) (\d+)`)
		j = re.ReplaceAll(j, []byte("$1$2"))

		var houseMap map[string]House
		if err := json.Unmarshal(j, &houseMap); err != nil {
			ecb(err)
			return
		}

		houses := []House{}
		for _, h := range houseMap {
			h.Stats = nil
			houses = append(houses, h)
		}

		cb(houses)
	})

	c.Visit(Url)

	return wg
}

func FetchAllHouseStats(c *colly.Collector, cb func([]Stats), ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(func(r *colly.Response) {
		defer wg.Done()

		var stats []Stats
		if err := json.Unmarshal(r.Body, &stats); err != nil {
			ecb(err)
			return
		}

		cb(stats)
	})

	q := url.Values{}
	q.Add("q", "all")
	q.Add("d", "fulldata")
	q.Add("f", "json")

	c.Visit(UrlApi + "?" + q.Encode())

	return wg
}

func FetchHouseStats(c *colly.Collector, ident string, cb func(Stats), ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(func(r *colly.Response) {
		defer wg.Done()

		var stats Stats
		if err := json.Unmarshal(r.Body, &stats); err != nil {
			ecb(err)
			return
		}

		cb(stats)
	})

	q := url.Values{}
	q.Add("q", ident)
	q.Add("d", "fulldata")
	q.Add("f", "json")

	c.Visit(UrlApi + "?" + q.Encode())

	return wg
}

func FetchHousesWithStats(c *colly.Collector, cb func(House), ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchAllHouseStats(c.Clone(), func(stats []Stats) {
		FetchAllHouses(c, func(houses []House) {
			for _, s := range stats {
				for _, h := range houses {
					if s.Ident == h.Ident {
						h.Stats = &s
						cb(h)
					}
				}
			}
		}, ecb)
	}, ecb)
}

func (h *House) FetchStats(c *colly.Collector, cb func(Stats), ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchHouseStats(c, h.Ident, cb, ecb)
}
