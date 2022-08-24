package sv_linkedapp

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_linkedapp"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type LinkedApp interface {
	List() (apps []*mo_linkedapp.LinkedApp, err error)
}

func New(ctx dbx_client.Client) LinkedApp {
	return &linkedAppImpl{
		ctx: ctx,
	}
}

type linkedAppImpl struct {
	ctx dbx_client.Client
}

func (z *linkedAppImpl) List() (apps []*mo_linkedapp.LinkedApp, err error) {
	apps = make([]*mo_linkedapp.LinkedApp, 0)

	res := z.ctx.List("team/linked_apps/list_members_linked_apps").Call(
		dbx_list.Continue("team/linked_apps/list_members_linked_apps"),
		dbx_list.UseHasMore(),
		dbx_list.ResultTag("apps"),
		dbx_list.OnEntry(func(entry es_json.Json) error {
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
				if err := a.Model(apiApp); err != nil {
					return err
				}
				apiApp.TeamMemberId = memberId
				apps = append(apps, apiApp)
			}
			return nil
		}),
	)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return apps, nil
}
