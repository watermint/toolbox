package workflow

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

const (
	WORKDER_DONE = "done"
)

func (w *DoneWorker) SetPipeline(p *Pipeline) {
	// nop
}
func (w *DoneWorker) Prefix() string {
	return WORKDER_DONE
}
