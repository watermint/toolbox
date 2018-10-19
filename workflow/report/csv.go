package report

import (
	"encoding/csv"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/workflow"
	"io"
	"os"
)

const (
	WORKER_REPORT_CSV = "report/csv"
)

type WorkerReportCsv struct {
	workflow.ReduceWorkerImpl

	ReportPath  string
	DataHeaders []string
}

func (w *WorkerReportCsv) Prefix() string {
	return WORKER_REPORT_CSV
}

func (w *WorkerReportCsv) Reduce(taskIter *workflow.TaskIterator) {
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

func (w *WorkerReportCsv) report(out io.Writer, taskIter *workflow.TaskIterator) {
	cw := csv.NewWriter(out)

	cw.Write(w.DataHeaders)
	defer cw.Flush()

	count := 0

	for taskIter.Next() {
		_, task := taskIter.Task()

		w.reportLine(cw, string(task.Context))
		count++

		w.Pipeline.MarkAsDone(task.TaskPrefix, task.TaskId)
	}

	seelog.Infof("%d Record(s)", count)
}

func (w *WorkerReportCsv) reportLine(out *csv.Writer, jsonData string) {
	data := make([]string, len(w.DataHeaders))
	for i, h := range w.DataHeaders {
		col := gjson.Get(jsonData, h)
		data[i] = col.String()
	}
	out.Write(data)
}
