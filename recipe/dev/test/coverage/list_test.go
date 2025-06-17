package coverage

import (
	"testing"
)

func TestList_parseCoverageOutput(t *testing.T) {
	c := &List{}
	
	testOutput := `ok  	github.com/watermint/toolbox	1.209s	coverage: 84.6% of statements
ok  	github.com/watermint/toolbox/catalogue	0.906s	coverage: 92.9% of statements
?   	github.com/watermint/toolbox/essentials/api/api_client	[no test files]
ok  	github.com/watermint/toolbox/essentials/api/api_auth_basic_test	1.571s	coverage: [no statements]
	github.com/watermint/toolbox/essentials/api/api_auth		coverage: 0.0% of statements
FAIL	github.com/watermint/toolbox/some/package	1.234s	coverage: 45.5% of statements`

	packages := c.parseCoverageOutput(testOutput)

	if len(packages) != 6 {
		t.Errorf("Expected 6 packages, got %d", len(packages))
	}

	// Test coverage parsing
	expectedCoverage := map[string]float64{
		"github.com/watermint/toolbox":                            84.6,
		"github.com/watermint/toolbox/catalogue":                  92.9,
		"github.com/watermint/toolbox/essentials/api/api_client": 0.0,
		"github.com/watermint/toolbox/essentials/api/api_auth_basic_test": 100.0,
		"github.com/watermint/toolbox/essentials/api/api_auth":    0.0,
		"github.com/watermint/toolbox/some/package":               45.5,
	}

	for _, pkg := range packages {
		if expected, ok := expectedCoverage[pkg.Package]; ok {
			if pkg.Coverage != expected {
				t.Errorf("Package %s: expected coverage %.1f%%, got %.1f%%", 
					pkg.Package, expected, pkg.Coverage)
			}
		}
	}

	// Test NoTest flag
	for _, pkg := range packages {
		if pkg.Package == "github.com/watermint/toolbox/essentials/api/api_client" {
			if !pkg.NoTest {
				t.Errorf("Package %s should have NoTest=true", pkg.Package)
			}
		}
	}
}

func TestList_Exec(t *testing.T) {
	// Test is handled by the Test() method in the main file
}