// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package wfs

import (
	"time"

	"github.com/stv0g/cfac/pkg/std/ietf/geojson"
)

type FeatureCollection struct {
	geojson.FeatureCollection

	TotalFeatures  int       `json:"totalFeatures"`
	NumberMatched  int       `json:"numberMatched"`
	NumberReturned int       `json:"numberReturned"`
	Timestamp      time.Time `json:"timeStamp"`
}
