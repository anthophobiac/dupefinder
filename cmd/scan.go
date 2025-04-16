package cmd

import (
	"dupefinder/internal/scanner"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var includeExt []string
var excludeExt []string

var scanCmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "Scan a directory for duplicate files",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		dupes, err := scanner.FindDuplicates(path, includeExt, excludeExt)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		found := false
		for hash, files := range dupes {
			if len(files) > 1 {
				found = true
				fmt.Printf("Duplicate group (%s):\n", hash)
				for _, f := range files {
					fmt.Println("  ", f)
				}
				fmt.Println()
			}
		}

		if !found {
			fmt.Println("🎉 No duplicate files found.")
		}
	},
}

func init() {
	scanCmd.Flags().StringSliceVarP(&includeExt, "include-ext", "i", nil, "Only include files with these extensions (e.g. .txt,.jpg)")
	scanCmd.Flags().StringSliceVarP(&excludeExt, "exclude-ext", "e", nil, "Exclude files with these extensions (e.g. .log,.tmp)")
	rootCmd.AddCommand(scanCmd)
}
