package app_user

import "github.com/watermint/toolbox/infra/control/app_workspace"

type User interface {
	// User identifier
	UserHash() string

	// Workspace of current user
	Workspace() app_workspace.MultiUser
}
