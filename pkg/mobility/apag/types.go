package apag

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
