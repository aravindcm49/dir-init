package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aravindcm49/dir-init/cmd/tui/models"
	"github.com/aravindcm49/dir-init/internal/categories"
	"github.com/aravindcm49/dir-init/internal/config"
	"github.com/aravindcm49/dir-init/internal/generator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func interactive() {
	green := color.New(color.FgGreen).Add(color.Bold)
	yellow := color.New(color.FgYellow).Add(color.Bold)

	// Step 1: Tech Stack Selection with Bubble Tea
	yellow.Printf("Step 1/4: Select Tech Stack + Language\n")

	techItems := buildTechStackItems()
	techModel := models.NewSelector("", techItems)

	p := tea.NewProgram(techModel)
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	m := finalModel.(models.SelectorModel)
	selectedTech := m.GetSelected()

	if selectedTech == nil {
		return // User quit
	}

	// Save if custom and user chose to save
	if selectedTech.IsCustom && m.ShouldSave() {
		if err := config.SaveTechStack(selectedTech.Code, selectedTech.Description); err != nil {
			// Silent fail
		}
	}

	selectedTechStack := selectedTech.Code
	selectedTechDesc := selectedTech.Description

	// Step 2: Framework Selection with Bubble Tea
	yellow.Printf("\nStep 2/4: Select Framework (%s)\n", selectedTechDesc)

	frameworkItems := buildFrameworkItems(selectedTechStack)
	frameworkModel := models.NewSelector("", frameworkItems)

	p = tea.NewProgram(frameworkModel)
	finalModel, err = p.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fm := finalModel.(models.SelectorModel)
	selectedFw := fm.GetSelected()

	if selectedFw == nil {
		return
	}

	// Save if custom and user chose to save
	if selectedFw.IsCustom && fm.ShouldSave() {
		if err := config.SaveFramework(selectedTechStack, selectedFw.Code, selectedFw.Description); err != nil {
			// Silent fail
		}
	}

	selectedFramework := selectedFw.Code

	// Step 3: Category Selection (using old promptui for now)
	yellow.Printf("\nStep 3/4: Select Category\n")

	var categoryItems []string
	categoryItems = append(categoryItems, "1. food", "2. animals", "3. pop", "4. silly", "5. dev", "6. all", "7. custom word (one-time)")

	categoryPrompt := promptui.Select{
		Label: "",
		Items: categoryItems,
		Templates: &promptui.SelectTemplates{
			Active:   "❯ {{ . | cyan }}",
			Inactive: "  {{ . }}",
			Selected: "✓ {{ . | green }}",
		},
		Searcher: func(input string, index int) bool {
			item := categoryItems[index]
			return strings.Contains(strings.ToLower(item), strings.ToLower(input))
		},
	}

	categoryIdx, _, err := categoryPrompt.Run()
	if err != nil {
		fmt.Printf("Error selecting category: %v\n", err)
		return
	}

	var selectedCategory string
	var customWord string
	var useCustomWord bool

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
	case 6:
		// Custom word
		word, save, err := promptCustomWord()
		if err != nil {
			fmt.Printf("Error getting custom word: %v\n", err)
			return
		}

		customWord = word
		useCustomWord = true
		selectedCategory = "food" // Default category for saving

		if save {
			if err := config.SaveCategoryWord(selectedCategory, word); err != nil {
				// Silent fail
			}
		}
	}

	// Step 4: Suffix Type Selection
	yellow.Printf("\nStep 4/4: Select Suffix Type\n")

	suffixItems := []string{
		"1. alphabetic (abc)",
		"2. numeric (123)",
		"3. mixed (a1b2)",
		"4. timestamp",
	}

	suffixPrompt := promptui.Select{
		Label: "",
		Items: suffixItems,
		Templates: &promptui.SelectTemplates{
			Active:   "❯ {{ . | cyan }}",
			Inactive: "  {{ . }}",
			Selected: "✓ {{ . | green }}",
		},
		Searcher: func(input string, index int) bool {
			item := suffixItems[index]
			return strings.Contains(strings.ToLower(item), strings.ToLower(input))
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
	for i := 0; i < count; i++ {
		var name string
		var err error

		if useCustomWord {
			// Use custom word directly
			name = fmt.Sprintf("%s-%s-%s-%s", selectedTechStack, selectedFramework, customWord, generateSimpleSuffix(suffixType))
		} else {
			// Generate name normally
			gen := generator.NewGenerator(generator.DefaultConfig())
			name, err = gen.GenerateEnhancedName(selectedTechStack, selectedFramework, selectedCategory, suffixType, 4)
			if err != nil {
				fmt.Printf("Error generating name: %v\n", err)
				continue
			}
		}

		// Create directory
		err = os.MkdirAll(name, 0755)
		if err != nil {
			fmt.Printf("❌ Failed to create directory '%s': %v\n", name, err)
			continue
		}

		fmt.Printf("\n")
		green.Printf("✓ Created: %s\n", name)
	}
}

// buildTechStackItems builds the list of tech stack items
func buildTechStackItems() []models.Item {
	items := []models.Item{}

	// Add built-in items
	techDescriptions := categories.TechStackDescriptions()
	for _, code := range categories.TechStackWords {
		items = append(items, models.Item{
			Code:        code,
			Description: techDescriptions[code],
			IsCustom:    false,
		})
	}

	// Add custom items from config
	cfg, _ := config.LoadConfig()
	for _, ts := range cfg.TechStacks {
		items = append(items, models.Item{
			Code:        ts.Code,
			Description: ts.Description,
			IsCustom:    false, // Already saved
		})
	}

	return items
}

// buildFrameworkItems builds the list of framework items for a tech stack
func buildFrameworkItems(techStack string) []models.Item {
	items := []models.Item{}

	// Add built-in frameworks
	frameworkDescriptions := categories.FrameworkDescriptions()
	availableFrameworks := categories.GetFrameworksForTech(techStack)

	for _, code := range availableFrameworks {
		items = append(items, models.Item{
			Code:        code,
			Description: frameworkDescriptions[code],
			IsCustom:    false,
		})
	}

	// Add custom frameworks from config
	cfg, _ := config.LoadConfig()
	if cfg.Frameworks[techStack] != nil {
		for _, fw := range cfg.Frameworks[techStack] {
			items = append(items, models.Item{
				Code:        fw.Code,
				Description: fw.Description,
				IsCustom:    false, // Already saved
			})
		}
	}

	return items
}

// generateSimpleSuffix generates a simple suffix for custom words
func generateSimpleSuffix(suffixType generator.SuffixType) string {
	switch suffixType {
	case generator.SuffixAlpha:
		return "abcd"
	case generator.SuffixNumeric:
		return "1234"
	case generator.SuffixTimestamp:
		return fmt.Sprintf("%d", os.Getpid()%10000)
	default:
		return "a1b2"
	}
}
