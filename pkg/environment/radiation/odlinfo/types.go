// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package odlinfo

import (
	"strings"
	"time"
)

type ResponseStammInfo map[string]StationInfo

type Statistics struct {
	OperationalCount int `json:"betriebsbereit"`
	MeasurementAvg   struct {
		Value float64    `json:"mw"`
		Time  CustomDate `json:"t"`
	} `json:"mwavg"`
	MeasurementMaximum struct {
		Value float64 `json:"mw"` // in µSv/h
		ID    string  `json:"kenn"`
	} `json:"mwmax"`
	MeasurementMin struct {
		Value float64 `json:"mw"` // in µSv/h
		Kenn  string  `json:"kenn"`
	} `json:"mwmin"`
}

type StationInfo struct {
	Ort       string  `json:"ort"`
	ID        string  `json:"kenn"`
	PLZ       int     `json:"plz,string"`
	Status    int     `json:"status"`
	KID       int     `json:"kid"`
	Altitude  int     `json:"hoehe"`
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
	Value     float64 `json:"mw"` // in µSv/h
	IMIS      string  `json:"imis"`
}

type StationData struct {
	Info   StationInfo        `json:"stamm"`
	Avg1h  TimeseriesExtended `json:"mw1"`
	Avg24h Timeseries         `json:"mw24"`
}

type Measurement struct {
	Time        time.Time
	Value       float64
	CheckStatus bool
	RainProp    float64
}

type TimeseriesExtended struct {
	Time         []CustomTime `json:"t"`
	Value        []float64    `json:"mw"` // in µSv/h
	CheckStatus  []int        `json:"ps"`
	TimeRainProp []CustomTime `json:"tr"`
	RainProp     []float64    `json:"r"`
}

type Timeseries struct {
	Time  []CustomDate `json:"t"`
	Value []float64    `json:"mw"` // in µSv/h
}

type CustomDate struct {
	time.Time
}

type CustomTime struct {
	time.Time
}

func (c *CustomDate) UnmarshalJSON(b []byte) error {
	var err error
	s := strings.Trim(string(b), "\"")
	c.Time, err = time.Parse("2006-01-02", s)
	return err
}

func (c *CustomTime) UnmarshalJSON(b []byte) error {
	var err error
	s := strings.Trim(string(b), "\"")
	c.Time, err = time.Parse("2006-01-02 15:04", s)
	return err
}
