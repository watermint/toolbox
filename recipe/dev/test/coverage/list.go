package coverage

import (
	"fmt"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"

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

	// Run go test with coverage
	ui.Info(z.MsgRunningCoverage.With("M", app_msg.Raw(fmt.Sprintf("Running coverage analysis (threshold: %d%%)...", z.Threshold))))

	cmd := exec.Command("go", "test", "-coverprofile=coverage.out", "./...")
	output, err := cmd.CombinedOutput()
	if err != nil {
		l.Debug("Coverage command failed", esl.Error(err), esl.String("output", string(output)))
		// Continue processing even if some tests fail
	}

	// Parse coverage output
	packages := z.parseCoverageOutput(string(output))

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

	return nil
}

func (z *List) parseCoverageOutput(output string) []PackageCoverage {
	packages := []PackageCoverage{}
	lines := strings.Split(output, "\n")

	// Patterns to match coverage output
	coveragePattern := regexp.MustCompile(`^(ok|FAIL)\s+(\S+)\s+.*coverage:\s+(\d+\.\d+)%\s+of\s+statements`)
	noTestPattern := regexp.MustCompile(`^\?\s+(\S+)\s+\[no test files\]`)
	noStatementsPattern := regexp.MustCompile(`^(ok|FAIL)\s+(\S+)\s+.*coverage:\s+\[no statements\]`)
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