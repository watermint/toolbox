package group_members

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_group"
	"github.com/watermint/toolbox/model/dbx_rpc"
	"go.uber.org/zap"
)

type Add struct {
	OnError   func(err error) bool
	OnSuccess func(group dbx_group.Group)
}

func (z *Add) AddMembers(c *dbx_api.DbxContext, groupId string, teamMemberIds []string) error {
	type GS struct {
		Tag     string `json:".tag"`
		GroupId string `json:"group_id"`
	}
	type U struct {
		Tag          string `json:".tag"`
		TeamMemberId string `json:"team_member_id"`
	}
	type M struct {
		User       U      `json:"user"`
		AccessType string `json:"access_type"`
	}

	members := make([]*M, 0)
	for _, m := range teamMemberIds {
		members = append(members, &M{
			AccessType: "member",
			User: U{
				Tag:          "team_member_id",
				TeamMemberId: m,
			},
		})
	}
	p := struct {
		Group         GS   `json:"group"`
		Members       []*M `json:"members"`
		ReturnMembers bool `json:"return_members"`
	}{
		Group: GS{
			Tag:     "group_id",
			GroupId: groupId,
		},
		Members:       members,
		ReturnMembers: false,
	}

	req := dbx_rpc.RpcRequest{
		Endpoint: "team/groups/members/add",
		Param:    p,
	}
	res, err := req.Call(c)
	if err != nil {
		z.OnError(err)
		return err
	}
	as := dbx_rpc.AsyncStatus{
		Endpoint: "team/groups/job_status/get",
		OnError: func(err error) bool {
			c.Log().Error("error", zap.Error(err))
			return true
		},
		OnComplete: func(complete gjson.Result) bool {
			c.Log().Debug("complete", zap.Any("complete", complete))
			return true
		},
	}
	if as.Poll(c, res) {
		return nil
	} else {
		return errors.New("unable to add members") // TODO: with replace meaningful err
	}
}
