// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package ford_carsharing

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

// curl 'https://www.ford-carsharing.de/de/rpc' --data-raw '{"method":"Search.byGeoPosition","params":[{"geoPos":{"radius":"7000","lat":50.790061643,"lng":6.060413354},"maxItems":"200","dateTimeStart":"2021-09-30 12:15:00","dateTimeEnd":"2021-09-30 13:15:00","address":"Aachen, Deutschland","vehicleTypeIds":[],"equipment":[]}],"id":1632996963387}'

const (
	UrlApi = "https://www.ford-carsharing.de/de/rpc"
)

func Fetch(coord cfac.Coordinate, radius int, c *colly.Collector, cb func(Station), ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var resp Response

		if err := json.Unmarshal(r.Body, &resp); err != nil {
			ecb(err)
			return
		}

		for _, s := range resp.Results {
			cb(s)
		}
	})

	end := time.Now()
	start := end.AddDate(0, 0, -1)

	req := Request{
		ID:     0,
		Method: "Search.byGeoPosition",
		Params: []RequestParameter{
			{
				Position: Position{
					Latitude:  coord.Latitude,
					Longitude: coord.Longitude,
					Radius:    strconv.Itoa(radius),
				},
				MaxItems:      strconv.Itoa(200),
				DateTimeStart: start.Format("2006-01-02 15:04:05"),
				DateTimeEnd:   end.Format("2006-01-02 15:04:05"),
				Address:       "Aachen, Deutschland",
			},
		},
	}

	reqRaw, err := json.Marshal(&req)
	if err != nil {
		ecb(err)
		return
	}

	c.PostRaw(UrlApi, reqRaw)
}
