package replay

import (
	"github.com/watermint/toolbox/essentials/ambient/ea_indicator"
	"github.com/watermint/toolbox/essentials/io/es_zip"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_replay"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

type Bundle struct {
	rc_recipe.RemarkSecret
	ReplayPath mo_string.OptionalString
}

func (z *Bundle) Preset() {
}

func (z *Bundle) Exec(c app_control.Control) error {
	replayPath, err := rc_replay.ReplayPath(z.ReplayPath)
	if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(replayPath)
	if err != nil {
		return err
	}

	ea_indicator.SuppressIndicatorForce()

	for _, entry := range entries {
		entryLower := strings.ToLower(entry.Name())
		l := c.Log().With(esl.String("entry", entryLower))
		replay := rc_replay.New(c.Log())
		if entry.IsDir() || !strings.HasSuffix(entryLower, ".zip") {
			l.Debug("Skip entry", esl.String("entry", entry.Name()))
			continue
		}

		entryName := strings.ReplaceAll(entryLower, ".zip", "")
		forkBundle, err := app_workspace.ForkBundle(c.WorkBundle(), entryName)
		if err != nil {
			l.Debug("Unable to fork bundle", esl.Error(err))
			return err
		}

		err = es_zip.Extract(l, filepath.Join(replayPath, entry.Name()), forkBundle.Workspace().Job())
		if err != nil {
			l.Debug("Unable to extract", esl.Error(err))
			return err
		}

		forkCtl := c.WithBundle(forkBundle).WithFeature(c.Feature().AsQuiet())
		start := time.Now()
		err = replay.Replay(forkBundle.Workspace(), forkCtl)
		if err != nil {
			l.Warn("Error on replay", esl.Error(err))
			return err
		}
		duration := time.Now().Sub(start).Truncate(time.Millisecond)
		l.Info("Success", esl.Duration("duration", duration))
	}
	return nil
}

func (z *Bundle) Test(c app_control.Control) error {
	panic("implement me")
}
