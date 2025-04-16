package cmd

import (
	"dupefinder/internal/scanner"
	"encoding/json"
	"fmt"
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
			fmt.Println("Error:", err)
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
				fmt.Println("🎉 No duplicate files found.")
			} else {
				if err := saveResultsToFile(filtered, outputFile); err != nil {
					fmt.Println("Error writing output file:", err)
					os.Exit(1)
				}
				fmt.Printf("📝 %d duplicate group(s) written to %s\n", len(filtered), outputFile)
			}
			return
		}

		if len(filtered) == 0 {
			fmt.Println("🎉 No duplicate files found.")
			return
		}

		for hash, files := range filtered {
			fmt.Printf("Duplicate group (%s):\n", hash)
			for _, f := range files {
				fmt.Println("  ", f)
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
