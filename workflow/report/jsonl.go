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
	if !taskIter.Next() {
		seelog.Debugf("Task not found for prefix[%s]", w.Prefix())
		return
	}
	// rewind to first element
	taskIter.Prev()
	seelog.Flush()

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
	count := 0
	for taskIter.Next() {
		_, task := taskIter.Task()

		fmt.Fprintln(out, string(task.Context))
		count++

		w.Pipeline.MarkAsDone(task.TaskPrefix, task.TaskId)
	}

	seelog.Infof("%d Record(s)", count)
}
