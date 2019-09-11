package mo_device

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
)

const (
	DeviceTypeWeb     = "web_session"
	DeviceTypeDesktop = "desktop_client"
	DeviceTypeMobile  = "mobile_client"
)

type Session interface {
	EntryRaw() json.RawMessage

	// Pseudo field for identify type of session.
	// This field is not appear in API definition.
	// "web_session", "desktop_client", or "mobile_client".
	EntryTag() string

	// Pseudo field for identify member.
	// This field should appear parent of session data.
	EntryTeamMemberId() string
	SessionId() string
	CreatedAt() string
	UpdatedAt() string
	SessionIPAddress() string
	SessionCountry() string

	Web() (web *Web, e bool)
	Desktop() (desktop *Desktop, e bool)
	Mobile() (mobile *Mobile, e bool)
}

type Metadata struct {
	Raw json.RawMessage
	// Pseudo field for identify type of session.
	// This field is not appear in API definition.
	// "web_session", "desktop_client", or "mobile_client".
	Tag string `path:"-" json:"tag"`

	// Pseudo field for identify member.
	// This field should appear parent of session data.
	TeamMemberId string `path:"-" json:"team_member_id"`
	Id           string `path:"session_id" json:"id"`
	IpAddress    string `path:"ip_address" json:"ip_address"`
	Country      string `path:"country" json:"country"`
	Created      string `path:"created" json:"created"`
	Updated      string `path:"updated" json:"updated"`
}

func (z *Metadata) EntryTeamMemberId() string {
	return z.TeamMemberId
}

func (z *Metadata) EntryRaw() json.RawMessage {
	return z.Raw
}

func (z *Metadata) EntryTag() string {
	return z.Tag
}

func (z *Metadata) SessionId() string {
	return z.Id
}

func (z *Metadata) CreatedAt() string {
	return z.Created
}

func (z *Metadata) UpdatedAt() string {
	return z.Updated
}

func (z *Metadata) SessionIPAddress() string {
	return z.IpAddress
}

func (z *Metadata) SessionCountry() string {
	return z.Country
}

func (z *Metadata) Web() (web *Web, e bool) {
	if z.Tag != DeviceTypeWeb {
		return nil, false
	}
	web = &Web{}
	if err := api_parser.ParseModelRaw(web, z.Raw); err != nil {
		app_root.Log().Debug("parse error", zap.Error(err))
		return nil, false
	}
	web.Tag = DeviceTypeWeb
	web.TeamMemberId = z.TeamMemberId
	return web, true
}

func (z *Metadata) Desktop() (desktop *Desktop, e bool) {
	if z.Tag != DeviceTypeDesktop {
		return nil, false
	}
	desktop = &Desktop{}
	if err := api_parser.ParseModelRaw(desktop, z.Raw); err != nil {
		app_root.Log().Debug("parse error", zap.Error(err))
		return nil, false
	}
	desktop.Tag = DeviceTypeDesktop
	desktop.TeamMemberId = z.TeamMemberId
	return desktop, true
}

func (z *Metadata) Mobile() (mobile *Mobile, e bool) {
	if z.Tag != "desktop" {
		return nil, false
	}
	mobile = &Mobile{}
	if err := api_parser.ParseModelRaw(mobile, z.Raw); err != nil {
		app_root.Log().Debug("parse error", zap.Error(err))
		return nil, false
	}
	mobile.Tag = DeviceTypeMobile
	mobile.TeamMemberId = z.TeamMemberId
	return mobile, true
}

type Web struct {
	Raw json.RawMessage

	// Pseudo field for identify type of session.
	// This field is not appear in API definition.
	// "web_session", "desktop_client", or "mobile_client".
	Tag string `path:"-" json:"tag"`

	// Pseudo field for identify member.
	// This field should appear parent of session data.
	TeamMemberId string `path:"-" json:"team_member_id"`

	Id        string `path:"session_id" json:"id"`
	UserAgent string `path:"user_agent" json:"user_agent"`
	Os        string `path:"os" json:"os"`
	Browser   string `path:"browser" json:"browser"`
	IpAddress string `path:"ip_address" json:"ip_address"`
	Country   string `path:"country" json:"country"`
	Created   string `path:"created" json:"created"`
	Updated   string `path:"updated" json:"updated"`
	Expires   string `path:"expires" json:"expires"`
}

func (z *Web) EntryTeamMemberId() string {
	return z.TeamMemberId
}

func (z *Web) EntryTag() string {
	return DeviceTypeWeb
}

func (z *Web) Web() (web *Web, e bool) {
	return z, true
}

func (z *Web) Desktop() (desktop *Desktop, e bool) {
	return nil, false
}

func (z *Web) Mobile() (mobile *Mobile, e bool) {
	return nil, false
}

func (z *Web) EntryRaw() json.RawMessage {
	return z.Raw
}

func (z *Web) SessionId() string {
	return z.Id
}

func (z *Web) CreatedAt() string {
	return z.Created
}

func (z *Web) UpdatedAt() string {
	return z.Updated
}

func (z *Web) SessionIPAddress() string {
	return z.IpAddress
}

func (z *Web) SessionCountry() string {
	return z.Country
}

type Desktop struct {
	Raw json.RawMessage
	// Pseudo field for identify type of session.
	// This field is not appear in API definition.
	// "web_session", "desktop_client", or "mobile_client".
	Tag string `path:"-" json:"tag"`

	// Pseudo field for identify member.
	// This field should appear parent of session data.
	TeamMemberId string `path:"-" json:"team_member_id"`

	Id                        string `path:"session_id" json:"id"`
	HostName                  string `path:"host_name" json:"host_name"`
	ClientType                string `path:"client_type.\\.tag" json:"client_type"`
	ClientVersion             string `path:"client_version" json:"client_version"`
	Platform                  string `path:"platform" json:"platform"`
	IsDeleteOnUnlinkSupported bool   `path:"is_delete_on_unlink_supported" json:"is_delete_on_unlink_supported"`
	IpAddress                 string `path:"ip_address" json:"ip_address"`
	Country                   string `path:"country" json:"country"`
	Created                   string `path:"created" json:"created"`
	Updated                   string `path:"updated" json:"updated"`
}

func (z *Desktop) EntryTeamMemberId() string {
	return z.TeamMemberId
}

func (z *Desktop) EntryTag() string {
	return DeviceTypeDesktop
}

func (z *Desktop) Web() (web *Web, e bool) {
	return nil, false
}

func (z *Desktop) Desktop() (desktop *Desktop, e bool) {
	return z, true
}

func (z *Desktop) Mobile() (mobile *Mobile, e bool) {
	return nil, false
}

func (z *Desktop) EntryRaw() json.RawMessage {
	return z.Raw
}

func (z *Desktop) SessionId() string {
	return z.Id
}

func (z *Desktop) CreatedAt() string {
	return z.Created
}

func (z *Desktop) UpdatedAt() string {
	return z.Updated
}

func (z *Desktop) SessionIPAddress() string {
	return z.IpAddress
}

func (z *Desktop) SessionCountry() string {
	return z.Country
}

type Mobile struct {
	Raw json.RawMessage
	// Pseudo field for identify type of session.
	// This field is not appear in API definition.
	// "web_session", "desktop_client", or "mobile_client".
	Tag string `path:"-" json:"tag"`

	// Pseudo field for identify member.
	// This field should appear parent of session data.
	TeamMemberId string `path:"-" json:"team_member_id"`

	Id            string `path:"session_id" json:"id"`
	DeviceName    string `path:"device_name" json:"device_name"`
	ClientType    string `path:"client_type.\\.tag" json:"client_type"`
	IpAddress     string `path:"ip_address" json:"ip_address"`
	Country       string `path:"country" json:"country"`
	Created       string `path:"created" json:"created"`
	Updated       string `path:"updated" json:"updated"`
	ClientVersion string `path:"client_version" json:"client_version"`
	OsVersion     string `path:"os_version" json:"os_version"`
	LastCarrier   string `path:"last_carrier" json:"last_carrier"`
}

func (z *Mobile) EntryTeamMemberId() string {
	return z.TeamMemberId
}

func (z *Mobile) EntryTag() string {
	return DeviceTypeMobile
}

func (z *Mobile) Web() (web *Web, e bool) {
	return nil, false
}

func (z *Mobile) Desktop() (desktop *Desktop, e bool) {
	return nil, false
}

func (z *Mobile) Mobile() (mobile *Mobile, e bool) {
	return z, true
}

func (z *Mobile) EntryRaw() json.RawMessage {
	return z.Raw
}

func (z *Mobile) SessionId() string {
	return z.Id
}

func (z *Mobile) CreatedAt() string {
	return z.Created
}

func (z *Mobile) UpdatedAt() string {
	return z.Updated
}

func (z *Mobile) SessionIPAddress() string {
	return z.IpAddress
}

func (z *Mobile) SessionCountry() string {
	return z.Country
}

type MemberSession struct {
	Raw                       json.RawMessage
	TeamMemberId              string `path:"profile.team_member_id" json:"team_member_id"`
	Email                     string `path:"profile.email" json:"email"`
	Status                    string `path:"profile.status.\\.tag" json:"status"`
	GivenName                 string `path:"profile.name.given_name" json:"given_name"`
	Surname                   string `path:"profile.name.surname" json:"surname"`
	FamiliarName              string `path:"profile.name.familiar_name" json:"familiar_name"`
	DisplayName               string `path:"profile.name.display_name" json:"display_name"`
	AbbreviatedName           string `path:"profile.name.abbreviated_name" json:"abbreviated_name"`
	ExternalId                string `path:"profile.external_id" json:"external_id"`
	AccountId                 string `path:"profile.account_id" json:"account_id"`
	DeviceTag                 string `path:"device_tag" json:"device_tag"`
	Id                        string `path:"session.session_id" json:"id"`
	UserAgent                 string `path:"session.user_agent" json:"user_agent"`
	Os                        string `path:"session.os" json:"os"`
	Browser                   string `path:"session.browser" json:"browser"`
	IpAddress                 string `path:"session.ip_address" json:"ip_address"`
	Country                   string `path:"session.country" json:"country"`
	Created                   string `path:"session.created" json:"created"`
	Updated                   string `path:"session.updated" json:"updated"`
	Expires                   string `path:"session.expires" json:"expires"`
	HostName                  string `path:"session.host_name" json:"host_name"`
	ClientType                string `path:"session.client_type.\\.tag" json:"client_type"`
	ClientVersion             string `path:"session.client_version" json:"client_version"`
	Platform                  string `path:"session.platform" json:"platform"`
	IsDeleteOnUnlinkSupported bool   `path:"session.is_delete_on_unlink_supported" json:"is_delete_on_unlink_supported"`
	DeviceName                string `path:"session.device_name" json:"device_name"`
	OsVersion                 string `path:"session.os_version" json:"os_version"`
	LastCarrier               string `path:"session.last_carrier" json:"last_carrier"`
}

func (z *MemberSession) Session() Session {
	session := &Metadata{}
	if err := api_parser.ParseModelPathRaw(session, z.Raw, "session"); err != nil {
		app_root.Log().Warn("unexpected data format", zap.String("entry", string(z.Raw)), zap.Error(err))
		// return empty
		return session
	}
	session.TeamMemberId = z.TeamMemberId
	session.Tag = z.DeviceTag
	return session
}

func NewMemberSession(member *mo_member.Member, session Session) *MemberSession {
	tag, err := json.Marshal(session.EntryTag())
	if err != nil {
		// should not fail
		app_root.Log().Error("unable to create MemberSession", zap.Error(err))
	}
	prof := gjson.ParseBytes(member.Raw).Get("profile")
	raws := make(map[string]json.RawMessage)
	raws["device_tag"] = tag
	raws["session"] = session.EntryRaw()
	raws["profile"] = json.RawMessage(prof.Raw)
	raw := api_parser.CombineRaw(raws)

	ms := &MemberSession{}
	if err := api_parser.ParseModelRaw(ms, raw); err != nil {
		// should not fail
		app_root.Log().Error("unable to create MemberSession", zap.Error(err))
	}
	return ms
}
