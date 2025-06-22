package scanner

import (
	"encoding/json"
)

type Vulnerability struct {
	VulnerabilityID string `json:"VulnerabilityID"`
	Title           string `json:"Title"`
	Severity        string `json:"Severity"`
	PkgName         string `json:"PkgName"`
}

type Result struct {
	Target          string           `json:"Target"`
	Vulnerabilities []Vulnerability `json:"Vulnerabilities"`
}

type TrivyReport struct {
	Results []Result `json:"Results"`
}

func ParseReport(data []byte) (*TrivyReport, error) {
	var report TrivyReport
	err := json.Unmarshal(data, &report)
	return &report, err
}
