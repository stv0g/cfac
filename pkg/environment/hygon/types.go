package hygon

import "time"

type Messwert struct {
	Name  string
	Date  time.Time
	Value float64
}
