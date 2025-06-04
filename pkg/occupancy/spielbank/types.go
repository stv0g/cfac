// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package spielbank

import "time"

type Occupancy struct {
	Utilization float64
	LastUpdated time.Time
}
