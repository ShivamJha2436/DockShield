package scanner

import (
	"dockshield/api/trivy"
	"encoding/json"
	"fmt"
)

// PrintReport displays only HIGH or CRITICAL vulnerabilities in CLI
func PrintReport(report *trivy.TrivyReport) {
	fmt.Println("📦 Image Scan Report")

	for _, result := range report.Results {
		fmt.Printf("\n🔍 Target: %s\n", result.Target)
		for _, vuln := range result.Vulnerabilities {
			if vuln.Severity == "HIGH" || vuln.Severity == "CRITICAL" {
				fmt.Printf("  [%s] %s - %s (%s)\n",
					vuln.Severity, vuln.VulnerabilityID, vuln.Title, vuln.PkgName)
			}
		}
	}
}

// FormatJSON returns the report in pretty JSON (for saving to file)
func FormatJSON(report *trivy.TrivyReport) ([]byte, error) {
	return json.MarshalIndent(report, "", "  ")
}
