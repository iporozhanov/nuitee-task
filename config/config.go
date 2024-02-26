package config

import "time"

type Config struct {
	HTTP     HTTP     `yaml:"http"`
	Apitude  Apitude  `yaml:"apitude"`
	Coinbase Coinbase `yaml:"coinbase"`
}

func (c *Config) Default() {
	c.HTTP.Port = "8080"
}

type HTTP struct {
	Port string `yaml:"port"`
}

type Apitude struct {
	APIKey string `yaml:"apiKey"`
	Secret string `yaml:"secret"`
	APIUrl string `yaml:"apiUrl"`
}

type Coinbase struct {
	BaseURL         string        `yaml:"baseUrl"`
	RefreshInterval time.Duration `yaml:"refreshInterval"`
}
