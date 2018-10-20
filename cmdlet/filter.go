package cmdlet

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/dbx_task/task/sharedlink"
	"github.com/watermint/toolbox/workflow"
	"strings"
)

type SharedLinkFilter struct {
	FilterByPath       string
	FilterByVisibility string
}

func (s *SharedLinkFilter) FlagConfig(f *flag.FlagSet) {
	descFilterPath := "Filter by file path (default: no filter)"
	f.StringVar(&s.FilterByPath, "filter-path", "", descFilterPath)

	descFilterVisibility := fmt.Sprintf(
		"Filter by visibility (default: no filter, {%s})",
		strings.Join(s.SupportedVisibility(), ", "),
	)
	f.StringVar(&s.FilterByVisibility, "filter-visibility", "", descFilterVisibility)
}

func (s *SharedLinkFilter) FilterStages(nextTask string) (firstFilter string, stages []workflow.Worker, err error) {
	stages = make([]workflow.Worker, 0)
	first := ""

	if s.FilterByPath != "" {
		seelog.Debugf("FilterByPath[%s]", s.FilterByPath)
		nt := nextTask
		if s.FilterByVisibility != "" {
			nt = sharedlink.WORKER_SHAREDLINK_FILTER_VISIBILITY
		}
		first = sharedlink.WORKER_SHAREDLINK_FILTER_PATH
		stages = append(
			stages,
			&sharedlink.WorkerSharedLinkFilterByPath{
				Path:     s.FilterByPath,
				NextTask: nt,
			},
		)
	}

	if s.FilterByVisibility != "" {
		seelog.Debugf("FilterByVisibility[%s]", s.FilterByVisibility)
		if first == "" {
			first = sharedlink.WORKER_SHAREDLINK_FILTER_VISIBILITY
		}
		found := false
		for _, v := range s.SupportedVisibility() {
			if v == s.FilterByVisibility {
				found = true
			}
		}
		if !found {
			seelog.Warnf("Unsupported visibility [%s] for filtering shared link", s.FilterByVisibility)
			return "", nil, errors.New("unsupported visibility")
		}
		stages = append(
			stages,
			&sharedlink.WorkerSharedLinkFilterByVisibility{
				Visibility: s.FilterByVisibility,
				NextTask:   nextTask,
			},
		)
	}

	seelog.Debugf("First filter stage[%s] NextTask[%s]", first, nextTask)

	return first, stages, nil
}

func (s *SharedLinkFilter) SupportedVisibility() []string {
	return []string{
		"public",
		"team_only",
		"password",
		"team_and_password",
		"shared_folder_only",
	}
}
