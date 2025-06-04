// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package buienalarm

import (
	"encoding/json"
	"errors"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	UrlApi = "https://cdn-secure.buienalarm.nl/api/3.4/forecast.php?lat={lat}&lon={lon}&region={region}&unit=mm/u"
)

type Callback func(Forecast)

func Fetch(coord cfac.Coordinate, c *colly.Collector, cb Callback, ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var fc Forecast

		if err := json.Unmarshal(r.Body, &fc); err != nil {
			ecb(err)
			return
		}

		if !fc.Success {
			ecb(errors.New("request failed"))
			return
		}

		cb(fc)
	})

	url := cfac.PrepareUrl(UrlApi, cfac.UrlArgs{
		"lat":    coord.Latitude,
		"lon":    coord.Longitude,
		"region": "de", // TODO: do we need to provide also other regions here?
	})

	c.Visit(url)
}
