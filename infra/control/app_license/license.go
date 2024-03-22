package app_license

import (
	"encoding/base32"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/security/sc_obfuscate"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"golang.org/x/crypto/sha3"
	"time"
)

type License interface {
	// IsValid returns true if the license is valid.
	// The second return value is the license expiration date.
	IsValid() (bool, time.Time)

	// HasLifecycleException returns true if the license has lifecycle exception.
	// The second return value is the new expiration date.
	HasLifecycleException() (bool, time.Time)

	// IsScopeEnabled returns true if the scope is enabled.
	IsScopeEnabled(scope string) bool

	// IsRecipeEnabled returns true if the recipe is enabled.
	IsRecipeEnabled(recipePath string) bool
}

const (
	LicenseVersionV1      = "74b495ab-051e-4bea-93d4-669fb0a671b1"
	LicenseVersionCurrent = LicenseVersionV1
	LicenseScopeBase      = "690797b2-ff61-46a0-a28d-3b0f7a5b49ed"

	// LicenseBlockSize is the block size of the license.
	LicenseBlockSize = 32 * 1024 // 32KiB

	// LicenseKeySize is the size of the license key.
	// must be divisible by 5 (need to eliminate padding characters in base32 encoding)
	LicenseKeySize = 35

	// MaxLicenseYears is the maximum years of the license.
	MaxLicenseYears = 16
)

var (
	KnownLicenseVersions = []string{
		LicenseVersionV1,
	}
	KnownLicenseScopes = []string{
		LicenseScopeBase,
	}
)

var (
	ErrorExpired            = errors.New("license expired")
	ErrorUnknownLicenseType = errors.New("unknown license type")
)

type LicenseRecipe struct {
	// Allow is the list of allowed recipes in recipe path.
	Allow []string `json:"allow"`
}

type LicenseLifecycle struct {
	// ExceptionUntil is the exception expiration date, in RFC3339 format.
	ExceptionUntil string `json:"exception_until"`
}

type LicenseData struct {
	// Version is the version identifier of the license.
	Version string `json:"version,omitempty"`

	// AppName is the name of the application which the license is for.
	AppName string `json:"app_name,omitempty"`

	// Scope is the scope of the license.
	Scope string `json:"scope,omitempty"`

	// Expiration is the expiration date of the license, in RFC3339 format.
	Expiration string `json:"expiration,omitempty"`

	// Lifecycle is the lifecycle information attached to the license.
	Lifecycle *LicenseLifecycle `json:"lifecycle,omitempty"`

	// Recipe is the recipe information attached to the license.
	Recipe *LicenseRecipe `json:"recipe,omitempty"`

	// Padding is the padding string for the license.
	Padding string `json:"padding,omitempty"`
}

func (z LicenseData) IsValid() (bool, time.Time) {
	if z.Expiration == "" {
		return false, time.Time{}
	} else {
		expiration, err := time.Parse(time.RFC3339, z.Expiration)
		if err != nil {
			return false, time.Time{}
		}
		// Check if the license is matching the application name
		if z.AppName != app_definitions.Name {
			return false, expiration
		}
		return expiration.After(time.Now()), expiration
	}
}

func (z LicenseData) HasLifecycleException() (bool, time.Time) {
	if ok, _ := z.IsValid(); !ok {
		return false, time.Time{}
	}
	if z.Lifecycle == nil {
		return false, time.Time{}
	}
	exception, err := time.Parse(time.RFC3339, z.Lifecycle.ExceptionUntil)
	if err != nil {
		return false, time.Time{}
	}
	return exception.After(time.Now()), exception
}

func (z LicenseData) IsScopeEnabled(scope string) bool {
	if ok, _ := z.IsValid(); !ok {
		return false
	}
	return z.Scope == scope
}

func (z LicenseData) IsRecipeEnabled(recipePath string) bool {
	if ok, _ := z.IsValid(); !ok {
		return false
	}
	if z.Recipe == nil {
		return false
	}
	for _, allow := range z.Recipe.Allow {
		if allow == recipePath {
			return true
		}
	}
	return false
}

func (z LicenseData) WithExpiration(expiration time.Time) LicenseData {
	// Limit the expiration date to the maximum license years
	if expiration.After(time.Now().AddDate(MaxLicenseYears, 0, 0)) {
		expiration = time.Now().AddDate(MaxLicenseYears, 0, 0)
	}
	z.Expiration = expiration.Format(time.RFC3339)
	return z
}

func (z LicenseData) WithLifecycle(lc *LicenseLifecycle) LicenseData {
	z.Lifecycle = lc
	return z
}

func (z LicenseData) WithRecipe(rc *LicenseRecipe) LicenseData {
	z.Recipe = rc
	return z
}

// Issue issues a new license data.
func (z LicenseData) Issue() (data []byte, key string, err error) {
	l := esl.Default()
	if z.Expiration == "" {
		l.Debug("Expiration date is not set, set to the maximum expiration date")
		z.Expiration = time.Now().AddDate(MaxLicenseYears, 0, 0).Format(time.RFC3339)
	} else {
		expiration, err := time.Parse(time.RFC3339, z.Expiration)
		if err != nil {
			l.Debug("Unable to parse the expiration date", esl.Error(err))
			return nil, "", err
		}

		// Check if the license is expired at the time of issue
		if expiration.Before(time.Now()) {
			l.Debug("License is expired", esl.Time("expiration", expiration), esl.Time("now", time.Now()))
			return nil, "", ErrorExpired
		}

		// Limit the expiration date to the maximum license years
		if expiration.After(time.Now().AddDate(MaxLicenseYears, 0, 0)) {
			l.Debug("Expiration date is beyond the maximum license years", esl.Time("expiration", expiration), esl.Time("now", time.Now()))
			expiration = time.Now().AddDate(MaxLicenseYears, 0, 0)
		}
		l.Debug("Expiration date", esl.Time("expiration", expiration))
		z.Expiration = expiration.Format(time.RFC3339)
	}

	p0, err := json.Marshal(z)
	if err != nil {
		l.Debug("Unable to marshal the data", esl.Error(err))
		return nil, "", err
	}
	paddingSize := LicenseBlockSize - len(p0)%LicenseBlockSize
	z.Padding = sc_random.MustGetSecureRandomString(paddingSize)

	p1, err := json.Marshal(z)
	if err != nil {
		l.Debug("Unable to marshal the data", esl.Error(err))
		return nil, "", err
	}

	key = sc_random.MustGetSecureRandomString(LicenseKeySize)
	keySalt := app_definitions.BuildInfo.LicenseSalt + key

	data, err = sc_obfuscate.Obfuscate(l, []byte(keySalt), p1)
	if err != nil {
		l.Debug("Unable to obfuscate the data", esl.Error(err))
		return nil, "", err
	}
	return
}

// NewLicense creates a new license data with the scope in the current license version.
func NewLicense(scope string) LicenseData {
	return LicenseData{
		Version:    LicenseVersionV1,
		AppName:    app_definitions.Name,
		Scope:      scope,
		Expiration: time.Now().AddDate(MaxLicenseYears, 0, 0).Format(time.RFC3339),
	}
}

func LicenseName(key string) string {
	digest := sha3.Sum512([]byte(app_definitions.BuildInfo.LicenseSalt + key))
	return base32.HexEncoding.EncodeToString(digest[:])[:LicenseKeySize]
}

func ParseLicense(data []byte, license string) (ld *LicenseData, err error) {
	l := esl.Default()
	keySalt := app_definitions.BuildInfo.LicenseSalt + license
	p1, err := sc_obfuscate.Deobfuscate(l, []byte(keySalt), data)
	if err != nil {
		l.Debug("Unable to de-obfuscate the data", esl.Error(err))
		return nil, err
	}
	ld = &LicenseData{}
	err = json.Unmarshal(p1, &ld)
	if err != nil {
		l.Debug("Unable to unmarshal the data", esl.Error(err))
		return nil, err
	}

	knownLicense := false
	knownScope := false
	for _, v := range KnownLicenseVersions {
		if ld.Version == v {
			knownLicense = true
			break
		}
	}
	for _, s := range KnownLicenseScopes {
		if ld.Scope == s {
			knownScope = true
			break
		}
	}
	if !knownLicense || !knownScope {
		l.Debug("Unknown license or scope", esl.Any("license", ld.Version), esl.Any("scope", ld.Scope))
		return nil, ErrorUnknownLicenseType
	}
	return
}
