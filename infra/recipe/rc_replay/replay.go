package rc_replay

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/io/es_file_copy"
	"github.com/watermint/toolbox/essentials/io/es_file_read"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_capture"
	"github.com/watermint/toolbox/essentials/network/nw_request"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_value"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Replay interface {
	// Preserve the specified Job to dest path.
	Preserve(target app_workspace.Job, destPath string) error

	// Replay the recipe.
	Replay(target app_workspace.Job, ctl app_control.Control) error

	// Compare plays.
	Compare(preserved app_workspace.Job, replay app_workspace.Job) (diffs int, err error)
}

type Capture struct {
	Req nw_request.Req `json:"req"`
	Res nw_capture.Res `json:"res"`
}

func New(logger esl.Logger) Replay {
	return &rpImpl{
		logger: logger,
	}
}

type rpImpl struct {
	logger esl.Logger
}

var (
	PreserveLogFilePrefixes = []string{
		rc_value.FeedBackupFilePrefix,
		app.LogCapture,
		app.LogNameStart,
	}
)

func (z rpImpl) Preserve(target app_workspace.Job, destPath string) error {
	l := z.logger.With(esl.String("targetJob", target.Job()), esl.String("destPath", destPath))
	destReportPath := filepath.Join(destPath, app_workspace.NameReport)
	destLogPath := filepath.Join(destPath, app_workspace.NameLogs)

	l.Debug("Preserve the job")

	if err := os.MkdirAll(destReportPath, 0755); err != nil {
		l.Debug("Unable to create", esl.Error(err), esl.String("path", destReportPath))
		return err
	}
	if err := os.MkdirAll(destLogPath, 0755); err != nil {
		l.Debug("Unable to create", esl.Error(err), esl.String("path", destLogPath))
		return err
	}

	reportEntries, err := ioutil.ReadDir(target.Report())
	if err != nil {
		l.Debug("Unable to read dir", esl.Error(err))
		return err
	}

	for _, re := range reportEntries {
		srcPath := filepath.Join(target.Report(), re.Name())
		dstPath := filepath.Join(destReportPath, re.Name())
		ll := l.With(esl.String("srcPath", srcPath), esl.String("dstPath", dstPath))
		if re.IsDir() {
			ll.Debug("Skip folder")
		} else {
			ll.Debug("Copy")
			if err := es_file_copy.Copy(srcPath, dstPath); err != nil {
				l.Debug("Unable to copy", esl.Error(err))
				return err
			}
		}
	}

	logEntries, err := ioutil.ReadDir(target.Log())
	if err != nil {
		l.Debug("Unable to read dir", esl.Error(err))
		return err
	}

	isPreserveFile := func(name string) bool {
		nameLower := strings.ToLower(name)
		for _, pf := range PreserveLogFilePrefixes {
			pfl := strings.ToLower(pf)
			if strings.HasPrefix(nameLower, pfl) {
				return true
			}
		}
		return false
	}

	for _, le := range logEntries {
		srcPath := filepath.Join(target.Log(), le.Name())
		dstPath := filepath.Join(destLogPath, le.Name())
		ll := l.With(esl.String("srcPath", srcPath), esl.String("dstPath", dstPath))

		switch {
		case le.IsDir():
			ll.Debug("Skip folder")
		case isPreserveFile(le.Name()):
			ll.Debug("Target file found, copy")
			if err := es_file_copy.Copy(srcPath, dstPath); err != nil {
				l.Debug("Unable to copy", esl.Error(err))
				return err
			}
		default:
			ll.Debug("Skip")
		}
	}
	return nil
}

func (z rpImpl) Replay(target app_workspace.Job, ctl app_control.Control) error {
	l := ctl.Log().With(esl.String("targetJobPath", target.Job()))
	captureData, err := ctl.NewKvs("capture" + sc_random.MustGenerateRandomString(6))
	if err != nil {
		l.Debug("Unable to create kvs", esl.Error(err))
		return err
	}
	defer func() {
		captureData.Close()
	}()

	ctlWithReplay := ctl.WithFeature(ctl.Feature().AsReplayTest(captureData))

	targetHistory, found := app_job_impl.NewOrphanHistory(target.Log())
	if !found {
		l.Debug("Target path does not look like a job path")
		return errors.New("target path does not look like a job path")
	}

	targetRecipeSpec, found := targetHistory.Recipe()
	if !found {
		l.Debug("Target recipe spec is not found", esl.String("name", targetHistory.RecipeName()))
		return errors.New("target recipe spec is not found")
	}

	targetLogs, err := targetHistory.Logs()
	if err != nil {
		l.Debug("Unable to retrieve logs", esl.Error(err))
		return err
	}

	l.Debug("Load capture logs")
	for _, tl := range targetLogs {
		ll := l.With(esl.String("log", tl.Name()))
		if tl.Type() == app_job.LogFileTypeCapture {
			capture := &bytes.Buffer{}
			if err := tl.CopyTo(capture); err != nil {
				ll.Debug("Unable to read capture log", esl.Error(err))
				return err
			}

			err = es_file_read.ReadLines(capture, func(line []byte) error {
				capLine := &Capture{}
				if err := json.Unmarshal(line, capLine); err != nil {
					l.Debug("Unable to unmarshal", esl.Error(err))
					return err
				}

				err = captureData.Update(func(kvs kv_kvs.Kvs) error {
					existingRecord := &nw_capture.Res{}
					switch kvs.GetJsonModel(capLine.Req.RequestHash, existingRecord) {
					case nil: // found
						switch {
						case existingRecord.ResponseCode < 0: // IO error
							l.Debug("Overwrite record with new", esl.Any("existing", existingRecord), esl.Any("capture", capLine))
							return kvs.PutJsonModel(capLine.Req.RequestHash, capLine.Res)
						case existingRecord.ResponseCode/100 == 2: // 2xx
							l.Debug("Skip updating record", esl.Any("existing", existingRecord), esl.Any("capture", capLine))
							return nil
						default:
							l.Debug("Overwrite record with new", esl.Any("existing", existingRecord), esl.Any("capture", capLine))
							return kvs.PutJsonModel(capLine.Req.RequestHash, capLine.Res)
						}
					default: // not found
						return kvs.PutJsonModel(capLine.Req.RequestHash, capLine.Res)
					}
				})
				if err != nil {
					l.Debug("Unable to update replay data", esl.Error(err))
					return err
				}
				return nil
			})
		}
	}

	l.Debug("Copy backup files")
	for _, tl := range targetLogs {
		switch {
		case strings.HasPrefix(tl.Name(), rc_value.FeedBackupFilePrefix):
			ll := l.With(esl.String("name", tl.Name()))

			dstPath := filepath.Join(ctlWithReplay.Workspace().Log(), tl.Name())
			ll.Debug("Copying a backup file")

			err := es_file_copy.Copy(tl.Path(), dstPath)
			if err != nil {
				ll.Debug("Unable to copy a backup file", esl.Error(err))
				return err
			}

		default:
			l.Debug("Skip copy backup", esl.String("name", tl.Name()))
		}
	}

	targetRecipeValuesData, err := json.Marshal(targetHistory.StartLog().RecipeValues)
	if err != nil {
		l.Debug("Unable to reconstruct recipe value", esl.Error(err))
		return err
	}
	targetRecipeValue, err := es_json.Parse(targetRecipeValuesData)
	if err != nil {
		l.Debug("Unable to reconstruct recipe value", esl.Error(err))
		return err
	}

	rcp, err := targetRecipeSpec.Restore(targetRecipeValue, ctlWithReplay)
	if err != nil {
		l.Debug("Unable to restore recipe", esl.Error(err))
		return err
	}

	l.Debug("Execute recipe")

	err = rcp.Exec(ctlWithReplay)
	if err != nil {
		l.Debug("Exec failed", esl.Error(err))
		return err
	}
	if err = targetRecipeSpec.SpinDown(ctlWithReplay); err != nil {
		l.Debug("Spin down error", esl.Error(err))
		return err
	}
	return nil
}

func (z rpImpl) Compare(preserved app_workspace.Job, replay app_workspace.Job) (diffs int, err error) {
	panic("implement me")
}
