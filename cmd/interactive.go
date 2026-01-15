package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aravindcm49/dir-init/cmd/tui/models"
	"github.com/aravindcm49/dir-init/internal/config"
	"github.com/aravindcm49/dir-init/internal/generator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color" // This line is kept as removing it would cause compilation errors due to its usage later in the file.
	"golang.org/x/term"
)

func interactive(verbose bool) {
	green := color.New(color.FgGreen).Add(color.Bold)
	yellow := color.New(color.FgYellow).Add(color.Bold)

	if verbose {
		fmt.Println("[verbose] Starting interactive mode")
	}

	// Clear screen once at start
	// fmt.Print("\033[H\033[2J")

	fmt.Println("========")
	fmt.Println("dir-init")
	fmt.Println("========")

	// Step 1: Nickname Input
	yellow.Printf("Step 1/4: Enter Nickname\n")

	nicknameModel := models.NewInput("", "Let's start with a nick name!")

	p := tea.NewProgram(nicknameModel)
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	nm := finalModel.(models.SelectorModel)
	selectedNicknameItem := nm.GetSelected()
	if selectedNicknameItem == nil {
		return
	}

	selectedNickname := strings.TrimSpace(selectedNicknameItem.Code)
	if selectedNickname == "" {
		return
	}

	if verbose {
		fmt.Printf("[verbose] Nickname: %s\n", selectedNickname)
	}

	// Reprint full line with selection
	fmt.Print("\033[A\033[K") // Move up and clear line
	fmt.Printf("Step 1/4: Enter Nickname >> %s\n", selectedNickname)

	// Step 2: Tech Stack Selection
	yellow.Printf("Step 2/4: Select Tech Stack\n")

	techStackItems := buildTechStackItems()
	techStackModel := models.NewSelector("", techStackItems)

	p = tea.NewProgram(techStackModel)
	finalModel, err = p.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	tsm := finalModel.(models.SelectorModel)
	selectedTechStack := tsm.GetSelected()

	if selectedTechStack == nil {
		return
	}

	if selectedTechStack.IsCustom && tsm.ShouldSave() {
		if err := config.SaveTechStack(selectedTechStack.Code, selectedTechStack.Description); err != nil {
			// Silent fail
		}
	}

	selectedTechStackCode := selectedTechStack.Code

	if verbose {
		fmt.Printf("[verbose] Selected tech stack: %s\n", selectedTechStackCode)
	}

	// Reprint full line with selection
	fmt.Print("\033[A\033[K") // Move up and clear line
	fmt.Printf("Step 2/4: Select Tech Stack >> %s\n", selectedTechStackCode)

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
	if verbose {
		fmt.Printf("[verbose] Selected category: %s\n", selectedCategory)
	}

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

	if verbose {
		fmt.Printf("[verbose] Selected suffix type: %s\n", selectedSuf.Code)
	}

	// Reprint full line with selection
	fmt.Print("\033[A\033[K") // Move up and clear line
	fmt.Printf("Step 4/4: Select Suffix Type >> %s\n", selectedSuf.Code)

	grey := color.New(color.FgHiBlack)

	// Arrow/jk controlled count input
	countHint := "<use ↑/k to increase, ↓/j to decrease, Enter to confirm>: "
	countPrompt := fmt.Sprintf("How many directories to create? %s", grey.Sprint(countHint))
	count := readArrowCount(1, 10, countPrompt)
	fmt.Printf("\r\x1b[2KHow many directories to create? %d\n", count)

	if verbose {
		fmt.Printf("[verbose] Selected count: %d\n", count)
	}

	// Generate and create directories
	// fmt.Println() // Single newline before output
	for i := 0; i < count; i++ {
		var name string
		var err error

		if useCustomWord {
			name = fmt.Sprintf("%s-%s-%s-%s", selectedNickname, selectedTechStackCode, customWord, generateSimpleSuffix(suffixType))
		} else {
			cfg, _ := config.LoadConfig()
			genConfig := generator.DefaultConfig()
			genConfig.Categories = cfg.Categories

			gen := generator.NewGenerator(genConfig)
			name, err = gen.GenerateEnhancedName(selectedNickname, selectedTechStackCode, selectedCategory, suffixType, 4)
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

		if verbose {
			fmt.Printf("[verbose] Created directory: %s\n", name)
		}

		green.Printf("%s created!\n", name)
	}
}

func buildTechStackItems() []models.Item {
	items := []models.Item{}

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return items
	}

	for _, ts := range cfg.TechStacks {
		items = append(items, models.Item{
			Code:        ts.Code,
			Description: ts.Description,
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

// readArrowCount reads arrow keys (or j/k) to increment/decrement a count value
func readArrowCount(min, max int, prompt string) int {
	// Save current terminal settings
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Failed to set raw mode:", err)
		return min
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	scanner := bufio.NewReader(os.Stdin)
	count := min

	render := func() {
		fmt.Printf("\r\x1b[2K%s%d", prompt, count)
	}
	render()

	for {
		// Read first byte
		b, err := scanner.ReadByte()
		if err != nil {
			return count
		}

		switch b {
		// Check for escape sequence (arrow keys start with 27)
		case 27:
			// Read next two bytes of escape sequence
			b2, _ := scanner.ReadByte()
			b3, _ := scanner.ReadByte()

			// \x1b[A = Up, \x1b[B = Down
			if b2 == 91 && b3 == 65 { // Up
				if count < max {
					count++
				}
				render()
			} else if b2 == 91 && b3 == 66 { // Down
				if count > min {
					count--
				}
				render()
			}

		case 'k':
			if count < max {
				count++
			}
			render()

		case 'j':
			if count > min {
				count--
			}
			render()

		case 13, 10: // Enter key
			return count
		}
	}
}
