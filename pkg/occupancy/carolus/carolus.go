package carolus

import (
	"strconv"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	Url = "https://www.carolus-thermen.de/en/thermalbath/#occupation"
)

type Callback func(o Occupancy)

func FetchOccupancy(c *colly.Collector, cb Callback, ecb cfac.ErrorCallback) {
	c.OnHTML("ul.occupancy", func(h *colly.HTMLElement) {

		occupancyThermalBathStr := h.ChildAttrs("div[data-name=\"Thermal Bath\"]", "data-percent")[0]
		occupancySaunaStr := h.ChildAttrs("div[data-name=Sauna]", "data-percent")[0]
		occupancyParkingStr := h.ChildAttrs("div[data-name=Parking]", "data-percent")[0]

		occupancyThermalBath, err := strconv.Atoi(occupancyThermalBathStr)
		occupancySauna, err := strconv.Atoi(occupancySaunaStr)
		occupancyParking, err := strconv.Atoi(occupancyParkingStr)

		lastUpdated, err := cfac.LastUpdated(h.Response)
		if err != nil {
			ecb(err)
			return
		}

		cb(Occupancy{
			ThermalBath: cfac.Percent(occupancyThermalBath),
			Sauna:       cfac.Percent(occupancySauna),
			Parking:     cfac.Percent(occupancyParking),
			LastUpdated: lastUpdated,
		})

	})

	c.Visit(Url)
}
