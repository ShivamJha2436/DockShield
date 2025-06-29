package tests

import (
	"dockshield/api/trivy"
	"encoding/json"
	"testing"
)

func TestParseTrivyOutput(t *testing.T) {
	rawJSON := `{
		"Results": [{
			"Target": "nginx",
			"Vulnerabilities": [
				{"VulnerabilityID": "CVE-123", "Severity": "HIGH", "Title": "Test", "PkgName": "curl"}
			]
		}]
	}`

	var report trivy.TrivyReport
	err := json.Unmarshal([]byte(rawJSON), &report)
	if err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}
	if report.Results[0].Vulnerabilities[0].VulnerabilityID != "CVE-123" {
		t.Error("Unexpected vulnerability ID")
	}
}
