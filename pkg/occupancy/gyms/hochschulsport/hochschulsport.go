package hochschulsport

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/otiai10/gosseract/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

const (
	Url = "https://buchung.hsz.rwth-aachen.de/cgi/studio.cgi?size=30"
	UrlReferer
)

var (
	requiredHeaders = map[string]string{
		"Referer": UrlReferer,
		"Accept":  "image/*",
	}
)

type Occupancy struct {
	Utilization float64
	LastUpdated time.Time
}

type Callback func(u Occupancy)

func FetchOccupancy(c *colly.Collector, cb Callback, errCb cfac.ErrorCallback) {
	c.OnResponse(func(r *colly.Response) {
		client := gosseract.NewClient()
		defer client.Close()

		err := client.SetImageFromBytes(r.Body)
		if err != nil {
			errCb(err)
			return
		}

		text, err := client.Text()
		if err != nil {
			errCb(err)
			return
		}

		if text == "O" || text == "o" {
			text = "0"
		}

		fmt.Println(text)

		utilization, err := strconv.ParseUint(text, 10, 64)
		if err != nil {
			errCb(err)
			return
		}

		lastUpdatedStr := r.Headers.Get("Date")
		lastUpdated, err := time.Parse(time.RFC1123, lastUpdatedStr)
		if err != nil {
			errCb(err)
			return
		}

		cb(Occupancy{
			Utilization: float64(utilization),
			LastUpdated: lastUpdated,
		})
	})

	url, err := url.Parse(Url)
	if err != nil {
		errCb(err)
		return
	}

	c.OnRequest(func(r *colly.Request) {
		if r.URL.Host == url.Host {
			for key, value := range requiredHeaders {
				r.Headers.Set(key, value)
			}
		}
	})

	c.Visit(Url)
}
