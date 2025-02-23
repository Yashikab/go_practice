package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Endpoint struct {
	Path   string  `yaml:"path"`
	Weight float64 `yaml:"weight"`
}

type Config struct {
	Endpoints  []Endpoint `yaml:"endpoints"`
	BackendURL string     `yaml:"backend_url"`
}

func LoadConfig(filename string) (*Config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
