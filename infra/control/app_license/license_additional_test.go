package app_license

import (
	"strings"
	"testing"
	"time"
)

func TestLicenseData_IsInvalid(t *testing.T) {
	// Test with CopyTypeCachedNotFound
	ld := &LicenseData{
		Version: LicenseVersionV1,
		CopyType: CopyTypeCachedNotFound,
	}
	if !ld.IsInvalid() {
		t.Error("Expected license with CopyTypeCachedNotFound to be invalid")
	}

	// Test with expired lifecycle
	ld = &LicenseData{
		Version: LicenseVersionV1,
		Lifecycle: &LicenseLifecycle{
			AvailableAfter:  -86400, // 1 day in the past from build time
			WarningAfter: 0,
		},
	}
	if !ld.IsInvalid() {
		t.Error("Expected expired license to be invalid")
	}

	// Test with valid license
	ld = &LicenseData{
		Version: LicenseVersionV1,
		Lifecycle: &LicenseLifecycle{
			AvailableAfter:  86400 * 365, // 1 year in future from build time
			WarningAfter: 86400 * 300,
		},
	}
	if ld.IsInvalid() {
		t.Error("Expected valid license to not be invalid")
	}
}

func TestLicenseData_IsCacheTimeout(t *testing.T) {
	// Test with recent cache
	ld := &LicenseData{
		CachedAt: time.Now().Format(time.RFC3339),
	}
	if ld.IsCacheTimeout() {
		t.Error("Expected recent cache to not be timed out")
	}

	// Test with old cache
	ld = &LicenseData{
		CachedAt: time.Now().Add(-31 * 24 * time.Hour).Format(time.RFC3339), // Older than CacheTimeout
	}
	if !ld.IsCacheTimeout() {
		t.Error("Expected old cache to be timed out")
	}

	// Test with invalid date format
	ld = &LicenseData{
		CachedAt: "invalid-date",
	}
	if !ld.IsCacheTimeout() {
		t.Error("Expected invalid date to be treated as timed out")
	}
}

func TestLicenseData_WithMethods(t *testing.T) {
	// Test WithBinding
	ld := NewLicense(LicenseScopeBase)
	ld = ld.WithBinding(100, 200)
	if ld.Binding == nil {
		t.Error("Expected binding to be set")
	}
	if ld.Binding.ReleaseMinimum != 100 || ld.Binding.ReleaseMaximum != 200 {
		t.Error("Expected binding values to match")
	}

	// Test WithLicensee
	ld = ld.WithLicensee("Test User", "test@example.com")
	if ld.LicenseeName != "Test User" {
		t.Error("Expected licensee name to match")
	}
	if ld.LicenseeEmail != "test@example.com" {
		t.Error("Expected licensee email to match")
	}

	// Test WithLicensee with long name
	longName := strings.Repeat("a", 150)
	longEmail := strings.Repeat("b", 150) + "@example.com"
	ld = ld.WithLicensee(longName, longEmail)
	if len(ld.LicenseeName) > MaxLicenseeNameLength {
		t.Error("Expected licensee name to be truncated")
	}
	if len(ld.LicenseeEmail) > MaxLicenseeNameLength {
		t.Error("Expected licensee email to be truncated")
	}

	// Test Cache
	ld = ld.Cache()
	if ld.CachedAt == "" {
		t.Error("Expected cached time to be set")
	}
	if ld.CopyType != CopyTypeCachedValidLicense {
		t.Error("Expected copy type to be cached valid license")
	}
}

func TestLicenseData_IsScopeEnabled(t *testing.T) {
	// Test with no scope
	ld := &LicenseData{
		Version: LicenseVersionV1,
		Scope:  "",
	}
	if ld.IsScopeEnabled(LicenseScopeBase) {
		t.Error("Expected license with no scope to have no enabled scopes")
	}

	// Test with scope
	ld = &LicenseData{
		Version: LicenseVersionV1,
		Scope:  LicenseScopeBase,
	}
	if !ld.IsScopeEnabled(LicenseScopeBase) {
		t.Error("Expected base scope to be enabled")
	}
	if ld.IsScopeEnabled("unknown-scope") {
		t.Error("Expected unknown scope to be disabled")
	}
}

func TestLicenseData_IsRecipeEnabled(t *testing.T) {
	// Test with no recipes
	ld := &LicenseData{
		Version: LicenseVersionV1,
		Recipe: nil,
	}
	if ld.IsRecipeEnabled("any-recipe") {
		t.Error("Expected license with no recipes to have no enabled recipes")
	}

	// Test with recipes
	ld = &LicenseData{
		Version: LicenseVersionV1,
		Recipe: &LicenseRecipe{
			Allow: []string{"dropbox file list", "dropbox team info"},
		},
	}
	if !ld.IsRecipeEnabled("dropbox file list") {
		t.Error("Expected 'dropbox file list' to be enabled")
	}
	if !ld.IsRecipeEnabled("dropbox team info") {
		t.Error("Expected 'dropbox team info' to be enabled")
	}
	if ld.IsRecipeEnabled("unknown recipe") {
		t.Error("Expected unknown recipe to be disabled")
	}
}

func TestDefaultWarningPeriodAdditional(t *testing.T) {
	// Test short lifecycle (less than minimum)
	shortLifecycle := 3 * 24 * time.Hour
	warningPeriod := DefaultWarningPeriod(shortLifecycle)
	if warningPeriod != DefaultWarningMinimumPeriod {
		t.Errorf("Expected warning period to be minimum for short lifecycle, got %v", warningPeriod)
	}

	// Test medium lifecycle
	mediumLifecycle := 30 * 24 * time.Hour
	warningPeriod = DefaultWarningPeriod(mediumLifecycle)
	expectedPeriod := time.Duration(float64(mediumLifecycle) * DefaultWarningPeriodFraction)
	if warningPeriod != expectedPeriod {
		t.Errorf("Expected warning period to be %v for medium lifecycle, got %v", expectedPeriod, warningPeriod)
	}

	// Test long lifecycle (more than maximum)
	longLifecycle := 1000 * 24 * time.Hour
	warningPeriod = DefaultWarningPeriod(longLifecycle)
	if warningPeriod != DefaultWarningMaximumPeriod {
		t.Errorf("Expected warning period to be maximum for long lifecycle, got %v", warningPeriod)
	}
}


func TestNewLicenseBundleFromKeys(t *testing.T) {
	// Test with empty keys
	bundle := NewLicenseBundleFromKeys([]string{}, "/tmp")
	if bundle.IsValid() {
		t.Error("Expected bundle from empty keys to be invalid")
	}

	// Test with invalid keys - should skip invalid licenses
	bundle = NewLicenseBundleFromKeys([]string{"invalid-key-1", "invalid-key-2"}, "/tmp")
	// The function will try to load and cache, but with invalid keys it should result in empty valid licenses
	if bundle.IsValid() {
		t.Error("Expected bundle from invalid keys to be invalid")
	}
}