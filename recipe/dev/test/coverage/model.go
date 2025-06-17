package coverage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
	
	"github.com/watermint/toolbox/infra/control/app_workspace"
)

// CoverageData represents the coverage information for the entire project
type CoverageData struct {
	LastUpdate      string                     `json:"last_update"`
	TotalPackages   int                        `json:"total_packages"`
	TestedPackages  int                        `json:"tested_packages"`
	TotalStatements int                        `json:"total_statements"`
	CoveredStatements int                      `json:"covered_statements"`
	OverallCoverage float64                    `json:"overall_coverage"`
	Packages        map[string]*PackageData    `json:"packages"`
}

// PackageData represents coverage information for a single package
type PackageData struct {
	Package         string    `json:"package"`
	Coverage        float64   `json:"coverage"`
	Statements      int       `json:"statements"`
	CoveredStatements int     `json:"covered_statements"`
	NoTest          bool      `json:"no_test"`
	LastUpdate      string    `json:"last_update"`
	TestDuration    string    `json:"test_duration,omitempty"`
	Error           string    `json:"error,omitempty"`
	Impact          float64   `json:"-"` // Calculated field, not persisted
}

// LoadCoverageData loads the coverage data from test/coverage.json
func LoadCoverageData(ws app_workspace.Workspace) (*CoverageData, error) {
	// Get project root
	projectRoot := getProjectRoot(ws)
	coverageFile := filepath.Join(projectRoot, "test", "coverage.json")
	
	// Check if file exists
	if _, err := os.Stat(coverageFile); os.IsNotExist(err) {
		// Return empty data if file doesn't exist
		return &CoverageData{
			Packages: make(map[string]*PackageData),
		}, nil
	}
	
	// Read the file
	data, err := os.ReadFile(coverageFile)
	if err != nil {
		return nil, err
	}
	
	// Parse JSON
	var coverage CoverageData
	if err := json.Unmarshal(data, &coverage); err != nil {
		return nil, err
	}
	
	// Initialize map if nil
	if coverage.Packages == nil {
		coverage.Packages = make(map[string]*PackageData)
	}
	
	return &coverage, nil
}

// SaveCoverageData saves the coverage data to test/coverage.json
func SaveCoverageData(ws app_workspace.Workspace, data *CoverageData) error {
	// Get project root
	projectRoot := getProjectRoot(ws)
	testDir := filepath.Join(projectRoot, "test")
	coverageFile := filepath.Join(testDir, "coverage.json")
	
	// Create test directory if it doesn't exist
	if err := os.MkdirAll(testDir, 0755); err != nil {
		return err
	}
	
	// Update timestamp
	data.LastUpdate = time.Now().Format(time.RFC3339)
	
	// Marshal to JSON with indentation
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	
	// Write to file
	return os.WriteFile(coverageFile, jsonData, 0644)
}

// SaveCoverageReport saves the detailed coverage report to build/coverage_<timestamp>.out
func SaveCoverageReport(ws app_workspace.Workspace, coverageData []byte) (string, error) {
	// Get project root
	projectRoot := getProjectRoot(ws)
	buildDir := filepath.Join(projectRoot, "build")
	
	// Create build directory if it doesn't exist
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		return "", err
	}
	
	// Create filename with timestamp
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("coverage_%s.out", timestamp)
	filepath := filepath.Join(buildDir, filename)
	
	// Write coverage data
	if err := os.WriteFile(filepath, coverageData, 0644); err != nil {
		return "", err
	}
	
	return filepath, nil
}

// getProjectRoot returns the project root directory
func getProjectRoot(ws app_workspace.Workspace) string {
	// Start from current working directory, not the job workspace
	cwd, err := os.Getwd()
	if err != nil {
		// Fallback to workspace if we can't get cwd
		return ws.Job()
	}
	
	// Look for project root by finding go.mod
	current := cwd
	for {
		if _, err := os.Stat(filepath.Join(current, "go.mod")); err == nil {
			return current
		}
		
		parent := filepath.Dir(current)
		if parent == current {
			// Reached root, fallback to current directory
			return cwd
		}
		current = parent
	}
}

// CalculateOverallCoverage calculates the overall project coverage
func CalculateOverallCoverage(data *CoverageData) {
	totalStatements := 0
	coveredStatements := 0
	testedPackages := 0
	
	for _, pkg := range data.Packages {
		if !pkg.NoTest && pkg.Statements > 0 {
			testedPackages++
			totalStatements += pkg.Statements
			coveredStatements += pkg.CoveredStatements
		}
	}
	
	data.TotalPackages = len(data.Packages)
	data.TestedPackages = testedPackages
	data.TotalStatements = totalStatements
	data.CoveredStatements = coveredStatements
	
	if totalStatements > 0 {
		data.OverallCoverage = float64(coveredStatements) / float64(totalStatements) * 100
	} else {
		data.OverallCoverage = 0
	}
}

// GetPackagesSortedByCoverage returns packages sorted by coverage (lowest first)
func GetPackagesSortedByCoverage(data *CoverageData) []*PackageData {
	packages := make([]*PackageData, 0, len(data.Packages))
	for _, pkg := range data.Packages {
		packages = append(packages, pkg)
	}
	
	sort.Slice(packages, func(i, j int) bool {
		// No test packages first
		if packages[i].NoTest && !packages[j].NoTest {
			return true
		}
		if !packages[i].NoTest && packages[j].NoTest {
			return false
		}
		// Then by coverage
		return packages[i].Coverage < packages[j].Coverage
	})
	
	return packages
}