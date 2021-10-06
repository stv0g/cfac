package cfac

import (
	"time"

	"github.com/gocolly/colly/v2"
)

func LastUpdated(r *colly.Response) (time.Time, error) {
	lastUpdatedStr := r.Headers.Get("Date")
	return time.Parse(time.RFC1123, lastUpdatedStr)
}
