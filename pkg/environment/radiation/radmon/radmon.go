// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package radmon

import (
	"errors"
	"net/url"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"
)

// https://radmon.org/index.php/forum/howtos-and-faqs/864-radmon-org-api

const (
	UrlApi = "https://radmon.org/radmon.php" // + ?function=showuserpage&user= oder ?function=lastreading
)

var Users []string = []string{
	"Rotter",
	"jokri",
}

var re = regexp.MustCompile(`(?m)(\d+) CPM on (\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})UTC at ([^<]*)`)

func FetchReading(c *colly.Collector, user string, cb func(Reading), ecb cfac.ErrorCallback) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c.OnResponse(func(r *colly.Response) {
		defer wg.Done()

		m := re.FindStringSubmatch(string(r.Body))
		if len(m) != 4 {
			ecb(errors.New("failed to parse response"))
			return
		}

		t, err := time.ParseInLocation("2006-01-02 15:04:05", m[2], time.UTC)
		if err != nil {
			ecb(err)
			return
		}

		c, err := strconv.Atoi(m[1])
		if err != nil {
			ecb(err)
			return
		}

		cb(Reading{
			CPM:      float32(c),
			Time:     t,
			Location: m[3],
			User:     user,
		})
	})

	q := url.Values{}
	q.Add("function", "lastreading")
	q.Add("user", user)

	c.Visit(UrlApi + "?" + q.Encode())

	return wg
}
