package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type GreetingMethod struct {
	Name   string  `yaml:"name"`
	Weight float64 `yaml:"weight"`
}

type Config struct {
	Greeting struct {
		Methods []GreetingMethod `yaml:"methods"`
	} `yaml:"greeting"`
}

func LoadConfig() (*Config, error) {
	f, err := os.Open("config/config.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	return &cfg, err
}
