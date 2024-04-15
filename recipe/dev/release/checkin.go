package release

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/service/sv_content"
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_license"
	"github.com/watermint/toolbox/infra/control/app_license_registry"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/resources"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Checkin struct {
	Peer             gh_conn.ConnGithubRepo
	Owner            string
	Repo             string
	Branch           string
	SupplementOwner  string
	SupplementRepo   string
	SupplementBranch string
}

func (z *Checkin) Preset() {
	z.SupplementOwner = app_definitions.SupplementRepositoryOwner
	z.SupplementRepo = app_definitions.SupplementRepositoryName
	z.SupplementBranch = "main"
	z.Owner = app_definitions.ApplicationRepositoryOwner
	z.Repo = app_definitions.ApplicationRepositoryName
	z.Branch = "main"
}

func (z *Checkin) repoReleaseNumber(c app_control.Control) (release uint64, err error) {
	l := c.Log()

	svc := sv_content.New(z.Peer.Client(), z.Owner, z.Repo)
	releaseContent, err := svc.Get("resources/release/release", sv_content.Branch(z.Branch))
	if err != nil {
		l.Debug("Unable to get release file", esl.Error(err))
		return 0, err
	}
	relFile, found := releaseContent.File()
	if !found {
		l.Debug("Unable to get release file", esl.Error(err))
		return 0, err
	}

	relContent, err := base64.StdEncoding.DecodeString(relFile.Content)
	if err != nil {
		l.Debug("Unable to decode release file", esl.Error(err))
		return 0, err
	}
	releaseText := strings.TrimSpace(string(relContent))

	l.Debug("Release text", esl.String("text", releaseText))
	if strings.TrimSpace(releaseText) == "" {
		l.Debug("Empty release number")
		return 0, nil
	}
	release, err = strconv.ParseUint(releaseText, 10, 64)
	if err != nil {
		l.Debug("Unable to parse release number", esl.Error(err))
		return 0, err
	}
	l.Debug("Release number", esl.Uint64("release", release))
	return release, nil
}

func (z *Checkin) updateReleaseNumber(c app_control.Control, repoRelease uint64) error {
	l := c.Log()

	repoRoot, err := es_project.DetectRepositoryRoot()
	if err != nil {
		l.Debug("Unable to detect repository root", esl.Error(err))
		return err
	}

	fileReleaseContent, err := os.ReadFile(filepath.Join(repoRoot, "resources", "release", "release"))
	if err != nil {
		l.Debug("Unable to read release file", esl.Error(err))
		return err
	}
	fileRelease, err := strconv.ParseUint(string(fileReleaseContent), 10, 64)
	if err != nil {
		l.Debug("Unable to parse release number", esl.Error(err))
		return err
	}

	checkinRelease := repoRelease + 1

	if fileRelease == checkinRelease {
		l.Debug("No need to update release number")
		return nil
	}

	l.Debug("Update release number", esl.Uint64("current", fileRelease), esl.Uint64("next", checkinRelease))
	return os.WriteFile(filepath.Join(repoRoot, "resources", "release", "release"), []byte(strconv.FormatUint(checkinRelease, 10)), 0644)
}

func (z *Checkin) releaseLicenseKey(c app_control.Control, targetRelease uint64) error {
	l := c.Log()

	repoRoot, err := es_project.DetectRepositoryRoot()
	if err != nil {
		l.Debug("Unable to detect repository root", esl.Error(err))
		return err
	}

	licenses := resources.ReleaseLicenses()
	currentHashedSalt := app_definitions.BuildInfo.HashedSalt()
	for _, license := range licenses.Licenses {
		if license.Release == targetRelease && license.HashedSalt == currentHashedSalt {
			l.Debug("Found license", esl.Uint64("release", license.Release))
			return nil
		}
	}
	l.Debug("No license found for the release", esl.Uint64("release", targetRelease))

	registryDatabasePath := app_license_registry.DefaultRegistryPath(c.Workspace().Secrets())
	registryDb, err := c.NewOrm(registryDatabasePath)
	if err != nil {
		l.Debug("Unable to open the license registry", esl.Error(err))
		return err
	}

	license := app_license.NewLicense(app_license.LicenseScopeBase).
		WithLifecycle(&app_license.LicenseLifecycle{
			AvailableAfter: int64(app_definitions.LifecycleExpirationCritical / time.Second),
			WarningAfter:   int64(app_definitions.LifecycleExpirationWarning / time.Second),
			IsEOL:          false,
			ReasonEOL:      "",
		}).
		WithLicensee(
			fmt.Sprintf("watermint toolbox, Release %d", targetRelease),
			"",
		).
		WithBinding(
			app_definitions.Version.Major,
			app_definitions.Version.Major,
		)

	registry := app_license_registry.NewRegistry(
		z.Peer.Client(),
		z.SupplementOwner,
		z.SupplementRepo,
		z.SupplementBranch,
		registryDb,
	)
	key, err := registry.Issue(&license)
	if err != nil {
		l.Debug("Unable to issue license", esl.Error(err))
		return err
	}

	licenses.Licenses = append(licenses.Licenses, &resources.License{
		Release:    targetRelease,
		Key:        key,
		HashedSalt: app_definitions.BuildInfo.HashedSalt(),
	})

	licenseData, err := json.MarshalIndent(licenses, "", "  ")
	if err != nil {
		l.Debug("Unable to marshal licenses", esl.Error(err))
		return err
	}
	licenseDataPath := filepath.Join(repoRoot, "resources", "release", "release_license")
	if err = os.WriteFile(licenseDataPath, licenseData, 0644); err != nil {
		l.Debug("Unable to write license data", esl.Error(err))
		return err
	}

	return nil
}

func (z *Checkin) Exec(c app_control.Control) error {
	l := c.Log()
	repoRelease, err := z.repoReleaseNumber(c)
	if err != nil {
		return err
	}
	l.Info("Current release number", esl.Uint64("release", repoRelease))
	if err := z.updateReleaseNumber(c, repoRelease); err != nil {
		return err
	}
	targetRelease := repoRelease + 1
	l.Info("Next release number", esl.Uint64("release", targetRelease))
	l.Info("Issue release license key", esl.Uint64("release", targetRelease))
	if err := z.releaseLicenseKey(c, targetRelease); err != nil {
		return err
	}
	return nil
}

func (z *Checkin) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
}
