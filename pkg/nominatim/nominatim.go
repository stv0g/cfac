package nominatim

import (
	"encoding/json"
	"net/url"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	Url       = "https://nominatim.openstreetmap.org"
	UrlSearch = Url + "/search"
)

type ReponseSearch []Place

type Place struct {
	PlaceID     int      `json:"place_id"`
	Licence     string   `json:"licence"`
	OsmType     string   `json:"osm_type"`
	OsmID       int      `json:"osm_id"`
	Boundingbox []string `json:"boundingbox"`
	Latitude    float64  `json:"lat,string"`
	Longitude   float64  `json:"lon,string"`
	DisplayName string   `json:"display_name"`
	Class       string   `json:"class"`
	Type        string   `json:"type"`
	Importance  float64  `json:"importance"`
	Icon        string   `json:"icon"`
	Address     Address  `json:"address"`
}

type Address struct {
	City        string `json:"city"`
	County      string `json:"county"`
	State       string `json:"state"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

type SearchCallback func(p []Place)

func Search(c *colly.Collector, query string, cb SearchCallback, errCb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var resp ReponseSearch
		if err := json.Unmarshal(r.Body, &resp); err != nil {
			cfac.DumpResponse(r)
			errCb(err)
			return
		}

		cb(resp)
	})

	q := url.Values{}
	q.Add("q", query)
	q.Add("addressdetails", "1")
	q.Add("format", "json")

	c.Visit(UrlSearch + "?" + q.Encode())
}
