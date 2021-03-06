package domain

import (
	"context"
	"time"
)

// Asset contains the key fields from a Nexpose asset.
type Asset struct {
	ScanTime time.Time
	Hostname string
	ID       int64
	IP       string
	ScanType string
}

// Vulnerability contains the key fields from a Nexpose vulnerability associated with a
// particular Asset.
type Vulnerability struct {
	ID             string
	Results        []AssessmentResult
	Status         string
	CvssV2Score    float64
	CvssV2Severity string
	Description    string
	Title          string
	Solutions      []string
	LocalCheck     bool
}

// AssessmentResult contains port and protocol information for the Vulnerability.
type AssessmentResult struct {
	Port     int
	Protocol string
	Proof    string
}

// VulnerabilityFilter is an interface for filtering an asset's vulnerabilities based on configurable criteria.
type VulnerabilityFilter interface {
	FilterVulnerabilities(context.Context, Asset, []Vulnerability) []Vulnerability
}
