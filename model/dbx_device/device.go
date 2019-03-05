package dbx_device

import "github.com/watermint/toolbox/model/dbx_device/flat_device"

type ClientType struct {
	Tag string `json:".tag"`
}

type WebSession struct {
	SessionId string `json:"session_id"`
	UserAgent string `json:"user_agent"`
	Os        string `json:"os"`
	Browser   string `json:"browser"`
	IpAddress string `json:"ip_address,omitempty"`
	Country   string `json:"country,omitempty"`
	Created   string `json:"created,omitempty"`
	Updated   string `json:"updated,omitempty"`
	Expires   string `json:"expires,omitempty"`
}

type DesktopClient struct {
	SessionId                 string     `json:"session_id"`
	HostName                  string     `json:"host_name"`
	ClientType                ClientType `json:"client_type"`
	ClientVersion             string     `json:"client_version"`
	Platform                  string     `json:"platform"`
	IsDeleteOnUnlinkSupported bool       `json:"is_delete_on_unlink_supported"`
	IpAddress                 string     `json:"ip_address,omitempty"`
	Country                   string     `json:"country,omitempty"`
	Created                   string     `json:"created,omitempty"`
	Updated                   string     `json:"updated,omitempty"`
}

type MobileClient struct {
	SessionId     string     `json:"session_id"`
	DeviceName    string     `json:"device_name"`
	ClientType    ClientType `json:"client_type"`
	IpAddress     string     `json:"ip_address,omitempty"`
	Country       string     `json:"country,omitempty"`
	Created       string     `json:"created,omitempty"`
	Updated       string     `json:"updated,omitempty"`
	ClientVersion string     `json:"client_version,omitempty"`
	OsVersion     string     `json:"os_version,omitempty"`
	LastCarrier   string     `json:"last_carrier,omitempty"`
}

type Device struct {
	TeamMemberId   string           `json:"team_member_id"`
	WebSessions    []*WebSession    `json:"web_sessions,omitempty"`
	DesktopClients []*DesktopClient `json:"desktop_clients,omitempty"`
	MobileClients  []*MobileClient  `json:"mobile_clients,omitempty"`
}

type Devices struct {
	Devices []*Device `json:"devices"`
	HasMore bool      `json:"has_more"`
	Cursor  string    `json:"cursor,omitempty"`
}

// Return flattened `WebSession`
func (z Device) Web() []*flat_device.WebSession {
	s := make([]*flat_device.WebSession, 0)
	for _, w := range z.WebSessions {
		x := flat_device.WebSession{
			Tag:          "web_session",
			TeamMemberId: z.TeamMemberId,
			SessionId:    w.SessionId,
			UserAgent:    w.UserAgent,
			Os:           w.Os,
			Browser:      w.Browser,
			IpAddress:    w.IpAddress,
			Country:      w.Country,
			Created:      w.Created,
			Updated:      w.Updated,
			Expires:      w.Expires,
		}
		s = append(s, &x)
	}
	return s
}

// Return flattened `DesktopClient`
func (z Device) Desktop() []*flat_device.DesktopClient {
	s := make([]*flat_device.DesktopClient, 0)
	for _, w := range z.DesktopClients {
		x := flat_device.DesktopClient{
			Tag:                       "desktop_client",
			TeamMemberId:              z.TeamMemberId,
			SessionId:                 w.SessionId,
			HostName:                  w.HostName,
			ClientType:                w.ClientType.Tag,
			ClientVersion:             w.ClientVersion,
			Platform:                  w.Platform,
			IsDeleteOnUnlinkSupported: w.IsDeleteOnUnlinkSupported,
			IpAddress:                 w.IpAddress,
			Country:                   w.Country,
			Created:                   w.Created,
			Updated:                   w.Updated,
		}
		s = append(s, &x)
	}
	return s
}

// Return flattened `MobileClient`
func (z Device) Mobile() []*flat_device.MobileClient {
	s := make([]*flat_device.MobileClient, 0)
	for _, w := range z.MobileClients {
		x := flat_device.MobileClient{
			Tag:           "mobile_client",
			TeamMemberId:  z.TeamMemberId,
			SessionId:     w.SessionId,
			DeviceName:    w.DeviceName,
			ClientType:    w.ClientType.Tag,
			IpAddress:     w.IpAddress,
			Country:       w.Country,
			Created:       w.Created,
			Updated:       w.Updated,
			ClientVersion: w.ClientVersion,
			OsVersion:     w.OsVersion,
			LastCarrier:   w.LastCarrier,
		}
		s = append(s, &x)
	}
	return s
}
