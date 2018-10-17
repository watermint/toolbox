package workflow

import (
	"bytes"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/infra"
	"path/filepath"
	"strings"
	"time"
)

type Pipeline struct {
	Infra        *infra.InfraContext
	dbStatus     *leveldb.DB
	currentStage int
	allStages    []Worker

	Stages []Worker
}

func (p *Pipeline) Init() error {
	p.allStages = make([]Worker, len(p.Stages)+1)
	for i, w := range p.Stages {
		p.allStages[i] = w
		w.SetPipeline(p)
	}
	genErr := &WorkerGeneralErrorReport{}
	genErr.SetPipeline(p)
	p.allStages[len(p.Stages)] = genErr

	dbName := fmt.Sprintf("p_%x", time.Now().UnixNano())
	db, err := leveldb.OpenFile(filepath.Join(p.Infra.WorkPath, dbName), nil)
	if err != nil {
		seelog.Warnf("Unable to open pipeline database : error[%s]", err)
		return err
	}

	p.dbStatus = db
	return nil
}

const (
	TASK_STATE_WAITING = "W"
	TASK_STATE_RUNNING = "R"
	TASK_STATE_DONE    = "D"
	TASK_SEPARATOR     = "\000"
)

func taskKey(state, prefix, taskId string) []byte {
	return []byte(state + TASK_SEPARATOR + prefix + TASK_SEPARATOR + taskId)
}

func parseKey(key []byte) (state, prefix, taskId string) {
	k := bytes.Split(key, []byte(TASK_SEPARATOR))
	if len(k) != 3 {
		panic(fmt.Sprintf("Incorrect key format: [%s]", string(key)))
	}
	state = string(k[0])
	prefix = string(k[1])
	taskId = string(k[2])
	return
}

func (p *Pipeline) Enqueue(task *Task) {
	seelog.Debugf("EnqueueTask: Prefix[%s] TaskId[%s] Context[%s]", task.TaskPrefix, task.TaskId, string(task.Context))
	err := p.dbStatus.Put(taskKey(TASK_STATE_WAITING, task.TaskPrefix, task.TaskId), NewTaskValueContainer(task).Value(), nil)
	if err != nil {
		seelog.Errorf("Unable to put task: Prefix[%s] TaskId[%s]", task.TaskPrefix, task.TaskId)
		panic("Unable to put task")
	}
}
func (p *Pipeline) TaskIterator(state, taskPrefix string) *TaskIterator {
	taskRange := util.Range{Start: []byte(state + TASK_SEPARATOR + taskPrefix + TASK_SEPARATOR)}
	iter := p.dbStatus.NewIterator(&taskRange, nil)

	return &TaskIterator{
		iter: iter,
	}
}

func (p *Pipeline) currentWorker() Worker {
	if p.currentStage < len(p.allStages) {
		return p.allStages[p.currentStage]
	} else {
		return &DoneWorker{}
	}
}

func (p *Pipeline) nextStage() bool {
	prefix := p.currentWorker().Prefix()
	if prefix == WORKER_DONE {
		return false
	}

	waiting := p.countTasks(TASK_STATE_WAITING, prefix)
	running := p.countTasks(TASK_STATE_RUNNING, prefix)

	seelog.Debugf("Current stage[%s] waiting[%d] running[%d]", prefix, waiting, running)
	if waiting == 0 && running == 0 {
		seelog.Debugf("Going to next stage from [%s]", prefix)
		p.currentStage++
		return true
	} else {
		return false
	}
}

func (p *Pipeline) countTasks(state, prefix string) int {
	iter := p.TaskIterator(state, prefix)
	defer iter.Release()
	count := 0
	for iter.Next() {
		s, t := iter.Task()
		if s == state && t.TaskPrefix == prefix {
			count++
		}
	}
	return count
}

func (p *Pipeline) isRunning() bool {
	cw := p.currentWorker()
	return cw.Prefix() != WORKER_DONE
}

func (p *Pipeline) dispatch() (deferUntil int64) {
	cw := p.currentWorker()
	if cw.Prefix() == WORKER_DONE {
		seelog.Debugf("Done")
		return 0
	}

	switch w := cw.(type) {
	case SimpleWorker:
		return p.dispatchSingle(w)

	case BatchWorker:
		return p.dispatchBatch(w)

	case ReduceWorker:
		return p.dispatchReduce(w)

	default:
		seelog.Errorf("Unsupported Worker type: prefix[%s]", cw.Prefix())
		panic("Unsupported worker type found")
	}

	return 0
}

func (p *Pipeline) DumpPipeline() {
	dbIter := p.dbStatus.NewIterator(nil, nil)
	taskIter := &TaskIterator{
		iter: dbIter,
	}

	seelog.Debugf("DumpPipeline")
	for taskIter.Next() {
		state, task := taskIter.Task()
		seelog.Debugf("State[%s] TaskPrefix[%s] TaskId[%s] Context[%s]", state, task.TaskPrefix, task.TaskId, string(task.Context))
	}
}

func (p *Pipeline) dispatchReduce(cw ReduceWorker) (deferUntil int64) {
	iter := p.TaskIterator(TASK_STATE_WAITING, cw.Prefix())
	defer iter.Release()

	cw.Reduce(iter)

	return 0
}

func (p *Pipeline) dispatchSingle(cw SimpleWorker) (deferUntil int64) {
	iter := p.TaskIterator(TASK_STATE_WAITING, cw.Prefix())
	defer iter.Release()

	var task *Task
	now := time.Now().Unix()
	for iter.Next() {
		s, t := iter.Task()
		if s != TASK_STATE_WAITING {
			seelog.Debugf("Invalid state task found: State[%s] TaskPrefix[%s] TaskId[%s]", s, t.TaskPrefix, t.TaskId)
			continue
		}
		if t.DeferUntil == 0 || t.DeferUntil < now {
			task = t
			break
		} else {
			if deferUntil < t.DeferUntil {
				deferUntil = t.DeferUntil
			}
		}
	}
	if task == nil {
		seelog.Debugf("No tasks to dispatch")
		return deferUntil
	}

	tran, err := p.dbStatus.OpenTransaction()
	if err != nil {
		seelog.Errorf("Unable to open transaction: error[%s]", err)
		panic("Unable to open transaction")
	}

	tran.Delete(taskKey(TASK_STATE_WAITING, task.TaskPrefix, task.TaskId), nil)
	tran.Put(taskKey(TASK_STATE_RUNNING, task.TaskPrefix, task.TaskId), NewTaskValueContainer(task).Value(), nil)
	err = tran.Commit()
	if err != nil {
		seelog.Errorf("Unable to commit transaction: error[%s]", err)
		panic("Unable to commit transaction")
	}

	seelog.Debugf("Exec: Prefix[%s] TaskId[%s] Context[%s]", task.TaskPrefix, task.TaskId, string(task.Context))
	cw.Exec(task)

	tran, err = p.dbStatus.OpenTransaction()
	if err != nil {
		seelog.Errorf("Unable to open transaction: error[%s]", err)
		panic("Unable to open transaction")
	}

	tran.Delete(taskKey(TASK_STATE_RUNNING, task.TaskPrefix, task.TaskId), nil)
	tran.Put(taskKey(TASK_STATE_DONE, task.TaskPrefix, task.TaskId), NewTaskValueContainer(task).Value(), nil)
	err = tran.Commit()
	if err != nil {
		seelog.Errorf("Unable to commit transaction: error[%s]", err)
		panic("Unable to commit transaction")
	}
	return 0
}

func (p *Pipeline) MarkAsDone(taskPrefix, taskId string) {
	runningKey := taskKey(TASK_STATE_WAITING, taskPrefix, taskId)
	taskValue, err := p.dbStatus.Get(runningKey, nil)
	if err != nil {
		seelog.Debugf("Unable to find task for markAsDone: Prefix[%s] TaskId[%s]", taskPrefix, taskId)
		return
	}
	_, task := NewTaskFromKeyAndValue(runningKey, taskValue)

	tran, err := p.dbStatus.OpenTransaction()
	if err != nil {
		seelog.Errorf("Unable to open transaction: error[%s]", err)
		panic("Unable to open transaction")
	}

	tran.Delete(taskKey(TASK_STATE_WAITING, taskPrefix, taskId), nil)
	tran.Put(taskKey(TASK_STATE_DONE, taskPrefix, taskId), NewTaskValueContainer(task).Value(), nil)
	err = tran.Commit()
	if err != nil {
		seelog.Errorf("Unable to commit transaction: error[%s]", err)
		panic("Unable to commit transaction")
	}
}

func (p *Pipeline) TasksRpc(tasks []*Task, apiContext *api.ApiContext, route string, arg interface{}) (cont bool, apiRes *api.ApiRpcResponse, specificErr error) {
	apiRes, err := apiContext.CallRpc(route, arg)
	if err == nil {
		return true, apiRes, nil
	}

	prefix := ""
	if len(tasks) < 1 {
		seelog.Debugf("No tasks specified: Route[%s]", route)
		prefix = "unknown"
	} else {
		prefix = tasks[0].TaskPrefix
	}

	apiRes.Error = err

	switch e := err.(type) {
	case api.ApiErrorRateLimit:
		for _, task := range tasks {
			p.RetryAfter(task, time.Now().Unix()+int64(e.RetryAfter))
			seelog.Debugf("Route[%s] Retrying Task due to ApiErrorRateLimit: TaskPrefix[%s] TaskId[%s] RetryAfter[%d]", route, task.TaskPrefix, task.TaskId, e.RetryAfter)
		}
		return false, apiRes, nil

	case api.ApiInvalidTokenError:
		seelog.Debugf("Route[%s] Invalid Token: TaskPrefix[%s] Error[%s]", route, prefix, e.Error())
		p.GeneralError("invalid_token", fmt.Sprintf("Task[%s] failed due to bad or expired token", prefix))
		return false, apiRes, nil

	case api.ApiAccessError:
		seelog.Debugf("Route[%s] Access Error: TaskPrefix[%s] Error[%s]", route, prefix, e.Error())
		p.GeneralError("access_error", fmt.Sprintf("Task[%s] failed due to access error", prefix))
		return false, apiRes, nil

	case api.ApiBadInputParamError:
		seelog.Debugf("Route[%s] Bad Input Param: TaskPrefix[%s] Error[%s]", route, prefix, e.Error())
		p.GeneralError("bad_input_param", fmt.Sprintf("Task[%s] failed due to bad input parameter", prefix))
		return false, apiRes, nil

	case api.ApiEndpointSpecificError:
		seelog.Debugf("Route[%s] API Specific: TaskPrefix[%s] Error[%s]", route, prefix, e.Error())
		return false, apiRes, e

	case api.ApiServerError:
		seelog.Debugf("Route[%s] Server Error: TaskPrefix[%s] Error[%s]", route, prefix, e.Error())
		p.GeneralError("server_error", fmt.Sprintf("Task[%s] failed due to server error. Check status.dropbox.com for announcements about Dropbox service issues.", prefix))
		return false, apiRes, nil
	}

	seelog.Debugf("Route[%s] Error: TaskPrefix[%s] Error[%s]", route, prefix, err.Error())
	p.GeneralError("error", fmt.Sprintf("Task[%s] failed due to error [%s]", prefix, err.Error()))
	return false, apiRes, nil
}

func (p *Pipeline) TaskRpc(task *Task, apiContext *api.ApiContext, route string, arg interface{}) (cont bool, apiRes *api.ApiRpcResponse, specificErr error) {
	apiRes, err := apiContext.CallRpc(route, arg)
	if err == nil {
		return true, apiRes, nil
	}

	apiRes.Error = err

	switch e := err.(type) {
	case api.ApiErrorRateLimit:
		seelog.Debugf("Route[%s] Retrying Task due to ApiErrorRateLimit: TaskPrefix[%s] TaskId[%s] RetryAfter[%d]", route, task.TaskPrefix, task.TaskId, e.RetryAfter)
		p.RetryAfter(task, time.Now().Unix()+int64(e.RetryAfter))
		return false, apiRes, nil

	case api.ApiInvalidTokenError:
		seelog.Debugf("Route[%s] Invalid Token: TaskPrefix[%s] TaskId[%s] Error[%s]", route, task.TaskPrefix, task.TaskId, e.Error())
		p.GeneralError("invalid_token", fmt.Sprintf("Task[%s] failed due to bad or expired token", task.TaskPrefix))
		return false, apiRes, nil

	case api.ApiAccessError:
		seelog.Debugf("Route[%s] Access Error: TaskPrefix[%s] TaskId[%s] Error[%s]", route, task.TaskPrefix, task.TaskId, e.Error())
		p.GeneralError("access_error", fmt.Sprintf("Task[%s] failed due to access error", task.TaskPrefix))
		return false, apiRes, nil

	case api.ApiBadInputParamError:
		seelog.Debugf("Route[%s] Bad Input Param: TaskPrefix[%s] TaskId[%s] Error[%s]", route, task.TaskPrefix, task.TaskId, e.Error())
		p.GeneralError("bad_input_param", fmt.Sprintf("Task[%s] failed due to bad input parameter", task.TaskPrefix))
		return false, apiRes, nil

	case api.ApiEndpointSpecificError:
		seelog.Debugf("Route[%s] API Specific: TaskPrefix[%s] TaskId[%s] Error[%s]", route, task.TaskPrefix, task.TaskId, e.Error())
		return false, apiRes, e

	case api.ApiServerError:
		seelog.Debugf("Route[%s] Server Error: TaskPrefix[%s] TaskId[%s] Error[%s]", route, task.TaskPrefix, task.TaskId, e.Error())
		p.GeneralError("server_error", fmt.Sprintf("Task[%s] failed due to server error. Check status.dropbox.com for announcements about Dropbox service issues.", task.TaskPrefix))
		return false, apiRes, nil
	}

	seelog.Debugf("Route[%s] Error: TaskPrefix[%s] TaskId[%s] Error[%s]", route, task.TaskPrefix, task.TaskId, err.Error())
	p.GeneralError("error", fmt.Sprintf("Task[%s] failed due to error [%s]", task.TaskPrefix, err.Error()))
	return false, apiRes, nil
}

func (p *Pipeline) RetryAfter(task *Task, deferUntil int64) {
	retry := &Task{
		TaskPrefix: task.TaskPrefix,
		TaskId:     task.TaskId,
		Context:    task.Context,
		DeferUntil: deferUntil,
	}
	p.Enqueue(retry)
}

func (p *Pipeline) dispatchBatch(cw BatchWorker) (deferUntil int64) {
	iter := p.TaskIterator(TASK_STATE_WAITING, cw.Prefix())
	defer iter.Release()

	tasks := make([]*Task, 0)
	count := 0
	now := time.Now().Unix()
	for iter.Next() {
		_, task := iter.Task()
		if task.DeferUntil == 0 || task.DeferUntil < now {
			tasks = append(tasks, task)
			count++
		} else {
			if deferUntil < task.DeferUntil {
				deferUntil = task.DeferUntil
			}
		}
		if count >= cw.BatchMaxSize() {
			break
		}
	}
	if count < 1 {
		return deferUntil
	}

	tran, err := p.dbStatus.OpenTransaction()
	if err != nil {
		seelog.Errorf("Unable to open transaction: error[%s]", err)
		panic("Unable to open transaction")
	}

	taskIds := make([]string, 0)
	for _, t := range tasks {
		taskIds = append(taskIds, t.TaskId)
		tran.Delete(taskKey(TASK_STATE_WAITING, t.TaskPrefix, t.TaskId), nil)
		tran.Put(taskKey(TASK_STATE_RUNNING, t.TaskPrefix, t.TaskId), NewTaskValueContainer(t).Value(), nil)
	}
	err = tran.Commit()
	if err != nil {
		seelog.Errorf("Unable to commit transaction: error[%s]", err)
		panic("Unable to commit transaction")
	}

	seelog.Debugf("BatchExec: Prefix[%s] TaskIds[%s]", cw.Prefix(), strings.Join(taskIds, ","))
	cw.BatchExec(tasks)

	tran, err = p.dbStatus.OpenTransaction()
	if err != nil {
		seelog.Errorf("Unable to open transaction: error[%s]", err)
		panic("Unable to open transaction")
	}

	for _, t := range tasks {
		tran.Delete(taskKey(TASK_STATE_RUNNING, t.TaskPrefix, t.TaskId), nil)
		tran.Put(taskKey(TASK_STATE_DONE, t.TaskPrefix, t.TaskId), NewTaskValueContainer(t).Value(), nil)
	}
	err = tran.Commit()
	if err != nil {
		seelog.Errorf("Unable to commit transaction: error[%s]", err)
		panic("Unable to commit transaction")
	}
	return 0
}

func (p *Pipeline) Loop() {
	for p.isRunning() {
		deferUntil := p.dispatch()
		now := time.Now().Unix()
		sleepSec := deferUntil - now
		seelog.Debugf("Dispatched: deferUntil[%d] sleepSec[%d] now[%d]", deferUntil, sleepSec, now)
		if deferUntil > 0 {
			if sleepSec > 0 {
				seelog.Debugf("Sleep until: %d (%d seconds)", deferUntil, sleepSec)
				time.Sleep(time.Duration(sleepSec)*time.Second + 100*time.Millisecond)
			}
		}
		p.nextStage()

		// TODO: fix: workaround for spin
		time.Sleep(100 * time.Millisecond)
	}
	p.DumpPipeline()
}

func (p *Pipeline) Close() {
	p.dbStatus.Close()
}

func (p *Pipeline) GeneralError(errorTag, errorDesc string) {
	p.Enqueue(
		MarshalTask(
			WORKER_GENERAL_ERROR_REPORT,
			time.Now().String(),
			ContextGeneralErrorReport{
				ErrorTag:         errorTag,
				ErrorDescription: errorDesc,
			},
		),
	)
}
