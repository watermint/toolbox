package dbx_auth

import (
	"github.com/watermint/toolbox/infra/api/api_appkey"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"golang.org/x/oauth2"
)

// Individual Scopes
const (
	// Read Dropbox account information (Required)
	ScopeAccountInfoRead = "account_info.read"

	// Create, modify, and delete Dropbox account information
	ScopeAccountInfoWrite = "account_info.write"

	// Read file data
	ScopeFilesContentRead = "files.content.read"

	// Create, modify, and delete file data
	ScopeFilesContentWrite = "files.content.write"

	// Read file metadata
	ScopeFilesMetadataRead = "files.metadata.read"

	// Create, modify, and delete file metadata
	ScopeFilesMetadataWrite = "files.metadata.write"

	// Permanently delete files
	ScopeFilesPermanentDelete = "files.permanent_delete"

	// Read sharing policies and memberships
	ScopeSharingRead = "sharing.read"

	// Create, modify, or delete sharing policies and memberships
	ScopeSharingWrite = "sharing.write"

	// Read a user's file requests
	ScopeFileRequestsRead = "file_requests.read"

	// Create, modify, or delete a user's file requests
	ScopeFileRequestsWrite = "file_requests.write"

	// Delete a user's manually added contacts
	ScopeContactsWrite = "contacts.write"
)

// Team Scopes
const (
	// Read basic team settings
	ScopeTeamInfoRead = "team_info.read"

	// Read team memberships and member settings
	ScopeMembersRead = "members.read"

	// Add, modify, and suspend team members
	ScopeMembersWrite = "members.write"

	// Delete team members
	ScopeMembersDelete = "members.delete"

	// Read groups and their memberships
	ScopeGroupsRead = "groups.read"

	// Create, modify, and delete groups
	ScopeGroupsWrite = "groups.write"

	// View linked web, device, and app sessions
	ScopeSessionsList = "sessions.list"

	// Unlink web, device, and app sessions
	ScopeSessionsModify = "sessions.modify"

	// Access data of other team members
	ScopeTeamDataMember = "team_data.member"

	// Manage the team space
	ScopeTeamDataTeamSpace = "team_data.team_space"

	// Read the team event log
	ScopeEventsRead = "events.read"
)

type Scoped struct {
	appType string
	ctl     app_control.Control
	res     api_appkey.Resource
}

func (z Scoped) UsePKCE() bool {
	return true
}

func (z Scoped) Config(scopes []string) *oauth2.Config {
	key, secret := z.res.Key(z.appType)
	return &oauth2.Config{
		ClientID:     key,
		ClientSecret: secret,
		Endpoint:     DropboxOAuthEndpoint(),
		Scopes:       scopes,
	}
}

func NewScopedIndividual(ctl app_control.Control) api_auth.App {
	return &Scoped{
		appType: api_auth.DropboxScopedIndividual,
		ctl:     ctl,
		res:     api_appkey.New(ctl),
	}
}

func NewScopedTeam(ctl app_control.Control) api_auth.App {
	return &Scoped{
		appType: api_auth.DropboxScopedTeam,
		ctl:     ctl,
		res:     api_appkey.New(ctl),
	}
}
