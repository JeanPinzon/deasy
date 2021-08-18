package cmd

import (
	"embed"

	"github.com/spf13/cobra"
)

var ApiTemplateFs embed.FS

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "deasy",
	Short: "CLI to easily manage dev tools",
	Long: `Deasy is a CLI to helps developers to create and manage usual things, like apis, forms and projects setups.
It helps to keep the projects updated and let you to handle complex problems. deasy manages the simple things for you.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
