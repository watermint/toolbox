package report

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/workflow"
	"io"
	"os"
)

const (
	WORKER_REPORT_JSONL = "report/jsonl"
)

type WorkerReportJsonl struct {
	workflow.ReduceWorkerImpl

	ReportPath string
}

func (w *WorkerReportJsonl) Prefix() string {
	return WORKER_REPORT_JSONL
}

func (w *WorkerReportJsonl) Reduce(taskIter *workflow.TaskIterator) {
	if w.ReportPath != "" {
		wr, err := os.Create(w.ReportPath)
		if err != nil {
			w.Pipeline.GeneralError("cant_write_file", fmt.Sprintf("Couldn't write reports into the file [%s]", w.ReportPath))
			w.report(os.Stdout, taskIter)
		} else {
			seelog.Debugf("Writing report to [%s]", w.ReportPath)
			defer wr.Close()
			w.report(wr, taskIter)
		}
	} else {
		seelog.Debug("Writing report to STDOUT")
		w.report(os.Stdout, taskIter)
	}
}

func (w *WorkerReportJsonl) report(out io.Writer, taskIter *workflow.TaskIterator) {
	for taskIter.Next() {
		_, task := taskIter.Task()

		fmt.Fprintln(out, string(task.Context))

		w.Pipeline.MarkAsDone(task.TaskPrefix, task.TaskId)
	}
}
