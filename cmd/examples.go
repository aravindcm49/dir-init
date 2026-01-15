package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	exampleCategory string
	exampleCount    int
)

func init() {
	rootCmd.AddCommand(examplesCmd)

	examplesCmd.Flags().StringVarP(&exampleCategory, "category", "c", "", "Show examples for specific category")
	examplesCmd.Flags().IntVarP(&exampleCount, "count", "n", 3, "Number of examples to show")
}

var examplesCmd = &cobra.Command{
	Use:   "examples [flags]",
	Short: "Show example folder names",
	Long: `Show example folder names for demonstration.

Examples:
  dir-init examples
  dir-init examples -c tech -n 5
  dir-init examples -c silly`,
	Run: func(cmd *cobra.Command, args []string) {
		printExamples()
	},
}

func printExamples() {
	fmt.Println()
	color.Cyan("Example Folder Names:\n")

	if exampleCategory != "" {
		printCategoryExamples(exampleCategory, exampleCount)
	} else {
		printAllExamples(exampleCount)
	}

	fmt.Println()
	color.Green("Try 'dir-init generate -c <category>' to create your own names!")
}

func printCategoryExamples(category string, count int) {
	color.Yellow("Category: %s\n", category)

	examples := getCategoryExamples(category)

	for i := 0; i < min(count, len(examples)); i++ {
		fmt.Printf("• %s\n", examples[i])
	}
}

func printAllExamples(count int) {
	categories := []string{"tech", "food", "animals", "pop", "silly", "dev"}

	for _, category := range categories {
		color.Yellow("%s:\n", strings.Title(category))
		examples := getCategoryExamples(category)

		for i := 0; i < min(min(count, len(examples)), 2); i++ {
			fmt.Printf("  • %s\n", examples[i])
		}
		if len(examples) > 2 {
			fmt.Printf("  • ...and %d more\n", len(examples)-2)
		}
		fmt.Println()
	}
}

func getCategoryExamples(category string) []string {
	examples := map[string][]string{
		"tech": {
			"code-1234",
			"debug-abc2",
			"api-v2beta",
			"service-x42",
			"binary-test",
		},
		"food": {
			"pizza-fresh",
			"burger-delicious",
			"taco-hot",
			"pasta-cold",
			"donut-sweet",
		},
		"animals": {
			"penguin-cute",
			"koala-gentle",
			"dolphin-smart",
			"eagle-mighty",
			"turtle-slow",
		},
		"pop": {
			"ninja-epic",
			"wizard-magical",
			"knight-brave",
			"astronaut-space",
			"robot-future",
		},
		"silly": {
			"potato-silly",
			"banana-goofy",
			"unicorn-odd",
			"noodle-weird",
			"pickle-strange",
		},
		"dev": {
			"github-v1",
			"docker-prod",
			"react-ui",
			"python-dev",
			"k8s-cluster",
		},
	}

	if ex, exists := examples[category]; exists {
		return ex
	}

	return []string{"example-1", "example-2", "example-3"}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
