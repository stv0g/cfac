package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"sync"

	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
	"github.com/stv0g/cfac/internal/helper"
	cfac "github.com/stv0g/cfac/pkg"

	_ "github.com/stv0g/cfac/pkg/all"
)

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
	c.Async = cfg.UBool("scraper.async", true)
	c.IgnoreRobotsTxt = cfg.UBool("scraper.ignore_robots_txt", true)

	wg := sync.WaitGroup{}
	wg.Add(1)

	meas.Fetch(c.Clone(), func(measurements []cfac.Measurement) {
		if payload, err := json.MarshalIndent(measurements, "", "  "); err == nil {
			os.Stdout.Write(payload)
			wg.Done()
		}
	}, func(err error) {
		log.WithError(err).Error("Failed to fetch measurement")
	})

	wg.Wait()
}
