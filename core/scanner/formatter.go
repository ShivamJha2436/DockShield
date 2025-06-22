package scanner

import (
	"encoding/json"
	"fmt"
)

func PrintReport(report *TrivyReport) {
	for _, result := range report.Results {
		fmt.Println("📦 Image:", result.Target)
		for _, v := range result.Vulnerabilities {
			if v.Severity == "HIGH" || v.Severity == "CRITICAL" {
				fmt.Printf("  🔥 [%s] %s - %s (%s)\n", v.Severity, v.VulnerabilityID, v.Title, v.PkgName)
			}
		}
	}
}

func FormatJSON(report *TrivyReport) ([]byte, error) {
	return json.MarshalIndent(report, "", "  ")
}
