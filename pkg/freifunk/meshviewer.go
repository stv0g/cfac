// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package freifunk

import (
	"encoding/json"
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

// See https://data.aachen.freifunk.net/

const (
	Url = "https://data.aachen.freifunk.net/"

	UrlNodelist          = Url + "/nodelist.json"
	UrlNodes             = Url + "/nodes.json"
	UrlGraph             = Url + "/graph.json"
	UrlMeshviewer        = Url + "/meshviewer.json"
	UrlDistricStatistics = Url + "/ffac-district-statistics.json"
)

func FetchNodeList(c *colly.Collector, cb func(nl ResponseNodeList), ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(func(r *colly.Response) {
		defer wg.Done()

		var resp ResponseNodeList
		if err := json.Unmarshal(r.Body, &resp); err != nil {
			ecb(err)
			return
		}

		cb(resp)
	})

	c.Visit(UrlNodelist)

	return wg
}
