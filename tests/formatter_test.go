package tests

import (
	"dockshield/api/trivy"
	"dockshield/core/scanner"
	"testing"
)

func TestPrintReport(t *testing.T) {
	report := &trivy.TrivyReport{
		Results: []trivy.Result{
			{
				Target: "nginx",
				Vulnerabilities: []trivy.Vulnerability{
					{VulnerabilityID: "CVE-123", Severity: "HIGH", Title: "Test Vulnerability", PkgName: "openssl"},
				},
			},
		},
	}
	scanner.PrintReport(report)
}
