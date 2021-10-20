package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gocolly/colly/v2"
	cfac "github.com/stv0g/cfac/pkg"

	"github.com/olebedev/config"

	_ "github.com/stv0g/cfac/pkg/mobility/apag"
)

func setupLogging() {

}

func setupConfig() *config.Config {
	var configFile string

	flag.StringVar(&configFile, "config", "config.yaml", "Configuration file")
	flag.Parse()

	cfg, err := config.ParseYamlFile(configFile)
	if err != nil {
		panic(err)
	}

	cfg.Flag()
	cfg.EnvPrefix("CFAC")

	return cfg
}

func main() {
	setupLogging()
	cfg := setupConfig()

	// Signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Ticker and Backup
	intervalStr := cfg.UString("interval", "10s")
	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(interval)

	// AMQP session
	session := NewSession("cfac", cfg.UString("amqp.url"))

	component := cfg.UString("component", "apag")

	f, err := cfac.GetMeasurable(component)
	if err != nil {
		panic(err)
	}

	c := colly.NewCollector()
	defer c.Wait()

	c.AllowURLRevisit = true
	c.Async = cfg.UBool("scraper.async", true)
	c.IgnoreRobotsTxt = cfg.UBool("scraper.ignore_robots_txt", true)

loop:
	for {
		select {
		case <-ticker.C:
			log.Println("Tick")

			f(c, func(measurements []cfac.Measurement) {
				if payload, err := json.Marshal(measurements); err == nil {
					session.Push(payload)
				}
			}, func(err error) {
				log.Printf("%s", err)
			})

		case sig := <-sigs:
			if sig == syscall.SIGINT || sig == syscall.SIGTERM {
				ticker.Stop()
				break loop
			}
		}
	}

	session.Close()

	log.Printf("Bye")
}
