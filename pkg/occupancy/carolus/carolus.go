package carolus

import (
	"regexp"
	"strconv"
	"sync"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	Url = "https://carolus-thermen.de/auslastung/"
)

var re = regexp.MustCompile(`(?m)(\d+) %`)

type Callback func(o Occupancy)

func FetchOccupancy(c *colly.Collector, cb Callback, ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnScraped(func(r *colly.Response) {
		wg.Done()
	})

	c.OnHTML("div.modulAuslastugsGrid", func(h *colly.HTMLElement) {
		occupancyThermalBathStr := h.ChildText("div.item-aus:nth-child(1) div div div")
		occupancyThermalBathStr = re.FindStringSubmatch(occupancyThermalBathStr)[1]

		occupancySaunaStr := h.ChildText("div.item-aus:nth-child(2) div div div")
		occupancySaunaStr = re.FindStringSubmatch(occupancySaunaStr)[1]

		occupancyParkingStr := h.ChildText("div.item-aus:nth-child(3) div div div")
		occupancyParkingStr = re.FindStringSubmatch(occupancyParkingStr)[1]

		occupancyThermalBath, _ := strconv.Atoi(occupancyThermalBathStr)
		occupancySauna, _ := strconv.Atoi(occupancySaunaStr)
		occupancyParking, _ := strconv.Atoi(occupancyParkingStr)

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

	return wg
}
