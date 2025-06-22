package scanner

import (
	"dockshield/api/trivy"
)

func RunScan(image string) ([]byte, error) {
	return trivy.ScanImage(image)
}
