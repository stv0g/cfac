// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package buienalarm

type ResponseForecast Forecast

type Forecast struct {
	Success    bool   `json:"success"`
	Start      int    `json:"start"`
	StartHuman string `json:"start_human"`
	Temp       int    `json:"temp"`
	Delta      int    `json:"delta"`
	Precip     []int  `json:"precip"`
	Levels     Level  `json:"levels"`
	Grid       Grid   `json:"grid"`
	Source     string `json:"source"`
	Bounds     Bounds `json:"bounds"`
}

type Bounds struct {
	N float64 `json:"N"`
	E float64 `json:"E"`
	S float64 `json:"S"`
	W float64 `json:"W"`
}

type Grid struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Level struct {
	Light    float64 `json:"light"`
	Moderate float64 `json:"moderate"`
	Heavy    float64 `json:"heavy"`
}
