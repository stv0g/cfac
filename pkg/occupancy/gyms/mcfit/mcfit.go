// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package mcfit

// https://github.com/vaaski/openmagicline

import (
	"encoding/json"
	"net/url"
	"sort"
	"strconv"
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	UrlStudios   = "https://rsg-group.api.magicline.com/connect/v1/studio" // + "?studioTags=AKTIV-391B8025C1714FB9B15BB02F2F8AC0B2"
	UrlStudios2  = "https://www.mcfit.com/typo3conf/ext/bra_studioprofiles_mcfitcom/Resources/Public/Json/studios_de.json?origLat=50.7753455&origLng=6.083886800000001&origAddress=aachen"
	UrlOccupancy = "https://www.mcfit.com/de/auslastung/antwort/request.json" // + "?tx_brastudioprofilesmcfitcom_brastudioprofiles%5BstudioId%5D=1536269110"
)

var StudioIDs = []int{1536266890, 1536269110}

func (s *Studio) DistanceTo(to cfac.Coordinate) float64 {
	loc := cfac.Coordinate{
		Latitude:  s.Address.Latitude,
		Longitude: s.Address.Longitude,
	}

	return loc.DistanceTo(to)
}

func FetchStudios(c *colly.Collector, cb func(s []Studio), ecb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		var resp ResponseStudios
		if err := json.Unmarshal(r.Body, &resp); err != nil {
			ecb(err)
			return
		}

		cb(resp)
	})

	c.Visit(UrlStudios)
}

func FetchStudiosByCoordinates(c *colly.Collector, co cfac.Coordinate, dist float64, cb func(s []Studio), ecb cfac.ErrorCallback) {
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
	}, ecb)
}

func FetchOccupancy(c *colly.Collector, studioID int, cb func(o ResponseOccupancy), ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(func(r *colly.Response) {
		defer wg.Done()

		var o ResponseOccupancy
		if err := json.Unmarshal(r.Body, &o); err != nil {
			ecb(err)
			return
		}

		cb(o)
	})

	q := url.Values{}
	q.Add("tx_brastudioprofilesmcfitcom_brastudioprofiles[studioId]", strconv.Itoa(studioID))

	c.Visit(UrlOccupancy + "?" + q.Encode())

	return wg
}

func FetchCurrentOccupancy(c *colly.Collector, studioID int, cb func(o Occupancy), ecb cfac.ErrorCallback) *sync.WaitGroup {
	return FetchOccupancy(c, studioID, func(ol ResponseOccupancy) {
		for _, o := range ol.Items {
			if o.IsCurrent {
				cb(o)
			}
		}
	}, ecb)
}
