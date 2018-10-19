package sharedlink

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/workflow"
	"time"
)

const (
	WORKER_SHAREDLINK_UPDATE_EXPIRES = "/sharedlink/update/expires"
)

type WorkerSharedLinkUpdateExpires struct {
	workflow.SimpleWorkerImpl
	Api      *api.ApiContext
	NextTask string
	Days     int
}

type ContextSharedLinkUpdateExpiresResult struct {
	AsMemberId    string          `json:"as_member_id"`
	AsMemberEmail string          `json:"as_member_email"`
	SharedLinkId  string          `json:"shared_link_id"`
	ExpiresOld    string          `json:"expires_old"`
	ExpiresNew    string          `json:"expires_new"`
	Link          json.RawMessage `json:"link"`
}

func (w *WorkerSharedLinkUpdateExpires) Prefix() string {
	return WORKER_SHAREDLINK_UPDATE_EXPIRES
}

func (w *WorkerSharedLinkUpdateExpires) Exec(task *workflow.Task) {
	tc := &ContextSharedLinkResult{}
	workflow.UnmarshalContext(task, tc)

	link := string(tc.Link)
	expires := gjson.Get(link, "expires").String()
	var origTime time.Time

	if expires != "" {
		var err error
		origTime, err = time.Parse(api.API_DATE_TIME_FORMAT, expires)
		if err != nil {
			seelog.Warnf("SharedLinkId[%s] Unable to parse time [%s]", tc.SharedLinkId, expires)
			return
		}
	}

	targetExpire := api.RebaseTimeForAPI(time.Now().Add(time.Duration(w.Days*24) * time.Hour))
	seelog.Debugf("LinkId[%s] Link's expire time[%s] Target[%s]", tc.SharedLinkId, origTime.String(), targetExpire.String())
	if origTime.IsZero() || origTime.After(targetExpire) {
		w.update(targetExpire, origTime, tc, task)
	} else {
		seelog.Debugf("Skip LinkId[%s] Expire[%s]", tc.SharedLinkId, origTime.String())
	}
}

func (w *WorkerSharedLinkUpdateExpires) update(targetTime time.Time, origTime time.Time, tc *ContextSharedLinkResult, task *workflow.Task) {
	link := string(tc.Link)
	url := gjson.Get(link, "url").String()

	type SettingsParam struct {
		Expires string `json:"expires"`
	}
	type UpdateParam struct {
		Url      string        `json:"url"`
		Settings SettingsParam `json:"settings"`
	}

	up := UpdateParam{
		Url: url,
		Settings: SettingsParam{
			Expires: targetTime.Format(api.API_DATE_TIME_FORMAT),
		},
	}

	oldTime := origTime.Format(api.API_DATE_TIME_FORMAT)
	newTime := targetTime.Format(api.API_DATE_TIME_FORMAT)

	seelog.Infof("Updating: SharedLinkId[%s] MemberEmail[%s]: Old[%s] -> New[%s]", tc.SharedLinkId, tc.AsMemberEmail, oldTime, newTime)
	cont, res, _ := w.Pipeline.TaskRpcAsMemberId(
		task,
		w.Api,
		"sharing/modify_shared_link_settings",
		up,
		tc.AsMemberId,
	)
	if !cont {
		return
	}

	w.Pipeline.Enqueue(
		workflow.MarshalTask(
			w.NextTask,
			tc.SharedLinkId,
			ContextSharedLinkUpdateExpiresResult{
				AsMemberId:    tc.AsMemberId,
				AsMemberEmail: tc.AsMemberEmail,
				SharedLinkId:  tc.SharedLinkId,
				ExpiresOld:    oldTime,
				ExpiresNew:    newTime,
				Link:          json.RawMessage(res.Body),
			},
		),
	)
}
