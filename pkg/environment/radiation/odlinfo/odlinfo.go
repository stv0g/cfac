package odlinfo

import (
	"encoding/json"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	UrlApi            = "https://odlinfo.bfs.de/json"
	UrlApiStatistics  = UrlApi + "/stat.json"
	UrlApiStationList = UrlApi + "/stamm.json"
	UrlApiStationData = UrlApi + "/{kenn}.json"
)

type Callback func(Statistics)

func FetchStatistics(c *colly.Collector, cb func(Statistics), ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var stat Statistics

		if err := json.Unmarshal(r.Body, &stat); err != nil {
			ecb(err)
			return
		}

		cb(stat)
	})

	c.Visit(UrlApiStatistics)
}

func FetchStations(c *colly.Collector, cb func(StationInfo), ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var stamm ResponseStammInfo

		if err := json.Unmarshal(r.Body, &stamm); err != nil {
			ecb(err)
			return
		}

		for id, s := range stamm {
			s.ID = id
			cb(s)
		}
	})

	c.Visit(UrlApiStationList)
}

func FetchStationData(kenn string, c *colly.Collector, cb func(StationData), ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var data StationData

		if err := json.Unmarshal(r.Body, &data); err != nil {
			ecb(err)
			return
		}

		cb(data)
	})

	url := cfac.PrepareUrl(UrlApiStationData, cfac.UrlArgs{
		"kenn": kenn,
	})

	c.Visit(url)
}
