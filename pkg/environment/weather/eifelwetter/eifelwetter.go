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

type StationListCallback func([]StationInfo)
type StationCallback func(Station)

func FetchStations(c *colly.Collector, cb StationListCallback, ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		stations := []StationInfo{}

		if err := json.Unmarshal(r.Body, &stations); err != nil {
			ecb(err)
		}

		cb(stations)
	})

	c.Visit(UrlApiStationList)
}

func FetchStation(sid string, c *colly.Collector, cb StationCallback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(func(r *colly.Response) {
		defer wg.Done()

		var station Station

		if err := json.Unmarshal(r.Body, &station); err != nil {
			ecb(err)
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
	wg := &sync.WaitGroup{}
	wg.Add(1)

	FetchStations(c, func(stations []StationInfo) {
		defer wg.Done()

		wg2 := &sync.WaitGroup{}
		wg2.Add(len(stations))

		for _, si := range stations {
			FetchStation(si.ID, c.Clone(), func(s Station) {
				defer wg2.Done()

				s.StationInfo = si

				cb(s)
			}, ecb)
		}

		wg2.Wait()
	}, ecb)

	return wg
}
