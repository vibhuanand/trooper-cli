package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config %s: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, fmt.Errorf("parse yaml %s: %w", path, err)
	}

	if cfg.Version == "" {
		return nil, fmt.Errorf("config missing 'version'")
	}

	return &cfg, nil
}

func (c *Config) FindWorkflow(name string) (*Workflow, bool) {
	for i := range c.Workflows {
		if c.Workflows[i].Name == name {
			return &c.Workflows[i], true
		}
	}
	return nil, false
}
