// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package webcams

import (
	"net/url"
	"time"
)

type WebcamType int

type Webcam struct {
	Type       WebcamType
	StillImage *url.URL
	LiveStream *url.URL

	Title    string
	SubTitle string

	Source *url.URL

	UpdateRate time.Duration
	Delay      time.Duration
}
