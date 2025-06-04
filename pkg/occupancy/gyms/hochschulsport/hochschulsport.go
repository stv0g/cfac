// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

//go:build gosseract
// +build gosseract

package hochschulsport

import (
	"fmt"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/otiai10/gosseract/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	Url = "https://buchung.hsz.rwth-aachen.de/cgi/studio.cgi?size=30"
	UrlReferer
)

var (
	requiredHeaders = map[string]string{
		"Referer": UrlReferer,
		"Accept":  "image/*",
	}
	RWTHGym = cfac.Object{
		Name: "RWTHgym",
		Location: &cfac.Coordinate{
			Latitude:  50.779244,
			Longitude: 6.068629,
		},
	}
)

type Callback func(u Occupancy)

func FetchOccupancy(c *colly.Collector, cb Callback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(func(r *colly.Response) {
		client := gosseract.NewClient()
		defer client.Close()
		defer wg.Done()

		if err := client.SetWhitelist("0123456789"); err != nil {
			ecb(err)
			return
		}

		if err := client.SetImageFromBytes(r.Body); err != nil {
			ecb(err)
			return
		}

		text, err := client.Text()
		if err != nil {
			ecb(err)
			return
		}

		fmt.Println(text)

		utilization, err := strconv.ParseUint(text, 10, 64)
		if err != nil {
			ecb(err)
			return
		}

		lastUpdatedStr := r.Headers.Get("Date")
		lastUpdated, err := time.Parse(time.RFC1123, lastUpdatedStr)
		if err != nil {
			ecb(err)
			return
		}

		cb(Occupancy{
			Occupancy:   float64(utilization),
			LastUpdated: lastUpdated,
		})
	})

	url, err := url.Parse(Url)
	if err != nil {
		ecb(err)
		return wg
	}

	c.OnRequest(func(r *colly.Request) {
		if r.URL.Host == url.Host {
			for key, value := range requiredHeaders {
				r.Headers.Set(key, value)
			}
		}
	})

	c.Visit(Url)

	return wg
}
