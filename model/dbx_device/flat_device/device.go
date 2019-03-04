package flat_device

// Flattened structure for `WebSession`
type WebSession struct {
	TeamMemberId string `json:"team_member_id"`
	SessionId    string `json:"session_id"`
	UserAgent    string `json:"user_agent"`
	Os           string `json:"os"`
	Browser      string `json:"browser"`
	IpAddress    string `json:"ip_address,omitempty"`
	Country      string `json:"country,omitempty"`
	Created      string `json:"created,omitempty"`
	Updated      string `json:"updated,omitempty"`
	Expires      string `json:"expires,omitempty"`
}

// Flattened structure for `DesktopClient`
type DesktopClient struct {
	TeamMemberId              string `json:"team_member_id"`
	SessionId                 string `json:"session_id"`
	HostName                  string `json:"host_name"`
	ClientType                string `json:"client_type"`
	ClientVersion             string `json:"client_version"`
	Platform                  string `json:"platform"`
	IsDeleteOnUnlinkSupported bool   `json:"is_delete_on_unlink_supported"`
	IpAddress                 string `json:"ip_address,omitempty"`
	Country                   string `json:"country,omitempty"`
	Created                   string `json:"created,omitempty"`
	Updated                   string `json:"updated,omitempty"`
}

// Flattened structure for `MobileClient`
type MobileClient struct {
	TeamMemberId  string `json:"team_member_id"`
	SessionId     string `json:"session_id"`
	DeviceName    string `json:"device_name"`
	ClientType    string `json:"client_type"`
	IpAddress     string `json:"ip_address,omitempty"`
	Country       string `json:"country,omitempty"`
	Created       string `json:"created,omitempty"`
	Updated       string `json:"updated,omitempty"`
	ClientVersion string `json:"client_version,omitempty"`
	OsVersion     string `json:"os_version,omitempty"`
	LastCarrier   string `json:"last_carrier,omitempty"`
}
