package cccac_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/occupancy/cccac"
)

func TestFetchStatus(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	cccac.FetchStatus(c.Collector, func(sts cccac.Status) {
		t.Logf("Current status: %+v", sts)

		c.MarkHandled()
	}, c.ErrorCallback())
}
