package apag

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	Url        = "https://www.apag.de"
	UrlApi     = Url + "/qad_restAPI.php"
	UrlChartXS = UrlApi + "?q={ident}&d=chartxs&max={max}"
)

type House struct {
	NID       uint    `json:"nid,string"`
	Ident     string  `json:"ident"`
	Url       string  `json:"url"`
	Title     string  `json:"title"`
	Latitude  float32 `json:"lat,string"`
	Longitude float32 `json:"lon,string"`
	Capacity  uint    `json:"capacity,string"`
	Open      string  `json:"open"`
	Height    string  `json:"height"`
	Type      string  `json:"type"`

	Stats *HouseStats
}

type CustomTime struct {
	time.Time
}

func (c *CustomTime) UnmarshalJSON(b []byte) error {
	var err error
	loc, _ := time.LoadLocation("Europe/Berlin")
	s := strings.Trim(string(b), "\"")
	c.Time, err = time.ParseInLocation("2006-01-02 15:04:05", s, loc)
	return err
}

type HouseStats struct {
	Ident   string     `json:"ident"`
	Date    CustomTime `json:"date"`
	Max     uint       `json:"max,string"`
	Percent float32    `json:"percent,string"`
	Count   uint       `json:"count,string"`
	Free    uint       `json:"free,string"`
	Full    int        `json:"full,string"`
	Trend   string     `json:"trend"`
}

func FetchAllHouses(c *colly.Collector, cb func([]House), errCb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var re = regexp.MustCompile(`(?m)var houses = (.*);$`)

		m := re.FindSubmatch(r.Body)
		if m == nil {
			errCb(errors.New("failed to find house list"))
			return
		}

		j := m[1]

		// Remove spaces in numbers
		re = regexp.MustCompile(`(\d+) (\d+)`)
		j = re.ReplaceAll(j, []byte("$1$2"))

		var houseMap map[string]House
		if err := json.Unmarshal(j, &houseMap); err != nil {
			errCb(err)
			return
		}

		var houses = []House{}
		for _, h := range houseMap {
			h.Stats = nil
			houses = append(houses, h)
		}

		cb(houses)
	})

	c.Visit(Url)
}

func FetchAllHouseStats(c *colly.Collector, cb func([]HouseStats), errCb cfac.ErrorCallback) {
	FetchHouseStats(c, "all", cb, errCb)
}

func FetchHouseStats(c *colly.Collector, ident string, cb func([]HouseStats), errCb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var stats []HouseStats
		if err := json.Unmarshal(r.Body, &stats); err != nil {
			errCb(err)
			return
		}

		cb(stats)
	})

	if ident == "" {
		ident = "all"
	}

	q := url.Values{}
	q.Add("q", ident)
	q.Add("d", "fulldata")
	q.Add("f", "json")

	c.Visit(UrlApi + "?" + q.Encode())
}

func FetchAllHousesWithStats(c *colly.Collector, cb func([]House), errCb cfac.ErrorCallback) {
	FetchAllHouses(c, func(houses []House) {
		FetchAllHouseStats(c, func(houseStats []HouseStats) {
			houseMap := map[string]*House{}
			for _, h := range houses {
				houseMap[h.Ident] = &h
			}

			for _, hs := range houseStats {
				if h, ok := houseMap[hs.Ident]; ok {
					h.Stats = &hs
				} else {
					log.Printf("failed to find matching house for ident: %s", hs.Ident)
				}
			}

			cb(houses)
		}, errCb)
	}, errCb)
}

func (h *House) FetchStats(c *colly.Collector, cb func([]HouseStats), errCb cfac.ErrorCallback) {
	FetchHouseStats(c, h.Ident, cb, errCb)
}
