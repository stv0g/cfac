package gammasense

import (
	"encoding/json"
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

type StationListCallback func([]Station)

type MeasurementCallback func(Measurement)

func FetchStations(c *colly.Collector, cb StationListCallback, ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		stations := []Station{}

		if err := json.Unmarshal(r.Body, &stations); err != nil {
			ecb(err)
			return
		}

		cb(stations)
	})

	c.Visit(UrlApiStations)
}

func FetchRecent(c *colly.Collector, cb MeasurementCallback, ecb cfac.ErrorCallback) {
	c.OnResponse(onMeasurementListResponse(cb, ecb))

	c.Visit(UrlApiRecent)
}

func FetchHourly(sid string, start, end time.Time, c *colly.Collector, cb MeasurementCallback, ecb cfac.ErrorCallback) {
	c.OnResponse(onMeasurementListResponse(cb, ecb))

	url := cfac.PrepareUrl(UrlApiHourly, cfac.UrlArgs{
		"sensor_id": sid,
		"start":     start.Format(time.RFC3339),
		"end":       end.Format(time.RFC3339),
	})

	c.Visit(url)
}

func onMeasurementListResponse(cb MeasurementCallback, ecb cfac.ErrorCallback) colly.ResponseCallback {
	return func(r *colly.Response) {
		var resp []Measurement

		if err := json.Unmarshal(r.Body, &resp); err != nil {
			ecb(err)
			return
		}

		for _, m := range resp {
			cb(m)
		}
	}
}
