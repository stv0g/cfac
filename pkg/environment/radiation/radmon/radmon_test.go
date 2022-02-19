package radmon_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/environment/radiation/radmon"
)

func TestFetchReading(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	for _, user := range radmon.Users {
		radmon.FetchReading(c.Collector.Clone(), user, func(reading radmon.Reading) {
			t.Logf("Reading = %+#v", reading)
			c.MarkHandled()
		}, c.ErrorCallback())
	}
}
