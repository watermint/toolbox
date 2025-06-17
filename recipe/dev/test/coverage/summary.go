package coverage

import (
	"fmt"
	"strings"

	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Summary struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkTransient
	SuggestCount           int
	RecommendationReport   rp_model.RowReport
	MsgOverallCoverage     app_msg.Message
	MsgPackageStats        app_msg.Message
	MsgRecommendations     app_msg.Message
	MsgNoCoverageData      app_msg.Message
	MsgTargetCoverage      app_msg.Message
}

type RecommendationReport struct {
	Priority   int     `json:"priority"`
	Package    string  `json:"package"`
	Coverage   float64 `json:"coverage"`
	Statements int     `json:"statements"`
	Impact     float64 `json:"impact"`
	NoTest     bool    `json:"no_test"`
}

func (z *Summary) Exec(c app_control.Control) error {
	ui := c.UI()

	// Open report
	if err := z.RecommendationReport.Open(); err != nil {
		return err
	}

	// Load coverage data
	coverageData, err := LoadCoverageData(c.Workspace())
	if err != nil {
		ui.Error(z.MsgNoCoverageData)
		return err
	}

	// Check if we have any data
	if len(coverageData.Packages) == 0 {
		ui.Error(z.MsgNoCoverageData)
		return fmt.Errorf("no coverage data found. Please run 'dev test coverage list' first")
	}

	// Display overall coverage
	ui.Info(z.MsgOverallCoverage.With("Coverage", fmt.Sprintf("%.1f%%", coverageData.OverallCoverage)))
	ui.Info(app_msg.Raw(strings.Repeat("=", 80)))
	ui.Info(app_msg.Raw(fmt.Sprintf("Total packages: %d", coverageData.TotalPackages)))
	ui.Info(app_msg.Raw(fmt.Sprintf("Tested packages: %d", coverageData.TestedPackages)))
	ui.Info(app_msg.Raw(fmt.Sprintf("Coverage: %.1f%% (%d/%d statements)",
		coverageData.OverallCoverage,
		coverageData.CoveredStatements,
		coverageData.TotalStatements)))
	ui.Info(app_msg.Raw(strings.Repeat("=", 80)))

	// Calculate target coverage (50%)
	targetCoverage := 50.0
	requiredStatements := int(float64(coverageData.TotalStatements) * targetCoverage / 100)
	statementsNeeded := requiredStatements - coverageData.CoveredStatements

	if coverageData.OverallCoverage >= targetCoverage {
		ui.Success(z.MsgTargetCoverage.With("Target", fmt.Sprintf("%.0f%%", targetCoverage)))
		return nil
	}

	ui.Info(z.MsgTargetCoverage.With("Target", fmt.Sprintf("%.0f%%", targetCoverage)))
	ui.Info(app_msg.Raw(fmt.Sprintf("Statements needed to reach target: %d", statementsNeeded)))

	// Get sorted packages
	sortedPackages := GetPackagesSortedByCoverage(coverageData)

	// Calculate recommendations based on impact
	recommendations := z.calculateRecommendations(sortedPackages, statementsNeeded)

	// Display recommendations
	ui.Info(z.MsgRecommendations)
	ui.Info(app_msg.Raw(strings.Repeat("-", 80)))
	ui.Info(app_msg.Raw(fmt.Sprintf("%-5s %-60s %10s %10s %10s", "Pri", "Package", "Coverage", "Statements", "Impact")))
	ui.Info(app_msg.Raw(strings.Repeat("-", 80)))

	displayCount := z.SuggestCount
	if displayCount > len(recommendations) {
		displayCount = len(recommendations)
	}

	for i := 0; i < displayCount; i++ {
		rec := recommendations[i]

		// Write to report
		z.RecommendationReport.Row(&RecommendationReport{
			Priority:   i + 1,
			Package:    rec.Package,
			Coverage:   rec.Coverage,
			Statements: rec.Statements,
			Impact:     rec.Impact,
			NoTest:     rec.NoTest,
		})

		// Display
		status := fmt.Sprintf("%.1f%%", rec.Coverage)
		if rec.NoTest {
			status = "NO TESTS"
		}
		impact := fmt.Sprintf("%.1f%%", rec.Impact)

		ui.Info(app_msg.Raw(fmt.Sprintf("%-5d %-60s %10s %10d %10s",
			i+1,
			z.truncatePackageName(rec.Package, 60),
			status,
			rec.Statements,
			impact)))
	}

	ui.Info(app_msg.Raw(strings.Repeat("-", 80)))

	// Show next steps
	if displayCount > 0 {
		ui.Info(app_msg.Raw("\nNext steps:"))
		for i := 0; i < displayCount && i < 3; i++ {
			ui.Info(app_msg.Raw(fmt.Sprintf("%d. Run: go run . dev test coverage pkg -package %s",
				i+1, recommendations[i].Package)))
		}
	}

	return nil
}

func (z *Summary) calculateRecommendations(packages []*PackageData, statementsNeeded int) []*PackageData {
	// Calculate impact for each package
	// Impact = potential statements that could be covered / total statements needed
	for _, pkg := range packages {
		if pkg.Statements > 0 && pkg.Coverage < 100 {
			// Potential statements = statements that are not covered
			uncoveredStatements := pkg.Statements - pkg.CoveredStatements
			pkg.Impact = float64(uncoveredStatements) / float64(statementsNeeded) * 100
		} else {
			pkg.Impact = 0
		}
	}

	// Sort by impact (highest first) for packages with low coverage
	recommendations := make([]*PackageData, 0)
	
	// First, add packages with no tests
	for _, pkg := range packages {
		if pkg.NoTest && pkg.Statements > 0 {
			recommendations = append(recommendations, pkg)
		}
	}

	// Then add packages with coverage < 50%, sorted by impact
	lowCoveragePackages := make([]*PackageData, 0)
	for _, pkg := range packages {
		if !pkg.NoTest && pkg.Coverage < 50 && pkg.Statements > 0 {
			lowCoveragePackages = append(lowCoveragePackages, pkg)
		}
	}

	// Sort by impact
	for i := 0; i < len(lowCoveragePackages)-1; i++ {
		for j := i + 1; j < len(lowCoveragePackages); j++ {
			if lowCoveragePackages[i].Impact < lowCoveragePackages[j].Impact {
				lowCoveragePackages[i], lowCoveragePackages[j] = lowCoveragePackages[j], lowCoveragePackages[i]
			}
		}
	}

	recommendations = append(recommendations, lowCoveragePackages...)

	return recommendations
}

func (z *Summary) truncatePackageName(name string, maxLen int) string {
	if len(name) <= maxLen {
		return name
	}
	
	// Try to intelligently truncate
	parts := strings.Split(name, "/")
	if len(parts) > 3 {
		// Keep first part and last 2 parts
		return parts[0] + "/.../" + parts[len(parts)-2] + "/" + parts[len(parts)-1]
	}
	
	// Simple truncate
	return name[:maxLen-3] + "..."
}

func (z *Summary) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Summary{}, func(r rc_recipe.Recipe) {
		m := r.(*Summary)
		m.SuggestCount = 5
	})
}

func (z *Summary) Preset() {
	z.SuggestCount = 10
	z.RecommendationReport.SetModel(&RecommendationReport{})
}