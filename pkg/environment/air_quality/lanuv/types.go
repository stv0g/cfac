// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package lanuv

import "time"

type AirQuality struct {
	Station   string
	Timestamp time.Time
	ID        string
	Ozon      float64
	SO2       float64
	NO2       float64
	PM10      float64
}
