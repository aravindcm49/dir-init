package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	configMutex sync.RWMutex
	configPath  string
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		configPath = ".dir-init/config.yaml"
	} else {
		configPath = filepath.Join(home, ".dir-init", "config.yaml")
	}
}

// GetConfigPath returns the path to the config file
func GetConfigPath() string {
	return configPath
}

// LoadConfig loads the user configuration from ~/.dir-init/config.yaml
func LoadConfig() (*Config, error) {
	configMutex.RLock()
	defer configMutex.RUnlock()

	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Return empty config if file doesn't exist
		return NewConfig(), nil
	}

	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse YAML
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Initialize maps if nil
	if config.Frameworks == nil {
		config.Frameworks = make(map[string][]Framework)
	}
	if config.Categories == nil {
		config.Categories = make(map[string][]string)
	}

	return &config, nil
}

// SaveConfig saves the configuration to ~/.dir-init/config.yaml
func SaveConfig(config *Config) error {
	configMutex.Lock()
	defer configMutex.Unlock()

	// Create directory if it doesn't exist
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Marshal to YAML
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Add header comment
	header := "# dir-init Custom Collections\n# Auto-generated and manually editable\n\n"
	finalData := append([]byte(header), data...)

	// Write to temp file first (atomic write)
	tempPath := configPath + ".tmp"
	if err := os.WriteFile(tempPath, finalData, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	// Rename temp file to actual config file
	if err := os.Rename(tempPath, configPath); err != nil {
		os.Remove(tempPath) // Clean up temp file
		return fmt.Errorf("failed to save config file: %w", err)
	}

	return nil
}

// SaveTechStack adds a tech stack to the config
func SaveTechStack(code, description string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// Check if already exists
	for _, ts := range config.TechStacks {
		if ts.Code == code {
			return fmt.Errorf("tech stack '%s' already exists", code)
		}
	}

	// Add new tech stack
	config.TechStacks = append(config.TechStacks, TechStack{
		Code:        code,
		Description: description,
	})

	return SaveConfig(config)
}

// SaveFramework adds a framework to the config
func SaveFramework(techStack, code, description string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// Initialize frameworks map for tech stack if needed
	if config.Frameworks[techStack] == nil {
		config.Frameworks[techStack] = []Framework{}
	}

	// Check if already exists
	for _, fw := range config.Frameworks[techStack] {
		if fw.Code == code {
			return fmt.Errorf("framework '%s' already exists for tech stack '%s'", code, techStack)
		}
	}

	// Add new framework
	config.Frameworks[techStack] = append(config.Frameworks[techStack], Framework{
		Code:        code,
		Description: description,
	})

	return SaveConfig(config)
}

// SaveCategoryWord adds a word to a category in the config
func SaveCategoryWord(category, word string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// Initialize category if needed
	if config.Categories[category] == nil {
		config.Categories[category] = []string{}
	}

	// Check if already exists
	for _, w := range config.Categories[category] {
		if w == word {
			return fmt.Errorf("word '%s' already exists in category '%s'", word, category)
		}
	}

	// Add new word
	config.Categories[category] = append(config.Categories[category], word)

	return SaveConfig(config)
}

// InitConfig creates an example config file
func InitConfig() error {
	// Check if config already exists
	if _, err := os.Stat(configPath); err == nil {
		return fmt.Errorf("config file already exists at %s", configPath)
	}

	// Create example config
	exampleConfig := &Config{
		TechStacks: []TechStack{
			{Code: "mystack", Description: "My Custom Stack"},
		},
		Frameworks: map[string][]Framework{
			"mystack": {
				{Code: "myframework", Description: "My Custom Framework"},
			},
		},
		Categories: map[string][]string{
			"food": {"ramen", "sushi"},
		},
	}

	return SaveConfig(exampleConfig)
}
