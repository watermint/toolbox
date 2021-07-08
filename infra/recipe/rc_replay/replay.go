package rc_replay

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/io/es_file_copy"
	"github.com/watermint/toolbox/essentials/io/es_file_read"
	"github.com/watermint/toolbox/essentials/io/es_zip"
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
	"reflect"
	"sort"
	"strings"
)

type Replay interface {
	// Preserve the specified Job to dest path as the zip archive.
	// destPath requires file name of the archive
	Preserve(target app_workspace.Job, destPath string) error

	// Replay the recipe.
	Replay(target app_workspace.Job, ctl app_control.Control) error

	// Compare plays.
	Compare(preserved app_workspace.Job, replay app_workspace.Job) (err error)
}

var (
	ErrorReportDiffFound = errors.New("report diff found")
)

type Capture struct {
	Req nw_request.Req `json:"req"`
	Res nw_capture.Res `json:"res"`
}

type Opts struct {
	reportDiffs bool
}

func (z Opts) Apply(opts []Opt) Opts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
}

type Opt func(o Opts) Opts

func ReportDiffs(enabled bool) Opt {
	return func(o Opts) Opts {
		o.reportDiffs = enabled
		return o
	}
}

func New(logger esl.Logger, opts ...Opt) Replay {
	return &rpImpl{
		logger: logger,
		opt:    Opts{}.Apply(opts),
	}
}

type rpImpl struct {
	logger esl.Logger
	opt    Opts
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
	destReportPath := app_workspace.NameReport
	destLogPath := app_workspace.NameLogs
	zw := es_zip.NewWriter(z.logger)
	success := false
	if err := zw.Open(destPath); err != nil {
		l.Debug("Unable to create the archive", esl.Error(err))
		return err
	}
	defer func() {
		_ = zw.Close()
		if !success {
			l.Debug("Remove incomplete archive file")
			_ = os.RemoveAll(destPath)
		}
	}()

	l.Debug("Preserve the job")

	reportEntries, err := ioutil.ReadDir(target.Report())
	if err != nil {
		l.Debug("Unable to read dir", esl.Error(err))
		return err
	}

	for _, re := range reportEntries {
		srcPath := filepath.Join(target.Report(), re.Name())
		ll := l.With(esl.String("srcPath", srcPath))
		if re.IsDir() {
			ll.Debug("Skip folder")
		} else {
			ll.Debug("Copy")
			if err := zw.AddFile(srcPath, destReportPath); err != nil {
				ll.Debug("Unable to add the file", esl.Error(err))
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
		ll := l.With(esl.String("srcPath", srcPath))

		switch {
		case le.IsDir():
			ll.Debug("Skip folder")
		case isPreserveFile(le.Name()):
			ll.Debug("Target file found, copy")
			if err := zw.AddFile(srcPath, destLogPath); err != nil {
				l.Debug("Unable to copy", esl.Error(err))
				return err
			}
		default:
			ll.Debug("Skip")
		}
	}
	success = true
	return nil
}

func (z rpImpl) Replay(target app_workspace.Job, ctl app_control.Control) error {
	l := ctl.Log().With(esl.String("targetJobPath", target.Job()))
	captureData, err := ctl.NewKvs("capture" + sc_random.MustGetSecureRandomString(6))
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
					existingRecords := make([]nw_capture.Res, 0)
					capData, err := kvs.GetJson(capLine.Req.RequestHash)
					if err != nil {
						existingRecords = []nw_capture.Res{capLine.Res}
					} else {
						if err = json.Unmarshal(capData, &existingRecords); err != nil {
							l.Debug("Unable to unmarshall", esl.Error(err))
							existingRecords = append(existingRecords, capLine.Res)
						} else {
							existingRecords = []nw_capture.Res{capLine.Res}
						}
					}
					return kvs.PutJsonModel(capLine.Req.RequestHash, existingRecords)
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
	logEntries, err := ioutil.ReadDir(target.Log())
	if err != nil {
		l.Debug("Unable to read log folder", esl.Error(err))
		return err
	}
	for _, tl := range logEntries {
		path := filepath.Join(target.Log(), tl.Name())
		switch {
		case strings.HasPrefix(tl.Name(), rc_value.FeedBackupFilePrefix):
			ll := l.With(esl.String("name", tl.Name()))

			dstPath := filepath.Join(ctlWithReplay.Workspace().Log(), tl.Name())
			if dstPath == path {
				ll.Debug("Skip copy (in case of the file already copied by prior process)")
				continue
			}

			ll.Debug("Copying a backup file")
			err := es_file_copy.Copy(path, dstPath)
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

	// Recover panic
	defer func() {
		if rvr := recover(); rvr != nil {
			if rvrErr, ok := rvr.(error); ok {
				l.Warn("Recovery from panic with an error", esl.Error(rvrErr))
				err = rvrErr
			} else {
				l.Warn("Recovery form panic", esl.Any("recover", rvr))
				err = errors.New("panic")
			}
		}
	}()

	err = rcp.Exec(ctlWithReplay)
	if err != nil {
		l.Debug("Exec failed", esl.Error(err))
		return err
	}
	if err = targetRecipeSpec.SpinDown(ctlWithReplay); err != nil {
		l.Debug("Spin down error", esl.Error(err))
		return err
	}
	if cc, ok := ctlWithReplay.(app_control.ControlCloser); ok {
		l.Debug("Close control")
		cc.Close()
	}

	if err := z.Compare(target, ctl.Workspace()); err != nil {
		l.Debug("Report diff found", esl.Error(err))
		return err
	}

	return err
}

func (z rpImpl) compareTextReport(approved app_workspace.Job, reportName string, replay app_workspace.Job) (err error) {
	l := z.logger.With(
		esl.String("preserved", approved.Job()),
		esl.String("replay", replay.Job()),
		esl.String("reportName", reportName),
	)

	approvedReportPath := filepath.Join(approved.Report(), reportName)
	approvedLines := make([]string, 0)
	l.Debug("Read approved report", esl.String("approvedReportPath", approvedReportPath))
	err = es_file_read.ReadFileLines(approvedReportPath, func(line []byte) error {
		hash := sha256.Sum256(line)
		hashEncoded := base64.RawStdEncoding.EncodeToString(hash[:])
		approvedLines = append(approvedLines, hashEncoded)
		return nil
	})
	if err != nil {
		l.Debug("Unable to read file", esl.Error(err))
		return err
	}

	replayReportPath := filepath.Join(replay.Report(), reportName)
	replayLines := make([]string, 0)
	l.Debug("Read replay report", esl.String("replayReportPath", replayReportPath))
	err = es_file_read.ReadFileLines(replayReportPath, func(line []byte) error {
		hash := sha256.Sum256(line)
		hashEncoded := base64.RawStdEncoding.EncodeToString(hash[:])
		replayLines = append(replayLines, hashEncoded)
		return nil
	})
	if err != nil {
		l.Debug("Unable to read file", esl.Error(err))
		return err
	}

	sort.Strings(approvedLines)
	sort.Strings(replayLines)

	if reflect.DeepEqual(approvedLines, replayLines) {
		l.Debug("Approved")
		return nil
	}

	if z.opt.reportDiffs {
		l.Warn("Report diff found")
	}

	return ErrorReportDiffFound
}

func (z rpImpl) Compare(approved app_workspace.Job, replay app_workspace.Job) (err error) {
	l := z.logger.With(esl.String("preserved", approved.Job()), esl.String("replay", replay.Job()))
	preservedReportEntries, err := ioutil.ReadDir(approved.Report())
	if err != nil {
		l.Debug("Unable to read reports folder", esl.Error(err))
		return err
	}

	var lastErr error
	for _, entry := range preservedReportEntries {
		if entry.IsDir() {
			continue
		}
		switch strings.ToLower(filepath.Ext(entry.Name())) {
		case ".csv", ".json":
			fileErr := z.compareTextReport(approved, entry.Name(), replay)
			l.Debug("Report diffs",
				esl.String("report", entry.Name()),
				esl.Error(fileErr))
			if fileErr != nil {
				l.Debug("Unable to compare", esl.Error(err))
				lastErr = fileErr
			}
		}
	}

	return lastErr
}
