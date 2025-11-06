package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aravindcm49/dir-init/internal/categories"
	"github.com/aravindcm49/dir-init/internal/generator"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func interactive() {
	// Initialize color settings
	green := color.New(color.FgGreen).Add(color.Bold)
	yellow := color.New(color.FgYellow).Add(color.Bold)

	
	// Step 1: Tech Stack + Language Selection
	yellow.Printf("Step 1/4: Select Tech Stack + Language\n")

	var techStackItems []string
	techDescriptions := categories.TechStackDescriptions()
	for _, code := range categories.TechStackWords {
		description := techDescriptions[code]
		techStackItems = append(techStackItems, fmt.Sprintf("%s - %s", code, description))
	}

	techStackPrompt := promptui.Select{
		Label: "",
		Items: techStackItems,
		Templates: &promptui.SelectTemplates{
			Active:   "❯ {{ . | cyan }}",
			Inactive: "  {{ . }}",
			Selected: "✓ {{ . | green }}",
		},
	}

	techStackIdx, selectedTechItem, err := techStackPrompt.Run()
	if err != nil {
		fmt.Printf("Error selecting tech stack: %v\n", err)
		return
	}

	selectedTechStack := categories.TechStackWords[techStackIdx]
	selectedTechDesc := strings.Split(selectedTechItem, " - ")[1]

	
	// Step 2: Framework Selection (filtered by tech stack)
	yellow.Printf("Step 2/4: Select Framework (%s)\n", selectedTechDesc)

	var frameworkItems []string
	frameworkDescriptions := categories.FrameworkDescriptions()
	availableFrameworks := categories.GetFrameworksForTech(selectedTechStack)

	for _, frameworkCode := range availableFrameworks {
		description := frameworkDescriptions[frameworkCode]
		frameworkItems = append(frameworkItems, fmt.Sprintf("%s - %s", frameworkCode, description))
	}

	frameworkPrompt := promptui.Select{
		Label: "",
		Items: frameworkItems,
		Templates: &promptui.SelectTemplates{
			Active:   "❯ {{ . | cyan }}",
			Inactive: "  {{ . }}",
			Selected: "✓ {{ . | green }}",
		},
	}

	frameworkIdx, _, err := frameworkPrompt.Run()
	if err != nil {
		fmt.Printf("Error selecting framework: %v\n", err)
		return
	}

	selectedFramework := availableFrameworks[frameworkIdx]

	
	// Step 3: Category Selection (removed tech category)
	yellow.Printf("Step 3/4: Select Category\n")

	var categoryItems []string
	categoryItems = append(categoryItems, "1. food", "2. animals", "3. pop", "4. silly", "5. dev", "6. all")

	categoryPrompt := promptui.Select{
		Label: "",
		Items: categoryItems,
		Templates: &promptui.SelectTemplates{
			Active:   "❯ {{ . | cyan }}",
			Inactive: "  {{ . }}",
			Selected: "✓ {{ . | green }}",
		},
	}

	categoryIdx, _, err := categoryPrompt.Run()
	if err != nil {
		fmt.Printf("Error selecting category: %v\n", err)
		return
	}

	var selectedCategory string
	switch categoryIdx {
	case 0:
		selectedCategory = "food"
	case 1:
		selectedCategory = "animals"
	case 2:
		selectedCategory = "pop"
	case 3:
		selectedCategory = "silly"
	case 4:
		selectedCategory = "dev"
	case 5:
		selectedCategory = "all"
	}

	
	// Step 4: Suffix Type Selection
	yellow.Printf("Step 4/4: Select Suffix Type\n")

	suffixPrompt := promptui.Select{
		Label: "",
		Items: []string{
			"1. alphabetic (abc)",
			"2. numeric (123)",
			"3. mixed (a1b2)",
			"4. timestamp",
		},
		Templates: &promptui.SelectTemplates{
			Active:   "❯ {{ . | cyan }}",
			Inactive: "  {{ . }}",
			Selected: "✓ {{ . | green }}",
		},
	}

	suffixIdx, _, err := suffixPrompt.Run()
	if err != nil {
		fmt.Printf("Error selecting suffix type: %v\n", err)
		return
	}

	// Extract suffix type
	var suffixType generator.SuffixType
	switch suffixIdx {
	case 0:
		suffixType = generator.SuffixAlpha
	case 1:
		suffixType = generator.SuffixNumeric
	case 2:
		suffixType = generator.SuffixMixed
	case 3:
		suffixType = generator.SuffixTimestamp
	}

	// Count selection
	countPrompt := promptui.Prompt{
		Label:   "How many directories? (1-10)",
		Default: "1",
		Validate: func(input string) error {
			count, err := strconv.Atoi(input)
			if err != nil || count < 1 || count > 10 {
				return fmt.Errorf("please enter a number between 1 and 10")
			}
			return nil
		},
	}

	countStr, err := countPrompt.Run()
	if err != nil {
		fmt.Printf("Error entering count: %v\n", err)
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		fmt.Printf("Error parsing count: %v\n", err)
		return
	}

	// Generate and create directories
	gen := generator.NewGenerator(generator.DefaultConfig())

	for i := 0; i < count; i++ {
		name, err := gen.GenerateEnhancedName(selectedTechStack, selectedFramework, selectedCategory, suffixType, 4)
		if err != nil {
			fmt.Printf("Error generating name: %v\n", err)
			continue
		}

		// Create directory
		err = os.MkdirAll(name, 0755)
		if err != nil {
			fmt.Printf("❌ Failed to create directory '%s': %v\n", name, err)
			continue
		}

		fmt.Printf("\n")
		green.Printf(" ✓ Created: %s\n", name)

		// Change into the created directory
		err = os.Chdir(name)
		if err != nil {
			fmt.Printf("⚠️  Could not change directory: %v\n", err)
		} else {
			green.Printf(" → Changed into directory: %s\n", name)

			// Generate a script that can be sourced to change the shell's directory
			fmt.Printf("\n")
			yellow.Printf("💡 To change your shell's directory, run:\n")
			yellow.Printf("   cd %s\n", name)
			fmt.Printf("\n")
		}
	}
}