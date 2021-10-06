package testing_helper

import (
	"testing"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

func NewCollyCollector(t *testing.T) *colly.Collector {
	c := colly.NewCollector()

	c.OnError(func(r *colly.Response, e error) {
		t.Errorf("Colly error: %s", e)
	})

	return c
}

func ErrorHandler(t *testing.T) cfac.ErrorCallback {
	return func(err error) {
		t.Errorf("Error for handler: %s", err)
	}
}
