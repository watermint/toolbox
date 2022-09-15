package release

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_release"
	"github.com/watermint/toolbox/domain/github/model/mo_release_asset"
	"github.com/watermint/toolbox/domain/github/service/sv_release"
	"github.com/watermint/toolbox/domain/github/service/sv_release_asset"
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	InstallStateFileName = "tbx-install-state.json"
)

var (
	ErrorNoLicenseAgreement      = errors.New("no license agreement")
	ErrorCompatibleAssetNotFound = errors.New("compatible release asset not found")
)

type InstallData struct {
	Installed              *mo_release_asset.Asset `json:"installed"`
	LicenseAccepted        bool
	LicenseAcceptTimestamp string
	LicenseAcceptedBy      string
}

type Install struct {
	Peer                      gh_conn.ConnGithubPublic
	Release                   string
	Path                      mo_path.FileSystemPath
	AcceptLicenseAgreement    bool
	InfoReleaseName           app_msg.Message
	InfoReleaseAsset          app_msg.Message
	InfoAlreadyInstalled      app_msg.Message
	InfoOtherVersionInstalled app_msg.Message
	InfoDownloading           app_msg.Message
	InfoInstalling            app_msg.Message
	InfoComplete              app_msg.Message
	ErrorNoAgreement          app_msg.Message
}

func (z *Install) Preset() {
	z.Release = "latest"
}

func (z *Install) getRelease() (release *mo_release.Release, err error) {
	svr := sv_release.New(z.Peer.Client(), "watermint", "toolbox")
	if z.Release == "latest" {
		return svr.Latest()
	} else {
		return svr.Get(z.Release)
	}
}

func (z *Install) currentOs() string {
	return runtime.GOOS
}
func (z *Install) currentArch() string {
	return runtime.GOARCH
}

func (z *Install) compatibleSuffix() (suffixes []string) {
	ro := z.currentOs()
	ra := z.currentArch()
	switch {
	case ro == "darwin" && ra == "arm64":
		return []string{"mac-arm", "mac"}
	case ro == "darwin" && ra == "amd64":
		return []string{"mac"}
	case ro == "windows" && ra == "arm64":
		return []string{"win-arm", "win"}
	case ro == "windows" && ra == "amd64":
		return []string{"win"}
	case ro == "linux" && ra == "arm64":
		return []string{"linux-arm"}
	case ro == "linux" && ra == "amd64":
		return []string{"linux"}
	default:
		return []string{}
	}
}

func (z *Install) selectExecutableName() string {
	switch z.currentOs() {
	case "windows":
		return "tbx.exe"
	default:
		return "tbx"
	}
}

func (z *Install) selectAsset(assets []*mo_release_asset.Asset) (*mo_release_asset.Asset, error) {
	suffixes := z.compatibleSuffix()

	for _, asset := range assets {
		for _, suffix := range suffixes {
			s := suffix + ".zip"
			if strings.HasSuffix(asset.Name, s) {
				return asset, nil
			}
		}
	}

	return nil, ErrorCompatibleAssetNotFound
}

func (z *Install) Exec(c app_control.Control) error {
	l := c.Log()
	ui := c.UI()
	if !z.AcceptLicenseAgreement {
		ui.Error(z.ErrorNoAgreement)
		return ErrorNoLicenseAgreement
	}
	release, err := z.getRelease()
	if err != nil {
		return err
	}
	ui.Info(z.InfoReleaseName.With("Release", release.Name).With("Tag", release.TagName))

	assets, err := sv_release_asset.New(z.Peer.Client(), "watermint", "toolbox", release.Id).List()
	if err != nil {
		l.Debug("Unable to retrieve release assets", esl.Error(err))
		return err
	}
	asset, err := z.selectAsset(assets)
	if err != nil {
		l.Debug("Unable to find compatible asset", esl.Error(err))
		return err
	}
	ui.Info(z.InfoReleaseAsset.With("AssetName", asset.Name))

	// Verify install state
	installStateFilePath := filepath.Join(z.Path.Path(), InstallStateFileName)
	stateContent, err := os.ReadFile(installStateFilePath)
	switch {
	case os.IsNotExist(err):
		l.Debug("No prior installation found")
		// fall through

	case err != nil:
		return err

	default:
		l.Debug("State file found")
		state := &InstallData{}
		if err := json.Unmarshal(stateContent, state); err != nil {
			l.Debug("State file broken", esl.Error(err))
			// fall through, retry installation in such case
		} else {
			if state.Installed.Name == asset.Name {
				ui.Info(z.InfoAlreadyInstalled)
				return nil
			}
			ui.Info(z.InfoOtherVersionInstalled)
		}
	}

	dlPath, err := os.MkdirTemp("", "tbx-util-release-install")
	if err != nil {
		l.Debug("Unable to create download folder", esl.Error(err))
		return err
	}
	defer func() {
		_ = os.RemoveAll(dlPath)
	}()

	ui.Progress(z.InfoDownloading)
	dlFilePath := filepath.Join(dlPath, asset.Name)
	if err = es_download.Download(l, asset.DownloadUrl, dlFilePath); err != nil {
		l.Debug("Unable to download the file", esl.Error(err), esl.String("url", asset.DownloadUrl))
		return err
	}

	pkg, err := zip.OpenReader(dlFilePath)
	if err != nil {
		l.Debug("Unable to open the archive file", esl.Error(err))
		return err
	}
	defer func() {
		_ = pkg.Close()
	}()

	ui.Progress(z.InfoInstalling)

	exeName := z.selectExecutableName()
	exeFile, err := pkg.Open(exeName)
	if err != nil {
		l.Debug("Unable to open the archive file", esl.Error(err))
		return err
	}

	if err := os.MkdirAll(z.Path.Path(), 0755); err != nil {
		l.Debug("Unable to create installation path", esl.Error(err))
		return err
	}

	installPath := filepath.Join(z.Path.Path(), exeName)
	installFile, err := os.Create(installPath)
	if err != nil {
		l.Debug("Unable to create the file in installation path", esl.Error(err))
		return err
	}
	if _, err := io.Copy(installFile, exeFile); err != nil {
		l.Debug("Unable to copy the file in installation path", esl.Error(err))
		_ = installFile.Close()
		return err
	}
	_ = installFile.Close()

	if err := os.Chmod(installPath, 0755); err != nil {
		l.Debug("Unable to set executable flag", esl.Error(err))
		_ = os.Remove(installPath)
		return err
	}

	var licenseAcceptedBy string
	if usr, err := user.Current(); err != nil {
		licenseAcceptedBy = "unknown"
	} else {
		licenseAcceptedBy = usr.Username
	}

	// write install state file
	installedStateContent, err := json.Marshal(&InstallData{
		Installed:              asset,
		LicenseAccepted:        z.AcceptLicenseAgreement,
		LicenseAcceptTimestamp: time.Now().Format(time.RFC3339),
		LicenseAcceptedBy:      licenseAcceptedBy,
	})
	if err != nil {
		l.Debug("Unable to create state file", esl.Error(err))
		return err
	}

	if err := os.WriteFile(installStateFilePath, installedStateContent, 0644); err != nil {
		l.Debug("Unable to write install state", esl.Error(err))
		return err
	}

	ui.Progress(z.InfoComplete)

	return nil
}

func (z *Install) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("install", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()

	return rc_exec.ExecMock(c, &Install{}, func(r rc_recipe.Recipe) {
		m := r.(*Install)
		m.Path = mo_path.NewFileSystemPath(f)
		m.AcceptLicenseAgreement = true
	})
}
