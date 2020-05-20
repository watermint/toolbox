package gh_conn

import (
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/infra/api/api_conn"
)

type ConnGithub interface {
	api_conn.Connection

	Context() gh_context.Context
}

type ConnGithubPublic interface {
	ConnGithub
	IsPublic() bool
}

type ConnGithubRepo interface {
	ConnGithub
	IsRepo() bool
}
