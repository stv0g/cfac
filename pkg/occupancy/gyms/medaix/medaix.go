package medaix

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	UrlVisitors = "https://www.medaix.de/counter/visitors.php"
	UrlCalendar = "https://portal.medaix.de/courseplan/frontend/fillCalendar"
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

type Callback func(v VisitorCounter)

func requestData(id int) map[string]string {
	return map[string]string{
		"id": strconv.Itoa(id),
	}
}

func FetchOccupancy(c *colly.Collector, cb Callback, errCb cfac.ErrorCallback) {
	id := 1

	c.OnResponse(func(r *colly.Response) {
		var counter VisitorCounter

		// Try to unmarshal to bool which indicates an invalid ID
		var b bool
		if err := json.Unmarshal(r.Body, &b); err == nil {
			return
		}

		err := json.Unmarshal(r.Body, &counter)
		if err != nil {
			errCb(err)
			return
		}

		cb(counter)

		id++
		c.Post(UrlVisitors, requestData(id))
	})

	c.Post(UrlVisitors, requestData(id))
}
