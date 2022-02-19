package tdrm_test

import (
	"fmt"
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/environment/radiation/tdrm"
)

func TestFetch(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	tdrm.Fetch(c.Collector, func(s tdrm.Station) {
		fmt.Printf("%#+v\n", s)
		c.MarkHandled()
	}, c.ErrorCallback())
}
