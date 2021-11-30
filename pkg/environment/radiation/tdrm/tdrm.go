package tdrm

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	Url = "https://tdrm2.fiff.de/index.php"
)

func Fetch(c *colly.Collector, cb func(Station), ecb cfac.ErrorCallback) {
	d := c.Clone()
	d.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	})

	d.OnHTML("table[id=TDRM_TAB]", func(h *colly.HTMLElement) {
		region := ""

		h.ForEach("tr[class^=TDRM_TAB_Stripe]", func(i int, h *colly.HTMLElement) {
			tdRegion := h.ChildText("td:nth-child(1) a")
			tdStation := h.ChildText("td:nth-child(2) a")
			tdAverage := h.ChildText("td:nth-child(3) a")

			stationURL, _ := url.Parse(h.ChildAttr("td:nth-child(2) a", "href"))
			stationURLQuery := stationURL.Query()

			valueURL, _ := url.Parse(h.ChildAttr("td:nth-child(3) a", "href"))
			valueURLQuery := valueURL.Query()

			lat, _ := strconv.ParseFloat(stationURLQuery["lat"][0], 32)
			lon, _ := strconv.ParseFloat(stationURLQuery["lon"][0], 32)

			id, _ := strconv.Atoi(valueURLQuery["stationid"][0])

			value, _ := strconv.ParseFloat(tdAverage, 32)

			station := Station{
				ID:   id,
				Name: tdStation,
				Coordinate: cfac.Coordinate{
					Latitude:  lat,
					Longitude: lon,
				},
				AvgValue: value,
			}

			if tdRegion != "" {
				station.Region = tdRegion
				region = tdRegion
			} else {
				station.Region = region
			}

			cb(station)
		})
	})

	d.Visit(Url)
}
