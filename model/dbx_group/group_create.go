package dbx_group

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
	"go.uber.org/zap"
)

type Create struct {
	OnError   func(err error) bool
	OnSuccess func(group Group)
}

func (z *Create) Create(c *dbx_api.Context, name, managementType string) error {
	p := struct {
		GroupName           string `json:"group_name"`
		GroupManagementType string `json:"group_management_type"`
	}{
		GroupName:           name,
		GroupManagementType: managementType,
	}
	switch managementType {
	case ManagementTypeUser, ManagementTypeSystem, ManagementTypeCompany:
		c.Log().Debug("create group", zap.String("name", name), zap.String("mgtType", managementType))

	default:
		c.Log().Error("Unsupported management type", zap.String("name", name), zap.String("mgtType", managementType))
		return errors.New("unsupported management type")
	}

	req := dbx_rpc.RpcRequest{
		Endpoint: "team/groups/create",
		Param:    p,
	}
	res, err := req.Call(c)
	if err != nil {
		z.OnError(err)
		return err
	}
	g := Group{}
	j := gjson.Parse(res.Body)
	if !j.Exists() {
		err = errors.New("unexpected data format")
		z.OnError(err)
		return err
	}
	err = c.ParseModel(&g, j)
	if err != nil {
		z.OnError(err)
		return err
	}
	z.OnSuccess(g)
	return nil
}
