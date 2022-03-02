package lanuv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
	"golang.org/x/text/encoding/charmap"
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

func Fetch(c *colly.Collector, cb Callback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(func(r *colly.Response) {
		defer wg.Done()

		isoDecoder := charmap.ISO8859_1.NewDecoder()
		utf8Body, _ := isoDecoder.Bytes(r.Body)

		reader := bytes.NewReader(utf8Body)
		csv_reader := csv.NewReader(reader)
		csv_reader.Comma = ';'
		csv_reader.FieldsPerRecord = 2

		var ts time.Time

		row := 0
		for {
			row++

			// Read each record from csv
			record, err := csv_reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				ecb(err)
			}

			switch {
			case row == 1: // Timestamp: Date;Time
				t := fmt.Sprintf("%s-%s", record[0], record[1])
				ts, err = time.Parse("02.01.2006-15:04", t)
				if err != nil {
					ecb(fmt.Errorf("failed to parse time '%s': %w", t, err))
					return
				}
			case row == 2: // Header
				// ignore
				csv_reader.FieldsPerRecord = 7
			case row > 2:
				aq, err := parseCols(record)
				if err != nil {
					ecb(err)
				}

				aq.Timestamp = ts

				cb(aq)
			}
		}

	})

	c.Visit(UrlCsv)

	return wg
}

func parseCols(row []string) (AirQuality, error) {
	var err error

	if len(row) < 6 {
		return AirQuality{}, fmt.Errorf("invalid number of columns: %s", strings.Join(row, ";"))
	}

	aq := AirQuality{
		Station: row[0],
		ID:      row[1],
	}

	if aq.Ozon, err = parseValue(row[2]); err != nil {
		return aq, err
	}

	if aq.SO2, err = parseValue(row[3]); err != nil {
		return aq, err
	}

	if aq.NO2, err = parseValue(row[4]); err != nil {
		return aq, err
	}

	if aq.PM10, err = parseValue(row[5]); err != nil {
		return aq, err
	}

	return aq, nil
}

func parseValue(s string) (float64, error) {
	if s == "*" || s == "-" {
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
