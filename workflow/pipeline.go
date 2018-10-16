package workflow

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	ldbutil "github.com/syndtr/goleveldb/leveldb/util"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/api/api_team"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
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

	Exec(task *Task)
	BatchMaxSize() int
	BatchExec(tasks []*Task)
}

func UnmarshalContext(t *Task, c interface{}) {
	err := json.Unmarshal(t.Context, c)
	if err != nil {
		seelog.Errorf("Unable to unmarshal context: TaskPrefix[%s]", t.TaskPrefix)
		panic("Unable to unmarshal context")
	}
}
func MarshalTask(prefix, taskId string, context interface{}) *Task {
	c, err := json.Marshal(context)
	if err != nil {
		seelog.Errorf("Unable to unmarshal context: TaskPrefix[%s]", prefix)
		panic("Unable to unmarshal context")
	}
	return &Task{
		TaskPrefix: prefix,
		TaskId:     taskId,
		Context:    c,
	}
}

type BatchWorker struct {
}

func (w *BatchWorker) Exec(task *Task) {
	seelog.Debugf("Ignore `Exec`: TaskPrefix[%s] TaskId[%s]", task.TaskPrefix, task.TaskId)
}

type SimpleWorker struct {
}

func (w *SimpleWorker) BatchMaxSize() int {
	return 0
}
func (w *SimpleWorker) BatchExec(tasks []*Task) {
	seelog.Debugf("Ignore `BatchExec`")
}

type DoneWorker struct {
	*SimpleWorker
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
func (w *DoneWorker) Exec(t *Task) {
	seelog.Tracef("Done")
}

type WorkerTeamMemberInviteLoaderCsv struct {
	*SimpleWorker
	Pipeline *Pipeline
}

func (w *WorkerTeamMemberInviteLoaderCsv) SetPipeline(p *Pipeline) {
	w.Pipeline = p
}

func (w *WorkerTeamMemberInviteLoaderCsv) Prefix() string {
	return WORKER_TEAM_MEMBER_INVITE_LOADER_CSV
}

func (w *WorkerTeamMemberInviteLoaderCsv) Exec(task *Task) {
	tc := &ContextTeamMemberInviteLoaderCsv{}
	UnmarshalContext(task, tc)

	f, err := os.Open(tc.Path)
	if err != nil {
		seelog.Warnf("Unable to open file[%s] : error[%s]", tc.Path, err)
		//TODO Error report
		return
	}
	csv := util.NewBomAwareCsvReader(f)

	for {
		cols, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			seelog.Warnf("Unable to read CSV file [%s] : error[%s]", tc.Path, err)
			return
		}
		if len(cols) < 1 {
			seelog.Warnf("Skip line: [%v]", cols)
			continue
		}
		var email, givenName, surName string
		email = cols[0]
		if len(cols) >= 2 {
			givenName = cols[1]
		}
		if len(cols) >= 3 {
			surName = cols[2]
		}

		w.Pipeline.Enqueue(NewTaskTeamMemberInvite(email, givenName, surName))
	}
}

type ContextTeamMemberInviteLoaderCsv struct {
	Path string `json:"path"`
}

const (
	WORKER_TEAM_MEMBER_INVITE                   = "team/member/invite"
	WORKER_TEAM_MEMBER_INVITE_LOADER_CSV        = "team/member/invite/loader/csv"
	WORKER_TEAM_MEMBER_INVITE_RESULT_ASYNC      = "team/member/invite/result/async"
	WORKER_TEAM_MEMBER_INVITE_RESULT_INDIVIDUAL = "team/member/invite/result/individual"
)

func NewTaskTeamMemberInviteLoaderCsv(path string) *Task {
	return MarshalTask(
		WORKER_TEAM_MEMBER_INVITE_LOADER_CSV,
		path,
		&ContextTeamMemberInviteLoaderCsv{
			Path: path,
		},
	)
}

type WorkerTeamMemberInvite struct {
	*BatchWorker
	Pipeline      *Pipeline
	ApiManagement *api.ApiContext
	Silent        bool
}

type ContextTeamMemberInvite struct {
	Email     string `json:"email"`
	GivenName string `json:"givenName"`
	SurName   string `json:"surName"`
}

func NewTaskTeamMemberInvite(email, givenName, surName string) *Task {
	return MarshalTask(
		WORKER_TEAM_MEMBER_INVITE,
		email,
		&ContextTeamMemberInvite{
			Email:     email,
			GivenName: givenName,
			SurName:   surName,
		},
	)
}
func (w *WorkerTeamMemberInvite) SetPipeline(p *Pipeline) {
	w.Pipeline = p
}
func (w *WorkerTeamMemberInvite) Prefix() string {
	return WORKER_TEAM_MEMBER_INVITE
}
func (w *WorkerTeamMemberInvite) BatchMaxSize() int {
	return 20
}
func (w *WorkerTeamMemberInvite) BatchExec(tasks []*Task) {
	invites := make([]api_team.ArgMemberAdd, 0)
	for _, t := range tasks {
		tc := ContextTeamMemberInvite{}
		UnmarshalContext(t, &tc)

		invite := api_team.ArgMemberAdd{
			MemberEmail:     tc.Email,
			MemberSurname:   tc.SurName,
			MemberGivenName: tc.GivenName,
		}
		if w.Silent {
			invite.SendWelcomeEmail = true
		}
		invites = append(invites, invite)
	}

	arg := api_team.ArgMembersAdd{
		ForceAsync: true,
		NewMembers: invites,
	}
	seelog.Debugf("AddMembersAdd Arg: [%s]", util.MarshalObjectToString(arg))

	res, err := w.ApiManagement.CallRpc("team/members/add", arg)

	if err != nil {
		seelog.Errorf("Error: %s", err)
		return
	}

	seelog.Infof("ResponseTag[%s] Response[%s]", res.Tag, res.Body)

	asyncJobId := gjson.Get(res.Body, "async_job_id")
	if asyncJobId.Exists() {
		w.Pipeline.Enqueue(NewTaskTeamMemberInviteResultAsync(asyncJobId.String()))
	}
}

type WorkerTeamMemberInviteResultAsync struct {
	*SimpleWorker
	Pipeline      *Pipeline
	ApiManagement *api.ApiContext
}
type ContextTeamMemberInviteResultAsync struct {
	AsyncJobId string `json:"asyncJobId"`
}

const ()

func NewTaskTeamMemberInviteResultAsync(asyncJobId string) *Task {
	return MarshalTask(
		WORKER_TEAM_MEMBER_INVITE_RESULT_ASYNC,
		asyncJobId,
		ContextTeamMemberInviteResultAsync{
			AsyncJobId: asyncJobId,
		},
	)
}
func (w *WorkerTeamMemberInviteResultAsync) SetPipeline(p *Pipeline) {
	w.Pipeline = p
}
func (w *WorkerTeamMemberInviteResultAsync) Prefix() string {
	return WORKER_TEAM_MEMBER_INVITE_RESULT_ASYNC
}
func (w *WorkerTeamMemberInviteResultAsync) Exec(task *Task) {
	tc := &ContextTeamMemberInviteResultAsync{}
	UnmarshalContext(task, tc)

	pa := api.ArgAsyncJobId{
		AsyncJobId: tc.AsyncJobId,
	}
	res, err := w.ApiManagement.CallRpc("team/members/add/job_status/get", pa)

	if err != nil {
		seelog.Errorf("Error: %s", err)
		return
	}
	seelog.Debugf("Tag: Tag[%s] Body", res.Tag, res.Body)
	switch res.Tag {
	case "in_progress":
		seelog.Debugf("In Progress Async[%s]", tc.AsyncJobId)

	case "complete":
		seelog.Debugf("Complete[%s]", tc.AsyncJobId)

	case "failed":
		seelog.Debugf("Failed [%s]", tc.AsyncJobId)
	}
}

type ContextTeamMemberInviteResultIndividual struct {
}

func NewTaskTeamMemberInviteResultIndividual(email string, errorTag, errorDetail string) *Task {
	return MarshalTask(
		WORKER_TEAM_MEMBER_INVITE_RESULT_INDIVIDUAL,
		email,
		&ContextTeamMemberInviteResultIndividual{},
	)
}

type TaskValueContainer struct {
	DeferUntil int64
	Context    []byte
}

func (t *TaskValueContainer) Value() []byte {
	buf := bytes.NewBuffer(nil)
	err := gob.NewEncoder(buf).Encode(t)
	if err != nil {
		seelog.Errorf("Unable to decode value: error[%s]", err)
		panic("Unable to encode value")
	}
	return buf.Bytes()
}

func NewTaskValueContainer(task *Task) *TaskValueContainer {
	return &TaskValueContainer{
		DeferUntil: task.DeferUntil,
		Context:    task.Context,
	}
}
func NewTaskFromKeyAndValue(key []byte, value []byte) (state string, task *Task) {
	state, taskPrefix, taskId := parseKey(key)
	buf := bytes.NewBuffer(value)
	tvc := TaskValueContainer{}
	err := gob.NewDecoder(buf).Decode(&tvc)
	if err != nil {
		seelog.Errorf("Unable to decode value: error[%s]", err)
		panic("Unable to decode value")
	}
	task = &Task{
		TaskPrefix: taskPrefix,
		TaskId:     taskId,
		Context:    tvc.Context,
		DeferUntil: tvc.DeferUntil,
	}
	return
}

type Pipeline struct {
	infra        *infra.InfraContext
	dbStatus     *leveldb.DB
	currentStage int

	Stages []Worker
}

func (p *Pipeline) Init() {
	for _, w := range p.Stages {
		w.SetPipeline(p)
	}
}

const (
	TASK_STATE_WAITING = "W"
	TASK_STATE_RUNNING = "R"
	TASK_STATE_DONE    = "D"
	TASK_SEPARATOR     = "\t"
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
	seelog.Infof("EnqueueTask: Prefix[%s] TaskId[%s] Context[%s]", task.TaskPrefix, task.TaskId, string(task.Context))
	err := p.dbStatus.Put(taskKey(TASK_STATE_WAITING, task.TaskPrefix, task.TaskId), NewTaskValueContainer(task).Value(), nil)
	if err != nil {
		seelog.Errorf("Unable to put task: Prefix[%s] TaskId[%s]", task.TaskPrefix, task.TaskId)
		panic("Unable to put task")
	}
}
func (p *Pipeline) TaskIterator(state, taskPrefix string) *TaskIterator {
	taskRange := ldbutil.Range{Start: []byte(state + TASK_SEPARATOR + taskPrefix + TASK_SEPARATOR)}
	iter := p.dbStatus.NewIterator(&taskRange, nil)

	return &TaskIterator{
		iter: iter,
	}
}

func (p *Pipeline) currentWorker() Worker {
	if p.currentStage < len(p.Stages) {
		return p.Stages[p.currentStage]
	} else {
		return &DoneWorker{}
	}
}

func (p *Pipeline) nextStage() bool {
	prefix := p.currentWorker().Prefix()
	if prefix == WORKDER_DONE {
		return false
	}

	waiting := p.countTasks(TASK_STATE_WAITING, prefix)
	running := p.countTasks(TASK_STATE_RUNNING, prefix)

	seelog.Infof("Current stage[%s] waiting[%d] running[%d]", prefix, waiting, running)
	if waiting == 0 && running == 0 {
		seelog.Infof("Going to next stage from [%s]", prefix)
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
	return cw.Prefix() != WORKDER_DONE
}

func (p *Pipeline) dispatch() (deferUntil int64) {
	cw := p.currentWorker()
	if cw.Prefix() == WORKDER_DONE {
		seelog.Debugf("Done")
		return 0
	}
	if cw.BatchMaxSize() <= 1 {
		return p.dispatchSingle(cw)
	} else {
		return p.dispatchBatch(cw)
	}
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

func (p *Pipeline) dispatchSingle(cw Worker) (deferUntil int64) {
	iter := p.TaskIterator(TASK_STATE_WAITING, cw.Prefix())
	defer iter.Release()

	var task *Task
	for iter.Next() {
		s, t := iter.Task()
		if s != TASK_STATE_WAITING {
			seelog.Debugf("Invalid state task found: State[%s] TaskPrefix[%s] TaskId[%s]", s, t.TaskPrefix, t.TaskId)
			continue
		}
		if t.DeferUntil == 0 || t.DeferUntil < time.Now().Unix() {
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

	seelog.Infof("Exec: Prefix[%s] TaskId[%s] Context[%s]", task.TaskPrefix, task.TaskId, string(task.Context))
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

func (p *Pipeline) dispatchBatch(cw Worker) (deferUntil int64) {
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

	seelog.Infof("BatchExec: Prefix[%s] TaskIds[%s]", cw.Prefix(), strings.Join(taskIds, ","))
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

func (p *Pipeline) dispatchLoop() {
	for p.isRunning() {
		p.dispatch()
		p.DumpPipeline()
		p.nextStage()
		time.Sleep(100 * time.Millisecond)
	}
}

type TaskIterator struct {
	iter iterator.Iterator
}

func (t *TaskIterator) Next() bool {
	return t.iter.Next()
}
func (t *TaskIterator) Prev() bool {
	return t.iter.Prev()
}
func (t *TaskIterator) First() bool {
	return t.iter.First()
}
func (t *TaskIterator) Last() bool {
	return t.iter.Last()
}
func (t *TaskIterator) Release() {
	t.iter.Release()
}
func (t *TaskIterator) Error() error {
	return t.iter.Error()
}
func (t *TaskIterator) Key() (state, prefix, taskId string) {
	return parseKey(t.iter.Key())
}
func (t *TaskIterator) Value() []byte {
	return t.iter.Value()
}
func (t *TaskIterator) Task() (state string, task *Task) {
	return NewTaskFromKeyAndValue(t.iter.Key(), t.Value())
}

func PipelinePoc() error {
	c := infra.InfraContext{}
	c.Startup()
	defer c.Shutdown()

	db, err := leveldb.OpenFile(filepath.Join(c.WorkPath, "workflow"), nil)
	if err != nil {
		seelog.Warnf("Unable to open pipeline database : error[%s]", err)
		return err
	}
	defer db.Close()

	apiMgmt, err := c.LoadOrAuthBusinessManagement()

	p := Pipeline{
		infra:    &c,
		dbStatus: db,
		Stages: []Worker{
			&WorkerTeamMemberInviteLoaderCsv{},
			&WorkerTeamMemberInvite{ApiManagement: apiMgmt, Silent: true},
			&WorkerTeamMemberInviteResultAsync{ApiManagement: apiMgmt},
		},
	}
	p.Init()

	p.Enqueue(NewTaskTeamMemberInviteLoaderCsv("invite.csv"))
	p.dispatchLoop()

	return nil
}
