package carolus

import (
	"time"

	cfac "github.com/stv0g/cfac/pkg"
)

type Occupancy struct {
	ThermalBath cfac.Percent
	Sauna       cfac.Percent
	Parking     cfac.Percent
	LastUpdated time.Time
}
