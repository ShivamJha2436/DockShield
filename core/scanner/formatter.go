package scanner

import (
	"dockShield/api/trivy"
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func PrintReport(report *trivy.TrivyReport) {
	fmt.Println("📦 Image Scan Report")

	for _, result := range report.Results {
		fmt.Printf("\n🔍 Target: %s\n\n", result.Target)

		// Create data slice
		var rows [][]string
		for _, vuln := range result.Vulnerabilities {
			if vuln.Severity == "HIGH" || vuln.Severity == "CRITICAL" {
				rows = append(rows, []string{
					vuln.Severity,
					vuln.VulnerabilityID,
					vuln.PkgName,
					truncate(vuln.Title, 60),
				})
			}
		}

		// Render the table
		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{"Severity", "CVE ID", "Package", "Title"})
		table.Bulk(rows)
		table.Render()
	}
}

func truncate(s string, max int) string {
	if len(s) > max {
		return s[:max-3] + "..."
	}
	return s
}

func FormatJSON(report *trivy.TrivyReport) ([]byte, error) {
	return json.MarshalIndent(report, "", "  ")
}
