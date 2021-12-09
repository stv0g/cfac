package medaix

import (
	"strings"
	"time"
)

type ResponseVisitors VisitorCounter

type CustomTime struct {
	time.Time
}

func (c *CustomTime) UnmarshalJSON(b []byte) error {
	var err error
	s := strings.Trim(string(b), "\"")
	c.Time, err = time.Parse("2006-01-02 15:04:05", s)
	return err
}

type VisitorCounter struct {
	ID         int        `json:"id,string"`
	Name       string     `json:"name"`
	Badge      string     `json:"badge"`
	BadgeClass string     `json:"badgeClass"`
	ShowBadge  string     `json:"showBadge"`
	Value      int        `json:"value,string"`
	Max        int        `json:"max,string"`
	A          string     `json:"a"`
	B          string     `json:"b"`
	Active     string     `json:"active"`
	Timestamp  CustomTime `json:"timestamp"`
}
