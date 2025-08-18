package build

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/watermint/toolbox/essentials/es_go/es_project"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"github.com/watermint/toolbox/resources"
)

func TestInfo_Preset(t *testing.T) {
	info := &Info{}
	info.Preset()
	// Preset does nothing, but we test it for coverage
}

func TestInfo_Exec_NoGitRepo(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(c app_control.Control) {
		// Create a temporary directory without git
		tempDir := t.TempDir()
		oldWd, _ := os.Getwd()
		defer os.Chdir(oldWd)

		if err := os.Chdir(tempDir); err != nil {
			t.Fatal(err)
		}

		info := &Info{}
		err := info.Exec(c)

		// Should fail because no git repository
		if err == nil {
			t.Error("Expected error when running outside git repository")
		}
	})
}

func TestInfo_Exec_MissingEnvVars(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(c app_control.Control) {
		// Save original env vars
		origBuilderKey := os.Getenv(app_definitions.EnvNameToolboxBuilderKey)
		origAppKeys := os.Getenv(app_definitions.EnvNameToolboxAppKeys)
		origLicenseSalt := os.Getenv(app_definitions.EnvNameToolboxLicenseSalt)

		// Unset env vars
		os.Unsetenv(app_definitions.EnvNameToolboxBuilderKey)
		os.Unsetenv(app_definitions.EnvNameToolboxAppKeys)
		os.Unsetenv(app_definitions.EnvNameToolboxLicenseSalt)

		defer func() {
			// Restore env vars
			if origBuilderKey != "" {
				os.Setenv(app_definitions.EnvNameToolboxBuilderKey, origBuilderKey)
			}
			if origAppKeys != "" {
				os.Setenv(app_definitions.EnvNameToolboxAppKeys, origAppKeys)
			}
			if origLicenseSalt != "" {
				os.Setenv(app_definitions.EnvNameToolboxLicenseSalt, origLicenseSalt)
			}
		}()

		// Should run in the actual project directory
		prjRoot, err := es_project.DetectRepositoryRoot()
		if err != nil {
			t.Skip("Not in a git repository")
		}

		// Change to project root
		oldWd, _ := os.Getwd()
		defer os.Chdir(oldWd)
		if err := os.Chdir(prjRoot); err != nil {
			t.Fatal(err)
		}

		info := &Info{FailFast: false}
		err = info.Exec(c)

		// Should succeed even without env vars when FailFast is false
		if err != nil {
			t.Errorf("Expected success without env vars when FailFast=false, got: %v", err)
		}

		// Verify info.json was created
		infoPath := filepath.Join(prjRoot, "resources/build", "info.json")
		if _, err := os.Stat(infoPath); os.IsNotExist(err) {
			t.Error("info.json was not created")
		}

		// Read and verify content
		data, err := os.ReadFile(infoPath)
		if err != nil {
			t.Fatal(err)
		}

		var buildInfo resources.BuildInfo
		if err := json.Unmarshal(data, &buildInfo); err != nil {
			t.Fatal(err)
		}

		// Verify production is false without env vars
		if buildInfo.Production {
			t.Error("Expected Production to be false without env vars")
		}

		// Verify empty values
		if buildInfo.Xap != "" {
			t.Error("Expected Xap to be empty without builder key")
		}
		if buildInfo.Zap != "" {
			t.Error("Expected Zap to be empty without app keys")
		}
		if buildInfo.LicenseSalt != "" {
			t.Error("Expected LicenseSalt to be empty without salt")
		}
	})
}

func TestInfo_Exec_FailFastMissingBuilderKey(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(c app_control.Control) {
		// Save and unset builder key
		origBuilderKey := os.Getenv(app_definitions.EnvNameToolboxBuilderKey)
		os.Unsetenv(app_definitions.EnvNameToolboxBuilderKey)
		defer func() {
			if origBuilderKey != "" {
				os.Setenv(app_definitions.EnvNameToolboxBuilderKey, origBuilderKey)
			}
		}()

		// Should run in the actual project directory
		prjRoot, err := es_project.DetectRepositoryRoot()
		if err != nil {
			t.Skip("Not in a git repository")
		}

		// Change to project root
		oldWd, _ := os.Getwd()
		defer os.Chdir(oldWd)
		if err := os.Chdir(prjRoot); err != nil {
			t.Fatal(err)
		}

		info := &Info{FailFast: true}
		err = info.Exec(c)

		// Should fail when FailFast is true and builder key is missing
		if err == nil {
			t.Error("Expected error when FailFast=true and builder key is missing")
		}
		if err != nil && err.Error() != "builder key not found" {
			t.Errorf("Expected 'builder key not found' error, got: %v", err)
		}
	})
}

func TestInfo_Exec_FailFastMissingAppKeys(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(c app_control.Control) {
		// Save and set builder key, unset app keys
		origBuilderKey := os.Getenv(app_definitions.EnvNameToolboxBuilderKey)
		origAppKeys := os.Getenv(app_definitions.EnvNameToolboxAppKeys)

		os.Setenv(app_definitions.EnvNameToolboxBuilderKey, "test-builder-key-12345")
		os.Unsetenv(app_definitions.EnvNameToolboxAppKeys)

		defer func() {
			if origBuilderKey != "" {
				os.Setenv(app_definitions.EnvNameToolboxBuilderKey, origBuilderKey)
			} else {
				os.Unsetenv(app_definitions.EnvNameToolboxBuilderKey)
			}
			if origAppKeys != "" {
				os.Setenv(app_definitions.EnvNameToolboxAppKeys, origAppKeys)
			}
		}()

		// Should run in the actual project directory
		prjRoot, err := es_project.DetectRepositoryRoot()
		if err != nil {
			t.Skip("Not in a git repository")
		}

		// Change to project root
		oldWd, _ := os.Getwd()
		defer os.Chdir(oldWd)
		if err := os.Chdir(prjRoot); err != nil {
			t.Fatal(err)
		}

		info := &Info{FailFast: true}
		err = info.Exec(c)

		// Should fail when FailFast is true and app keys are missing
		if err == nil {
			t.Error("Expected error when FailFast=true and app keys are missing")
		}
		if err != nil && err.Error() != "app key data not found" {
			t.Errorf("Expected 'app key data not found' error, got: %v", err)
		}
	})
}

func TestInfo_Exec_InvalidAppKeys(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(c app_control.Control) {
		// Save and set invalid app keys
		origAppKeys := os.Getenv(app_definitions.EnvNameToolboxAppKeys)
		os.Setenv(app_definitions.EnvNameToolboxAppKeys, "invalid-json-{")

		defer func() {
			if origAppKeys != "" {
				os.Setenv(app_definitions.EnvNameToolboxAppKeys, origAppKeys)
			} else {
				os.Unsetenv(app_definitions.EnvNameToolboxAppKeys)
			}
		}()

		// Should run in the actual project directory
		prjRoot, err := es_project.DetectRepositoryRoot()
		if err != nil {
			t.Skip("Not in a git repository")
		}

		// Change to project root
		oldWd, _ := os.Getwd()
		defer os.Chdir(oldWd)
		if err := os.Chdir(prjRoot); err != nil {
			t.Fatal(err)
		}

		info := &Info{FailFast: false}
		err = info.Exec(c)

		// Should fail with invalid JSON
		if err == nil {
			t.Error("Expected error with invalid JSON app keys")
		}
		if err != nil && err.Error() != "invalid app key data format" {
			t.Errorf("Expected 'invalid app key data format' error, got: %v", err)
		}
	})
}

func TestInfo_Exec_ShortBuilderKey(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(c app_control.Control) {
		// Save and set short builder key
		origBuilderKey := os.Getenv(app_definitions.EnvNameToolboxBuilderKey)
		os.Setenv(app_definitions.EnvNameToolboxBuilderKey, "short")

		defer func() {
			if origBuilderKey != "" {
				os.Setenv(app_definitions.EnvNameToolboxBuilderKey, origBuilderKey)
			} else {
				os.Unsetenv(app_definitions.EnvNameToolboxBuilderKey)
			}
		}()

		// Should run in the actual project directory
		prjRoot, err := es_project.DetectRepositoryRoot()
		if err != nil {
			t.Skip("Not in a git repository")
		}

		// Change to project root
		oldWd, _ := os.Getwd()
		defer os.Chdir(oldWd)
		if err := os.Chdir(prjRoot); err != nil {
			t.Fatal(err)
		}

		info := &Info{FailFast: false}
		err = info.Exec(c)

		// Should succeed but mark as not production ready
		if err != nil {
			t.Errorf("Should succeed with short builder key when FailFast=false, got: %v", err)
		}

		// Verify info.json was created
		infoPath := filepath.Join(prjRoot, "resources/build", "info.json")
		data, err := os.ReadFile(infoPath)
		if err != nil {
			t.Fatal(err)
		}

		var buildInfo resources.BuildInfo
		if err := json.Unmarshal(data, &buildInfo); err != nil {
			t.Fatal(err)
		}

		// Verify production is false with short key
		if buildInfo.Production {
			t.Error("Expected Production to be false with short builder key")
		}

		// Verify Xap is empty with short key
		if buildInfo.Xap != "" {
			t.Error("Expected Xap to be empty with short builder key")
		}
	})
}
