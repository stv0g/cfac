// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package hygon

// https://luadb.lds.nrw.de/LUA/hygon/hygon-datenmodell.pdf

const (
	Url = "https://luadb.lds.nrw.de/LUA/hygon/messwerte"

	UrlMesswerte    = Url + "/messwerte.tar.gz"
	UrlTemperatur   = Url + "/temperatur.tar.gz"
	UrlNiederschlag = Url + "/niederschlag.tar.gz"
)
