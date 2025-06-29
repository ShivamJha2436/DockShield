package trivy

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"
)

type Vulnerability struct {
	VulnerabilityID string `json:"VulnerabilityID"`
	Title           string `json:"Title"`
	Severity        string `json:"Severity"`
	PkgName         string `json:"PkgName"`
}

type Result struct {
	Target          string          `json:"Target"`
	Vulnerabilities []Vulnerability `json:"Vulnerabilities"`
}

type TrivyReport struct {
	Results []Result `json:"Results"`
}

// ScanImage runs the Trivy CLI tool and parses the output into TrivyReport
func ScanImage(image string) (*TrivyReport, error) {
	cmd := exec.Command("trivy", "image", "--format", "json", image)
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, errors.New(stderr.String())
	}

	var report TrivyReport
	err = json.Unmarshal(out.Bytes(), &report)
	if err != nil {
		return nil, err
	}

	return &report, nil
}
