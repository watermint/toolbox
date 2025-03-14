package dbx_auth

import "github.com/watermint/toolbox/essentials/api/api_auth"

// Individual Scopes
const (
	// ScopeAccountInfoRead Read Dropbox account information
	ScopeAccountInfoRead = "account_info.read"

	// ScopeAccountInfoWrite Create, modify, and delete Dropbox account information
	ScopeAccountInfoWrite = "account_info.write"

	// ScopeFilesContentRead Read file data
	ScopeFilesContentRead = "files.content.read"

	// ScopeFilesContentWrite Create, modify, and delete file data
	ScopeFilesContentWrite = "files.content.write"

	// ScopeFilesMetadataRead Read file metadata
	ScopeFilesMetadataRead = "files.metadata.read"

	// ScopeFilesMetadataWrite Create, modify, and delete file metadata
	ScopeFilesMetadataWrite = "files.metadata.write"

	// ScopeFilesPermanentDelete Permanently delete files
	ScopeFilesPermanentDelete = "files.permanent_delete"

	// ScopeSharingRead Read sharing policies and memberships
	ScopeSharingRead = "sharing.read"

	// ScopeSharingWrite Create, modify, or delete sharing policies and memberships
	ScopeSharingWrite = "sharing.write"

	// ScopeFileRequestsRead Read a user's file requests
	ScopeFileRequestsRead = "file_requests.read"

	// ScopeFileRequestsWrite Create, modify, or delete a user's file requests
	ScopeFileRequestsWrite = "file_requests.write"

	// ScopeContactsWrite Delete a user's manually added contacts
	ScopeContactsWrite = "contacts.write"
)

var (
	ScopeIndividual = []string{
		ScopeAccountInfoRead,
		ScopeAccountInfoWrite,
		ScopeContactsWrite,
		ScopeFileRequestsRead,
		ScopeFileRequestsWrite,
		ScopeFilesContentRead,
		ScopeFilesContentWrite,
		ScopeFilesMetadataRead,
		ScopeFilesMetadataWrite,
		ScopeFilesPermanentDelete,
		ScopeSharingRead,
		ScopeSharingWrite,
	}
)

// Team Scopes
const (
	// ScopeTeamInfoRead Read basic team settings
	ScopeTeamInfoRead = "team_info.read"

	// ScopeMembersRead Read team memberships and member settings
	ScopeMembersRead = "members.read"

	// ScopeMembersWrite Add, modify, and suspend team members
	ScopeMembersWrite = "members.write"

	// ScopeMembersDelete Delete team members
	ScopeMembersDelete = "members.delete"

	// ScopeGroupsRead Read groups and their memberships
	ScopeGroupsRead = "groups.read"

	// ScopeGroupsWrite Create, modify, and delete groups
	ScopeGroupsWrite = "groups.write"

	// ScopeSessionsList View linked web, device, and app sessions
	ScopeSessionsList = "sessions.list"

	// ScopeSessionsModify Unlink web, device, and app sessions
	ScopeSessionsModify = "sessions.modify"

	// ScopeTeamDataMember Access data of other team members
	ScopeTeamDataMember = "team_data.member"

	// ScopeTeamDataTeamSpace View and edit content of your team's files and folders
	ScopeTeamDataTeamSpace = "team_data.team_space"

	// ScopeTeamDataGovernanceWrite View and edit governance data of your team's files and folders
	ScopeTeamDataGovernanceWrite = "team_data.governance.write"

	// ScopeTeamDataGovernanceRead View governance data of your team's files and folders
	ScopeTeamDataGovernanceRead = "team_data.governance.read"

	// ScopeTeamDataContentRead View content of your team's files and folders
	ScopeTeamDataContentRead = "team_data.content.read"

	// ScopeTeamDataContentWrite View and edit content of your team's files and folders
	ScopeTeamDataContentWrite = "team_data.content.write"

	// ScopeEventsRead Read the team event log
	ScopeEventsRead = "events.read"
)

func IsTeamScope(scope string) bool {
	for _, s := range ScopeTeam {
		if s == scope {
			return true
		}
	}
	return false
}

func IsTeamScopeSession(session api_auth.OAuthSessionData) bool {
	for _, scope := range session.Scopes {
		if IsTeamScope(scope) {
			return true
		}
	}
	return false
}

func HasAccountInfoRead(scopes []string) bool {
	// If no scope is specified, that means all scopes are granted.
	if len(scopes) == 0 {
		return true
	}
	for _, scope := range scopes {
		if scope == ScopeAccountInfoRead {
			return true
		}
	}
	return false
}

func HasTeamInfoRead(scopes []string) bool {
	// If no scope is specified, that means all scopes are granted.
	if len(scopes) == 0 {
		return true
	}
	for _, scope := range scopes {
		if scope == ScopeTeamInfoRead {
			return true
		}
	}
	return false
}

var (
	ScopeTeam = []string{
		ScopeEventsRead,
		ScopeGroupsRead,
		ScopeGroupsWrite,
		ScopeMembersDelete,
		ScopeMembersRead,
		ScopeMembersWrite,
		ScopeSessionsList,
		ScopeSessionsModify,
		ScopeTeamDataContentRead,
		ScopeTeamDataContentWrite,
		ScopeTeamDataGovernanceRead,
		ScopeTeamDataGovernanceWrite,
		ScopeTeamDataMember,
		ScopeTeamDataTeamSpace,
		ScopeTeamInfoRead,
	}
)
