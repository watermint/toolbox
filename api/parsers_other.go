package api

import (
	"encoding/json"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team_log"
)

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
