package cmd

import (
	"fmt"
	"regexp"

	"github.com/aravindcm49/dir-init/internal/config"
	"github.com/manifoldco/promptui"
)

// promptCustomTechStack prompts for custom tech stack input
func promptCustomTechStack() (code, description string, save bool, err error) {
	// Prompt for code
	codePrompt := promptui.Prompt{
		Label: "Tech stack code",
		Validate: func(input string) error {
			if len(input) < 2 || len(input) > 10 {
				return fmt.Errorf("code must be 2-10 characters")
			}
			if !regexp.MustCompile(`^[a-z0-9-]+$`).MatchString(input) {
				return fmt.Errorf("code must be lowercase alphanumeric (hyphens allowed)")
			}
			return nil
		},
	}

	code, err = codePrompt.Run()
	if err != nil {
		return "", "", false, err
	}

	// Prompt for description
	descPrompt := promptui.Prompt{
		Label: "Description",
		Validate: func(input string) error {
			if len(input) < 2 {
				return fmt.Errorf("description must be at least 2 characters")
			}
			return nil
		},
	}

	description, err = descPrompt.Run()
	if err != nil {
		return "", "", false, err
	}

	// Ask if they want to save
	savePrompt := promptui.Select{
		Label: "Save?",
		Items: []string{"No", "Yes"},
		Templates: &promptui.SelectTemplates{
			Active:   "{{ . | cyan }}",
			Inactive: "{{ . }}",
		},
	}

	saveIdx, _, err := savePrompt.Run()
	if err != nil {
		return "", "", false, err
	}
	save = (saveIdx == 1)

	return code, description, save, nil
}

// promptCustomFramework prompts for custom framework input
func promptCustomFramework() (code, description string, save bool, err error) {
	// Prompt for code
	codePrompt := promptui.Prompt{
		Label: "Framework code",
		Validate: func(input string) error {
			if len(input) < 2 || len(input) > 20 {
				return fmt.Errorf("code must be 2-20 characters")
			}
			if !regexp.MustCompile(`^[a-z0-9-]+$`).MatchString(input) {
				return fmt.Errorf("code must be lowercase alphanumeric (hyphens allowed)")
			}
			return nil
		},
	}

	code, err = codePrompt.Run()
	if err != nil {
		return "", "", false, err
	}

	// Prompt for description
	descPrompt := promptui.Prompt{
		Label: "Description",
		Validate: func(input string) error {
			if len(input) < 2 {
				return fmt.Errorf("description must be at least 2 characters")
			}
			return nil
		},
	}

	description, err = descPrompt.Run()
	if err != nil {
		return "", "", false, err
	}

	// Ask if they want to save
	savePrompt := promptui.Select{
		Label: "Save?",
		Items: []string{"No", "Yes"},
		Templates: &promptui.SelectTemplates{
			Active:   "{{ . | cyan }}",
			Inactive: "{{ . }}",
		},
	}

	saveIdx, _, err := savePrompt.Run()
	if err != nil {
		return "", "", false, err
	}
	save = (saveIdx == 1)

	return code, description, save, nil
}

// promptCustomWord prompts for custom category word input
func promptCustomWord() (word string, save bool, err error) {
	// Prompt for word
	wordPrompt := promptui.Prompt{
		Label: "Custom word",
		Validate: func(input string) error {
			if len(input) < 2 || len(input) > 20 {
				return fmt.Errorf("word must be 2-20 characters")
			}
			if !regexp.MustCompile(`^[a-z0-9]+$`).MatchString(input) {
				return fmt.Errorf("word must be lowercase alphanumeric (no spaces or special chars)")
			}
			return nil
		},
	}

	word, err = wordPrompt.Run()
	if err != nil {
		return "", false, err
	}

	// Ask if they want to save
	savePrompt := promptui.Select{
		Label: "Save?",
		Items: []string{"No", "Yes"},
		Templates: &promptui.SelectTemplates{
			Active:   "{{ . | cyan }}",
			Inactive: "{{ . }}",
		},
	}

	saveIdx, _, err := savePrompt.Run()
	if err != nil {
		return "", false, err
	}
	save = (saveIdx == 1)

	return word, save, nil
}

// mergeWithCustomConfig merges built-in items with custom config
func mergeWithCustomConfig() (techStacks map[string]string, frameworks map[string]map[string]string, err error) {
	// Load custom config
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, nil, err
	}

	// Initialize maps
	techStacks = make(map[string]string)
	frameworks = make(map[string]map[string]string)

	// Add custom tech stacks
	for _, ts := range cfg.TechStacks {
		techStacks[ts.Code] = ts.Description
	}

	// Add custom frameworks
	for techStack, fws := range cfg.Frameworks {
		if frameworks[techStack] == nil {
			frameworks[techStack] = make(map[string]string)
		}
		for _, fw := range fws {
			frameworks[techStack][fw.Code] = fw.Description
		}
	}

	return techStacks, frameworks, nil
}
