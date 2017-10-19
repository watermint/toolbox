package sdk

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"net/http"
)

type ZRelocationBatchLaunch struct {
	dropbox.Tagged

	AsyncJobId string `json:"async_job_id,omitempty"`

	// Complete : has no documentation (yet)
	Complete *files.RelocationBatchResult `json:"complete,omitempty"`
}

func (u *ZRelocationBatchLaunch) UnmarshalJSON(body []byte) error {
	type wrap struct {
		dropbox.Tagged

		AsyncJobId string `json:"async_job_id,omitempty"`

		// Complete : has no documentation (yet)
		Complete json.RawMessage `json:"complete,omitempty"`
	}
	var w wrap
	var err error
	if err = json.Unmarshal(body, &w); err != nil {
		return err
	}
	u.AsyncJobId = w.AsyncJobId
	u.Tag = w.Tag
	switch u.Tag {
	case "complete":
		err = json.Unmarshal(body, &u.Complete)

		if err != nil {
			return err
		}
	}
	return nil
}

func ZMoveBatch(cfg dropbox.Config, arg *files.RelocationBatchArg) (res ZRelocationBatchLaunch, err error) {
	ctx := dropbox.NewContext(cfg)
	dbx := &ctx
	cli := dbx.Client

	dbx.Config.LogDebug("arg: %v", arg)
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	if dbx.Config.AsMemberID != "" {
		headers["Dropbox-API-Select-User"] = dbx.Config.AsMemberID
	}

	req, err := (*dropbox.Context)(dbx).NewRequest("api", "rpc", true, "files", "move_batch", headers, bytes.NewReader(b))
	if err != nil {
		return
	}
	dbx.Config.LogInfo("req: %v", req)

	resp, err := cli.Do(req)
	if err != nil {
		return
	}

	dbx.Config.LogInfo("resp: %v", resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	dbx.Config.LogDebug("body: %v", body)
	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(body, &res)
		if err != nil {
			return
		}

		return
	}
	if resp.StatusCode == http.StatusConflict {
		var apiError files.MoveBatchAPIError
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	var apiError dropbox.APIError
	if resp.StatusCode == http.StatusBadRequest {
		apiError.ErrorSummary = string(body)
		err = apiError
		return
	}
	err = json.Unmarshal(body, &apiError)
	if err != nil {
		return
	}
	err = apiError
	return
}
