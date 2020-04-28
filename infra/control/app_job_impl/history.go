package app_job_impl

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/file/es_archive"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/control/app_job"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type HistoryMetadata struct {
	JobId string `json:"job_id"`
}

func parse(path string, model interface{}) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(content, model); err != nil {
		return err
	}
	return nil
}

func newHistory(ctl app_control.Control, jobId string) (h app_job.History, found bool) {
	l := ctl.Log()
	start := &app_job.StartLog{}
	finish := &app_job.ResultLog{}
	logPath := filepath.Join(ctl.Workspace().Home(), "jobs", jobId)
	startLogPath := filepath.Join(logPath, "logs", app_job.StartLogName)
	finishLogPath := filepath.Join(logPath, "logs", app_job.FinishLogName)
	if err := parse(startLogPath, start); err != nil {
		l.Debug("Unable to load start log", zap.Error(err))
		return nil, false
	}
	if err := parse(finishLogPath, finish); err != nil {
		l.Debug("Unable to load finish log", zap.Error(err))
	}

	return &History{
		ctl:    ctl,
		jobId:  jobId,
		start:  start,
		finish: finish,
	}, true
}

type History struct {
	ctl    app_control.Control
	jobId  string
	start  *app_job.StartLog
	finish *app_job.ResultLog
}

func (z *History) JobId() string {
	return z.jobId
}

func (z *History) RecipeName() string {
	return z.start.Name
}

func (z *History) Recipe() (r rc_recipe.Spec, found bool) {
	cat, ok := z.ctl.(app_control_launcher.ControlLauncher)
	if !ok {
		return nil, false
	}
	_, r, _, err := cat.Catalogue().RootGroup().Select(strings.Split(z.start.Name, " "))
	if err != nil {
		return nil, false
	}
	return r, true
}

func (z *History) AppName() string {
	return z.start.AppName
}

func (z *History) AppVersion() string {
	return z.start.AppVersion
}

func (z *History) TimeStart() (t time.Time, found bool) {
	if z.start == nil || z.start.TimeStart == "" {
		if t, err := time.Parse(app_workspace.JobIdFormat, z.jobId); err == nil {
			return t, true
		}
		return time.Time{}, false
	}
	t, err := time.Parse(time.RFC3339, z.start.TimeStart)
	if err != nil {
		z.ctl.Log().Debug("Unable to parse time", zap.Error(err), zap.String("time", z.start.TimeStart))
		return time.Time{}, false
	}
	return t, true
}

func (z *History) TimeFinish() (t time.Time, found bool) {
	if z.finish == nil || z.finish.TimeFinish == "" {
		return time.Time{}, false
	}
	t, err := time.Parse(time.RFC3339, z.finish.TimeFinish)
	if err != nil {
		z.ctl.Log().Debug("Unable to parse time", zap.Error(err), zap.String("time", z.finish.TimeFinish))
		return time.Time{}, false
	}
	return t, true
}

func (z *History) Delete() error {
	l := z.ctl.Log()
	logPath := filepath.Join(z.ctl.Workspace().Home(), "jobs", z.jobId)
	l.Debug("Trying remove history", zap.String("path", logPath))
	if err := os.RemoveAll(logPath); err != nil {
		l.Debug("Unable to remove", zap.Error(err))
		return err
	}
	return nil
}

func (z *History) Archive() (path string, err error) {
	l := z.ctl.Log()
	logPath := filepath.Join(z.ctl.Workspace().Home(), "jobs", z.jobId)
	arcPath := filepath.Join(z.ctl.Workspace().Home(), "jobs", z.jobId+".zip")

	meta := &HistoryMetadata{JobId: z.jobId}
	metaMarshal, err := json.Marshal(meta)
	if err != nil {
		metaMarshal = []byte("{}")
	}

	if err := es_archive.Create(arcPath, logPath, string(metaMarshal)); err != nil {
		l.Debug("Unable to create archive", zap.Error(err), zap.String("arcPath", arcPath))
		return "", err
	}

	l.Debug("Try removing processed path", zap.String("logPath", logPath))
	err = os.RemoveAll(logPath)
	l.Debug("Remove result", zap.Error(err))
	if err != nil {
		l.Debug("Unable to remove", zap.Error(err))
		return "", err
	}
	return arcPath, nil
}

func NewHistorian(ctl app_control.Control) app_job.Historian {
	return &Historian{ctl: ctl}
}

type Historian struct {
	ctl app_control.Control
}

func (z *Historian) Histories() (histories []app_job.History) {
	l := z.ctl.Log()
	histories = make([]app_job.History, 0)

	path := filepath.Join(z.ctl.Workspace().Home(), "jobs")
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		l.Debug("Unable to read dir", zap.Error(err))
		return
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		h, found := newHistory(z.ctl, e.Name())
		if found {
			histories = append(histories, h)
		}
	}
	return histories
}
