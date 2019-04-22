package sequence

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_test"
	"github.com/watermint/toolbox/domain/sequence/sq_group"
	"github.com/watermint/toolbox/domain/sequence/sq_sharedfolder"
	"github.com/watermint/toolbox/domain/service"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"time"
)

type Task interface {
}

type TeamTask interface {
	Task
	Do(svc service.Business) error
}

type InterTeamTask interface {
	Task
	Do(src, dst service.Business) error
}

type Sequence interface {
	Load(path string) error
	Run(opts ...RunOpt) error
}

type RunOpt func(opts *runOpts) *runOpts
type runOpts struct {
	retryable bool
	maxRetry  int
}

func Retryable() RunOpt {
	return func(opts *runOpts) *runOpts {
		opts.retryable = true
		return opts
	}
}
func MaxRetry(max int) RunOpt {
	return func(opts *runOpts) *runOpts {
		opts.maxRetry = max
		return opts
	}
}

func New(ec *app.ExecContext) Sequence {
	return &sequenceImpl{
		ec: ec,
	}
}

type sequenceImpl struct {
	ec      *app.ExecContext
	seqPath string
	seqName string
	runId   int
}

func (z *sequenceImpl) backlogPath(runId int) string {
	return filepath.Join(z.seqPath, fmt.Sprintf("%03d.json", runId))
}

func (z *sequenceImpl) Load(path string) error {
	l := z.ec.Log().With(zap.String("filepath", path))
	z.seqName = fmt.Sprintf("%x", time.Now().Unix())
	z.seqPath = filepath.Join(z.ec.JobsPath(), "sequence", z.seqName)
	z.runId = 0
	if err := os.MkdirAll(z.seqPath, 0755); err != nil {
		l.Error("Unable to create sequence folder", zap.Error(err))
		return err
	}

	l.Debug("Opening sequence", zap.String("path", z.seqPath))
	f, err := os.Open(path)
	if err != nil {
		l.Error("Unable to open file", zap.Error(err))
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			l.Error("Unable to close", zap.Error(err))
		}
	}()

	backlog, err := os.Create(z.backlogPath(z.runId))
	if err != nil {
		l.Error("Unable to open backlog file", zap.Error(err))
		return err
	}
	defer func() {
		if err := backlog.Close(); err != nil {
			l.Error("Unable to close backlog", zap.Error(err))
		}
	}()

	scanner := bufio.NewScanner(f)
	var lastErr error = nil
	for scanner.Scan() {
		t := scanner.Text()
		_, _, _, err := Parse(t)
		if err != nil {
			l.Error("Unable to parse line", zap.Error(err), zap.String("line", t))
			lastErr = err
		}
		_, err = backlog.WriteString(t + "\n")
		if err != nil {
			l.Error("Unable to write into backlog", zap.Error(err))
			lastErr = err
		}
	}
	if lastErr != nil {
		return lastErr
	}

	return nil
}
func (z *sequenceImpl) runWithRunId(runId int) (numBacklog int, err error) {
	l := z.ec.Log()
	l.Info("Run", zap.Int("runId", runId))

	rep := app_report.Factory{
		ExecContext: z.ec,
		Path:        filepath.Join(z.seqPath, fmt.Sprintf("%03d", z.runId)),
	}
	if err := rep.Init(z.ec); err != nil {
		l.Error("Unable to prepare report", zap.Error(err))
		return 0, err
	}
	defer rep.Close()

	l.Debug("Opening backlog of runId", zap.String("backlog", z.backlogPath(z.runId)))
	f, err := os.Open(z.backlogPath(z.runId))
	if err != nil {
		l.Error("Unable to open backlog file", zap.Error(err))
		return 0, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			l.Error("Unable to close", zap.Error(err))
		}
	}()

	backlog, err := os.Create(z.backlogPath(runId + 1))
	if err != nil {
		l.Error("Unable to open next backlog file", zap.Error(err))
		return 0, err
	}
	defer func() {
		if err := backlog.Close(); err != nil {
			l.Error("Unable to close backlog", zap.Error(err))
		}
	}()

	// Poison message queue
	poison, err := os.Create(filepath.Join(z.seqPath, "poison.json"))
	if err != nil {
		l.Error("Unable to open next poison file", zap.Error(err))
		return 0, err
	}
	defer func() {
		if err := poison.Close(); err != nil {
			l.Error("Unable to close poison file", zap.Error(err))
		}
	}()

	scanner := bufio.NewScanner(f)
	numBacklog = 0
	numPoison := 0
	numSuccess := 0

	defer func() {
		l.Debug("RunId result",
			zap.Int("poison", numPoison),
			zap.Int("backlog", numBacklog),
			zap.Int("success", numSuccess),
		)
	}()

	type RunReport struct {
		TaskResult string      `json:"result"`
		TaskType   string      `json:"task"`
		Peer       interface{} `json:"peer"`
		TaskParam  interface{} `json:"task_param"`
		Reason     string      `json:"reason"`
	}

	enqueuePoison := func(line string) {
		numPoison++
		if _, err := poison.WriteString(line + "\n"); err != nil {
			l.Warn("Unable to write line into poison message queue", zap.Error(err))
		}
	}
	enqueueBacklog := func(line string, name string, peer, task interface{}, err error) {
		numBacklog++
		if _, err := backlog.WriteString(line + "\n"); err != nil {
			l.Warn("Unable to write line into poison message queue", zap.Error(err))
		}
		rr := RunReport{
			TaskResult: "failure",
			TaskType:   name,
			Peer:       peer,
			TaskParam:  task,
			Reason:     err.Error(),
		}
		if err := rep.Report(rr); err != nil {
			l.Warn("Unable to write report", zap.Error(err))
		}
	}
	reportSuccess := func(line string, name string, peer, task interface{}) {
		numSuccess++
		rr := RunReport{
			TaskResult: "success",
			TaskType:   name,
			Peer:       peer,
			TaskParam:  task,
		}
		if err := rep.Report(rr); err != nil {
			l.Warn("Unable to write report", zap.Error(err))
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		peer, name, task, err := Parse(line)
		if err != nil {
			l.Error("Unable to parse line", zap.Error(err), zap.String("line", line))
		}

		switch t := task.(type) {
		case TeamTask:
			p, ok := peer.(*PeerTeam)
			if !ok {
				l.Error("Invalid peer type")
				enqueuePoison(line)
				continue
			}
			l.Debug("Authentication with peer name", zap.String("peerName", p.PeerName))
			ctxMgmt, err := api_auth_impl.Auth(z.ec, api_auth_impl.PeerName(p.PeerName), api_auth_impl.BusinessManagement())
			if err != nil {
				enqueuePoison(line)
				continue
			}
			ctxFile, err := api_auth_impl.Auth(z.ec, api_auth_impl.PeerName(p.PeerName), api_auth_impl.BusinessFile())
			if err != nil {
				enqueuePoison(line)
				continue
			}
			biz, err := service.New(ctxMgmt.NoRetryOnError(), ctxFile.NoRetryOnError())
			if err != nil {
				enqueuePoison(line)
				continue
			}
			if err := t.Do(biz); err != nil {
				enqueueBacklog(line, name, peer, t, err)
				continue
			}
			reportSuccess(line, name, peer, t)
		}
	}

	return numBacklog, nil
}

func (z *sequenceImpl) Run(opts ...RunOpt) error {
	ro := &runOpts{}
	for _, o := range opts {
		o(ro)
	}

	if !ro.retryable {
		ro.maxRetry = 1
	}
	l := z.ec.Log().With(zap.Int("maxRetry", ro.maxRetry))

	for runId := z.runId; runId < ro.maxRetry; runId++ {
		ll := l.With(zap.Int("runId", runId))
		ll.Debug("Run")
		if num, err := z.runWithRunId(runId); err != nil {
			ll.Debug("Abort retry", zap.Error(err))
			return err
		} else {
			if num == 0 {
				ll.Debug("No more backlogs")
				return nil
			}
		}
	}
	return errors.New("one or more backlogs exists")
}

type Metadata struct {
	TaskName string          `json:"task_name"`
	TaskData json.RawMessage `json:"task_data"`
	Peer     json.RawMessage `json:"peer"`
}

type PeerTeam struct {
	PeerName string `json:"peer_name"`
}
type PeerInterTeam struct {
	PeerSource      string `json:"peer_src"`
	PeerDestination string `json:"peer_dst"`
}

func Parse(jsonString string) (peer interface{}, name string, t Task, err error) {
	l := app.Root().Log()

	meta := &Metadata{}
	if err := json.Unmarshal([]byte(jsonString), meta); err != nil {
		l.Debug("Unable to unmarshal data", zap.Error(err))
		return nil, "", nil, err
	}

	switch meta.TaskName {
	case "group/add_member":
		t = &sq_group.AddMember{}
		peer = &PeerTeam{}
	case "shared_folder/add_group":
		t = &sq_sharedfolder.AddGroup{}
		peer = &PeerTeam{}
	case "shared_folder/add_user":
		t = &sq_sharedfolder.AddUser{}
		peer = &PeerTeam{}
	case "shared_folder/mount":
		t = &sq_sharedfolder.Mount{}
		peer = &PeerTeam{}
	default:
		l.Debug("Unknown task name", zap.String("taskName", meta.TaskName))
		return nil, "", nil, errors.New("unknown task name (" + meta.TaskName + ")")
	}

	if err := json.Unmarshal(meta.Peer, peer); err != nil {
		l.Debug("Unable to unmarshal peer", zap.Error(err))
		return nil, "", nil, err
	}
	if err := json.Unmarshal(meta.TaskData, t); err != nil {
		l.Debug("Unable to unmarshal task data", zap.Error(err))
		return nil, "", nil, err
	}

	return peer, meta.TaskName, t, nil
}

func DoTestTeamTask(test func(biz service.Business)) {
	peerName := api_test.TestPeerName
	ec := app.NewExecContextForTest()
	defer ec.Shutdown()
	if !api_auth_impl.IsCacheAvailable(ec, peerName) {
		return
	}

	ctxMgmt, err := api_auth_impl.Auth(ec, api_auth_impl.PeerName(peerName), api_auth_impl.BusinessManagement())
	if err != nil {
		return
	}
	ctxFile, err := api_auth_impl.Auth(ec, api_auth_impl.PeerName(peerName), api_auth_impl.BusinessFile())
	if err != nil {
		return
	}
	biz, err := service.New(ctxMgmt, ctxFile)
	if err != nil {
		return
	}
	test(biz)
}
