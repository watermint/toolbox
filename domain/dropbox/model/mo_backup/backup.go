package mo_backup

type TeamDeviceBackupCoreInfo struct {
	ActorUserTeamMemberId    string `path:"actor.user.team_member_id" json:"actor_user_team_member_id"`
	ActorUserEmail           string `path:"actor.user.email" json:"actor_user_email"`
	ActorUserDisplayName     string `path:"actor.user.display_name" json:"actor_user_display_name"`
	SessionInfoSessionId     string `path:"details.desktop_device_session_info.session_info.session_id" json:"session_info_session_id" gorm:"primaryKey"`
	SessionInfoIpAddress     string `path:"details.desktop_device_session_info.ip_address" json:"session_info_ip_address"`
	SessionInfoHostName      string `path:"details.desktop_device_session_info.host_name" json:"session_info_host_name"`
	SessionInfoUpdated       string `path:"details.desktop_device_session_info.updated" json:"session_info_updated"`
	SessionInfoClientType    string `path:"details.desktop_device_session_info.client_type.\\.tag" json:"session_info_client_type"`
	SessionInfoClientVersion string `path:"details.desktop_device_session_info.client_version" json:"session_info_client_version"`
	SessionInfoPlatform      string `path:"details.desktop_device_session_info.platform.\\.tag" json:"session_info_platform"`
}

type TeamDeviceBackupStatus struct {
	ActorUserTeamMemberId    string `path:"actor.user.team_member_id" json:"actor_user_team_member_id"`
	ActorUserEmail           string `path:"actor.user.email" json:"actor_user_email"`
	ActorUserDisplayName     string `path:"actor.user.display_name" json:"actor_user_display_name"`
	SessionInfoSessionId     string `path:"details.desktop_device_session_info.session_info.session_id" json:"session_info_session_id" gorm:"primaryKey"`
	SessionInfoIpAddress     string `path:"details.desktop_device_session_info.ip_address" json:"session_info_ip_address"`
	SessionInfoHostName      string `path:"details.desktop_device_session_info.host_name" json:"session_info_host_name"`
	SessionInfoUpdated       string `path:"details.desktop_device_session_info.updated" json:"session_info_updated"`
	SessionInfoClientType    string `path:"details.desktop_device_session_info.client_type.\\.tag" json:"session_info_client_type"`
	SessionInfoClientVersion string `path:"details.desktop_device_session_info.client_version" json:"session_info_client_version"`
	SessionInfoPlatform      string `path:"details.desktop_device_session_info.platform" json:"session_info_platform"`

	Timestamp    string `json:"timestamp"`
	LatestStatus string `json:"latest_status"`
}

type TeamActivityDeviceBackupEvent struct {
	Timestamp     string `path:"timestamp" json:"timestamp" gorm:"primaryKey"`
	EventCategory string `path:"event_category.\\.tag" json:"event_category"`
	EventType     string `path:"event_type.\\.tag" json:"event_type"`

	ActorUserTeamMemberId    string `path:"actor.user.team_member_id" json:"actor_user_team_member_id"`
	ActorUserEmail           string `path:"actor.user.email" json:"actor_user_email"`
	ActorUserDisplayName     string `path:"actor.user.display_name" json:"actor_user_display_name"`
	SessionInfoSessionId     string `path:"details.desktop_device_session_info.session_info.session_id" json:"session_info_session_id" gorm:"primaryKey"`
	SessionInfoIpAddress     string `path:"details.desktop_device_session_info.ip_address" json:"session_info_ip_address"`
	SessionInfoHostName      string `path:"details.desktop_device_session_info.host_name" json:"session_info_host_name"`
	SessionInfoUpdated       string `path:"details.desktop_device_session_info.updated" json:"session_info_updated"`
	SessionInfoClientType    string `path:"details.desktop_device_session_info.client_type.\\.tag" json:"session_info_client_type"`
	SessionInfoClientVersion string `path:"details.desktop_device_session_info.client_version" json:"session_info_client_version"`
	SessionInfoPlatform      string `path:"details.desktop_device_session_info.platform" json:"session_info_platform"`

	PreviousValue string `path:"details.previous_value.\\.tag" json:"previous_value"`
	NewValue      string `path:"details.new_value.\\.tag" json:"new_value"`
}
