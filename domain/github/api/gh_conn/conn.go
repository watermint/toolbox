package gh_conn

import (
	"github.com/watermint/toolbox/domain/github/api/gh_client"
	"github.com/watermint/toolbox/essentials/api/api_conn"
)

type ConnGithub interface {
	api_conn.Connection

	Context() gh_client.Client
}

type ConnGithubPublic interface {
	ConnGithub
	IsPublic() bool
}

type ConnGithubRepo interface {
	ConnGithub
	IsRepo() bool
}
