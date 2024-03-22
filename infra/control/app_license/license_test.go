package app_license

import (
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"testing"
	"time"
)

func TestNewLicense(t *testing.T) {
	lic := NewLicense(LicenseScopeBase).WithExpiration(time.Now().AddDate(1, 0, 0))
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
		licExceptionUntil := time.Now().Add(24 * time.Hour)
		lic1 := lic.WithLifecycle(&LicenseLifecycle{
			ExceptionUntil: licExceptionUntil.Format(time.RFC3339),
		})
		if h, e := lic1.HasLifecycleException(); h && e.Equal(licExceptionUntil) {
			t.Error("Invalid lifecycle")
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
	lic = lic.WithExpiration(time.Now().AddDate(1, 0, 0))
	lic = lic.WithLifecycle(&LicenseLifecycle{
		ExceptionUntil: time.Now().AddDate(1, 0, 0).Format(time.RFC3339),
	})
	lic = lic.WithRecipe(&LicenseRecipe{
		Allow: []string{
			"allow1",
		},
	})
	licData, key, err := lic.Issue()
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
	if lic.Expiration != li2.Expiration {
		t.Error("Invalid expiration")
	}
	if lic.Lifecycle.ExceptionUntil != li2.Lifecycle.ExceptionUntil {
		t.Error("Invalid lifecycle")
	}
	if lic.Recipe.Allow[0] != li2.Recipe.Allow[0] {
		t.Error("Invalid recipe")
	}
}

func TestLicenseName(t *testing.T) {
	lic := NewLicense(LicenseScopeBase)
	_, key, err := lic.Issue()
	if err != nil {
		t.Error(err)
	}
	name := LicenseName(key)
	if len(name) < 1 {
		t.Error("Invalid name")
	}
}
