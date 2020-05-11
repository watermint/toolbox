package app_job_impl

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/file/es_zip"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_job"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
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

func newHistory(app app_workspace.Application, jobId string) (h app_job.History, found bool) {
	l := esl.Default()
	start := &app_job.StartLog{}
	finish := &app_job.ResultLog{}

	jobPath := filepath.Join(app.Home(), "jobs", jobId)
	ws, err := app_workspace.NewWorkspaceByJobPath(app, jobPath)
	if err != nil {
		l.Debug("Unable to determine the path as job", esl.Error(err))
		return nil, false
	}

	startLogPath := filepath.Join(ws.Log(), app_job.StartLogName)
	finishLogPath := filepath.Join(ws.Log(), app_job.FinishLogName)
	if err := parse(startLogPath, start); err != nil {
		l.Debug("Unable to load start log", esl.Error(err))
		return nil, false
	}
	if err := parse(finishLogPath, finish); err != nil {
		l.Debug("Unable to load finish log", esl.Error(err))
	}

	return &History{
		ws:     ws,
		jobId:  jobId,
		start:  start,
		finish: finish,
	}, true
}

type History struct {
	ws     app_workspace.Workspace
	jobId  string
	start  *app_job.StartLog
	finish *app_job.ResultLog
}

func (z History) JobPath() string {
	return z.ws.Job()
}

func (z History) JobId() string {
	return z.jobId
}

func (z History) RecipeName() string {
	return z.start.Name
}

func (z History) Recipe() (r rc_recipe.Spec, found bool) {
	cat := app_catalogue.Current()
	_, r, _, err := cat.RootGroup().Select(strings.Split(z.start.Name, " "))
	if err != nil {
		return nil, false
	}
	return r, true
}

func (z History) AppName() string {
	return z.start.AppName
}

func (z History) AppVersion() string {
	return z.start.AppVersion
}

func (z History) TimeStart() (t time.Time, found bool) {
	l := esl.Default()
	if z.start == nil || z.start.TimeStart == "" {
		if t, err := time.Parse(app_workspace.JobIdFormat, z.jobId); err == nil {
			return t, true
		}
		return time.Time{}, false
	}
	t, err := time.Parse(time.RFC3339, z.start.TimeStart)
	if err != nil {
		l.Debug("Unable to parse time", esl.Error(err), esl.String("time", z.start.TimeStart))
		return time.Time{}, false
	}
	return t, true
}

func (z History) TimeFinish() (t time.Time, found bool) {
	l := esl.Default()
	if z.finish == nil || z.finish.TimeFinish == "" {
		return time.Time{}, false
	}
	t, err := time.Parse(time.RFC3339, z.finish.TimeFinish)
	if err != nil {
		l.Debug("Unable to parse time", esl.Error(err), esl.String("time", z.finish.TimeFinish))
		return time.Time{}, false
	}
	return t, true
}

func (z History) Delete() error {
	l := esl.Default()
	logPath := filepath.Join(z.ws.Home(), "jobs", z.jobId)
	l.Debug("Trying remove history", esl.String("path", logPath))
	if err := os.RemoveAll(logPath); err != nil {
		l.Debug("Unable to remove", esl.Error(err))
		return err
	}
	return nil
}

func (z History) Archive() (path string, err error) {
	l := esl.Default()
	logPath := filepath.Join(z.ws.Home(), "jobs", z.jobId)
	arcPath := filepath.Join(z.ws.Home(), "jobs", z.jobId+".zip")

	meta := &HistoryMetadata{JobId: z.jobId}
	metaMarshal, err := json.Marshal(meta)
	if err != nil {
		metaMarshal = []byte("{}")
	}

	if err := es_zip.CompressPath(arcPath, logPath, string(metaMarshal)); err != nil {
		l.Debug("Unable to create archive", esl.Error(err), esl.String("arcPath", arcPath))
		return "", err
	}

	l.Debug("Try removing processed path", esl.String("logPath", logPath))
	err = os.RemoveAll(logPath)
	l.Debug("Remove result", esl.Error(err))
	if err != nil {
		l.Debug("Unable to remove", esl.Error(err))
		return "", err
	}
	return arcPath, nil
}

func (z History) Logs() (logs []app_job.LogFile, err error) {
	l := esl.Default()
	logs = make([]app_job.LogFile, 0)
	entries, err := ioutil.ReadDir(z.ws.Log())
	if err != nil {
		l.Debug("")
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		p := filepath.Join(z.ws.Log(), entry.Name())
		lf, err := newLogFile(p)
		if err != nil {
			l.Debug("the file is not a log", esl.Error(err), esl.String("name", entry.Name()))
			continue
		}

		logs = append(logs, lf)
	}
	sort.Slice(logs, func(i, j int) bool {
		return strings.Compare(logs[i].Name(), logs[j].Name()) < 0
	})
	l.Debug("logs found", esl.Int("entries", len(logs)))
	return
}
