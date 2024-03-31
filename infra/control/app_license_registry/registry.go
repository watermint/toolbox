package app_license_registry

import (
	"encoding/base32"
	"encoding/base64"
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_client"
	"github.com/watermint/toolbox/domain/github/model/mo_content"
	"github.com/watermint/toolbox/domain/github/service/sv_content"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_license"
	"golang.org/x/crypto/sha3"
	"gorm.io/gorm"
	"strings"
)

type Registry interface {
	// Issue a license.
	Issue(licenseData *app_license.LicenseData) (key string, err error)

	// Revoke a license.
	Revoke(key string) error
}

func DefaultRegistryPath(base string) string {
	salt := app_definitions.BuildInfo.LicenseSalt
	tag := sha3.Sum384([]byte("REGISTRY_PATH:" + salt))
	suffix := base32.HexEncoding.EncodeToString(tag[:])[:16]

	return base + "/licenses_" + strings.ToLower(suffix) + ".db"
}

var (
	LicensePath = "licenses"

	ErrorLicenseNotFound = errors.New("license not found")
)

// Record represents a license record.
type Record struct {
	Key           string `json:"key,omitempty" gorm:"primaryKey"`
	Version       string `json:"version,omitempty"`
	AppName       string `json:"app_name,omitempty"`
	Scope         string `json:"scope,omitempty"`
	LicenseeName  string `json:"licensee_name,omitempty"`
	LicenseeEmail string `json:"licensee_email,omitempty"`
	Expiration    string `json:"expiration,omitempty"`
}

// NewRegistry creates a new registry.
func NewRegistry(client gh_client.Client, owner, repo, branch string, registry *gorm.DB) Registry {
	return &registryImpl{
		client:   client,
		owner:    owner,
		repo:     repo,
		branch:   branch,
		registry: registry,
	}
}

type registryImpl struct {
	client   gh_client.Client
	owner    string
	repo     string
	branch   string
	registry *gorm.DB
}

func (z registryImpl) setup() error {
	return z.registry.AutoMigrate(&Record{})
}

func (z registryImpl) path(key string) string {
	return LicensePath + "/" + app_license.LicenseName(key)
}

func (z registryImpl) fetch(key string) (*mo_content.Content, error) {
	l := z.client.Log()
	path := z.path(key)
	ghc := sv_content.New(z.client, z.owner, z.repo)
	data, err := ghc.Get(path)
	if err != nil {
		l.Debug("Unable to get the license", esl.Error(err))
		return nil, err
	}

	file, found := data.File()
	if !found {
		l.Debug("License file not found")
		return nil, ErrorLicenseNotFound
	}

	return &file, nil
}

func (z registryImpl) Issue(licenseData *app_license.LicenseData) (key string, err error) {
	l := z.client.Log()
	if err := z.setup(); err != nil {
		l.Debug("Unable to setup the registry", esl.Error(err))
		return "", err
	}

	seal, key, err := licenseData.Seal()
	if err != nil {
		l.Debug("Unable to seal the license", esl.Error(err))
		return "", err
	}
	name := app_license.LicenseName(key)

	record := &Record{
		Key:           key,
		Version:       licenseData.Version,
		AppName:       licenseData.AppName,
		Scope:         licenseData.Scope,
		LicenseeName:  licenseData.LicenseeName,
		LicenseeEmail: licenseData.LicenseeEmail,
		Expiration:    licenseData.Expiration,
	}

	ghc := sv_content.New(z.client, z.owner, z.repo)
	_, _, err = ghc.Put(
		z.path(key),
		"LICENSE:"+name,
		base64.StdEncoding.EncodeToString(seal),
		sv_content.Branch(z.branch),
	)
	if err != nil {
		l.Debug("Unable to put the license", esl.Error(err))
		return "", err
	}

	if err := z.registry.Save(record).Error; err != nil {
		l.Debug("Unable to save the record", esl.Error(err))
		return "", err
	}

	return key, nil
}

func (z registryImpl) Revoke(key string) error {
	l := z.client.Log()
	if err := z.setup(); err != nil {
		l.Debug("Unable to setup the registry", esl.Error(err))
		return err
	}

	name := app_license.LicenseName(key)
	path := z.path(key)
	ghc := sv_content.New(z.client, z.owner, z.repo)
	file, err := z.fetch(key)
	if err != nil {
		l.Debug("Unable to get the license", esl.Error(err))
		return err
	}

	_, commit, err := ghc.Delete(
		path,
		"REVOKE:"+name,
		sv_content.Branch(z.branch),
		sv_content.Sha(file.Sha),
	)
	if err != nil {
		l.Debug("Unable to delete the license", esl.Error(err))
		return err
	}
	l.Debug("License revoked", esl.Any("commit", commit))

	if err := z.registry.Delete(&Record{}, "key = ?", key).Error; err != nil {
		l.Debug("Unable to delete the record", esl.Error(err))
		return err
	}
	return nil
}
