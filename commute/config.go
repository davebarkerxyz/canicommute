package commute

import (
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	ArrivalTime struct {
		Hour int `yaml:"hour"`
		Min  int `yaml:"min"`
	} `yaml:"arrival_time"`
	Locations  []string `yaml:"locations"`
	AutoSuffix string   `yaml:"auto_suffix"`
	ApiKey     string   `yaml:"api_key"`
}

// Read config from config.yaml and return populated Config struct
func GetConfig() Config {
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		Die("Error reading config file: %s", err)
	}

	config := Config{}
	err = yaml.UnmarshalStrict(configFile, &config)
	if err != nil {
		Die("Error parsing config file: %s", err)
	}

	return config
}
