package release

import (
	"strings"
	"testing"

	"github.com/watermint/toolbox/domain/github/model/mo_release_asset"
	"github.com/watermint/toolbox/infra/control/app_definitions"
)

func TestAssetPlatformConstants(t *testing.T) {
	// Test that constants are defined as expected
	expectedConstants := map[string]string{
		"AssetPlatformUnknown":      "unknown",
		"AssetPlatformMacIntel":     "mac-intel",
		"AssetPlatformMacArm":       "mac-arm",
		"AssetPlatformLinuxIntel":   "linux-intel",
		"AssetPlatformLinuxArm":     "linux-arm",
		"AssetPlatformWindowsIntel": "win-intel",
	}

	actualConstants := map[string]string{
		"AssetPlatformUnknown":      AssetPlatformUnknown,
		"AssetPlatformMacIntel":     AssetPlatformMacIntel,
		"AssetPlatformMacArm":       AssetPlatformMacArm,
		"AssetPlatformLinuxIntel":   AssetPlatformLinuxIntel,
		"AssetPlatformLinuxArm":     AssetPlatformLinuxArm,
		"AssetPlatformWindowsIntel": AssetPlatformWindowsIntel,
	}

	for name, expected := range expectedConstants {
		if actual, ok := actualConstants[name]; !ok || actual != expected {
			t.Errorf("Constant %s: expected '%s', got '%s'", name, expected, actual)
		}
	}
}

func TestIdentifyPlatform(t *testing.T) {
	testCases := []struct {
		name     string
		asset    *mo_release_asset.Asset
		expected string
	}{
		// Mac Intel variants
		{
			name:     "Mac Intel - mac-intel.zip",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-mac-intel.zip"},
			expected: AssetPlatformMacIntel,
		},
		{
			name:     "Mac Intel - mac-amd64.zip",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-mac-amd64.zip"},
			expected: AssetPlatformMacIntel,
		},
		{
			name:     "Mac Intel - mac-x86_64.zip",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-mac-x86_64.zip"},
			expected: AssetPlatformMacIntel,
		},

		// Mac ARM variants
		{
			name:     "Mac ARM - mac-applesilicon.zip",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-mac-applesilicon.zip"},
			expected: AssetPlatformMacArm,
		},
		{
			name:     "Mac ARM - mac-arm64.zip",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-mac-arm64.zip"},
			expected: AssetPlatformMacArm,
		},
		{
			name:     "Mac ARM - mac-arm.zip",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-mac-arm.zip"},
			expected: AssetPlatformMacArm,
		},

		// Linux Intel variants
		{
			name:     "Linux Intel - linux-intel.zip",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-linux-intel.zip"},
			expected: AssetPlatformLinuxIntel,
		},
		{
			name:     "Linux Intel - linux-amd64.zip",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-linux-amd64.zip"},
			expected: AssetPlatformLinuxIntel,
		},
		{
			name:     "Linux Intel - linux-x86_64.zip",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-linux-x86_64.zip"},
			expected: AssetPlatformLinuxIntel,
		},

		// Linux ARM variants
		{
			name:     "Linux ARM - linux-arm.zip",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-linux-arm.zip"},
			expected: AssetPlatformLinuxArm,
		},
		{
			name:     "Linux ARM - linux-arm64.zip",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-linux-arm64.zip"},
			expected: AssetPlatformLinuxArm,
		},

		// Windows variants
		{
			name:     "Windows - win.zip",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-win.zip"},
			expected: AssetPlatformWindowsIntel,
		},
		{
			name:     "Windows - win-intel.zip",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-win-intel.zip"},
			expected: AssetPlatformWindowsIntel,
		},
		{
			name:     "Windows - win-amd64.zip",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-win-amd64.zip"},
			expected: AssetPlatformWindowsIntel,
		},
		{
			name:     "Windows - win-x86_64.zip",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-win-x86_64.zip"},
			expected: AssetPlatformWindowsIntel,
		},

		// Unknown platforms
		{
			name:     "Unknown - no matching suffix",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-freebsd.zip"},
			expected: AssetPlatformUnknown,
		},
		{
			name:     "Unknown - wrong prefix",
			asset:    &mo_release_asset.Asset{Name: "othertool-1.0.0-mac-intel.zip"},
			expected: AssetPlatformUnknown,
		},
		{
			name:     "Unknown - no extension",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-mac-intel"},
			expected: AssetPlatformUnknown,
		},
		{
			name:     "Unknown - empty name",
			asset:    &mo_release_asset.Asset{Name: ""},
			expected: AssetPlatformUnknown,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IdentifyPlatform(tc.asset)
			if result != tc.expected {
				t.Errorf("Expected platform '%s', got '%s' for asset name '%s'", tc.expected, result, tc.asset.Name)
			}
		})
	}
}

func TestIdentifyPlatform_CaseInsensitive(t *testing.T) {
	// Test that the function is case-insensitive
	testCases := []struct {
		name     string
		asset    *mo_release_asset.Asset
		expected string
	}{
		{
			name:     "Uppercase executable name",
			asset:    &mo_release_asset.Asset{Name: strings.ToUpper(app_definitions.ExecutableName) + "-1.0.0-mac-intel.zip"},
			expected: AssetPlatformMacIntel,
		},
		{
			name:     "Mixed case",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-MAC-INTEL.ZIP"},
			expected: AssetPlatformMacIntel,
		},
		{
			name:     "All uppercase",
			asset:    &mo_release_asset.Asset{Name: strings.ToUpper(app_definitions.ExecutableName + "-1.0.0-LINUX-ARM64.ZIP")},
			expected: AssetPlatformLinuxArm,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IdentifyPlatform(tc.asset)
			if result != tc.expected {
				t.Errorf("Expected platform '%s', got '%s' for asset name '%s'", tc.expected, result, tc.asset.Name)
			}
		})
	}
}

func TestIdentifyPlatform_EdgeCases(t *testing.T) {
	// Test edge cases
	testCases := []struct {
		name     string
		asset    *mo_release_asset.Asset
		expected string
	}{
		{
			name:     "Nil asset",
			asset:    nil,
			expected: AssetPlatformUnknown,
		},
		{
			name:     "Asset with extra text after platform",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-mac-intel.zip.sig"},
			expected: AssetPlatformUnknown,
		},
		{
			name:     "Asset with platform in middle",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-mac-intel-1.0.0.zip"},
			expected: AssetPlatformUnknown,
		},
		{
			name:     "Very long version string",
			asset:    &mo_release_asset.Asset{Name: app_definitions.ExecutableName + "-1.0.0-beta1-rc2-snapshot-20240101-mac-intel.zip"},
			expected: AssetPlatformMacIntel,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Handle nil case
			if tc.asset == nil {
				// Should not panic
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for nil asset, but function completed normally")
					}
				}()
			}
			
			result := IdentifyPlatform(tc.asset)
			if result != tc.expected {
				t.Errorf("Expected platform '%s', got '%s'", tc.expected, result)
			}
		})
	}
}