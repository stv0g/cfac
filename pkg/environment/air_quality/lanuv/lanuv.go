package lanuv

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"math"
	"strconv"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

// https://luadb.lds.nrw.de/LUA/hygon/pegel.php?rohdaten=ja
// https://www.lanuv.nrw.de/fileadmin/lanuv/luft/temes/heut/VACW.htm
// https://github.com/mpfeil/qualitySCHU
// https://github.com/mpfeil/lanuv-now
// https://github.com/Piccar/NRW-Wetter/blob/master/new1.py
// https://github.com/smatt/LanuvParser/blob/master/LanuvParser.java

const (
	UrlBase      = "https://www.lanuv.nrw.de/fileadmin/lanuv/luft/immissionen/aktluftqual"
	UrlCsv       = UrlBase + "/eu_luftqualitaet.csv"
	UrlCsvHeader = UrlBase + "/header_eu_luftqualitaet.csv"
)

type Callback func(AirQuality)

func Fetch(c *colly.Collector, cb Callback, ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		reader := bytes.NewReader(r.Body)
		csv_reader := csv.NewReader(reader)
		csv_reader.Comma = ';'

		for {
			// Read each record from csv
			record, err := csv_reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				ecb(err)
			}

			aq, err := ParseCols(record)
			if err != nil {
				ecb(err)
			}

			cb(aq)
		}

	})

	c.Visit(UrlCsv)
}

func ParseCols(row []string) (AirQuality, error) {
	var err error

	if len(row) < 6 {
		return AirQuality{}, errors.New("invalid number of columns")
	}

	aq := AirQuality{
		Station: row[0],
		ID:      row[1],
	}

	if aq.Ozon, err = ParseValue(row[2]); err != nil {
		return aq, err
	}

	if aq.SO2, err = ParseValue(row[3]); err != nil {
		return aq, err
	}

	if aq.NO2, err = ParseValue(row[4]); err != nil {
		return aq, err
	}

	if aq.PM10, err = ParseValue(row[5]); err != nil {
		return aq, err
	}

	return aq, nil
}

func ParseValue(s string) (float64, error) {
	if s == "*" {
		return math.NaN(), nil
	}

	if s == "<10" {
		return math.Inf(-1), nil
	}

	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return -1, err
	}

	return v, nil
}
