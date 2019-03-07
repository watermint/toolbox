package dbx_teamfolder

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
)

type Create struct {
	SyncSetting string
	OnError     func(annotation dbx_api.ErrorAnnotation) bool
	OnSuccess   func(teamFolder TeamFolder)
}

func (z *Create) Create(c *dbx_api.Context, name string) error {
	p := struct {
		Name        string `json:"name"`
		SyncSetting string `json:"sync_setting,omitempty"`
	}{
		Name:        name,
		SyncSetting: z.SyncSetting,
	}

	req := dbx_rpc.RpcRequest{
		Endpoint: "team/team_folder/create",
		Param:    p,
	}
	res, ea, err := req.Call(c)
	if ea.IsFailure() {
		z.OnError(ea)
		return err
	}
	tf := TeamFolder{}
	j := gjson.Parse(res.Body)
	if !j.Exists() {
		err = errors.New("unexpected data format")
		z.OnError(dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorUnexpectedDataType,
			Error:     err,
		})
		return err
	}
	err = c.ParseModel(&tf, j)
	if err != nil {
		z.OnError(dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorUnexpectedDataType,
			Error:     err,
		})
		return err
	}
	return nil
}
