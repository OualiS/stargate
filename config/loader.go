package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Structs exportés
type Config struct {
	Routes []Route `yaml:"routes"`
}

type Route struct {
	Match     Match     `yaml:"match"`
	ForwardTo ForwardTo `yaml:"forward_to"`
}

type Match struct {
	Host       string `yaml:"host"`
	PathPrefix string `yaml:"path_prefix"`
}

type ForwardTo struct {
	Container string `yaml:"container"`
	Port      int    `yaml:"port"`
}

func Load(path string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return cfg, nil
}

func Validate(cfg Config) error {
	if len(cfg.Routes) == 0 {
		return errors.New("no routes found in config")
	}

	for _, route := range cfg.Routes {
		if route.Match.Host == "" {
			return errors.New("route has empty host")
		}
		if route.Match.PathPrefix == "" {
			return errors.New("route has empty path prefix")
		}
		if route.ForwardTo.Port < 1 || route.ForwardTo.Port > 65535 {
			return fmt.Errorf("route has invalid port: %d", route.ForwardTo.Port)
		}
		if route.ForwardTo.Container == "" {
			return errors.New("route has empty container name")
		}
	}

	return nil
}
