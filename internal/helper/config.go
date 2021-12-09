package helper

import (
	"flag"
	"os"

	"github.com/olebedev/config"
)

const defaultConfig = `
amqp:
  url: amqp://admin:admin@rabbitmq:5672//

interval: 10s

scraper:
  ignore_robots_txt: true
  async: true

measurable: apag
`

func SetupConfig() (*config.Config, error) {
	var cfgFile string = ""
	var cfg *config.Config
	var err error

	flag.StringVar(&cfgFile, "config", "", "Configuration file")
	flag.Parse()

	if cfgFile == "" {
		cfg, err = config.ParseYaml(defaultConfig)
	} else {
		cfg, err = config.ParseYamlFile(cfgFile)
	}
	if err != nil && err != os.ErrNotExist {
		return nil, err
	}

	cfg.EnvPrefix("CFAC")
	cfg.Flag()

	return cfg, nil
}
