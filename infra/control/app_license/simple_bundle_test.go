package app_license

import (
    "testing"
)

func TestLicenseBundle_BasicMethods(t *testing.T) {
    // Test seal methods return errors
    bundle := LicenseBundle{}

    data, err := bundle.SealWithKey("test-key")
    if err != ErrorBundleCannotBeSealed {
        t.Error("Expected ErrorBundleCannotBeSealed from SealWithKey")
    }
    if data != nil {
        t.Error("Expected nil data from SealWithKey")
    }

    data, key, err := bundle.Seal()
    if err != ErrorBundleCannotBeSealed {
        t.Error("Expected ErrorBundleCannotBeSealed from Seal")
    }
    if data != nil || key != "" {
        t.Error("Expected nil data and empty key from Seal")
    }
}
