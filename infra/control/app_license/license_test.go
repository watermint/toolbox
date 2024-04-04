package app_license

import (
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"testing"
	"time"
)

func TestNewLicense(t *testing.T) {
	lic := NewLicense(LicenseScopeBase)
	if lic.Version != LicenseVersionCurrent {
		t.Error("Invalid version")
	}
	if lic.AppName != app_definitions.Name {
		t.Error("Invalid app name")
	}
	if lic.Scope != LicenseScopeBase {
		t.Error("Invalid scope")
	}

	{
		availableAfter := int64(3600)
		warningAfter := int64(7200)

		lic1 := lic.WithLifecycle(&LicenseLifecycle{
			AvailableAfter: availableAfter,
			WarningAfter:   warningAfter,
		})
		if lic1.Lifecycle.AvailableAfter != availableAfter || lic1.Lifecycle.WarningAfter != warningAfter {
			t.Error("Invalid lifecycle")
		}
	}

	{
		lic2 := lic.WithLicensee("Scott", "scott@example.com")
		if lic2.LicenseeName != "Scott" || lic2.LicenseeEmail != "scott@example.com" {
			t.Error("Invalid licensee")
		}
	}

	{
		licRecipe := &LicenseRecipe{
			Allow: []string{
				"allow1",
				"allow2",
			},
		}
		lic1 := lic.WithRecipe(licRecipe)
		if lic1.Recipe.Allow[0] != "allow1" || lic1.Recipe.Allow[1] != "allow2" {
			t.Error("Invalid recipe")
		}

		if !lic1.IsRecipeEnabled("allow1") {
			t.Error("recipe not enabled")
		}
		if lic1.IsRecipeEnabled("allow999") {
			t.Error("recipe enabled")
		}
	}
}

func TestIssueParse(t *testing.T) {
	lic := NewLicense(LicenseScopeBase)
	availableAfter := int64(3600)
	warningAfter := int64(7200)
	lic = lic.WithLifecycle(&LicenseLifecycle{
		AvailableAfter: availableAfter,
		WarningAfter:   warningAfter,
	})
	lic = lic.WithRecipe(&LicenseRecipe{
		Allow: []string{
			"allow1",
		},
	})
	licData, key, err := lic.Seal()
	if err != nil {
		t.Error(err)
	}
	if len(key) < LicenseKeySize {
		t.Error("Invalid key")
	}
	li2, err := ParseLicense(licData, key)
	if err != nil {
		t.Error(err)
	}
	if lic.Scope != li2.Scope {
		t.Error("Invalid scope")
	}
	if lic.Lifecycle.AvailableAfter != li2.Lifecycle.AvailableAfter {
		t.Error("Invalid lifecycle")
	}
	if lic.Lifecycle.WarningAfter != li2.Lifecycle.WarningAfter {
		t.Error("Invalid lifecycle")
	}
	if lic.Recipe.Allow[0] != li2.Recipe.Allow[0] {
		t.Error("Invalid recipe")
	}
}

func TestLicenseName(t *testing.T) {
	lic := NewLicense(LicenseScopeBase)
	_, key, err := lic.Seal()
	if err != nil {
		t.Error(err)
	}
	name := LicenseName(key)
	if len(name) < 1 {
		t.Error("Invalid name")
	}
}

func TestDefaultWarningPeriod(t *testing.T) {
	if x := DefaultWarningPeriod(0); x != DefaultWarningMinimumPeriod {
		t.Error(x)
	}

	x0 := 365 * 24 * 3600 * time.Second
	if x := DefaultWarningPeriod(x0); x != time.Duration(float64(x0)*DefaultWarningPeriodFraction)*time.Second {
		t.Error(x)
	}

	x1 := 10 * 365 * 24 * 3600 * time.Second
	if x := DefaultWarningPeriod(x1); x != DefaultWarningMaximumPeriod {
		t.Error(x)
	}
}
