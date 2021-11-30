package cambio

import (
	"encoding/json"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	CityID = "AAC"

	UrlApi         = "https://cwapi.cambio-carsharing.com/pub/{city}"
	UrlApiStations = UrlApi + "/stations"
	UrlApiVehicles = UrlApi + "/vehicles"
)

func FetchStations(city string, c *colly.Collector, cb func(Station), ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var stations []Station

		if err := json.Unmarshal(r.Body, &stations); err != nil {
			ecb(err)
			return
		}

		for _, s := range stations {
			cb(s)
		}
	})

	url := cfac.PrepareUrl(UrlApiStations, cfac.UrlArgs{
		"city": city,
	})

	c.Visit(url)
}

func FetchVehicles(city string, c *colly.Collector, cb func(Vehicle), ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var vehicles []Vehicle

		if err := json.Unmarshal(r.Body, &vehicles); err != nil {
			ecb(err)
			return
		}

		for _, v := range vehicles {
			cb(v)
		}
	})

	url := cfac.PrepareUrl(UrlApiVehicles, cfac.UrlArgs{
		"city": city,
	})

	c.Visit(url)
}
