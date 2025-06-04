// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package sensor_community_test

import (
	"testing"

	th "github.com/stv0g/cfac/internal/testing"
	"github.com/stv0g/cfac/pkg/city"
	"github.com/stv0g/cfac/pkg/environment/sensor_community"
)

func TestFetchData(t *testing.T) {
	c := th.NewCollector(t)
	defer c.Close()

	sensor_community.FetchDataRadius(city.Aachen.Center, 10e3, c.Collector, func(sensor sensor_community.Sensor) {
		t.Logf("Sensor = %+#v", sensor)
		c.MarkHandled()
	}, c.ErrorCallback())
}
