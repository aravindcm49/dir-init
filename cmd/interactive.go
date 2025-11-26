package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aravindcm49/dir-init/cmd/tui/models"
	"github.com/aravindcm49/dir-init/internal/config"
	"github.com/aravindcm49/dir-init/internal/generator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color" // This line is kept as removing it would cause compilation errors due to its usage later in the file.
	"github.com/manifoldco/promptui"
)

func interactive() {
	green := color.New(color.FgGreen).Add(color.Bold)
	yellow := color.New(color.FgYellow).Add(color.Bold)

	// Clear screen once at start
	// fmt.Print("\033[H\033[2J")

	fmt.Println("========")
	fmt.Println("dir-init")
	fmt.Println("========")

	// Step 1: Frontend Selection
	yellow.Printf("Step 1/4: Select Frontend\n")

	frontendItems := buildFrontendItems()
	frontendModel := models.NewSelector("", frontendItems)

	p := tea.NewProgram(frontendModel)
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	m := finalModel.(models.SelectorModel)
	selectedFrontend := m.GetSelected()

	if selectedFrontend == nil {
		return
	}

	if selectedFrontend.IsCustom && m.ShouldSave() {
		if err := config.SaveFrontend(selectedFrontend.Code, selectedFrontend.Description); err != nil {
			// Silent fail
		}
	}

	selectedFrontendCode := selectedFrontend.Code

	// Reprint full line with selection
	fmt.Print("\033[A\033[K") // Move up and clear line
	fmt.Printf("Step 1/4: Select Frontend >> %s\n", selectedFrontendCode)

	// Step 2: Backend Selection
	yellow.Printf("Step 2/4: Select Backend\n")

	backendItems := buildBackendItems()
	backendModel := models.NewSelector("", backendItems)

	p = tea.NewProgram(backendModel)
	finalModel, err = p.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	bm := finalModel.(models.SelectorModel)
	selectedBackend := bm.GetSelected()

	if selectedBackend == nil {
		return
	}

	if selectedBackend.IsCustom && bm.ShouldSave() {
		if err := config.SaveBackend(selectedBackend.Code, selectedBackend.Description); err != nil {
			// Silent fail
		}
	}

	selectedBackendCode := selectedBackend.Code

	// Reprint full line with selection
	fmt.Print("\033[A\033[K") // Move up and clear line
	fmt.Printf("Step 2/4: Select Backend >> %s\n", selectedBackendCode)

	// Step 3: Category Selection
	yellow.Printf("Step 3/4: Select Category\n")

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	categoryItems := []models.Item{}
	for cat := range cfg.Categories {
		// Capitalize first letter for description
		desc := cat
		if len(cat) > 0 {
			desc = strings.ToUpper(cat[:1]) + cat[1:]
		}
		categoryItems = append(categoryItems, models.Item{
			Code:        cat,
			Description: desc,
		})
	}
	// Add "all" option
	categoryItems = append(categoryItems, models.Item{Code: "all", Description: "All Categories"})

	categoryModel := models.NewSelector("", categoryItems)
	p = tea.NewProgram(categoryModel)
	finalModel, err = p.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	cm := finalModel.(models.SelectorModel)
	selectedCat := cm.GetSelected()

	if selectedCat == nil {
		return
	}

	selectedCategory := selectedCat.Code
	var customWord string
	var useCustomWord bool

	// Handle custom word if needed
	if selectedCat.IsCustom {
		customWord = selectedCat.Code
		useCustomWord = true
		selectedCategory = "food" // Default

		if cm.ShouldSave() {
			if err := config.SaveCategoryWord(selectedCategory, customWord); err != nil {
				// Silent fail
			}
		}
	}

	// Reprint full line with selection
	fmt.Print("\033[A\033[K") // Move up and clear line
	fmt.Printf("Step 3/4: Select Category >> %s\n", selectedCategory)

	// Step 4: Suffix Type Selection
	yellow.Printf("Step 4/4: Select Suffix Type\n")

	suffixItems := []models.Item{
		{Code: "alphabetic", Description: "Alphabetic (abc)"},
		{Code: "numeric", Description: "Numeric (123)"},
		{Code: "mixed", Description: "Mixed (a1b2)"},
		{Code: "timestamp", Description: "Timestamp"},
	}

	suffixModel := models.NewSelector("", suffixItems)
	p = tea.NewProgram(suffixModel)
	finalModel, err = p.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	sm := finalModel.(models.SelectorModel)
	selectedSuf := sm.GetSelected()

	if selectedSuf == nil {
		return
	}

	var suffixType generator.SuffixType
	switch selectedSuf.Code {
	case "alphabetic":
		suffixType = generator.SuffixAlpha
	case "numeric":
		suffixType = generator.SuffixNumeric
	case "mixed":
		suffixType = generator.SuffixMixed
	case "timestamp":
		suffixType = generator.SuffixTimestamp
	}

	// Reprint full line with selection
	fmt.Print("\033[A\033[K") // Move up and clear line
	fmt.Printf("Step 4/4: Select Suffix Type >> %s\n", selectedSuf.Code)

	grey := color.New(color.FgHiBlack)

	countPrompt := promptui.Prompt{
		Label:   fmt.Sprintf("How many directories to create? %s", grey.Sprint("<default 1, enter number to change>")),
		Default: "1",
		Validate: func(input string) error {
			count, err := strconv.Atoi(input)
			if err != nil || count < 1 || count > 10 {
				return fmt.Errorf("please enter a number between 1 and 10")
			}
			return nil
		},
		Templates: &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Valid:   "{{ . }} ",
			Invalid: "{{ . }} ",
			Success: "{{ . }} ",
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
	// fmt.Println() // Single newline before output
	for i := 0; i < count; i++ {
		var name string
		var err error

		if useCustomWord {
			name = fmt.Sprintf("%s-%s-%s-%s", selectedFrontendCode, selectedBackendCode, customWord, generateSimpleSuffix(suffixType))
		} else {
			cfg, _ := config.LoadConfig()
			genConfig := generator.DefaultConfig()
			genConfig.Categories = cfg.Categories

			gen := generator.NewGenerator(genConfig)
			name, err = gen.GenerateEnhancedName(selectedFrontendCode, selectedBackendCode, selectedCategory, suffixType, 4)
			if err != nil {
				fmt.Printf("Error generating name: %v\n", err)
				continue
			}
		}

		err = os.MkdirAll(name, 0755)
		if err != nil {
			fmt.Printf("❌ Failed to create directory '%s': %v\n", name, err)
			continue
		}

		green.Printf("%s created!\n", name)
	}
}

func buildFrontendItems() []models.Item {
	items := []models.Item{}

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return items
	}

	for _, fe := range cfg.Frontends {
		items = append(items, models.Item{
			Code:        fe.Code,
			Description: fe.Description,
			IsCustom:    false,
		})
	}

	return items
}

func buildBackendItems() []models.Item {
	items := []models.Item{}

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return items
	}

	for _, be := range cfg.Backends {
		items = append(items, models.Item{
			Code:        be.Code,
			Description: be.Description,
			IsCustom:    false,
		})
	}

	return items
}

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
