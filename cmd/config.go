package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/aravindcm49/dir-init/internal/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configPathCmd)
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configValidateCmd)
	configCmd.AddCommand(configEditCmd)
	configCmd.AddCommand(configAddCmd)
	configCmd.AddCommand(configRemoveCmd)

	// Add subcommands for config add
	configAddCmd.AddCommand(configAddTechStackCmd)
	configAddCmd.AddCommand(configAddFrameworkCmd)
	configAddCmd.AddCommand(configAddWordCmd)

	// Add subcommands for config remove
	configRemoveCmd.AddCommand(configRemoveTechStackCmd)
	configRemoveCmd.AddCommand(configRemoveFrameworkCmd)
	configRemoveCmd.AddCommand(configRemoveWordCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage custom word collections",
	Long:  `Manage your custom word collections for tech stacks, frameworks, and categories.`,
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config file with examples",
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.InitConfig(); err != nil {
			color.Red("❌ Error: %v\n", err)
			return
		}
		color.Green("✓ Config file created at: %s\n", config.GetConfigPath())
		fmt.Println("  Format: {nickname}-{techstack}-{categoryword}-{suffix}")
	},
}

var configPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Show config file path",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.GetConfigPath())
	},
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show loaded collections",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			color.Red("❌ Error loading config: %v\n", err)
			return
		}

		green := color.New(color.FgGreen).Add(color.Bold)
		yellow := color.New(color.FgYellow).Add(color.Bold)

		// Show tech stacks
		if len(cfg.TechStacks) > 0 {
			green.Println("\nTech Stacks:")
			for _, ts := range cfg.TechStacks {
				fmt.Printf("  • %s - %s\n", ts.Code, ts.Description)
			}
		}

		// Show frameworks
		if len(cfg.Frameworks) > 0 {
			green.Println("\nFrameworks:")
			for techStack, frameworks := range cfg.Frameworks {
				yellow.Printf("  %s:\n", techStack)
				for _, fw := range frameworks {
					fmt.Printf("    • %s - %s\n", fw.Code, fw.Description)
				}
			}
		}

		// Show categories
		if len(cfg.Categories) > 0 {
			green.Println("\nCategory Words:")
			for category, words := range cfg.Categories {
				yellow.Printf("  %s:\n", category)
				for _, word := range words {
					fmt.Printf("    • %s\n", word)
				}
			}
		}

		if len(cfg.TechStacks) == 0 && len(cfg.Categories) == 0 {
			color.Yellow("No custom collections found. Run 'dir-init config init' to create an example config.\n")
		}
	},
}

var configValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate config syntax",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetConfigPath()

		// Check if file exists
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			color.Yellow("⚠️  Config file not found at: %s\n", configPath)
			return
		}

		// Try to load config
		_, err := config.LoadConfig()
		if err != nil {
			color.Red("❌ Config validation failed: %v\n", err)
			return
		}

		color.Green("✓ Config file is valid\n")
	},
}

var configEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit config in default editor",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetConfigPath()

		// Get editor from environment or use default
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "vim"
		}

		// Open editor
		editorCmd := exec.Command(editor, configPath)
		editorCmd.Stdin = os.Stdin
		editorCmd.Stdout = os.Stdout
		editorCmd.Stderr = os.Stderr

		if err := editorCmd.Run(); err != nil {
			color.Red("❌ Error opening editor: %v\n", err)
			return
		}
	},
}

var configAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add items to config",
	Long:  `Add tech stacks, frameworks, or words to your config.`,
}

var configAddTechStackCmd = &cobra.Command{
	Use:   "techstack <code> <description>",
	Short: "Add a tech stack",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		code := args[0]
		description := args[1]

		if err := config.SaveTechStack(code, description); err != nil {
			color.Red("❌ Error: %v\n", err)
			return
		}

		color.Green("✓ Added tech stack: %s - %s\n", code, description)
	},
}

var configAddFrameworkCmd = &cobra.Command{
	Use:   "framework <techstack> <code> <description>",
	Short: "Add a framework",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		techStack := args[0]
		code := args[1]
		description := args[2]

		if err := config.SaveFramework(techStack, code, description); err != nil {
			color.Red("❌ Error: %v\n", err)
			return
		}

		color.Green("✓ Added framework: %s - %s (for %s)\n", code, description, techStack)
	},
}

var configAddWordCmd = &cobra.Command{
	Use:   "word <category> <word>",
	Short: "Add a word to a category",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		category := args[0]
		word := args[1]

		if err := config.SaveCategoryWord(category, word); err != nil {
			color.Red("❌ Error: %v\n", err)
			return
		}

		color.Green("✓ Added word '%s' to category '%s'\n", word, category)
	},
}

var configRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove items from config",
	Long:  `Remove tech stacks, frameworks, or words from your config.`,
}

var configRemoveTechStackCmd = &cobra.Command{
	Use:   "techstack <code>",
	Short: "Remove a tech stack",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		code := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			color.Red("❌ Error loading config: %v\n", err)
			return
		}

		// Find and remove tech stack
		found := false
		newTechStacks := []config.TechStack{}
		for _, ts := range cfg.TechStacks {
			if ts.Code != code {
				newTechStacks = append(newTechStacks, ts)
			} else {
				found = true
			}
		}

		if !found {
			color.Yellow("⚠️  Tech stack '%s' not found\n", code)
			return
		}

		cfg.TechStacks = newTechStacks
		if err := config.SaveConfig(cfg); err != nil {
			color.Red("❌ Error saving config: %v\n", err)
			return
		}

		color.Green("✓ Removed tech stack: %s\n", code)
	},
}

var configRemoveFrameworkCmd = &cobra.Command{
	Use:   "framework <techstack> <code>",
	Short: "Remove a framework",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		techStack := args[0]
		code := args[1]

		cfg, err := config.LoadConfig()
		if err != nil {
			color.Red("❌ Error loading config: %v\n", err)
			return
		}

		// Check if tech stack exists
		if cfg.Frameworks[techStack] == nil {
			color.Yellow("⚠️  No frameworks found for tech stack '%s'\n", techStack)
			return
		}

		// Find and remove framework
		found := false
		newFrameworks := []config.Framework{}
		for _, fw := range cfg.Frameworks[techStack] {
			if fw.Code != code {
				newFrameworks = append(newFrameworks, fw)
			} else {
				found = true
			}
		}

		if !found {
			color.Yellow("⚠️  Framework '%s' not found for tech stack '%s'\n", code, techStack)
			return
		}

		cfg.Frameworks[techStack] = newFrameworks
		if err := config.SaveConfig(cfg); err != nil {
			color.Red("❌ Error saving config: %v\n", err)
			return
		}

		color.Green("✓ Removed framework: %s (from %s)\n", code, techStack)
	},
}

var configRemoveWordCmd = &cobra.Command{
	Use:   "word <category> <word>",
	Short: "Remove a word from a category",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		category := args[0]
		word := args[1]

		cfg, err := config.LoadConfig()
		if err != nil {
			color.Red("❌ Error loading config: %v\n", err)
			return
		}

		// Check if category exists
		if cfg.Categories[category] == nil {
			color.Yellow("⚠️  Category '%s' not found\n", category)
			return
		}

		// Find and remove word
		found := false
		newWords := []string{}
		for _, w := range cfg.Categories[category] {
			if w != word {
				newWords = append(newWords, w)
			} else {
				found = true
			}
		}

		if !found {
			color.Yellow("⚠️  Word '%s' not found in category '%s'\n", word, category)
			return
		}

		cfg.Categories[category] = newWords
		if err := config.SaveConfig(cfg); err != nil {
			color.Red("❌ Error saving config: %v\n", err)
			return
		}

		color.Green("✓ Removed word '%s' from category '%s'\n", word, category)
	},
}
