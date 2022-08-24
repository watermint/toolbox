package sv_adminrole

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_adminrole"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_user"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Role interface {
	List() (roles []*mo_adminrole.Role, err error)
	UpdateRole(user mo_user.UserSelector, roleIds []string) (roles []*mo_adminrole.Role, err error)
}

func New(ctx dbx_client.Client) Role {
	return &roleImpl{
		ctx: ctx,
	}
}

type RoleUpdate struct {
	User     mo_user.UserSelector `json:"user"`
	NewRoles []string             `json:"new_roles"`
}

type roleImpl struct {
	ctx dbx_client.Client
}

func (z roleImpl) parseRoles(res es_response.Response) (roles []*mo_adminrole.Role, err error) {
	roles = make([]*mo_adminrole.Role, 0)
	roleModel, arrayFound := res.Success().Json().FindArray("roles")
	if !arrayFound {
		return nil, errors.New("unexpected response format")
	}
	for _, rm := range roleModel {
		r := &mo_adminrole.Role{}
		if err := rm.Model(r); err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}
	return roles, nil
}

func (z roleImpl) List() (roles []*mo_adminrole.Role, err error) {
	res := z.ctx.Post("team/members/get_available_team_member_roles")
	if err, fail := res.Failure(); fail {
		return nil, err
	}

	return z.parseRoles(res)
}

func (z roleImpl) UpdateRole(user mo_user.UserSelector, roleIds []string) (roles []*mo_adminrole.Role, err error) {
	p := &RoleUpdate{
		User:     user,
		NewRoles: roleIds,
	}
	res := z.ctx.Post("team/members/set_admin_permissions_v2", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}

	return z.parseRoles(res)
}
