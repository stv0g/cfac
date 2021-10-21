package fitx

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

//

const (
	Url                     = "https://www.fitx.de"
	UrlWorkload             = Url + "/fitnessstudio/{studio_id}/workload"
	UrlStudiosByCoordinates = Url + "/studios/by-coordinates"

	UrlServices     = "https://services.fitx.de"
	UrlStudioPlan   = UrlServices + "/survey/nps_courseplan_display"
	UrlStudioDetail = UrlServices + "/survey/nps_studiodetail_v2"
)

var (
	StudioIDs = []int{
		38, // Aachen Europaplatz
	}
)

func FetchStudios(c *colly.Collector, coords cfac.Coordinate, cb func(s []Studio), ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var resp struct {
			Results []string `json:"result"`
		}
		if err := json.Unmarshal(r.Body, &resp); err != nil {
			cfac.DumpResponse(r)
			ecb(err)
			return
		}

		studios := []Studio{}
		for _, result := range resp.Results {
			var studio Studio
			if err := json.Unmarshal([]byte(result), &studio); err != nil {
				ecb(err)
				continue
			}

			studios = append(studios, studio)
		}

		cb(studios)
	})

	c.Post(UrlStudiosByCoordinates, map[string]string{
		"lat": strconv.FormatFloat(coords.Latitude, 'f', 6, 64),
		"lon": strconv.FormatFloat(coords.Longitude, 'f', 6, 64),
	})
}

func FetchStudioWorkload(c *colly.Collector, studioID int, cb func(s Studio), ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var strings []string
		var studio Studio

		if err := json.Unmarshal(r.Body, &strings); err != nil {
			ecb(err)
			return
		}

		if len(strings) != 1 {
			ecb(errors.New("malformed response"))
			return
		}

		if err := json.Unmarshal([]byte(strings[0]), &studio); err != nil {
			ecb(err)
			return
		}

		cb(studio)
	})

	c.Visit(cfac.PrepareUrl(UrlWorkload, cfac.UrlArgs{
		"studio_id": strconv.Itoa(studioID),
	}))
}
