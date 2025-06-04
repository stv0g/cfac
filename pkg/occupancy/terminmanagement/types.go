// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package terminmanagement

import (
	"strings"
	"time"
)

type Call string

type Response struct {
	WaitCircles  map[string]string         `json:"waitcircles"`
	WaitingTime  map[string]CustomDuration `json:"waitingtime"`
	VisitorCount map[string]string         `json:"visitorcount"`
	LastCalls    map[string]string         `json:"last_calls"`
	LastRefresh  CustomTime                `json:"last_refresh"`
	NextCalls    interface{}               `json:"next_calls"`
	// NextCalls    map[string][]Call         `json:"next_calls"`
}

type WaitCircle struct {
	Name         string
	WaitingTime  time.Duration
	VisitorCount int
	LastRefresh  time.Time
}

type CustomDuration struct {
	time.Duration
}

type CustomTime struct {
	time.Time
}

func (c *CustomDuration) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	z, _ := time.Parse("15:04:05", "00:00:00")
	t, err := time.Parse("15:04:05", s)
	c.Duration = t.Sub(z)
	return err
}

func (c *CustomTime) UnmarshalJSON(b []byte) error {
	var err error
	s := strings.Trim(string(b), "\"")
	c.Time, err = time.Parse("02.01.2006 15:04:05", s)
	return err
}
