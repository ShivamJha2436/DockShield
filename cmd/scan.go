package cmd

import (
	"dockshield/core/scanner"
	"fmt"
	"github.com/spf13/cobra"
)

var output string

var scanCmd = &cobra.Command{
	Use:   "scan [image]",
	Short: "Scan a Docker image for vulnerabilities",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		image := args[0]
		report, err := scanner.RunTrivyScan(image)
		if err != nil {
			fmt.Println("❌ Scan error:", err)
			return
		}
		scanner.PrintReport(report)
	},
}

func init() {
	scanCmd.Flags().StringVarP(&output, "output", "o", "", "Save output to file")
}
