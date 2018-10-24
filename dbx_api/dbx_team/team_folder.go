package dbx_team

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
)

type TeamFolder struct {
	TeamFolderId string          `json:"team_folder_id"`
	Name         string          `json:"name"`
	TeamFolder   json.RawMessage `json:"team_folder"`
}

func ParseTeamFolder(r gjson.Result) (tf *TeamFolder, annotation dbx_api.ErrorAnnotation, err error) {
	teamFolderId := r.Get("team_folder_id")
	if !teamFolderId.Exists() {
		err = errors.New("required field `team_folder_id` not found in the response")
		annotation = dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorUnexpectedDataType,
			Error:     err,
		}
		return nil, annotation, err
	}

	c := &TeamFolder{
		TeamFolderId: teamFolderId.String(),
		Name:         r.Get("name").String(),
		TeamFolder:   json.RawMessage(r.Raw),
	}
	return c, dbx_api.Success, nil
}

type TeamFolderList struct {
	OnError func(annotation dbx_api.ErrorAnnotation) bool
	OnEntry func(teamFolder *TeamFolder) bool
}

func (w *TeamFolderList) List(c *dbx_api.Context) bool {
	list := dbx_rpc.RpcList{
		EndpointList:         "team/team_folder/list",
		EndpointListContinue: "team/team_folder/list/continue",
		UseHasMore:           true,
		ResultTag:            "team_folders",
		OnError:              w.OnError,
		OnEntry: func(folder gjson.Result) bool {
			tf, ea, _ := ParseTeamFolder(folder)
			if ea.IsSuccess() {
				return w.OnEntry(tf)
			} else {
				if w.OnError != nil {
					return w.OnError(ea)
				}
				return false
			}
		},
	}

	return list.List(c, nil)
}
