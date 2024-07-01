package app_license

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/security/sc_obfuscate"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"golang.org/x/crypto/sha3"
	"os"
	"path/filepath"
	"time"
)

type License interface {
	// IsValid returns true if the license is valid.
	IsValid() bool

	// IsInvalid returns true if the license is invalid.
	IsInvalid() bool

	// IsCacheTimeout returns true if the license is cached and the cache is timed out.
	IsCacheTimeout() bool

	// IsLifecycleWithinLimit returns true if the license is active in terms of the expiration date.
	// The second return value is true if the license is warned before the expiration.
	IsLifecycleWithinLimit() (active bool, warning bool)

	// IsEOL returns true if the license is end-of-life.
	IsEOL() (eol bool, reason string)

	// LifecycleLimit returns the lifecycle limit of the license.
	LifecycleLimit() time.Time

	// IsScopeEnabled returns true if the scope is enabled.
	IsScopeEnabled(scope string) bool

	// IsRecipeEnabled returns true if the recipe is enabled.
	IsRecipeEnabled(recipePath string) bool

	// SealWithKey seals the license data with the key.
	SealWithKey(key string) (data []byte, err error)

	// Seal seals the license data, and returns the sealed data and the license key.
	Seal() (data []byte, key string, err error)
}

const (
	LicenseVersionV1      = "74b495ab-051e-4bea-93d4-669fb0a671b1"
	LicenseVersionCurrent = LicenseVersionV1
	LicenseScopeBase      = "690797b2-ff61-46a0-a28d-3b0f7a5b49ed"

	// LicenseBlockSize is the block size of the license.
	LicenseBlockSize = 16 * 1024 // 16KiB

	// LicenseKeySize is the size of the license key.
	// must be divisible by 5 (need to eliminate padding characters in base32 encoding)
	LicenseKeySize = 35

	// MaxLicenseYears is the maximum years of the license.
	MaxLicenseYears = 3
	MinLicenseHours = 1

	// MaxLicenseeNameLength is the maximum length of the licensee name.
	MaxLicenseeNameLength = 128

	DefaultLifecyclePeriod = MaxLicenseYears * 365 * 24 * time.Hour

	DefaultWarningPeriodFraction = 0.8
	DefaultWarningMinimumPeriod  = 7 * 24 * time.Hour
	DefaultWarningMaximumPeriod  = 365 * 24 * time.Hour

	CacheTimeout = 30 * 24 * time.Hour
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
	ErrorLicenseNotFound    = errors.New("license not found")
	ErrorUnknownLicenseType = errors.New("unknown license type")
	ErrorCacheNotFound      = errors.New("cache not found")
	ErrorCacheExpired       = errors.New("license expired")
)

const (
	CopyTypeOriginal = iota
	CopyTypeCachedValidLicense
	CopyTypeCachedNotFound
)

type LicenseRecipe struct {
	// Allow is the list of allowed recipes in recipe path.
	Allow []string `json:"allow"`
}

type LicenseLifecycle struct {
	// AvailableAfter is the time when the license is available after the expiration in seconds.
	AvailableAfter int64 `json:"available_after"`

	// WarningAfter is the time when the license is warned before the expiration in seconds.
	WarningAfter int64 `json:"warning_after"`

	// IsEOL is the flag to indicate the license is end-of-life.
	IsEOL bool `json:"is_eol"`

	// ReasonEOL is the reason of the end-of-life.
	ReasonEOL string `json:"reason_eol"`
}

// LicenseReleaseBinding is the binding of the release number.
type LicenseReleaseBinding struct {
	// ReleaseMinimum is the minimum release number (inclusive).
	ReleaseMinimum uint64 `json:"release_minimum"`

	// ReleaseMaximum is the maximum release number (inclusive).
	ReleaseMaximum uint64 `json:"release_maximum"`
}

type LicenseData struct {
	// Version is the version identifier of the license.
	Version string `json:"version,omitempty"`

	// AppName is the name of the application which the license is for.
	AppName string `json:"app_name,omitempty"`

	// Scope is the scope of the license.
	Scope string `json:"scope,omitempty"`

	// Binding is the binding of the license.
	Binding *LicenseReleaseBinding `json:"binding,omitempty"`

	// CopyType is the copy type of the license.
	CopyType int `json:"cash_status,omitempty"`

	// LicenseeName is the name of the licensee.
	LicenseeName string `json:"licensee_name,omitempty"`

	// LicenseeEmail is the email address of the licensee.
	LicenseeEmail string `json:"licensee_email,omitempty"`

	// CachedAt is the date when the license was cached, in RFC3339 format.
	CachedAt string `json:"cached_at,omitempty"`

	// Lifecycle is the lifecycle information attached to the license.
	Lifecycle *LicenseLifecycle `json:"lifecycle,omitempty"`

	// Recipe is the recipe information attached to the license.
	Recipe *LicenseRecipe `json:"recipe,omitempty"`

	// Padding is the padding string for the license.
	Padding string `json:"padding,omitempty"`
}

func (z LicenseData) buildTimestamp() time.Time {
	if app_definitions.BuildInfo.Timestamp == "" {
		return time.Now()
	}
	buildTime, err := time.Parse(time.RFC3339, app_definitions.BuildInfo.Timestamp)
	if err != nil {
		return time.Now()
	}
	return buildTime
}

func (z LicenseData) LifecycleLimit() time.Time {
	return z.buildTimestamp().Add(time.Duration(z.Lifecycle.AvailableAfter) * time.Second)
}

func (z LicenseData) IsLifecycleWithinLimit() (active bool, warning bool) {
	if z.CopyType == CopyTypeCachedNotFound {
		return false, false
	}

	if z.Binding != nil {
		if app_definitions.Version.Major < z.Binding.ReleaseMinimum || z.Binding.ReleaseMaximum < app_definitions.Version.Major {
			return false, false
		}
	}

	if z.Lifecycle == nil {
		return true, false
	}
	if z.Lifecycle.IsEOL {
		return false, false
	}

	buildTimestamp := app_definitions.BuildInfo.Timestamp
	if buildTimestamp == "" {
		buildTimestamp = time.Now().Format(time.RFC3339)
	}
	warningAfter := z.buildTimestamp().Add(time.Duration(z.Lifecycle.WarningAfter) * time.Second)
	lifecycleLimit := z.LifecycleLimit()
	active = time.Now().Before(lifecycleLimit)
	warning = time.Now().After(warningAfter)
	return
}

func (z LicenseData) IsCacheTimeout() bool {
	if z.CachedAt != "" {
		cachedAt, err := time.Parse(time.RFC3339, z.CachedAt)
		if err != nil {
			return true
		}
		if time.Now().Sub(cachedAt) > CacheTimeout {
			return true
		}
	}
	return false
}

func (z LicenseData) IsValid() bool {
	if z.CopyType == CopyTypeCachedNotFound {
		return false
	}

	if z.Binding != nil {
		if app_definitions.Version.Major < z.Binding.ReleaseMinimum || z.Binding.ReleaseMaximum < app_definitions.Version.Major {
			return false
		}
	}

	if lc, _ := z.IsLifecycleWithinLimit(); !lc {
		return false
	}

	return true
}

func (z LicenseData) IsInvalid() bool {
	return !z.IsValid()
}

func (z LicenseData) IsScopeEnabled(scope string) bool {
	if z.IsInvalid() {
		return false
	}
	return z.Scope == scope
}

func (z LicenseData) IsRecipeEnabled(recipePath string) bool {
	if z.IsInvalid() {
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

func (z LicenseData) WithBinding(minimum, maximum uint64) LicenseData {
	z.Binding = &LicenseReleaseBinding{
		ReleaseMinimum: minimum,
		ReleaseMaximum: maximum,
	}
	return z
}

func (z LicenseData) WithLicensee(name, email string) LicenseData {
	// Limit the length of the licensee name and email
	if len(name) > MaxLicenseeNameLength {
		name = name[:MaxLicenseeNameLength]
	}
	if len(email) > MaxLicenseeNameLength {
		email = email[:MaxLicenseeNameLength]
	}

	z.LicenseeName = name
	z.LicenseeEmail = email
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

func (z LicenseData) Cache() LicenseData {
	z.CachedAt = time.Now().Format(time.RFC3339)
	z.CopyType = CopyTypeCachedValidLicense
	return z
}

func (z LicenseData) SealWithKey(key string) (data []byte, err error) {
	l := esl.Default()
	keySalt := app_definitions.BuildInfo.LicenseSalt + key

	p0, err := json.Marshal(z)
	if err != nil {
		l.Debug("Unable to marshal the data", esl.Error(err))
		return nil, err
	}
	paddingSize := LicenseBlockSize - len(p0)%LicenseBlockSize
	z.Padding = sc_random.MustGetSecureRandomString(paddingSize)

	p1, err := json.Marshal(z)
	if err != nil {
		l.Debug("Unable to marshal the data", esl.Error(err))
		return nil, err
	}

	data, err = sc_obfuscate.Obfuscate(l, []byte(keySalt), p1)
	if err != nil {
		l.Debug("Unable to obfuscate the data", esl.Error(err))
		return nil, err
	}
	return
}

// Seal seals the license data.
func (z LicenseData) Seal() (data []byte, key string, err error) {
	l := esl.Default()
	if z.Lifecycle == nil {
		l.Debug("Expiration date is not set, set to the maximum expiration date")
		z.Lifecycle.AvailableAfter = int64(DefaultLifecyclePeriod.Seconds())
		z.Lifecycle.WarningAfter = int64(DefaultWarningPeriod(DefaultLifecyclePeriod).Seconds())
	} else {
		z.Lifecycle.AvailableAfter = min(int64(DefaultLifecyclePeriod.Seconds()),
			max(z.Lifecycle.AvailableAfter, int64(MinLicenseHours)*3600))
		z.Lifecycle.WarningAfter = min(
			max(
				z.Lifecycle.WarningAfter,
				0,
			),
			z.Lifecycle.AvailableAfter,
		)
	}

	key = sc_random.MustGetSecureRandomString(LicenseKeySize)
	data, err = z.SealWithKey(key)
	if err != nil {
		l.Debug("Unable to seal the data", esl.Error(err))
		return nil, "", err
	}
	return
}

// NewLicense creates a new license data with the scope in the current license version.
func NewLicense(scope string) LicenseData {
	return LicenseData{
		Version: LicenseVersionV1,
		AppName: app_definitions.Name,
		Scope:   scope,
		Lifecycle: &LicenseLifecycle{
			AvailableAfter: int64(DefaultLifecyclePeriod.Seconds()),
			WarningAfter:   int64(DefaultWarningPeriod(DefaultLifecyclePeriod).Seconds()),
		},
	}
}

func LicenseName(key string) string {
	digest := sha3.Sum512([]byte(app_definitions.BuildInfo.LicenseSalt + key))
	return base32.HexEncoding.EncodeToString(digest[:])[:LicenseKeySize]
}

func cacheName(key string) string {
	digest := sha3.Sum512([]byte("CACHE:" + app_definitions.BuildInfo.LicenseSalt + key))
	return "license_cache_" + base32.HexEncoding.EncodeToString(digest[:])[:LicenseKeySize]
}

func LoadAndCacheLicense(key, url, path string) (ld *LicenseData, err error) {
	l := esl.Default().With(esl.String("url", url), esl.String("path", path), esl.String("key", key))
	cache, err := loadLicenseFile(key, path)
	switch {
	case err == nil:
		if cache.IsCacheTimeout() {
			ld, err = loadLicenseUrl(key, url)
			if err != nil {
				l.Debug("Unable to load the license", esl.Error(err))
				return nil, err
			}
			cached := ld.Cache()
			cached.Padding = ""
			if err = cacheLicenseFile(key, path, &cached); err != nil {
				l.Debug("Unable to cache the license", esl.Error(err))
				return nil, err
			}
			return ld, nil
		}
		return cache, nil

	case errors.Is(err, ErrorCacheNotFound), errors.Is(err, ErrorCacheExpired):
		ld, err = loadLicenseUrl(key, url)

		if errors.Is(err, ErrorLicenseNotFound) {
			// Cache even if the license is not found, to avoid the repeated download
			_ = cacheLicenseFile(key, path, &LicenseData{
				CopyType: CopyTypeCachedNotFound,
			})
			return nil, err
		}
		if err != nil {
			l.Debug("Unable to load the license", esl.Error(err))
			return nil, err
		}

		cached := ld.Cache()
		cached.Padding = ""
		if err = cacheLicenseFile(key, path, &cached); err != nil {
			l.Debug("Unable to cache the license", esl.Error(err))
			return nil, err
		}
		return ld, nil

	default:
		l.Debug("Unable to load the license, mark this key as NOT_FOUND.", esl.Error(err))

		// Cache even if the license is not found, to avoid the repeated download
		_ = cacheLicenseFile(key, path, &LicenseData{
			CopyType: CopyTypeCachedNotFound,
		})
		return nil, err
	}
}

func loadLicenseUrl(key, url string) (ld *LicenseData, err error) {
	fileUrl := url + LicenseName(key)
	l := esl.Default().With(esl.String("url", fileUrl))
	dataBase64, err := es_download.DownloadText(l, fileUrl)
	if errors.Is(err, es_download.ErrorNotFound) {
		l.Debug("License not found", esl.String("url", fileUrl))
		return nil, ErrorLicenseNotFound
	}
	if err != nil {
		l.Debug("Unable to download the data", esl.Error(err))
		return nil, err
	}
	dataBin, err := base64.StdEncoding.DecodeString(dataBase64)
	if err != nil {
		l.Debug("Unable to decode the data", esl.Error(err))
		return nil, err
	}
	ld, err = ParseLicense(dataBin, key)
	if err != nil {
		l.Debug("Unable to parse the data", esl.Error(err))
		return nil, err
	}
	if ld.IsInvalid() {
		l.Debug("License is invalid", esl.String("url", fileUrl))
		return nil, ErrorExpired
	}
	return
}

func cacheLicenseFile(key, path string, ld *LicenseData) (err error) {
	filePath := filepath.Join(path, cacheName(key))
	l := esl.Default().With(esl.String("path", filePath))
	if err = os.MkdirAll(path, 0755); err != nil {
		l.Debug("Unable to create the directory", esl.Error(err))
		return err
	}
	data, err := ld.SealWithKey(key)
	if err != nil {
		l.Debug("Unable to seal the data", esl.Error(err))
		return err
	}
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		l.Debug("Unable to write the data", esl.Error(err))
		return err
	}
	return nil
}

func loadLicenseFile(key, path string) (ld *LicenseData, err error) {
	filePath := filepath.Join(path, cacheName(key))
	l := esl.Default().With(esl.String("path", filePath))
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			l.Debug("License file not found", esl.String("path", path))
			return nil, ErrorCacheNotFound
		}
		l.Debug("Unable to load the data", esl.Error(err))
		return nil, err
	}
	ld, err = ParseLicense(data, key)
	if err != nil {
		l.Debug("Unable to parse the data", esl.Error(err))
		return nil, err
	}
	if ld.IsCacheTimeout() {
		l.Debug("Cache timeout", esl.String("path", path))
		return nil, ErrorCacheExpired
	}
	return
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

// DefaultWarningPeriod returns the default warning period for the expiration in seconds.
func DefaultWarningPeriod(expiration time.Duration) time.Duration {
	warningPeriod := time.Duration(expiration.Seconds()*DefaultWarningPeriodFraction) * time.Second
	if warningPeriod < DefaultWarningMinimumPeriod {
		return DefaultWarningMinimumPeriod
	}
	if warningPeriod > DefaultWarningMaximumPeriod {
		return DefaultWarningMaximumPeriod
	}
	return warningPeriod
}
