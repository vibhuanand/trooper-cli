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
	Name    string `yaml:"name"`
	Workdir string `yaml:"workdir,omitempty"`
	Steps   []Step `yaml:"steps"`
}

type Step struct {
	// Workdir overrides the workflow workdir for this step (optional)
	Workdir string `yaml:"workdir,omitempty"`

	// Mutually exclusive step types (exactly one should be set)
	Run       string    `yaml:"run,omitempty"`
	Terraform *ToolStep `yaml:"terraform,omitempty"`
	Kubectl   *ToolStep `yaml:"kubectl,omitempty"`
}

type ToolStep struct {
	Args []string `yaml:"args"`
}
