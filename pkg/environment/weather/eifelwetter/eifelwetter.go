// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package eifelwetter

import (
	"encoding/json"
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	UrlApi               = "https://www.eifelwetter.de/api/v1"
	UrlApiPublic         = UrlApi + "/public"
	UrlApiCurrentComment = UrlApiPublic + "/comment/current"
	UrlApiStationList    = UrlApiPublic + "/stations"
	UrlApiStationData    = UrlApiPublic + "/station/{station_id}/data/latest"
)

type (
	StationListCallback func([]StationInfo)
	StationCallback     func(Station)
)

func FetchStations(c *colly.Collector, cb StationListCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(func(r *colly.Response) {
		defer wg.Done()

		stations := []StationInfo{}

		if err := json.Unmarshal(r.Body, &stations); err != nil {
			ecb(err)
			return
		}

		cb(stations)
	})

	c.Visit(UrlApiStationList)

	return wg
}

func FetchStation(sid string, c *colly.Collector, cb StationCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(func(r *colly.Response) {
		defer wg.Done()

		var station Station

		if err := json.Unmarshal(r.Body, &station); err != nil {
			ecb(err)
			return
		}

		cb(station)
	})

	url := cfac.PrepareUrl(UrlApiStationData, cfac.UrlArgs{
		"station_id": sid,
	})

	c.Visit(url)

	return wg
}

func FetchAllStations(c *colly.Collector, cb StationCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchStations(c, func(stations []StationInfo) {
		for _, si := range stations {
			FetchStation(si.ID, c.Clone(), func(s Station) {
				info := &si
				s.StationInfo = info

				cb(s)
			}, ecb).Wait()
		}
	}, ecb)
}
