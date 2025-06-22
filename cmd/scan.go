package cmd

import (
	"dockshield/internal/dockerutil"
	"dockshield/core/scanner"
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
		report, err := scanner.RunScan(image)
		if err != nil {
			fmt.Println("❌ Scan error:", err)
			return
		}
		scanner.PrintReport(report)

		if outputPath != "" {
			err := scanner.SaveReport(report, outputPath)
			if err != nil {
				fmt.Println("❌ Failed to save report:", err)
			}
		}
	},
}

func init() {
	scanCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output file to save scan results (JSON)")
}
