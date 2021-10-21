package cfac

import "github.com/gocolly/colly/v2"

type Measurement interface{}

type Measurable interface {
	Fetch(c *colly.Collector, cb MeasurementsCallback, ecb ErrorCallback)
}
