package sv_group

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_list"
	"github.com/watermint/toolbox/domain/model/mo_group"
)

type Group interface {
	Resolve(groupId string) (g *mo_group.Group, err error)
	List() (g []*mo_group.Group, err error)
	Create(name string, opt ...CreateOpt) (g *mo_group.Group, err error)
	Remove(groupId string) error
	Update(group *mo_group.Group) (g *mo_group.Group, err error)
}

type CreateOpt func(opt *createOpts) *createOpts
type createOpts struct {
	mgmtType string
}

func CompanyManaged() CreateOpt {
	return func(opt *createOpts) *createOpts {
		opt.mgmtType = "company_managed"
		return opt
	}
}
func UserManaged() CreateOpt {
	return func(opt *createOpts) *createOpts {
		opt.mgmtType = "user_managed"
		return opt
	}
}
func ManagementType(mgmtType string) CreateOpt {
	return func(opt *createOpts) *createOpts {
		opt.mgmtType = mgmtType
		return opt
	}
}

func New(ctx api_context.Context) Group {
	g := &implGroup{
		ctx: ctx,
	}
	return g
}

type implGroup struct {
	ctx   api_context.Context
	limit int
}

func (z *implGroup) Create(name string, opt ...CreateOpt) (g *mo_group.Group, err error) {
	co := &createOpts{}
	for _, o := range opt {
		o(co)
	}

	type MT struct {
		Tag string `json:".tag"`
	}
	p := struct {
		GroupName           string `json:"group_name"`
		GroupManagementType MT     `json:"group_management_type"`
	}{
		GroupName: name,
		GroupManagementType: MT{
			Tag: co.mgmtType,
		},
	}
	g = &mo_group.Group{}
	res, err := z.ctx.Request("team/groups/create").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err = res.Model(g); err != nil {
		return nil, err
	}
	return g, nil
}

func (z *implGroup) Remove(groupId string) error {
	p := struct {
		Tag     string `json:".tag"`
		GroupId string `json:"group_id"`
	}{
		Tag:     "group_id",
		GroupId: groupId,
	}
	_, err := z.ctx.Async("team/groups/delete").
		Status("team/groups/job_status/get").
		Param(p).Call()
	if err != nil {
		return err
	}
	return nil
}

func (z *implGroup) List() (groups []*mo_group.Group, err error) {
	groups = make([]*mo_group.Group, 0)
	p := struct {
		Limit int `json:"limit,omitempty"`
	}{
		Limit: z.limit,
	}

	req := z.ctx.List("team/groups/list").
		Continue("team/groups/list/continue").
		Param(p).
		UseHasMore(true).
		ResultTag("groups").
		OnEntry(func(entry api_list.ListEntry) error {
			g := &mo_group.Group{}
			if err := entry.Model(g); err != nil {
				return err
			}
			groups = append(groups, g)
			return nil
		})
	if err := req.Call(); err != nil {
		return nil, err
	}
	return groups, nil
}

func (z *implGroup) Resolve(groupId string) (g *mo_group.Group, err error) {
	p := struct {
		Tag      string   `json:".tag"`
		GroupIds []string `json:"group_ids"`
	}{
		Tag:      "group_ids",
		GroupIds: []string{groupId},
	}
	g = &mo_group.Group{}
	res, err := z.ctx.Request("team/groups/get_info").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err := res.ModelArrayFirst(g); err != nil {
		return nil, err
	}
	return g, nil
}

func (z *implGroup) Update(group *mo_group.Group) (g *mo_group.Group, err error) {
	panic("implement me")
}
