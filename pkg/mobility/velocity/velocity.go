package velocity

const (
	UrlApi          = "https://cms.velocity-aachen.de/backend/app"
	UrlStations     = UrlApi + "/stations"
	UrlStation      = UrlApi + "/stations/{station_id}"
	UrlStationSlots = UrlApi + "/stations/{station_id}/slots/full"
)

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
