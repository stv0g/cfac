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

type SearchCallback func(p []Place)

func Search(c *colly.Collector, query string, cb SearchCallback, ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var resp ReponseSearch
		if err := json.Unmarshal(r.Body, &resp); err != nil {
			cfac.DumpResponse(r)
			ecb(err)
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
