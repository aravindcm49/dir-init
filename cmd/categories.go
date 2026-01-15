package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "List all available categories",
	Long: `List all available categories with description and word count.

Examples:
  dir-init categories`,
	Run: func(cmd *cobra.Command, args []string) {
		printCategories()
	},
}

func init() {
	rootCmd.AddCommand(categoriesCmd)
}

func printCategories() {
	fmt.Println()
	color.Cyan("Available Categories:\n")

	categories := map[string]struct {
		Description string
		WordCount   int
	}{
		"tech": {
			Description: "Technology & programming related words",
			WordCount:   len(getTechWords()),
		},
		"food": {
			Description: "Food, cooking and beverage words",
			WordCount:   len(getFoodWords()),
		},
		"animals": {
			Description: "Animals and nature words",
			WordCount:   len(getAnimalWords()),
		},
		"pop": {
			Description: "Pop culture, fantasy and creative arts words",
			WordCount:   len(getPopWords()),
		},
		"silly": {
			Description: "Silly, funny and absurd words",
			WordCount:   len(getSillyWords()),
		},
		"dev": {
			Description: "Development tools and programming words",
			WordCount:   len(getDevWords()),
		},
		"all": {
			Description: "All categories combined",
			WordCount:   getTotalWordCount(),
		},
	}

	for name, info := range categories {
		fmt.Printf("%s%s%s", color.BlueString("• "), color.YellowString(name), color.BlueString(" - "))
		fmt.Printf("%s", info.Description)
		if name != "all" {
			fmt.Printf(" (%d words)", info.WordCount)
		}
		fmt.Println()

		if name != "all" {
			fmt.Printf("  %s\n", color.New(color.FgWhite).Sprint("Example: tech-1234, food-goofy, animals-abc2"))
		}
	}

	fmt.Println()
	color.Green("Use 'dir-init generate -c <category>' to generate names from a specific category.")
}

// Helper functions to get word counts (simplified for now)
func getTechWords() []string {
	// This would normally import from the categories package
	// For now, return a placeholder count
	return make([]string, 100) // Approximately 100 tech words
}

func getFoodWords() []string {
	return make([]string, 80) // Approximately 80 food words
}

func getAnimalWords() []string {
	return make([]string, 120) // Approximately 120 animal words
}

func getPopWords() []string {
	return make([]string, 90) // Approximately 90 pop words
}

func getSillyWords() []string {
	return make([]string, 110) // Approximately 110 silly words
}

func getDevWords() []string {
	return make([]string, 95) // Approximately 95 dev words
}

func getTotalWordCount() int {
	tech := len(getTechWords())
	food := len(getFoodWords())
	animals := len(getAnimalWords())
	pop := len(getPopWords())
	silly := len(getSillyWords())
	dev := len(getDevWords())

	return tech + food + animals + pop + silly + dev
}
