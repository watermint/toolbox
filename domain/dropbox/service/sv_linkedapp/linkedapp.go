package sv_linkedapp

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_linkedapp"
	"github.com/watermint/toolbox/essentials/format/tjson"
)

type LinkedApp interface {
	List() (apps []*mo_linkedapp.LinkedApp, err error)
}

func New(ctx dbx_context.Context) LinkedApp {
	return &linkedAppImpl{
		ctx: ctx,
	}
}

type linkedAppImpl struct {
	ctx dbx_context.Context
}

func (z *linkedAppImpl) List() (apps []*mo_linkedapp.LinkedApp, err error) {
	apps = make([]*mo_linkedapp.LinkedApp, 0)

	err = z.ctx.List("team/linked_apps/list_members_linked_apps").
		Continue("team/linked_apps/list_members_linked_apps").
		UseHasMore(true).
		ResultTag("apps").
		OnEntry(func(entry tjson.Json) error {
			memberId, found := entry.FindString("team_member_id")
			if !found {
				return nil
			}
			apiApps, found := entry.FindArray("linked_api_apps")
			if !found {
				return nil
			}

			for _, a := range apiApps {
				apiApp := &mo_linkedapp.LinkedApp{}
				if _, err := a.Model(apiApp); err != nil {
					return err
				}
				apiApp.TeamMemberId = memberId
				apps = append(apps, apiApp)
			}
			return nil
		}).Call()
	if err != nil {
		return nil, err
	}
	return apps, nil
}
