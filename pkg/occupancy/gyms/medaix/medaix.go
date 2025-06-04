// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package medaix

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	UrlVisitors = "https://www.medaix.de/counter/visitors.php"
	UrlCalendar = "https://portal.medaix.de/courseplan/frontend/fillCalendar"
)

type Callback func(v VisitorCounter)

func requestData(id int) map[string]string {
	return map[string]string{
		"id": strconv.Itoa(id),
	}
}

func FetchOccupancy(id int, c *colly.Collector, cb Callback, ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var counter VisitorCounter

		// Try to unmarshal to bool which indicates an invalid ID
		var b bool
		if err := json.Unmarshal(r.Body, &b); err == nil {
			ecb(fmt.Errorf("unknown studio: %d", id))
			return
		}

		if err := json.Unmarshal(r.Body, &counter); err != nil {
			ecb(err)
			return
		}

		cb(counter)
	})

	c.Post(UrlVisitors, requestData(id))
}
