package scanner

import (
	"dockshield/api/trivy"
)

// RunScan simply executes the Trivy scan logic from api layer
func RunScan(image string) (*trivy.TrivyReport, error) {
	return trivy.ScanImage(image)
}

