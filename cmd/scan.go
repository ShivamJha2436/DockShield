package cmd

import (
	"dockShield/api/trivy"
	"dockShield/core/scanner"
	"dockShield/internal/output"
	"fmt"

	"github.com/spf13/cobra"
)

var outputPath string

var scanCmd = &cobra.Command{
	Use:   "scan [image]",
	Short: "Scan a Docker image for vulnerabilities",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		image := args[0]

		// Step 1: Run Trivy scan
		report, err := trivy.ScanImage(image)
		if err != nil {
			fmt.Println("❌ Failed to scan:", err)
			return
		}

		// Step 2: Print to terminal
		scanner.PrintReport(report)

		// Step 3: Optionally write to file
		if outputPath != "" {
			data, err := scanner.FormatJSON(report)
			if err != nil {
				fmt.Println("❌ Could not format JSON:", err)
				return
			}
			err = output.WriteToFile(data, outputPath)
			if err != nil {
				fmt.Println("❌ Failed to write report:", err)
				return
			}
			fmt.Println("✅ Report saved to", outputPath)
		}
	},
}

func init() {
	scanCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Save JSON output to file")
}
