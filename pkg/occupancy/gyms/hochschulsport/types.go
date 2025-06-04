// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package hochschulsport

import "time"

type Occupancy struct {
	Occupancy   float64
	LastUpdated time.Time
}
