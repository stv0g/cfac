package radmon

import "time"

type Reading struct {
	Time     time.Time
	CPM      float32
	Location string
	User     string
}
