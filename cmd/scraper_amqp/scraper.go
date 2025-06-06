// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
	cfac "github.com/stv0g/cfac/pkg"

	_ "github.com/stv0g/cfac/pkg/all"

	"github.com/stv0g/cfac/internal/helper"
)

func main() {
	helper.SetupLogging()

	cfg, err := helper.SetupConfig()
	if err != nil {
		log.WithError(err).Fatal("Failed to parse config")
	}

	// Signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Ticker and Backup
	interval := cfg.GetDuration("interval")
	ticker := time.NewTicker(interval)

	// AMQP session
	session := NewSession("cfac", cfg.GetString("amqp.url"))
	defer session.Close()

	meas_name := cfg.GetString("measurable")

	new_meas, err := cfac.GetMeasurable(meas_name)
	if err != nil {
		panic(err)
	}

	meas := new_meas()

	c := colly.NewCollector()
	defer c.Wait()

	c.AllowURLRevisit = true
	c.Async = cfg.GetBool("scraper.async")
	c.IgnoreRobotsTxt = cfg.GetBool("scraper.ignore_robots_txt")

loop:
	for {
		select {
		case <-ticker.C:
			log.Info("Tick")

			meas.Fetch(c.Clone(), func(measurement cfac.Measurement) {
				if payload, err := json.Marshal(measurement); err == nil {
					session.Push(payload)
				}
			}, func(err error) {
				log.WithError(err).Error("Failed to fetch measurement")
			})

		case sig := <-sigs:
			if sig == syscall.SIGINT || sig == syscall.SIGTERM {
				ticker.Stop()
				break loop
			}
		}
	}

	log.Printf("Bye")
}
