package config

// TechStack represents a technology stack configuration
type TechStack struct {
	Code        string `yaml:"code"`
	Description string `yaml:"description"`
}

// Framework represents a framework configuration
type Framework struct {
	Code        string `yaml:"code"`
	Description string `yaml:"description"`
}

// Config represents the user's custom configuration
type Config struct {
	TechStacks []TechStack            `yaml:"tech-stacks,omitempty"`
	Frameworks map[string][]Framework `yaml:"frameworks,omitempty"`
	Categories map[string][]string    `yaml:"categories,omitempty"`
}

// NewConfig creates a new empty config
func NewConfig() *Config {
	return &Config{
		TechStacks: []TechStack{},
		Frameworks: make(map[string][]Framework),
		Categories: make(map[string][]string),
	}
}
