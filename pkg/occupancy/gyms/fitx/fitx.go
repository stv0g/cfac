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

type StudioLocation struct {
	Street     string  `json:"street"`
	Zip        string  `json:"zip"`
	Place      string  `json:"place"`
	Region     string  `json:"region"`
	Phone      string  `json:"phone"`
	Directions string  `json:"directions"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	PlaceID    string  `json:"placeId"`
}

type StudioTour struct {
	Enabled  bool   `json:"enabled"`
	URL      string `json:"url"`
	EmbedURL string `json:"embedUrl"`
}

type StudioDates struct {
	Opening                 int         `json:"opening"`
	OpeningRegular          int         `json:"openingRegular"`
	OpeningCourses          int         `json:"openingCourses"`
	TemporarilyClosedReason interface{} `json:"temporarilyClosedReason"`
	TemporarilyClosedFrom   interface{} `json:"temporarilyClosedFrom"`
	TemporarilyClosedTo     interface{} `json:"temporarilyClosedTo"`
	Presale                 interface{} `json:"presale"`
}

type StudioWorkload struct {
	Term       string `json:"term"`
	Percentage int    `json:"percentage"`
}

type Studio struct {
	Distance                               interface{}    `json:"distance"`
	Name                                   string         `json:"name"`
	Alias                                  string         `json:"alias"`
	Identifier                             string         `json:"identifier"`
	BranchID                               int            `json:"branchId"`
	Location                               StudioLocation `json:"location"`
	VirtualTour                            StudioTour     `json:"virtualTour"`
	Notice                                 interface{}    `json:"notice"`
	Dates                                  StudioDates    `json:"dates"`
	Badge                                  interface{}    `json:"badge"`
	Promotions                             []string       `json:"promotions"`
	Status                                 int            `json:"status"`
	DisablePreregistration                 interface{}    `json:"disablePreregistration"`
	VisitorData                            [][]int        `json:"visitorData"`
	TrialworkoutEnabled                    bool           `json:"trialworkoutEnabled"`
	TrialworkoutRestricted                 bool           `json:"trialworkoutRestricted"`
	PreregistrationAgeRestricted           bool           `json:"preregistrationAgeRestricted"`
	PreregistrationAccountHolderRestricted bool           `json:"preregistrationAccountHolderRestricted"`
	Images                                 []string       `json:"images"`
	ListText                               string         `json:"listText"`
	ListTextColor                          interface{}    `json:"listTextColor"`
	GoogleRating                           float64        `json:"googleRating"`
	GoogleRatingCount                      int            `json:"googleRatingCount"`
	Workload                               StudioWorkload `json:"workload"`
	// NewsletterForm                         string         `json:"newsletterForm"`
}

func FetchStudios(c *colly.Collector, coords cfac.Coordinate, cb func(s []Studio), errCb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var resp struct {
			Results []string `json:"result"`
		}
		if err := json.Unmarshal(r.Body, &resp); err != nil {
			cfac.DumpResponse(r)
			errCb(err)
			return
		}

		studios := []Studio{}
		for _, result := range resp.Results {
			var studio Studio
			if err := json.Unmarshal([]byte(result), &studio); err != nil {
				errCb(err)
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

func FetchStudioWorkload(c *colly.Collector, studioID int, cb func(s Studio), errCb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var strings []string
		var studio Studio

		if err := json.Unmarshal(r.Body, &strings); err != nil {
			errCb(err)
			return
		}

		if len(strings) != 1 {
			errCb(errors.New("malformed response"))
			return
		}

		if err := json.Unmarshal([]byte(strings[0]), &studio); err != nil {
			errCb(err)
			return
		}

		cb(studio)
	})

	c.Visit(cfac.PrepareUrl(UrlWorkload, cfac.UrlArgs{
		"studio_id": strconv.Itoa(studioID),
	}))
}
