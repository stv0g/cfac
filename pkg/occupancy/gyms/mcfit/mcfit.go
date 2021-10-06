package mcfit

// https://github.com/vaaski/openmagicline

import (
	"encoding/json"
	"net/url"
	"sort"
	"strconv"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	UrlStudios   = "https://rsg-group.api.magicline.com/connect/v1/studio" // + "?studioTags=AKTIV-391B8025C1714FB9B15BB02F2F8AC0B2"
	UrlStudios2  = "https://www.mcfit.com/typo3conf/ext/bra_studioprofiles_mcfitcom/Resources/Public/Json/studios_de.json?origLat=50.7753455&origLng=6.083886800000001&origAddress=aachen"
	UrlOccupancy = "https://www.mcfit.com/de/auslastung/antwort/request.json" // + "?tx_brastudioprofilesmcfitcom_brastudioprofiles%5BstudioId%5D=1536269110"
)

var (
	StudioIDs = []int{1536266890, 1536269110}
)

type ResponseStudios []Studio

type ResponseOccupancy struct {
	StartTime string      `json:"startTime"`
	EndTime   string      `json:"endTime"`
	Items     []Occupancy `json:"items"`
}

type StudioAddress struct {
	Street                       string      `json:"street"`
	SecondStreet                 interface{} `json:"secondStreet"`
	CityPart                     interface{} `json:"cityPart"`
	District                     interface{} `json:"district"`
	City                         string      `json:"city"`
	ZipCode                      string      `json:"zipCode"`
	StreetAddition               interface{} `json:"streetAddition"`
	HouseNumber                  string      `json:"houseNumber"`
	BuildingName                 interface{} `json:"buildingName"`
	CountryCode                  string      `json:"countryCode"`
	CountryCodeAlpha2            string      `json:"countryCodeAlpha2"`
	Longitude                    float64     `json:"longitude"`
	Latitude                     float64     `json:"latitude"`
	StreetType                   interface{} `json:"streetType"`
	StreetBlock                  interface{} `json:"streetBlock"`
	Portal                       interface{} `json:"portal"`
	Stairway                     interface{} `json:"stairway"`
	Door                         interface{} `json:"door"`
	Floor                        interface{} `json:"floor"`
	Province                     interface{} `json:"province"`
	AdditionalAddressInformation interface{} `json:"additionalAddressInformation"`
}

type StudioTag struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}

type OpeningHour struct {
	DayOfWeekFrom string `json:"dayOfWeekFrom"`
	DayOfWeekTo   string `json:"dayOfWeekTo"`
	TimeFrom      string `json:"timeFrom"`
	TimeTo        string `json:"timeTo"`
}

type Studio struct {
	ID                   int           `json:"id"`
	StudioName           string        `json:"studioName"`
	StudioPhone          string        `json:"studioPhone"`
	StudioEmail          string        `json:"studioEmail"`
	TrialSessionBookable bool          `json:"trialSessionBookable"`
	Address              StudioAddress `json:"address"`
	StudioTags           []StudioTag   `json:"studioTags"`
	MasterStudioID       interface{}   `json:"masterStudioId"`
	OpeningDate          string        `json:"openingDate"`
	ClosingDate          interface{}   `json:"closingDate"`
	OpeningHours         []OpeningHour `json:"openingHours"`
	ZoneID               string        `json:"zoneId"`
}

type Occupancy struct {
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	Percentage int    `json:"percentage"`
	Level      string `json:"level"`
	IsCurrent  bool   `json:"isCurrent"`
}

func (s *Studio) DistanceTo(to cfac.Coordinate) float64 {
	loc := cfac.Coordinate{
		Latitude:  s.Address.Latitude,
		Longitude: s.Address.Longitude,
	}

	return loc.DistanceTo(to)
}

func FetchStudios(c *colly.Collector, cb func(s []Studio), errCb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var resp ResponseStudios
		if err := json.Unmarshal(r.Body, &resp); err != nil {
			errCb(err)
			return
		}

		cb(resp)
	})

	c.Visit(UrlStudios)
}

func FetchStudiosByCoordinates(c *colly.Collector, co cfac.Coordinate, dist float64, cb func(s []Studio), errCb cfac.ErrorCallback) {
	FetchStudios(c, func(studios []Studio) {
		sort.Slice(studios, func(i, j int) bool {
			return studios[i].DistanceTo(co) < studios[j].DistanceTo(co)
		})

		if dist > 0 {
			filtStudios := []Studio{}
			for _, studio := range studios {
				if studio.DistanceTo(co) <= dist {
					filtStudios = append(filtStudios, studio)
				}
			}

			cb(filtStudios)
		} else {
			cb(studios)
		}
	}, errCb)
}

func FetchOccupancy(c *colly.Collector, studioID int, cb func(o ResponseOccupancy), errCb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var o ResponseOccupancy
		if err := json.Unmarshal(r.Body, &o); err != nil {
			errCb(err)
			return
		}

		cb(o)
	})

	q := url.Values{}

	q.Add("tx_brastudioprofilesmcfitcom_brastudioprofiles[studioId]", strconv.Itoa(studioID))

	c.Visit(UrlOccupancy + "?" + q.Encode())
}

func FetchCurrentOccupancy(c *colly.Collector, studioID int, cb func(o Occupancy), errCb cfac.ErrorCallback) {
	FetchOccupancy(c, studioID, func(ol ResponseOccupancy) {
		for _, o := range ol.Items {
			if o.IsCurrent {
				cb(o)
			}
		}
	}, errCb)
}
