package sharedlink

import (
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/workflow"
	"path/filepath"
	"strings"
)

const (
	WORKER_SHAREDLINK_FILTER_VISIBILITY = "sharedlink/filter/visibility"
	WORKER_SHAREDLINK_FILTER_PATH       = "sharedlink/filter/path"
)

type WorkerSharedLinkFilterByVisibility struct {
	workflow.SimpleWorkerImpl
	NextTask   string
	Visibility string
}

func (*WorkerSharedLinkFilterByVisibility) Prefix() string {
	return WORKER_SHAREDLINK_FILTER_VISIBILITY
}

func (w *WorkerSharedLinkFilterByVisibility) Exec(task *workflow.Task) {
	tc := &ContextSharedLinkResult{}
	workflow.UnmarshalContext(task, tc)

	link := string(tc.Link)
	visibility := gjson.Get(link, "link_permissions.resolved_visibility.\\.tag").String()

	seelog.Debugf("SharedLinkId[%s] ResolvedVisibility[%s] FilterBy[%s]", tc.SharedLinkId, visibility, w.Visibility)

	if visibility == w.Visibility {
		w.Pipeline.Enqueue(
			workflow.MarshalTask(
				w.NextTask,
				task.TaskId,
				tc,
			),
		)
	}
}

type WorkerSharedLinkFilterByPath struct {
	workflow.SimpleWorkerImpl
	NextTask string
	Path     string
}

func (*WorkerSharedLinkFilterByPath) Prefix() string {
	return WORKER_SHAREDLINK_FILTER_PATH
}

func (w *WorkerSharedLinkFilterByPath) Exec(task *workflow.Task) {
	tc := &ContextSharedLinkResult{}
	workflow.UnmarshalContext(task, tc)

	link := string(tc.Link)
	pathLower := gjson.Get(link, "path_lower").String()

	filterPath := filepath.ToSlash(strings.ToLower(w.Path))

	seelog.Debugf("SharedLinkId[%s] PathLower[%s] FilterBy[%s]", tc.SharedLinkId, pathLower, filterPath)

	if strings.HasPrefix(pathLower, filterPath) {
		w.Pipeline.Enqueue(
			workflow.MarshalTask(
				w.NextTask,
				task.TaskId,
				tc,
			),
		)
	}
}
