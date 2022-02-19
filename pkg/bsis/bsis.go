package bsis

import (
	"encoding/json"
	"net/url"
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	UrlApi = "https://bsis.aachen.de/geoserver/ows"
)

type Callback func(*ConstructionSiteList)

func FetchConstructionSites(c *colly.Collector, cb Callback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnScraped(func(r *colly.Response) {
		wg.Done()
	})

	c.OnResponse(func(r *colly.Response) {
		var resp ConstructionSiteList

		if err := json.Unmarshal(r.Body, &resp); err != nil {
			cfac.DumpResponse(r)
			ecb(err)
			return
		}

		cb(&resp)
	})

	q := url.Values{}
	q.Add("service", "WFS")
	q.Add("request", "GetFeature")
	q.Add("typeName", "BSIS:PUNKTE_ALLE")
	q.Add("version", "1.1.0")
	q.Add("outputFormat", "application/json")
	q.Add("srsName", "EPSG:3857")

	c.Visit(UrlApi + "?" + q.Encode())

	return wg
}
