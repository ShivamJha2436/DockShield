package trivy

import (
	"dockshield/core/scanner"
	"dockshield/core/scanner/parser"
	"os/exec"
)

func ScanImage(image string) (*scanner.TrivyReport, error) {
	cmd := exec.Command("trivy", "image", "--format", "json", image)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return parser.ParseReport(out)
}
