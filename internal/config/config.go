package config

type Config struct {
	Version   string     `yaml:"version"`
	Project   Project    `yaml:"project"`
	Workflows []Workflow `yaml:"workflows"`
}

type Project struct {
	Name string `yaml:"name"`
}

type Workflow struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Run string `yaml:"run"`
}
