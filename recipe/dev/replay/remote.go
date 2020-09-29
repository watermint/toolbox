package replay

import (
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/io/es_zip"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"os"
	"path/filepath"
	"regexp"
)

type Remote struct {
	rc_recipe.RemarkSecret
	ReplayUrl mo_string.OptionalString
}

func (z *Remote) Preset() {
}

func (z *Remote) Exec(c app_control.Control) error {
	url := os.Getenv(app.EnvNameReplayUrl)
	if z.ReplayUrl.IsExists() {
		url = z.ReplayUrl.Value()
	}
	l := c.Log().With(esl.String("replayUrl", url))
	if url == "" {
		l.Warn("No replay url. Skip")
		return nil
	}

	url = regexp.MustCompile(`\?.*$`).ReplaceAllString(url, "") + "?raw=1"
	archivePath := filepath.Join(c.Workspace().Job(), "replay.zip")
	l.Debug("Downloading replay data", esl.String("url", url), esl.String("path", archivePath))
	err := es_download.Download(l, url, archivePath)
	if err != nil {
		l.Debug("Unable to download", esl.Error(err))
		return err
	}

	replayPath := filepath.Join(c.Workspace().Job(), "replay")
	l.Debug("Extract archive", esl.String("archivePath", archivePath), esl.String("replayPath", replayPath))
	err = es_zip.Extract(l, archivePath, replayPath)
	if err != nil {
		l.Debug("Unable to extract", esl.Error(err))
		return err
	}

	l.Debug("Run replay bundle", esl.String("replayPath", replayPath))
	return rc_exec.Exec(c, &Bundle{}, func(r rc_recipe.Recipe) {
		m := r.(*Bundle)
		m.ReplayPath = mo_string.NewOptional(replayPath)
	})
}

func (z *Remote) Test(c app_control.Control) error {
	return qt_errors.ErrorScenarioTest
}
