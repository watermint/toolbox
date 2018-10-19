package sharedlink

import (
	"encoding/json"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/workflow"
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

