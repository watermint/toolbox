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

// Determine whether the path contains job history data or not.
func (z Historian) isHistory(jobIds []string) (app_job.History, bool) {
	h, found := newHistory(z.ws, jobIds)
	if !found {
		return nil, false
	}
	if logs, err := h.Logs(); err != nil {
		return nil, false
	} else {
		return h, len(logs) > 0
	}
}

func (z Historian) scanPath(path string, parentJobId []string) (histories []app_job.History, err error) {
	l := esl.Default()
	sp := path
	histories = make([]app_job.History, 0)
	if len(parentJobId) > 0 {
		sp = filepath.Join(path, strings.Join(parentJobId, "/"))
	}
	l.Debug("Reading entries", esl.String("path", sp))
	entries, err := ioutil.ReadDir(sp)
	if err != nil {
		l.Debug("Unable to read dir", esl.Error(err))
		return nil, err
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		switch e.Name() {
		case app_workspace.NameLogs,
			app_workspace.NameJobs,
			app_workspace.NameKvs,
			app_workspace.NameReport,
			app_workspace.NameTest:
			continue
		}

		jp := append(parentJobId, e.Name())
		if h, found := z.isHistory(jp); found {
			histories = append(histories, h)
		}
		children, err := z.scanPath(path, jp)
		if err != nil {
			l.Debug("No job history found in child due to an error. Ignore", esl.Error(err))
			continue
		}
		histories = append(histories, children...)
	}
	return
}

func (z Historian) Histories() (histories []app_job.History, err error) {
	histories = make([]app_job.History, 0)

	path := filepath.Join(z.ws.Home(), app_workspace.NameJobs)
	histories, err = z.scanPath(path, []string{})

	sort.Slice(histories, func(i, j int) bool {
		return strings.Compare(histories[i].JobId(), histories[j].JobId()) < 0
	})

	return histories, nil
}
