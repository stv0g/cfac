package helper

import (
	"flag"
	"time"

	"github.com/spf13/viper"
)

func SetupConfig() (*viper.Viper, error) {
	cfg := viper.New()

	cfg.SetDefault("amqp.url", "amqp://admin:admin@rabbitmq:5672//")
	cfg.SetDefault("interval", 10*time.Second)
	cfg.SetDefault("scraper.ignore_robots_txt", true)
	cfg.SetDefault("scraper.async", true)
	cfg.SetDefault("measurable", "apag")

	cfg.SetEnvPrefix("CFAC")

	cfgFile := flag.String("config", "", "Configuration file")
	flag.Parse()

	if cfgFile != nil {
		cfg.SetConfigFile(*cfgFile)
	} else {
		cfg.SetConfigName("config")
		cfg.SetConfigType("yaml")
		cfg.AddConfigPath("/etc/cfac/")
		cfg.AddConfigPath("$HOME/.cfac")
		cfg.AddConfigPath(".")
	}

	return cfg, cfg.ReadInConfig()
}
