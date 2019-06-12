package kite_user

import "github.com/watermint/toolbox/experimental/app_workspace"

type User interface {
	// User identifier
	UserHash() string

	// Workspace of current user
	Workspace() app_workspace.MultiUser
}
