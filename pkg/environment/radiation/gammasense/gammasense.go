// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package gammasense

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	UrlApi         = "https://data.waag.org/api/gammasense"
	UrlApiStations = UrlApi + "/stations"
	UrlApiHourly   = UrlApi + "/hourly?sensor_id={sensor_id}&start={start}&end={end}" // TS: 2021-07-07T20:00:00.000Z
	UrlApiRecent   = UrlApi + "/recent"
)

type (
	StationCallback         func(Station)
	StationListCallback     func([]Station)
	MeasurementListCallback func([]Measurement)
)

func FetchStations(c *colly.Collector, cb StationListCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(func(r *colly.Response) {
		defer wg.Done()

		stations := []Station{}
		if err := json.Unmarshal(r.Body, &stations); err != nil {
			ecb(err)
			return
		}

		cb(stations)
	})

	c.Visit(UrlApiStations)

	return wg
}

func FetchRecent(c *colly.Collector, cb MeasurementListCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(onMeasurementListResponse(cb, ecb, wg))
	c.Visit(UrlApiRecent)

	return wg
}

func FetchHourly(sid string, start, end time.Time, c *colly.Collector, cb MeasurementListCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(onMeasurementListResponse(cb, ecb, wg))

	url := cfac.PrepareUrl(UrlApiHourly, cfac.UrlArgs{
		"sensor_id": sid,
		"start":     start.Format(time.RFC3339),
		"end":       end.Format(time.RFC3339),
	})

	c.Visit(url)

	return wg
}

func onMeasurementListResponse(cb MeasurementListCallback, ecb cfac.ErrorCallback, wg *sync.WaitGroup) colly.ResponseCallback {
	return func(r *colly.Response) {
		defer wg.Done()

		var resp []Measurement
		if err := json.Unmarshal(r.Body, &resp); err != nil {
			ecb(err)
			return
		}

		for i := range resp {
			resp[i].Time = resp[i].Time.UTC()
		}

		cb(resp)
	}
}

func FetchStationMeasurements(c *colly.Collector, cb StationCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchRecent(c, func(meas []Measurement) {
		FetchStations(c.Clone(), func(sts []Station) {
			for _, m := range meas {
				for _, s := range sts {
					if m.ID == s.ID {
						s.Measurement = &m
						cb(s)
					}
				}
			}
		}, ecb).Wait()
	}, ecb)
}
