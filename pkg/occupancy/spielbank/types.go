package spielbank

import "time"

type Occupancy struct {
	Utilization float64
	LastUpdated time.Time
}
