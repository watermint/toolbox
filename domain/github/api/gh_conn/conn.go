package gh_conn

import (
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
)

type ConnGithub interface {
	rc_conn.Connection

	Context() gh_context.Context
}

type ConnGithubPublic interface {
	ConnGithub
	IsPublic() bool
}
