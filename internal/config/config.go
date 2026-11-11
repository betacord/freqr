package config

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	instance *Config
	once     sync.Once
	loadErr  error
)

type Config struct {
	Radio RadioConfig `yaml:"radio"`
}

type RadioConfig struct {
	PortName       string `yaml:"port_name"`
	BaudRate       int    `yaml:"baud_rate"`
	MaxRetries     int    `yaml:"max_retries"`
	ReadSleepTime  int    `yaml:"read_sleep_time"`
	WriteSleepTime int    `yaml:"write_sleep_time"`
	TotalChannels  int    `yaml:"total_channels"`
	BreakWhenErr   bool   `yaml:"break_when_err"`
}

type CLIConfig struct {
	ListFormat string
}

func (c *Config) Validate() error {
	if c.Radio.PortName == "" {
		return fmt.Errorf("Port name is required")
	}

	if c.Radio.BaudRate <= 0 {
		return fmt.Errorf("Baud rate is required")
	}

	if c.Radio.MaxRetries <= 0 {
		return fmt.Errorf("Max retries is required")
	}

	return nil
}

func Get() *Config {
	if instance == nil {
		panic("configuration not initialized")
	}

	return instance
}

func Initialize(path string) error {
	once.Do(func() {
		data, err := os.ReadFile(path)
		if err != nil {
			loadErr = err
			return
		}

		var cfg Config

		if err := yaml.Unmarshal(data, &cfg); err != nil {
			loadErr = err
			return
		}

		if err := cfg.Validate(); err != nil {
			loadErr = err
			return
		}

		instance = &cfg
	})

	return loadErr
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error while reading config: %w", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error while parsing config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config file: %w", err)
	}

	return &cfg, nil
}
