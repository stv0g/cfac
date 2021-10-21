/*
The cccac package fetches the current opening state of club room
of the Chaos Communication Club in Aachen.

The public API only provides access to a boolean state.

See: https://wiki.aachen.ccc.de/doku.php?id=projekte:clubstatus
*/
package cccac

import (
	"encoding/json"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	UrlApi              = "https://status.aachen.ccc.de/api/v0"
	UrlApiCurrentStatus = UrlApi + "/status/current?public"
)

func FetchStatus(c *colly.Collector, cb func(sts Status), ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var resp ResponseCurrentStatus
		if err := json.Unmarshal(r.Body, &resp); err != nil {
			ecb(err)
			return
		}

		cb(resp.Changed)
	})

	c.Visit(UrlApiCurrentStatus)
}
