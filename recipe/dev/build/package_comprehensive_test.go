package build

import (
	"archive/zip"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

func TestPackage_Preset(t *testing.T) {
	p := &Package{}
	p.Preset()

	if p.ExecutableName != app_definitions.ExecutableName {
		t.Errorf("Expected ExecutableName to be %s, got %s", app_definitions.ExecutableName, p.ExecutableName)
	}
}

func TestPackage_platformName(t *testing.T) {
	p := &Package{}

	testCases := []struct {
		envValue string
		expected string
	}{
		{"windows/amd64", "win"},
		{"linux/amd64", "linux-intel"},
		{"linux/arm64", "linux-arm"},
		{"darwin/amd64", "mac-intel"},
		{"darwin/arm64", "mac-applesilicon"},
		{"unknown/unknown", "unknown"},
		{"", "unknown"},
	}

	for _, tc := range testCases {
		if tc.envValue == "" {
			os.Unsetenv(app_definitions.EnvNameToolboxBuildTarget)
		} else {
			os.Setenv(app_definitions.EnvNameToolboxBuildTarget, tc.envValue)
		}

		result := p.platformName()
		if result != tc.expected {
			t.Errorf("For env value '%s', expected platform name '%s', got '%s'", tc.envValue, tc.expected, result)
		}
	}

	// Clean up
	os.Unsetenv(app_definitions.EnvNameToolboxBuildTarget)
}

func TestPackage_createPackage(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(c app_control.Control) {
		// Create test directories
		buildDir, err := qt_file.MakeTestFolder("build", false)
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(buildDir)

		distDir, err := qt_file.MakeTestFolder("dist", false)
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(distDir)

		// Create test binary
		testBinary := filepath.Join(buildDir, app_definitions.ExecutableName)
		err = os.WriteFile(testBinary, []byte("test binary content"), 0755)
		if err != nil {
			t.Fatal(err)
		}

		p := &Package{
			BuildPath:      mo_path.NewExistingFileSystemPath(buildDir),
			DistPath:       mo_path.NewFileSystemPath(distDir),
			ExecutableName: app_definitions.ExecutableName,
		}

		// Test package creation
		pkgPath, err := p.createPackage(c)
		if err != nil {
			t.Fatal(err)
		}

		// Verify package was created
		if _, err := os.Stat(pkgPath); os.IsNotExist(err) {
			t.Error("Package file was not created")
		}

		// Verify package contents
		reader, err := zip.OpenReader(pkgPath)
		if err != nil {
			t.Fatal(err)
		}
		defer reader.Close()

		expectedFiles := map[string]bool{
			"LICENSE.txt":                  false,
			"README.txt":                   false,
			app_definitions.ExecutableName: false,
		}

		for _, f := range reader.File {
			if _, ok := expectedFiles[f.Name]; ok {
				expectedFiles[f.Name] = true
			}
		}

		for name, found := range expectedFiles {
			if !found {
				t.Errorf("Expected file '%s' not found in package", name)
			}
		}
	})
}

func TestPackage_createPackage_withWindowsTarget(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(c app_control.Control) {
		// Set Windows target
		os.Setenv(app_definitions.EnvNameToolboxBuildTarget, "windows/amd64")
		defer os.Unsetenv(app_definitions.EnvNameToolboxBuildTarget)

		// Create test directories
		buildDir, err := qt_file.MakeTestFolder("build", false)
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(buildDir)

		distDir, err := qt_file.MakeTestFolder("dist", false)
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(distDir)

		// Create test binary with Windows naming
		testBinary := filepath.Join(buildDir, app_definitions.ExecutableName+"-windows-amd64.exe")
		err = os.WriteFile(testBinary, []byte("test binary content"), 0755)
		if err != nil {
			t.Fatal(err)
		}

		p := &Package{
			BuildPath:      mo_path.NewExistingFileSystemPath(buildDir),
			DistPath:       mo_path.NewFileSystemPath(distDir),
			ExecutableName: app_definitions.ExecutableName,
		}

		// Test package creation
		pkgPath, err := p.createPackage(c)
		if err != nil {
			t.Fatal(err)
		}

		// Verify package was created with correct name
		if _, err := os.Stat(pkgPath); os.IsNotExist(err) {
			t.Error("Package file was not created")
		}

		// Check that package name contains "win"
		if !strings.Contains(pkgPath, "-win.") {
			t.Errorf("Package name should contain '-win.' for Windows platform, got: %s", pkgPath)
		}
	})
}

func TestPackage_createPackage_invalidDistPath(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(c app_control.Control) {
		buildDir, err := qt_file.MakeTestFolder("build", false)
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(buildDir)

		// Use an invalid dist path
		p := &Package{
			BuildPath:      mo_path.NewExistingFileSystemPath(buildDir),
			DistPath:       mo_path.NewFileSystemPath("/invalid/path/that/cannot/be/created"),
			ExecutableName: app_definitions.ExecutableName,
		}

		_, err = p.createPackage(c)
		if err == nil {
			t.Error("Expected error for invalid dist path")
		}
	})
}

func TestPackage_Exec_withoutDeploy(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(c app_control.Control) {
		// Create test directories
		buildDir, err := qt_file.MakeTestFolder("build", false)
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(buildDir)

		distDir, err := qt_file.MakeTestFolder("dist", false)
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(distDir)

		// Create test binary
		testBinary := filepath.Join(buildDir, app_definitions.ExecutableName)
		err = os.WriteFile(testBinary, []byte("test binary content"), 0755)
		if err != nil {
			t.Fatal(err)
		}

		p := &Package{
			BuildPath:      mo_path.NewExistingFileSystemPath(buildDir),
			DistPath:       mo_path.NewFileSystemPath(distDir),
			ExecutableName: app_definitions.ExecutableName,
			DeployPath:     mo_string.NewOptional(""), // Empty deploy path
		}

		// Execute should succeed without deployment
		err = p.Exec(c)
		if err != nil {
			t.Fatal(err)
		}

		// Verify package was created
		files, err := os.ReadDir(distDir)
		if err != nil {
			t.Fatal(err)
		}

		if len(files) == 0 {
			t.Error("No package files created")
		}
	})
}

func TestPackage_binaryNaming(t *testing.T) {
	testCases := []struct {
		target         string
		expectedSuffix string
		expectedName   string
	}{
		{"windows/amd64", ".exe", app_definitions.ExecutableName + "-windows-amd64.exe"},
		{"linux/amd64", "", app_definitions.ExecutableName + "-linux-amd64"},
		{"linux/arm64", "", app_definitions.ExecutableName + "-linux-arm64"},
		{"darwin/amd64", "", app_definitions.ExecutableName + "-darwin-amd64"},
		{"darwin/arm64", "", app_definitions.ExecutableName + "-darwin-arm64"},
	}

	for _, tc := range testCases {
		t.Run(tc.target, func(t *testing.T) {
			os.Setenv(app_definitions.EnvNameToolboxBuildTarget, tc.target)
			defer os.Unsetenv(app_definitions.EnvNameToolboxBuildTarget)

			qtr_endtoend.TestWithControl(t, func(c app_control.Control) {
				buildDir, err := qt_file.MakeTestFolder("build", false)
				if err != nil {
					t.Fatal(err)
				}
				defer os.RemoveAll(buildDir)

				distDir, err := qt_file.MakeTestFolder("dist", false)
				if err != nil {
					t.Fatal(err)
				}
				defer os.RemoveAll(distDir)

				// Create test binary with expected name
				testBinary := filepath.Join(buildDir, tc.expectedName)
				err = os.WriteFile(testBinary, []byte("test binary content"), 0755)
				if err != nil {
					t.Fatal(err)
				}

				p := &Package{
					BuildPath:      mo_path.NewExistingFileSystemPath(buildDir),
					DistPath:       mo_path.NewFileSystemPath(distDir),
					ExecutableName: app_definitions.ExecutableName,
				}

				_, err = p.createPackage(c)
				if err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}
