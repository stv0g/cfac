package wof

import (
	"time"

	cfac "github.com/stv0g/cfac/pkg"
)

type Studio struct {
	Name        string
	Location    string
	Occupancy   cfac.Percent
	LastUpdated time.Time
}
