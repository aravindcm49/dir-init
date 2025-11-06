package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	enableInteractive bool
	avoidInteractive  bool
)

var rootCmd = &cobra.Command{
	Use:   "dir-init",
	Short: "Generate funny randomized folder names",
	Long: `dir-init is a CLI tool that generates funny, randomized folder names
with customizable categories and alphanumeric suffixes.

Interactive Mode (default):
- Tech stack selection (fejs, bepy, etc.)
- Framework selection
- Category selection (food, animals, pop, silly, dev)
- Suffix type selection

It comes with multiple categories of funny names including:
- Food & Cooking
- Animals & Nature
- Pop Culture
- Silly & Absurd
- Developer-related

Perfect for adding some humor to your development workflow!`,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		if !avoidInteractive {
			interactive()
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&avoidInteractive, "no-interactive", false, "Skip interactive mode")
	rootCmd.PersistentFlags().BoolVarP(&enableInteractive, "interactive", "i", false, "Start interactive mode (overrides --no-interactive)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}