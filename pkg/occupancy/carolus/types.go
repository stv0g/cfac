package carolus

import (
	"time"

	cfac "github.com/stv0g/cfac/pkg"
)

var Coordinate = cfac.Coordinate{
	Latitude:  50.782993,
	Longitude: 6.097059,
}

type Occupancy struct {
	ThermalBath cfac.Percent
	Sauna       cfac.Percent
	Parking     cfac.Percent
	LastUpdated time.Time
}
