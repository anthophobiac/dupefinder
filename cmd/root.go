package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dupefinder",
	Short: "Find duplicate files in a directory",
	Long:  `Recursively scan directories and detect duplicate files by comparing their hashes.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use `dupefinder scan <path>` to begin scanning for duplicates.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
