package apag

import (
	"strings"
	"time"
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

	Stats *Stats
}

type Stats struct {
	Ident   string     `json:"ident"`
	Date    CustomTime `json:"date"`
	Max     uint       `json:"max,string"`
	Percent float32    `json:"percent,string"`
	Count   uint       `json:"count,string"`
	Free    uint       `json:"free,string"`
	Full    int        `json:"full,string"`
	Trend   string     `json:"trend"`
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
