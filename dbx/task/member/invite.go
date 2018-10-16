package member

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/api/api_team"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/workflow"
	"io"
	"os"
	"time"
)

const (
	WORKER_TEAM_MEMBER_INVITE              = "team/member/invite"
	WORKER_TEAM_MEMBER_INVITE_LOADER_CSV   = "team/member/invite/loader/csv"
	WORKER_TEAM_MEMBER_INVITE_RESULT_ASYNC = "team/member/invite/result/async"
	WORKER_TEAM_MEMBER_INVITE_RESULT       = "team/member/invite/result"
)

type WorkerTeamMemberInviteLoaderCsv struct {
	workflow.SimpleWorkerImpl
}

func (w *WorkerTeamMemberInviteLoaderCsv) Prefix() string {
	return WORKER_TEAM_MEMBER_INVITE_LOADER_CSV
}

func (w *WorkerTeamMemberInviteLoaderCsv) Exec(task *workflow.Task) {
	tc := &ContextTeamMemberInviteLoaderCsv{}
	workflow.UnmarshalContext(task, tc)

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

		w.SimpleWorkerImpl.Pipeline.Enqueue(NewTaskTeamMemberInvite(email, givenName, surName))
	}
}

type ContextTeamMemberInviteLoaderCsv struct {
	Path string `json:"path"`
}

func NewTaskTeamMemberInviteLoaderCsv(path string) *workflow.Task {
	return workflow.MarshalTask(
		WORKER_TEAM_MEMBER_INVITE_LOADER_CSV,
		path,
		&ContextTeamMemberInviteLoaderCsv{
			Path: path,
		},
	)
}

type WorkerTeamMemberInvite struct {
	workflow.BatchWorkerImpl
	ApiManagement *api.ApiContext
	Silent        bool
}

type ContextTeamMemberInvite struct {
	Email     string `json:"email"`
	GivenName string `json:"givenName"`
	SurName   string `json:"surName"`
}

func NewTaskTeamMemberInvite(email, givenName, surName string) *workflow.Task {
	return workflow.MarshalTask(
		WORKER_TEAM_MEMBER_INVITE,
		email,
		&ContextTeamMemberInvite{
			Email:     email,
			GivenName: givenName,
			SurName:   surName,
		},
	)
}
func (w *WorkerTeamMemberInvite) Prefix() string {
	return WORKER_TEAM_MEMBER_INVITE
}
func (w *WorkerTeamMemberInvite) BatchMaxSize() int {
	return 20
}
func (w *WorkerTeamMemberInvite) BatchExec(tasks []*workflow.Task) {
	invites := make([]api_team.ArgMemberAdd, 0)
	emails := make([]string, 0)
	for _, t := range tasks {
		tc := ContextTeamMemberInvite{}
		workflow.UnmarshalContext(t, &tc)

		invite := api_team.ArgMemberAdd{
			MemberEmail:     tc.Email,
			MemberSurname:   tc.SurName,
			MemberGivenName: tc.GivenName,
		}
		if w.Silent {
			invite.SendWelcomeEmail = true
		}
		invites = append(invites, invite)
		emails = append(emails, tc.Email)
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
		w.BatchWorkerImpl.Pipeline.Enqueue(NewTaskTeamMemberInviteResultAsync(asyncJobId.String(), emails))
	} else {
		seelog.Errorf("Async Job Id not found in the response: Response[%s]", res.Body)
	}
}

type WorkerTeamMemberInviteResultAsync struct {
	workflow.SimpleWorkerImpl
	ApiManagement *api.ApiContext
}
type ContextTeamMemberInviteResultAsync struct {
	AsyncJobId   string   `json:"asyncJobId"`
	MemberEmails []string `json:"memberEmails"`
}

func NewTaskTeamMemberInviteResultAsync(asyncJobId string, memberEmails []string) *workflow.Task {
	return workflow.MarshalTask(
		WORKER_TEAM_MEMBER_INVITE_RESULT_ASYNC,
		asyncJobId,
		ContextTeamMemberInviteResultAsync{
			AsyncJobId:   asyncJobId,
			MemberEmails: memberEmails,
		},
	)
}
func (w *WorkerTeamMemberInviteResultAsync) Prefix() string {
	return WORKER_TEAM_MEMBER_INVITE_RESULT_ASYNC
}
func (w *WorkerTeamMemberInviteResultAsync) Exec(task *workflow.Task) {
	tc := &ContextTeamMemberInviteResultAsync{}
	workflow.UnmarshalContext(task, tc)

	pa := api.ArgAsyncJobId{
		AsyncJobId: tc.AsyncJobId,
	}
	res, err := w.ApiManagement.CallRpc("team/members/add/job_status/get", pa)
	if w.SimpleWorkerImpl.Pipeline.HandleRateLimit(err, task) {
		return
	}
	if err != nil {
		seelog.Errorf("Error: %s", err)
		return
	}
	seelog.Debugf("Tag: Tag[%s] Body", res.Tag, res.Body)
	switch res.Tag {
	case "in_progress":
		seelog.Debugf("In Progress Async[%s]", tc.AsyncJobId)
		w.reactInProgress(res, task, tc)

	case "complete":
		seelog.Debugf("Complete[%s]", tc.AsyncJobId)
		w.reactComplete(res, task, tc)

	case "failed":
		seelog.Debugf("Failed [%s]", tc.AsyncJobId)
		w.reactFailed(res, task, tc)
	}
}

func (w *WorkerTeamMemberInviteResultAsync) reactInProgress(res *api.ApiRpcResponse, task *workflow.Task, tc *ContextTeamMemberInviteResultAsync) {
	w.SimpleWorkerImpl.Pipeline.RetryAfter(task, time.Now().Unix()+5)
}

func (w *WorkerTeamMemberInviteResultAsync) reactComplete(res *api.ApiRpcResponse, task *workflow.Task, tc *ContextTeamMemberInviteResultAsync) {
	completeJson := gjson.Get(res.Body, "complete")
	if !completeJson.Exists() {
		// TODO Error handling
		seelog.Errorf("`complete` tag not found in the response: Response[%s]", res.Body)
		return
	}

	for _, complete := range completeJson.Array() {
		tagJson := complete.Get(api.API_RES_JSON_DOT_TAG)
		if !tagJson.Exists() {
			// TODO Error handling
			seelog.Errorf("`complete.\\.tag` not found: Response[%s]", res.Body)
			return
		}

		tag := tagJson.String()

		if tag == "success" {
			emailTag := complete.Get("profile.email")
			if !emailTag.Exists() {
				seelog.Debugf("Ignore unexpected result: `complete.profile.email` not found: Response[%s]", res.Body)
				continue
			}

			w.SimpleWorkerImpl.Pipeline.Enqueue(
				NewTaskTeamMemberInviteResult(
					&ContextTeamMemberInviteResult{
						Email:     emailTag.String(),
						IsSuccess: true,
						Success: ContextTeamMemberInviteResultSuccess{
							Success: json.RawMessage(complete.Raw),
						},
						Failure: ContextTeamMemberInviteResultFailure{},
					},
				),
			)
		} else {
			emailTag := complete.Get(tag)
			if !emailTag.Exists() {
				seelog.Debugf("Ignore unexpected result: `complete.%s` not found: Response[%s]", tag, res.Body)
				continue
			}

			w.SimpleWorkerImpl.Pipeline.Enqueue(
				NewTaskTeamMemberInviteResult(
					&ContextTeamMemberInviteResult{
						Email:     emailTag.String(),
						IsSuccess: false,
						Success:   ContextTeamMemberInviteResultSuccess{},
						Failure: ContextTeamMemberInviteResultFailure{
							ReasonTag:    tag,
							ReasonDetail: "",
						},
					},
				),
			)
		}
	}
}

func (w *WorkerTeamMemberInviteResultAsync) reactFailed(res *api.ApiRpcResponse, task *workflow.Task, tc *ContextTeamMemberInviteResultAsync) {
	reasonJson := gjson.Get(res.Body, "failed")
	reason := "unknown"
	if reasonJson.Exists() {
		reason = reasonJson.String()
	}

	for _, email := range tc.MemberEmails {
		w.SimpleWorkerImpl.Pipeline.Enqueue(
			NewTaskTeamMemberInviteResult(
				&ContextTeamMemberInviteResult{
					Email:     email,
					IsSuccess: false,
					Success:   ContextTeamMemberInviteResultSuccess{},
					Failure: ContextTeamMemberInviteResultFailure{
						ReasonTag:    reason,
						ReasonDetail: "",
					},
				},
			),
		)
	}
}

type ContextTeamMemberInviteResultSuccess struct {
	Success json.RawMessage `json:"raw"`
}
type ContextTeamMemberInviteResultFailure struct {
	ReasonTag    string `json:"reasonTag"`
	ReasonDetail string `json:"reasonDetail"`
}

type ContextTeamMemberInviteResult struct {
	Email     string                               `json:"email"`
	IsSuccess bool                                 `json:"isSuccess"`
	Success   ContextTeamMemberInviteResultSuccess `json:"success"`
	Failure   ContextTeamMemberInviteResultFailure `json:"failure"`
}

func NewTaskTeamMemberInviteResult(ctx *ContextTeamMemberInviteResult) *workflow.Task {
	return workflow.MarshalTask(
		WORKER_TEAM_MEMBER_INVITE_RESULT,
		ctx.Email,
		ctx,
	)
}

type WorkerTeamMemberInviteResultReduce struct {
	workflow.ReduceWorkerImpl
}

func (w *WorkerTeamMemberInviteResultReduce) Prefix() string {
	return WORKER_TEAM_MEMBER_INVITE_RESULT
}

func (w *WorkerTeamMemberInviteResultReduce) Reduce(taskIter *workflow.TaskIterator) {
	for taskIter.Next() {
		_, task := taskIter.Task()
		tc := ContextTeamMemberInviteResult{}
		workflow.UnmarshalContext(task, &tc)

		seelog.Infof("Reduce: Email[%s] IsSuccess[%t] Success[%s] Failure[%s]", tc.Email, tc.IsSuccess, tc.Success, tc.Failure)

		w.Pipeline.MarkAsDone(task.TaskPrefix, task.TaskId)
	}
}
