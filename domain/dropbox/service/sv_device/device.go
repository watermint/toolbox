package sv_device

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_device"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Session interface {
	List() (sessions []mo_device.Session, err error)
	Revoke(session mo_device.Session, opts ...RevokeOption) (err error)
}

func DeleteOnUnlink() RevokeOption {
	return func(opt *revokeOptions) *revokeOptions {
		opt.deleteOnUnlink = true
		return opt
	}
}

type RevokeOption func(opt *revokeOptions) *revokeOptions

type revokeOptions struct {
	deleteOnUnlink bool
}

func New(ctx dbx_context.Context) Session {
	return &sessionImpl{
		ctx: ctx,
	}
}

type sessionImpl struct {
	ctx dbx_context.Context
}

func (z *sessionImpl) List() (sessions []mo_device.Session, err error) {
	sessions = make([]mo_device.Session, 0)
	p := struct {
		IncludeWebSessions    bool `json:"include_web_sessions"`
		IncludeDesktopClients bool `json:"include_desktop_clients"`
		IncludeMobileClients  bool `json:"include_mobile_clients"`
	}{
		IncludeWebSessions:    true,
		IncludeDesktopClients: true,
		IncludeMobileClients:  true,
	}
	onEntry := func(entry es_json.Json) error {
		m, found := entry.FindString("team_member_id")
		if !found {
			z.ctx.Log().Debug("no `team_member_id` field found", es_log.ByteString("entry", entry.Raw()))
			return errors.New("team_member_id not found")
		}
		teamMemberId := m
		{
			ws, found := entry.FindArray("web_sessions")
			if found {
				for _, w := range ws {
					mw := &mo_device.Web{}
					if err := w.Model(mw); err != nil {
						z.ctx.Log().Debug("unable to parse web_session", es_log.Error(err), es_log.ByteString("entry", entry.Raw()))
						return err
					}
					mw.TeamMemberId = teamMemberId
					mw.Tag = mo_device.DeviceTypeWeb
					sessions = append(sessions, mw)
				}
			}
		}
		{
			ds, found := entry.FindArray("desktop_clients")
			if found {
				for _, d := range ds {
					md := &mo_device.Desktop{}
					if err := d.Model(md); err != nil {
						z.ctx.Log().Debug("unable to parse desktop_session", es_log.Error(err), es_log.ByteString("entry", entry.Raw()))
						return err
					}
					md.TeamMemberId = teamMemberId
					md.Tag = mo_device.DeviceTypeDesktop
					sessions = append(sessions, md)
				}
			}
		}
		{
			ms, found := entry.FindArray("mobile_clients")
			if found {
				for _, m := range ms {
					mm := &mo_device.Mobile{}
					if err := m.Model(mm); err != nil {
						z.ctx.Log().Debug("unable to parse desktop_session", es_log.Error(err), es_log.ByteString("entry", entry.Raw()))
						return err
					}
					mm.TeamMemberId = teamMemberId
					mm.Tag = mo_device.DeviceTypeMobile
					sessions = append(sessions, mm)
				}
			}
		}
		return nil
	}
	res := z.ctx.List("team/devices/list_members_devices", api_request.Param(p)).Call(
		dbx_list.Continue("team/devices/list_members_devices"),
		dbx_list.UseHasMore(),
		dbx_list.ResultTag("devices"),
		dbx_list.OnEntry(onEntry),
	)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return sessions, nil
}

func (z *sessionImpl) Revoke(session mo_device.Session, opts ...RevokeOption) (err error) {
	ro := &revokeOptions{}
	for _, o := range opts {
		o(ro)
	}
	p := struct {
		Tag            string `json:".tag"`
		SessionId      string `json:"session_id"`
		TeamMemberId   string `json:"team_member_id"`
		DeleteOnUnlink bool   `json:"delete_on_unlink,omitempty"`
	}{
		Tag:            session.EntryTag(),
		SessionId:      session.SessionId(),
		TeamMemberId:   session.EntryTeamMemberId(),
		DeleteOnUnlink: ro.deleteOnUnlink,
	}

	res := z.ctx.Post("team/devices/revoke_device_session", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return err
	}
	return nil
}
