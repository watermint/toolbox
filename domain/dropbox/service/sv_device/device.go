package sv_device

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_device"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"go.uber.org/zap"
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
	err = z.ctx.List("team/devices/list_members_devices").
		Continue("team/devices/list_members_devices").
		Param(p).
		UseHasMore(true).
		ResultTag("devices").
		OnEntry(func(entry tjson.Json) error {
			m, found := entry.FindString("team_member_id")
			if !found {
				z.ctx.Log().Debug("no `team_member_id` field found", zap.ByteString("entry", entry.Raw()))
				return errors.New("team_member_id not found")
			}
			teamMemberId := m
			{
				ws, found := entry.FindArray("web_sessions")
				if found {
					for _, w := range ws {
						mw := &mo_device.Web{}
						if _, err := w.Model(mw); err != nil {
							z.ctx.Log().Debug("unable to parse web_session", zap.Error(err), zap.ByteString("entry", entry.Raw()))
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
						if _, err := d.Model(md); err != nil {
							z.ctx.Log().Debug("unable to parse desktop_session", zap.Error(err), zap.ByteString("entry", entry.Raw()))
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
						if _, err := m.Model(mm); err != nil {
							z.ctx.Log().Debug("unable to parse desktop_session", zap.Error(err), zap.ByteString("entry", entry.Raw()))
							return err
						}
						mm.TeamMemberId = teamMemberId
						mm.Tag = mo_device.DeviceTypeMobile
						sessions = append(sessions, mm)
					}
				}
			}
			return nil
		}).
		Call()
	if err != nil {
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

	_, err = z.ctx.Post("team/devices/revoke_device_session").Param(p).Call()
	return err
}
