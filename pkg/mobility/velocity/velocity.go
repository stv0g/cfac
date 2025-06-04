// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package velocity

const (
	UrlApi          = "https://cms.velocity-aachen.de/backend/app"
	UrlStations     = UrlApi + "/stations"
	UrlStation      = UrlApi + "/stations/{station_id}"
	UrlStationSlots = UrlApi + "/stations/{station_id}/slots/full"
)
