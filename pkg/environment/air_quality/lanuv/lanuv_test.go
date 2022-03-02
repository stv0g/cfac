package lanuv_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"

	"github.com/stv0g/cfac/pkg/environment/air_quality/lanuv"
)

func TestFetch(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	lanuv.Fetch(c.Collector, func(aq lanuv.AirQuality) {
		t.Logf("Air quality: %+#v", aq)
		c.MarkHandled()
	}, c.ErrorCallback())
}
