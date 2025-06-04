// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package nominatim

import cfac "github.com/stv0g/cfac/pkg"

type ReponseSearch []Place

type Place struct {
	PlaceID     int      `json:"place_id"`
	Licence     string   `json:"licence"`
	OsmType     string   `json:"osm_type"`
	OsmID       int      `json:"osm_id"`
	Boundingbox []string `json:"boundingbox"`
	Latitude    float64  `json:"lat,string"`
	Longitude   float64  `json:"lon,string"`
	DisplayName string   `json:"display_name"`
	Class       string   `json:"class"`
	Type        string   `json:"type"`
	Importance  float64  `json:"importance"`
	Icon        string   `json:"icon"`
	Address     Address  `json:"address"`
}

func (p *Place) Coordinate() cfac.Coordinate {
	return cfac.Coordinate{
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
	}
}

type Address struct {
	City        string `json:"city"`
	County      string `json:"county"`
	State       string `json:"state"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}
