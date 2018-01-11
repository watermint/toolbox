package api

import (
	"encoding/json"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team_log"
)

type ApiTeamLog struct {
	Context *ApiContext
}

func (a *ApiTeamLog) Compat() team_log.Client {
	return team_log.New(a.Context.compatConfig())
}

func (a *ApiTeamLog) GetEvents(arg *team_log.GetTeamEventsArg) (res *team_log.GetTeamEventsResult, err error) {
	return a.Compat().GetEvents(arg)
}

func (a *ApiTeamLog) GetEventsContinue(arg *team_log.GetTeamEventsContinueArg) (res *team_log.GetTeamEventsResult, err error) {
	return a.Compat().GetEventsContinue(arg)
}

type GetTeamEventsRawResult struct {
	Cursor  string            `json:"cursor"`
	HasMore bool              `json:"has_more"`
	Events  []json.RawMessage `json:"events,omitempty"`
}

func parseGetTeamEventsRawResult(res *ApiRpcResponse) (r *GetTeamEventsRawResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}

func parseGetEventsAPIError(body []byte) error {
	var apiErr team_log.GetEventsAPIError
	if err := json.Unmarshal(body, &apiErr); err != nil {
		return err
	}
	return apiErr
}

func (a *ApiTeamLog) RawGetEvents(arg *team_log.GetTeamEventsArg) (r *GetTeamEventsRawResult, err error) {
	if res, err := a.Context.NewApiRpcRequest("team_log/get_events", parseGetEventsAPIError, arg).Call(); err != nil {
		return nil, err
	} else {
		return parseGetTeamEventsRawResult(res)
	}
}

func (a *ApiTeamLog) RawGetEventsContinue(arg *team_log.GetTeamEventsContinueArg) (res *GetTeamEventsRawResult, err error) {
	if res, err := a.Context.NewApiRpcRequest("team_log/get_events/continue", parseGetEventsAPIError, arg).Call(); err != nil {
		return nil, err
	} else {
		return parseGetTeamEventsRawResult(res)
	}
}
