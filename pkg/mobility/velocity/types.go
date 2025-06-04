// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package velocity

type Station struct {
	StationID         int         `json:"stationId"`
	Name              string      `json:"name"`
	LocationLatitude  float64     `json:"locationLatitude"`
	LocationLongitude float64     `json:"locationLongitude"`
	State             string      `json:"state"`
	Note              string      `json:"note"`
	NumFreeSlots      int         `json:"numFreeSlots"`
	NumAllSlots       int         `json:"numAllSlots"`
	Image             interface{} `json:"image"`
}

type PedelecInfo struct {
	StateOfCharge float64 `json:"stateOfCharge"`
	Availability  string  `json:"availability"`
	Name          string  `json:"name"`
}

type StationSlot struct {
	StationSlotID       int         `json:"stationSlotId"`
	StationSlotPosition int         `json:"stationSlotPosition"`
	State               string      `json:"state"`
	IsOccupied          bool        `json:"isOccupied"`
	StateOfCharge       float64     `json:"stateOfCharge"`
	PedelecInfo         PedelecInfo `json:"pedelecInfo,omitempty"`
}

type StationSlots struct {
	StationSlots []StationSlot `json:"stationSlots"`
}
