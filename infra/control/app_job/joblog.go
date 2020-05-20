package app_job

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"os"
	"path/filepath"
)

const (
	StartLogName  = "recipe.log"
	FinishLogName = "result.log"
)

type StartLog struct {
	Name        string                 `json:"name"`
	ValueObject map[string]interface{} `json:"value_object"`
	CommonOpts  map[string]interface{} `json:"common_opts"`
	TimeStart   string                 `json:"time_start,omitempty"`
	AppName     string                 `json:"app_name"`
	AppHash     string                 `json:"app_hash"`
	AppVersion  string                 `json:"app_version"`
}

func (z StartLog) Write(ws app_workspace.Workspace) error {
	return write(filepath.Join(ws.Log(), StartLogName), z)
}

type ResultLog struct {
	Success    bool   `json:"success"`
	TimeFinish string `json:"time_finish"`
	Error      string `json:"error"`
}

func (z ResultLog) Write(ws app_workspace.Workspace) error {
	return write(filepath.Join(ws.Log(), FinishLogName), z)
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
