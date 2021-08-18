package cmd

import (
	"github.com/spf13/cobra"
)

var apisCmd = &cobra.Command{
	Use:   "apis",
	Short: "Manage REST APIs",
}

func init() {
	rootCmd.AddCommand(apisCmd)
}
