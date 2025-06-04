// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"

	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
	"github.com/stv0g/cfac/internal/helper"
	cfac "github.com/stv0g/cfac/pkg"

	_ "github.com/stv0g/cfac/pkg/all"
)

func DumpMeasurement(measurement cfac.Measurement) {
	if payload, err := json.MarshalIndent(measurement, "", "  "); err == nil {
		os.Stdout.Write(payload)
		os.Stdout.WriteString("\n")
	}
}

func DumpError(err error) {
	log.WithError(err).Error("Failed to fetch measurement")
}

func main() {
	helper.SetupLogging()

	cfg, err := helper.SetupConfig()
	if err != nil {
		log.WithError(err).Fatal("Failed to parse config")
	}

	if len(flag.Args()) != 1 {
		prog := filepath.Base(os.Args[0])
		log.Fatalf("usage: %s MEASURABLE", prog)
	}

	component := os.Args[1]

	ctor, err := cfac.GetMeasurable(component)
	if err != nil {
		panic(err)
	}

	meas := ctor()

	c := colly.NewCollector()
	defer c.Wait()

	c.AllowURLRevisit = true
	c.Async = cfg.GetBool("scraper.async")
	c.IgnoreRobotsTxt = cfg.GetBool("scraper.ignore_robots_txt")

	meas.Fetch(c, DumpMeasurement, DumpError).Wait()
}
