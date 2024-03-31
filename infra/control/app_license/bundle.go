package app_license

import (
	"errors"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"time"
)

var (
	ErrorBundleCannotBeSealed = errors.New("bundle cannot be sealed")
)

func NewLicenseBundleFromKeys(keys []string, path string) LicenseBundle {
	licenses := make([]*LicenseData, 0)
	for _, k := range keys {
		l, err := LoadAndCacheLicense(k, app_definitions.RepositorySupplementLicenseUrl, path)
		if err != nil {
			continue
		}
		if v, _, _ := l.IsValid(); v {
			licenses = append(licenses, l)
		}
	}
	return LicenseBundle{
		licenses: licenses,
	}
}

type LicenseBundle struct {
	// licenses that are valid licenses
	licenses []*LicenseData
}

func (z LicenseBundle) LifecycleLimit() time.Time {
	var bestLicense *LicenseData
	var bestExpiration time.Time

	for _, l := range z.licenses {
		_, _, e := l.IsValid()
		if bestLicense == nil || e.After(bestExpiration) {
			bestLicense = l
			bestExpiration = e
		}
	}

	if bestLicense == nil {
		return time.Time{}
	}

	return bestLicense.LifecycleLimit()
}

func (z LicenseBundle) IsValid() (valid bool, cacheTimeout bool, expiration time.Time) {
	var bestLicense *LicenseData
	var bestExpiration time.Time

	for _, l := range z.licenses {
		_, _, e := l.IsValid()
		if bestLicense == nil || e.After(bestExpiration) {
			bestLicense = l
			bestExpiration = e
		}
	}

	if bestLicense == nil {
		return false, false, time.Time{}
	}

	return bestLicense.IsValid()
}

func (z LicenseBundle) IsLifecycleWithinLimit() (active bool, warning bool) {
	for _, l := range z.licenses {
		if a, w := l.IsLifecycleWithinLimit(); a && !w {
			return a, w
		}
	}
	for _, l := range z.licenses {
		if a, w := l.IsLifecycleWithinLimit(); a {
			return a, w
		}
	}
	return false, false
}

func (z LicenseBundle) IsScopeEnabled(scope string) bool {
	for _, l := range z.licenses {
		if l.IsScopeEnabled(scope) {
			return true
		}
	}
	return false
}

func (z LicenseBundle) IsRecipeEnabled(recipePath string) bool {
	for _, l := range z.licenses {
		if l.IsRecipeEnabled(recipePath) {
			return true
		}
	}
	return false
}

func (z LicenseBundle) SealWithKey(key string) (data []byte, err error) {
	return nil, ErrorBundleCannotBeSealed
}

func (z LicenseBundle) Seal() (data []byte, key string, err error) {
	return nil, "", ErrorBundleCannotBeSealed
}
