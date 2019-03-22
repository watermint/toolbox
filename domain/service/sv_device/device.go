package sv_device

import (
	"errors"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_list"
	"github.com/watermint/toolbox/domain/infra/api_parser"
	"github.com/watermint/toolbox/domain/model/mo_device"
	"go.uber.org/zap"
)

type Session interface {
	List() (sessions []mo_device.Session, err error)
	Revoke(session mo_device.Session, opts ...RevokeOption) (err error)
}

type RevokeOption func(opt *revokeOptions) *revokeOptions

type revokeOptions struct {
	deleteOnUnlink bool
}

func New(ctx api_context.Context) Session {
	return &sessionImpl{
		ctx: ctx,
	}
}

type sessionImpl struct {
	ctx api_context.Context
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
		OnEntry(func(entry api_list.ListEntry) error {
			ej, err := entry.Json()
			if err != nil {
				return err
			}
			m := ej.Get("team_member_id")
			if !m.Exists() {
				z.ctx.Log().Debug("no `team_member_id` field found", zap.String("entry", ej.Raw))
				return errors.New("team_member_id not found")
			}
			teamMemberId := m.String()
			{
				ws := ej.Get("web_sessions")
				if ws.Exists() && ws.IsArray() {
					for _, w := range ws.Array() {
						mw := &mo_device.Web{}
						if err := api_parser.ParseModel(mw, w); err != nil {
							z.ctx.Log().Debug("unable to parse web_session", zap.Error(err), zap.String("entry", w.Raw))
							return err
						}
						mw.TeamMemberId = teamMemberId
						mw.Tag = mo_device.DeviceTypeWeb
						sessions = append(sessions, mw)
					}
				}
			}
			{
				ds := ej.Get("desktop_clients")
				if ds.Exists() && ds.IsArray() {
					for _, d := range ds.Array() {
						md := &mo_device.Desktop{}
						if err := api_parser.ParseModel(md, d); err != nil {
							z.ctx.Log().Debug("unable to parse desktop_session", zap.Error(err), zap.String("entry", d.Raw))
							return err
						}
						md.TeamMemberId = teamMemberId
						md.Tag = mo_device.DeviceTypeDesktop
						sessions = append(sessions, md)
					}
				}
			}
			{
				ms := ej.Get("mobile_clients")
				if ms.Exists() && ms.IsArray() {
					for _, m := range ms.Array() {
						mm := &mo_device.Mobile{}
						if err := api_parser.ParseModel(mm, m); err != nil {
							z.ctx.Log().Debug("unable to parse desktop_session", zap.Error(err), zap.String("entry", m.Raw))
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

	_, err = z.ctx.Request("team/devices/revoke_device_session").Param(p).Call()
	return err
}
