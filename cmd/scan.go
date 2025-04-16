package cmd

import (
	"dupefinder/internal/scanner"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"os"

	"github.com/spf13/cobra"
)

var includeExt []string
var excludeExt []string
var outputFile string

var scanCmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "Scan a directory for duplicate files",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		dupes, err := scanner.FindDuplicates(path, includeExt, excludeExt)
		if err != nil {
			color.Red("Error:", err)
			os.Exit(1)
		}

		filtered := make(map[string][]string)
		for hash, files := range dupes {
			if len(files) > 1 {
				filtered[hash] = files
			}
		}

		if outputFile != "" {
			if len(filtered) == 0 {
				color.Green("🎉 No duplicate files found.")
			} else {
				if err := saveResultsToFile(filtered, outputFile); err != nil {
					color.Red("Error writing output file: %v", err)
					os.Exit(1)
				}
				color.Green("📝 %d duplicate group(s) written to %s", len(filtered), outputFile)
			}
			return
		}

		if len(filtered) == 0 {
			color.Green("🎉 No duplicate files found.")
			return
		}

		for hash, files := range filtered {
			fmt.Printf("Duplicate group (%s):\n", hash)
			for _, f := range files {
				color.Cyan("  %s", f)
			}
			fmt.Println()
		}
	},
}

func init() {
	scanCmd.Flags().StringSliceVarP(&includeExt, "include-ext", "i", nil, "Only include files with these extensions (e.g. .txt,.jpg)")
	scanCmd.Flags().StringSliceVarP(&excludeExt, "exclude-ext", "e", nil, "Exclude files with these extensions (e.g. .log,.tmp)")
	scanCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Write duplicate results to a JSON file")
	rootCmd.AddCommand(scanCmd)
}

func saveResultsToFile(results map[string][]string, filename string) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
