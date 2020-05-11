package app_job_impl

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_job"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

func NewHistorian(ws app_workspace.Workspace) app_job.Historian {
	return &Historian{ws: ws}
}

type Historian struct {
	ws app_workspace.Workspace
}

func (z Historian) Histories() (histories []app_job.History, err error) {
	l := esl.Default()
	histories = make([]app_job.History, 0)

	path := filepath.Join(z.ws.Home(), "jobs")
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		l.Debug("Unable to read dir", esl.Error(err))
		return nil, err
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		h, found := newHistory(z.ws, e.Name())
		if found {
			histories = append(histories, h)
		}
	}
	sort.Slice(histories, func(i, j int) bool {
		return strings.Compare(histories[i].JobId(), histories[j].JobId()) < 0
	})

	return histories, nil
}
