package app_job

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"os"
	"path/filepath"
	"time"
)

type TimingLog interface {
	Time() (t time.Time, ok bool)
}

func TimeFromLog(tl TimingLog, jobId string) (t time.Time, ok bool) {
	if tl != nil {
		if t, ok = tl.Time(); ok {
			return t, ok
		}
	}
	if t, err := time.Parse(app_workspace.JobIdFormat, jobId); err == nil {
		return t, true
	}
	return time.Time{}, false
}

func parseTime(ts string) (t time.Time, ok bool) {
	l := esl.Default()
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		l.Debug("Unable to parse time", esl.Error(err), esl.String("time", ts))
		return time.Time{}, false
	}
	return t, true
}

type StartLog struct {
	TimingLog
	Name         string                 `json:"name"`
	ValueObject  map[string]interface{} `json:"value_object"`
	CommonOpts   map[string]interface{} `json:"common_opts"`
	TimeStart    string                 `json:"time_start,omitempty"`
	JobId        string                 `json:"job_id"`
	AppName      string                 `json:"app_name"`
	AppHash      string                 `json:"app_hash"`
	AppVersion   string                 `json:"app_version"`
	RecipeValues interface{}            `json:"recipe_values"`
}

func (z StartLog) Write(ws app_workspace.Workspace) error {
	return write(filepath.Join(ws.Log(), app.LogNameStart), z)
}

func (z StartLog) Time() (t time.Time, ok bool) {
	return parseTime(z.TimeStart)
}

type ResultLog struct {
	Success    bool   `json:"success"`
	TimeFinish string `json:"time_finish"`
	Error      string `json:"error"`
}

func (z ResultLog) Write(ws app_workspace.Workspace) error {
	return write(filepath.Join(ws.Log(), app.LogNameFinish), z)
}

func (z ResultLog) Time() (t time.Time, ok bool) {
	return parseTime(z.TimeFinish)
}

func write(path string, d interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	rb, err := json.Marshal(d)
	if err != nil {
		return err
	}
	_, err = f.Write(rb)
	if err != nil {
		return err
	}
	return nil
}
