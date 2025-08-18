package build

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

func TestReadme_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Readme{})
}

func TestReadme_ExecWithDebug(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(c app_control.Control) {
		// Create test directory
		testDir, err := qt_file.MakeTestFolder("readme_debug", false)
		if err != nil {
			t.Fatalf("Failed to create test directory: %v", err)
		}
		defer func() {
			if err := os.RemoveAll(testDir); err != nil {
				t.Logf("Warning: Failed to clean up test directory: %v", err)
			}
		}()

		// Create readme instance
		readme := &Readme{
			Path: mo_path.NewFileSystemPath(filepath.Join(testDir, "README.txt")),
		}

		// Execute with debug info
		t.Logf("Starting readme generation in test mode")
		err = readme.Exec(c)
		if err != nil {
			t.Fatalf("Readme execution failed: %v", err)
		}

		t.Logf("Readme generation completed successfully")

		// Verify output when not in test mode
		if !c.Feature().IsTest() {
			// Check if file was created
			if _, err := os.Stat(readme.Path.Path()); os.IsNotExist(err) {
				t.Error("README file was not created")
			} else {
				t.Logf("README file created successfully at: %s", readme.Path.Path())
			}
		}
	})
}
