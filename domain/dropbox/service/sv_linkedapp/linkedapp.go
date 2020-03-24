package sv_linkedapp

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_linkedapp"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_list"
	"github.com/watermint/toolbox/infra/api/api_parser"
)

type LinkedApp interface {
	List() (apps []*mo_linkedapp.LinkedApp, err error)
}

func New(ctx api_context.DropboxApiContext) LinkedApp {
	return &linkedAppImpl{
		ctx: ctx,
	}
}

type linkedAppImpl struct {
	ctx api_context.DropboxApiContext
}

func (z *linkedAppImpl) List() (apps []*mo_linkedapp.LinkedApp, err error) {
	apps = make([]*mo_linkedapp.LinkedApp, 0)

	err = z.ctx.List("team/linked_apps/list_members_linked_apps").
		Continue("team/linked_apps/list_members_linked_apps").
		UseHasMore(true).
		ResultTag("apps").
		OnEntry(func(entry api_list.ListEntry) error {
			j, err := entry.Json()
			if err != nil {
				return err
			}
			memberId := j.Get("team_member_id").String()
			apiApps := j.Get("linked_api_apps")
			if !apiApps.Exists() || !apiApps.IsArray() {
				return nil
			}

			for _, a := range apiApps.Array() {
				apiApp := &mo_linkedapp.LinkedApp{}
				if err := api_parser.ParseModel(apiApp, a); err != nil {
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
