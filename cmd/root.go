package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dir-init",
	Short: "Generate funny randomized folder names",
	Long: `dir-init is a CLI tool that generates funny, randomized folder names
with customizable categories and alphanumeric suffixes.

It comes with multiple categories of funny names including:
- Tech & Programming
- Food & Cooking
- Animals & Nature
- Pop Culture
- Silly & Absurd
- Developer-related

Perfect for adding some humor to your development workflow!`,
	Version: "1.0.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}