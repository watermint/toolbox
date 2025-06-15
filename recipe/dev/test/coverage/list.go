package coverage

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type List struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkTransient
	Threshold              int
	MinPackage             int
	MaxPackage             int
	CoverageReport         rp_model.RowReport
	SummaryReport          rp_model.RowReport
	MsgRunningCoverage     app_msg.Message
	MsgLowCoveragePackages app_msg.Message
	MsgSummary             app_msg.Message
	MsgRecommendation      app_msg.Message
	MsgSavedCoverage       app_msg.Message
}

type PackageCoverage struct {
	Package    string
	Coverage   float64
	Statements int
	NoTest     bool
}

type CoverageReport struct {
	Package    string  `json:"package"`
	Coverage   float64 `json:"coverage"`
	Statements int     `json:"statements"`
	NoTest     bool    `json:"no_test"`
}

func (z *List) Exec(c app_control.Control) error {
	l := c.Log()
	ui := c.UI()

	// Open reports
	if err := z.CoverageReport.Open(); err != nil {
		return err
	}
	if err := z.SummaryReport.Open(); err != nil {
		return err
	}

	// Load existing coverage data
	coverageData, err := LoadCoverageData(c.Workspace())
	if err != nil {
		l.Debug("Unable to load existing coverage data", esl.Error(err))
		coverageData = &CoverageData{
			Packages: make(map[string]*PackageData),
		}
	}

	// Get project root for coverage file
	projectRoot := getProjectRoot(c.Workspace())
	buildDir := filepath.Join(projectRoot, "build")
	coverageFile := filepath.Join(buildDir, "coverage.out")
	
	// Ensure build directory exists
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		l.Debug("Unable to create build directory", esl.Error(err))
	}
	
	// Run go test with coverage
	ui.Info(z.MsgRunningCoverage.With("M", app_msg.Raw(fmt.Sprintf("Running coverage analysis (threshold: %d%%)...", z.Threshold))))

	startTime := time.Now()
	cmd := exec.Command("go", "test", fmt.Sprintf("-coverprofile=%s", coverageFile), "./...")
	output, err := cmd.CombinedOutput()
	duration := time.Since(startTime)

	if err != nil {
		l.Debug("Coverage command failed", esl.Error(err), esl.String("output", string(output)))
		// Continue processing even if some tests fail
	}

	// Save the coverage report to build directory
	if coverageBytes, err := os.ReadFile(coverageFile); err == nil {
		if savedPath, err := SaveCoverageReport(c.Workspace(), coverageBytes); err == nil {
			ui.Info(z.MsgSavedCoverage.With("Path", savedPath))
		} else {
			l.Debug("Unable to save coverage report", esl.Error(err))
		}
	}

	// Parse coverage output and coverage profile
	packages := z.parseCoverageOutput(string(output))
	if _, err := os.Stat(coverageFile); err == nil {
		packages = z.enrichWithCoverageProfile(packages, coverageFile)
	}

	// Update coverage data
	for _, pkg := range packages {
		coveredStmts := int(float64(pkg.Statements) * pkg.Coverage / 100)
		coverageData.Packages[pkg.Package] = &PackageData{
			Package:           pkg.Package,
			Coverage:          pkg.Coverage,
			Statements:        pkg.Statements,
			CoveredStatements: coveredStmts,
			NoTest:            pkg.NoTest,
			LastUpdate:        time.Now().Format(time.RFC3339),
			TestDuration:      duration.String(),
		}
	}

	// Calculate overall coverage
	CalculateOverallCoverage(coverageData)

	// Save updated coverage data
	if err := SaveCoverageData(c.Workspace(), coverageData); err != nil {
		l.Debug("Unable to save coverage data", esl.Error(err))
	}

	// Sort by coverage (ascending)
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

	// Write all packages to coverage report
	for _, pkg := range packages {
		z.CoverageReport.Row(&CoverageReport{
			Package:    pkg.Package,
			Coverage:   pkg.Coverage,
			Statements: pkg.Statements,
			NoTest:     pkg.NoTest,
		})
	}

	// Find packages below threshold
	lowCoveragePackages := []PackageCoverage{}
	for _, pkg := range packages {
		if pkg.Coverage < float64(z.Threshold) {
			lowCoveragePackages = append(lowCoveragePackages, pkg)
		}
	}

	// Limit to requested number of packages
	displayCount := len(lowCoveragePackages)
	if displayCount > z.MaxPackage {
		displayCount = z.MaxPackage
	}
	if displayCount < z.MinPackage && len(lowCoveragePackages) >= z.MinPackage {
		displayCount = z.MinPackage
	}

	// Display and write summary
	ui.Info(z.MsgLowCoveragePackages.With("M", app_msg.Raw(fmt.Sprintf("\nPackages with coverage below %d%%:", z.Threshold))))
	ui.Info(app_msg.Raw(strings.Repeat("-", 80)))

	for i := 0; i < displayCount && i < len(lowCoveragePackages); i++ {
		pkg := lowCoveragePackages[i]
		z.SummaryReport.Row(&CoverageReport{
			Package:    pkg.Package,
			Coverage:   pkg.Coverage,
			Statements: pkg.Statements,
			NoTest:     pkg.NoTest,
		})

		status := fmt.Sprintf("%.1f%%", pkg.Coverage)
		if pkg.NoTest {
			status = "NO TESTS"
		}
		ui.Info(app_msg.Raw(fmt.Sprintf("%-60s %10s", pkg.Package, status)))
	}

	ui.Info(app_msg.Raw(strings.Repeat("-", 80)))
	ui.Info(z.MsgSummary.With("M", app_msg.Raw(fmt.Sprintf("\nTotal packages below threshold: %d", len(lowCoveragePackages)))))
	ui.Info(z.MsgRecommendation.With("M", app_msg.Raw(fmt.Sprintf("Showing top %d packages that need test coverage improvements", displayCount))))
	ui.Info(app_msg.Raw(fmt.Sprintf("\nOverall project coverage: %.1f%% (%d/%d statements)",
		coverageData.OverallCoverage,
		coverageData.CoveredStatements,
		coverageData.TotalStatements)))

	// Clean up temporary coverage file
	os.Remove(coverageFile)

	return nil
}

func (z *List) parseCoverageOutput(output string) []PackageCoverage {
	packages := []PackageCoverage{}
	lines := strings.Split(output, "\n")

	// Patterns to match coverage output
	coveragePattern := regexp.MustCompile(`^(ok|FAIL)\s+(\S+)\s+.*coverage:\s+(\d+\.\d+)%\s+of\s+statements`)
	noTestPattern := regexp.MustCompile(`^\?\s+(\S+)\s+\[no test files\]`)
	noStatementsPattern := regexp.MustCompile(`^(ok|FAIL)\s+(\S+)\s+.*coverage:\s+\[no statements\]`)
	noValuePattern := regexp.MustCompile(`^(ok|FAIL)\s+(\S+)\s+.*coverage:\s+<no value>%`)
	noCoveragePattern := regexp.MustCompile(`^(\S+)\s+coverage:\s+(\d+\.\d+)%\s+of\s+statements`)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Match coverage with percentage
		if matches := coveragePattern.FindStringSubmatch(line); matches != nil {
			coverage, _ := strconv.ParseFloat(matches[3], 64)
			packages = append(packages, PackageCoverage{
				Package:    matches[2],
				Coverage:   coverage,
				Statements: 1, // We don't have exact count from this format
				NoTest:     false,
			})
			continue
		}

		// Match no test files
		if matches := noTestPattern.FindStringSubmatch(line); matches != nil {
			packages = append(packages, PackageCoverage{
				Package:    matches[1],
				Coverage:   0.0,
				Statements: 0,
				NoTest:     true,
			})
			continue
		}

		// Match no statements
		if matches := noStatementsPattern.FindStringSubmatch(line); matches != nil {
			packages = append(packages, PackageCoverage{
				Package:    matches[2],
				Coverage:   100.0, // No statements means 100% coverage
				Statements: 0,
				NoTest:     false,
			})
			continue
		}

		// Match no value coverage
		if matches := noValuePattern.FindStringSubmatch(line); matches != nil {
			packages = append(packages, PackageCoverage{
				Package:    matches[2],
				Coverage:   0.0, // No value means 0% coverage
				Statements: 0,
				NoTest:     false,
			})
			continue
		}

		// Match standalone coverage (no ok/FAIL prefix)
		if matches := noCoveragePattern.FindStringSubmatch(line); matches != nil {
			coverage, _ := strconv.ParseFloat(matches[2], 64)
			packages = append(packages, PackageCoverage{
				Package:    matches[1],
				Coverage:   coverage,
				Statements: 1,
				NoTest:     false,
			})
			continue
		}
	}

	return packages
}

func (z *List) enrichWithCoverageProfile(packages []PackageCoverage, coverageFile string) []PackageCoverage {
	// Read the coverage profile
	profileBytes, err := os.ReadFile(coverageFile)
	if err != nil {
		return packages
	}

	// Parse the coverage profile to get accurate statement counts
	packageStats := make(map[string]*PackageStats)
	
	lines := strings.Split(string(profileBytes), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "mode:") {
			continue
		}
		
		// Parse line format: package/file.go:startLine.startCol,endLine.endCol numStmt covered
		parts := strings.Fields(line)
		if len(parts) != 3 {
			continue
		}
		
		filePath := parts[0]
		numStmt, _ := strconv.Atoi(parts[1])
		covered, _ := strconv.Atoi(parts[2])
		
		// Extract package from file path
		pkgName := extractPackageFromPath(filePath)
		if pkgName == "" {
			continue
		}
		
		if packageStats[pkgName] == nil {
			packageStats[pkgName] = &PackageStats{
				Package: pkgName,
			}
		}
		
		packageStats[pkgName].TotalStatements += numStmt
		if covered > 0 {
			packageStats[pkgName].CoveredStatements += numStmt
		}
	}
	
	// Create a map for quick lookup of existing packages
	packageMap := make(map[string]*PackageCoverage)
	for i := range packages {
		packageMap[packages[i].Package] = &packages[i]
	}
	
	// Update existing packages and add new ones from profile
	for pkgName, stats := range packageStats {
		if stats.TotalStatements == 0 {
			continue
		}
		
		coverage := float64(stats.CoveredStatements) / float64(stats.TotalStatements) * 100
		
		if existingPkg, exists := packageMap[pkgName]; exists {
			// Update existing package with accurate counts
			existingPkg.Statements = stats.TotalStatements
			existingPkg.Coverage = coverage
		} else {
			// Add new package found in profile
			newPkg := PackageCoverage{
				Package:    pkgName,
				Coverage:   coverage,
				Statements: stats.TotalStatements,
				NoTest:     false,
			}
			packages = append(packages, newPkg)
		}
	}
	
	return packages
}

type PackageStats struct {
	Package           string
	TotalStatements   int
	CoveredStatements int
}

func extractPackageFromPath(filePath string) string {
	// Extract package name from file path like "github.com/watermint/toolbox/package/file.go"
	lastSlash := strings.LastIndex(filePath, "/")
	if lastSlash == -1 {
		return ""
	}
	
	pkgPath := filePath[:lastSlash]
	return pkgPath
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &List{}, func(r rc_recipe.Recipe) {
		m := r.(*List)
		m.Threshold = 30
		m.MinPackage = 5
		m.MaxPackage = 10
	})
}

func (z *List) Preset() {
	z.Threshold = 50
	z.MinPackage = 10
	z.MaxPackage = 30
	z.CoverageReport.SetModel(&CoverageReport{})
	z.SummaryReport.SetModel(&CoverageReport{})
}
