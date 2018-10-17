package workflow

import "github.com/cihub/seelog"

const (
	WORKER_DONE                 = "done"
	WORKER_GENERAL_ERROR_REPORT = "error"
)

type Task struct {
	TaskPrefix string
	TaskId     string
	Context    []byte
	DeferUntil int64
}

type Worker interface {
	Prefix() string
	SetPipeline(p *Pipeline)
}

type SimpleWorker interface {
	Worker
	Exec(task *Task)
}
type BatchWorker interface {
	Worker
	BatchMaxSize() int
	BatchExec(tasks []*Task)
}
type ReduceWorker interface {
	Worker
	Reduce(taskIter *TaskIterator)
}

type SimpleWorkerImpl struct {
	SimpleWorker
	Pipeline *Pipeline
}

func (w *SimpleWorkerImpl) SetPipeline(p *Pipeline) {
	w.Pipeline = p
}

type BatchWorkerImpl struct {
	BatchWorker
	Pipeline *Pipeline
}

func (w *BatchWorkerImpl) SetPipeline(p *Pipeline) {
	w.Pipeline = p
}

type ReduceWorkerImpl struct {
	ReduceWorker
	Pipeline *Pipeline
}

func (w *ReduceWorkerImpl) SetPipeline(p *Pipeline) {
	w.Pipeline = p
}

type DoneWorker struct {
}

func (w *DoneWorker) SetPipeline(p *Pipeline) {
	// nop
}
func (w *DoneWorker) Prefix() string {
	return WORKER_DONE
}

type WorkerGeneralErrorReport struct {
	ReduceWorkerImpl
}

func (w *WorkerGeneralErrorReport) Prefix() string {
	return WORKER_GENERAL_ERROR_REPORT
}

func (w *WorkerGeneralErrorReport) Reduce(taskIter *TaskIterator) {
	for taskIter.Next() {
		_, task := taskIter.Task()
		tc := ContextGeneralErrorReport{}

		UnmarshalContext(task, &tc)

		seelog.Debugf("GeneralError Prefix[%s] TaskId[%s] ErrorTag[%s] Desc[%s]", task.TaskPrefix, task.TaskId, tc.ErrorTag, tc.ErrorDescription)
		seelog.Errorf("Error[%s]: %s", tc.ErrorTag, tc.ErrorDescription)

		w.Pipeline.MarkAsDone(task.TaskPrefix, task.TaskId)
	}
}

type ContextGeneralErrorReport struct {
	ErrorTag         string `json:"errorTag"`
	ErrorDescription string `json:"errorDesc"`
}
