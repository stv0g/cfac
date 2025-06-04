// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package nominatim

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	Url       = "https://nominatim.openstreetmap.org"
	UrlSearch = Url + "/search"
)

type SearchCallback func(p []Place)

func Search(c *colly.Collector, query string, cb SearchCallback, ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var resp ReponseSearch
		if err := json.Unmarshal(r.Body, &resp); err != nil {
			cfac.DumpResponse(r)
			ecb(err)
			return
		}

		cb(resp)
	})

	q := url.Values{}
	q.Add("q", query)
	q.Add("addressdetails", "1")
	q.Add("format", "json")

	c.Visit(UrlSearch + "?" + q.Encode())
}

var cache = map[string]cfac.Coordinate{}

func SearchAndCache(query string) (cfac.Coordinate, error) {
	var err error = nil
	var coord cfac.Coordinate

	if coord, ok := cache[query]; ok {
		return coord, nil
	}

	c := colly.NewCollector()

	wg := sync.WaitGroup{}
	wg.Add(1)

	Search(c, query, func(p []Place) {
		if len(p) != 1 {
			err = fmt.Errorf("geo location found %d results", len(p))
		} else {
			coord = p[0].Coordinate()
		}

		wg.Done()
	}, func(e error) {
		err = e
		wg.Done()
	})

	wg.Wait()

	return coord, err
}
