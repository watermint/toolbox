package group_members

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_group"
	"github.com/watermint/toolbox/model/dbx_rpc"
	"go.uber.org/zap"
)

type Remove struct {
	OnError   func(err error) bool
	OnSuccess func(group dbx_group.Group)
}

func (z *Remove) RemoveMembers(c *dbx_api.Context, groupId string, teamMemberIds []string) error {
	type GS struct {
		Tag     string `json:".tag"`
		GroupId string `json:"group_id"`
	}
	type U struct {
		Tag          string `json:".tag"`
		TeamMemberId string `json:"team_member_id"`
	}

	users := make([]*U, 0)
	for _, m := range teamMemberIds {
		users = append(users, &U{
			Tag:          "team_member_id",
			TeamMemberId: m,
		})
	}
	p := struct {
		Group         GS   `json:"group"`
		Users         []*U `json:"users"`
		ReturnMembers bool `json:"return_members"`
	}{
		Group: GS{
			Tag:     "group_id",
			GroupId: groupId,
		},
		Users:         users,
		ReturnMembers: false,
	}

	req := dbx_rpc.RpcRequest{
		Endpoint: "team/groups/members/remove",
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
		return errors.New("unable to remove members") // TODO: with replace meaningful err
	}
}
