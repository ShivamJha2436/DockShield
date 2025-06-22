package cmd

import (
	"github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
	Use:   "dockshield",
	Short: "DockShield - Scan Docker images for vulnerabilities using Trivy",
	Long:  "DockShield is a CLI tool that scans Docker images for security vulnerabilities using Trivy.",
}
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
func init() {
	rootCmd.AddCommand(scanCmd)
}