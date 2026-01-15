package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aravindcm49/dir-init/internal/generator"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	category     string
	suffixType   string
	suffixLength int
	count        int
	seed         int64
	outputFormat string
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&category, "category", "c", "all", "Category to use (tech, food, animals, pop, silly, dev, all)")
	generateCmd.Flags().StringVarP(&suffixType, "suffix", "s", "mixed", "Suffix type (alpha, numeric, mixed, timestamp)")
	generateCmd.Flags().IntVarP(&suffixLength, "length", "l", 4, "Suffix length (1-8)")
	generateCmd.Flags().IntVarP(&count, "count", "n", 1, "Number of names to generate")
	generateCmd.Flags().Int64VarP(&seed, "seed", "S", 0, "Random seed for reproducible results")
	generateCmd.Flags().StringVarP(&outputFormat, "output", "o", "text", "Output format (text, json)")
}

var generateCmd = &cobra.Command{
	Use:   "generate [flags]",
	Short: "Generate funny folder names",
	Long: `Generate funny folder names with specified categories and options.

Examples:
  dir-init generate -c tech
  dir-init generate -c food -n 5
  dir-init generate -c silly -s numeric -l 6
  dir-init generate -c all -n 10 -o json`,
	Run: func(cmd *cobra.Command, args []string) {
		config := generator.DefaultConfig()
		config.Category = category
		config.Count = count
		config.Seed = seed

		// Parse suffix type
		switch strings.ToLower(suffixType) {
		case "alpha":
			config.SuffixType = generator.SuffixAlpha
		case "numeric":
			config.SuffixType = generator.SuffixNumeric
		case "mixed":
			config.SuffixType = generator.SuffixMixed
		case "timestamp":
			config.SuffixType = generator.SuffixTimestamp
		default:
			fmt.Printf("Invalid suffix type: %s. Using default.\n", suffixType)
		}

		// Validate suffix length
		if suffixLength < 1 || suffixLength > 8 {
			fmt.Printf("Suffix length must be between 1 and 8. Using default 4.\n")
			config.SuffixLength = 4
		} else {
			config.SuffixLength = suffixLength
		}

		// Create generator and generate names
		generator := generator.NewGenerator(config)
		names := generator.Generate()

		// Output results
		switch strings.ToLower(outputFormat) {
		case "json":
			outputJSON(names)
		default:
			outputText(names)
		}
	},
}

func outputText(names []string) {
	fmt.Println()
	if len(names) == 1 {
		color.Green("Generated folder name: %s\n", names[0])
	} else {
		color.Green("Generated folder names:\n")
		for i, name := range names {
			fmt.Printf("%d. %s\n", i+1, name)
		}
		fmt.Println()
	}
}

func outputJSON(names []string) {
	type Output struct {
		Count int      `json:"count"`
		Names []string `json:"names"`
	}

	output := Output{
		Count: len(names),
		Names: names,
	}

	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Printf("Error generating JSON: %v\n", err)
		return
	}

	fmt.Println(string(jsonData))
}
