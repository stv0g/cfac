// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package wof

import (
	"time"

	cfac "github.com/stv0g/cfac/pkg"
)

type Studio struct {
	Name        string
	Location    string
	Occupancy   cfac.Percent
	LastUpdated time.Time
}
