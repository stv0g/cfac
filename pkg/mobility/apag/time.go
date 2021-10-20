package apag

import (
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

func (c *CustomTime) UnmarshalJSON(b []byte) error {
	var err error
	loc, _ := time.LoadLocation("Europe/Berlin")
	s := strings.Trim(string(b), "\"")
	c.Time, err = time.ParseInLocation("2006-01-02 15:04:05", s, loc)
	return err
}
