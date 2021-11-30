package sensor_community

import (
	"encoding/json"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	UrlDataApi     = "https://data.sensor.community"
	UrlDataApiData = UrlDataApi + "/static/v2/data.json"

	UrlMapsApi       = "https://maps.sensor.community"
	UrlMapsApiData24 = UrlMapsApi + "/data/v2/data.24h.json"
	UrlMapsApiNoise  = UrlMapsApi + "/data/v1/data.noise.json"
	UrlMapsApiTemp   = UrlMapsApi + "/data/v1/data.temp.json"
	UrlMapsApiDust   = UrlMapsApi + "/data/v1/data.dust.json"
	UrlMapsApiLabs   = UrlMapsApi + "/local-labs/labs.json"
)

func FetchLabs(coord cfac.Coordinate, c *colly.Collector, cb func(Lab), ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var labs []Lab

		if err := json.Unmarshal(r.Body, &labs); err != nil {
			ecb(err)
			return
		}

		for _, l := range labs {
			cb(l)
		}
	})

	c.Visit(UrlMapsApiLabs)
}

func FetchDataURL(url string, c *colly.Collector, cb func(Sensor), ecb cfac.ErrorCallback) {
	d := c.Clone()
	d.MaxBodySize = 100 << 20 // 100 MiB

	d.OnResponse(func(r *colly.Response) {
		var sensors []Sensor

		if err := json.Unmarshal(r.Body, &sensors); err != nil {
			ecb(err)
			return
		}

		for _, s := range sensors {
			cb(s)
		}
	})

	d.Visit(url)
}

func FetchData(c *colly.Collector, cb func(Sensor), ecb cfac.ErrorCallback) {
	FetchDataURL(UrlDataApiData, c, cb, ecb)
}

func FetchDataRadius(center cfac.Coordinate, radius float64, c *colly.Collector, cb func(Sensor), ecb cfac.ErrorCallback) {
	FetchData(c, func(s Sensor) {
		c := cfac.Coordinate{
			Latitude:  s.Location.Latitude,
			Longitude: s.Location.Longitude,
		}

		if c.DistanceTo(center) < radius {
			cb(s)
		}
	}, ecb)
}

func FetchNoise(c *colly.Collector, cb func(Sensor), ecb cfac.ErrorCallback) {
	FetchDataURL(UrlMapsApiNoise, c, cb, ecb)
}

func FetchDust(c *colly.Collector, cb func(Sensor), ecb cfac.ErrorCallback) {
	FetchDataURL(UrlMapsApiDust, c, cb, ecb)
}

func FetchTemp(c *colly.Collector, cb func(Sensor), ecb cfac.ErrorCallback) {
	FetchDataURL(UrlMapsApiTemp, c, cb, ecb)
}
