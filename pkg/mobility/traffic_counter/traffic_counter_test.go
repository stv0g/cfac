// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package traffic_counter_test

import (
	"testing"

	"github.com/stv0g/cfac/pkg/mobility/traffic_counter"
)

func TestTrafficCounter(t *testing.T) {
	var cars chan traffic_counter.Car
	traffic_counter.TrafficCounter(traffic_counter.UrlFile, cars)
}
