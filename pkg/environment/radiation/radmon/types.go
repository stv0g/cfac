// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package radmon

import "time"

type Reading struct {
	Time     time.Time
	CPM      float32
	Location string
	User     string
}
