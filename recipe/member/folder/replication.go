package folder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_file_mirror"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Replication struct {
	rc_recipe.RemarkIrreversible
	Peer           dbx_conn.ConnScopedTeam
	SrcMemberEmail string
	SrcPath        mo_path.DropboxPath
	DstMemberEmail string
	DstPath        mo_path.DropboxPath
}

func (z *Replication) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesContentWrite,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeTeamDataMember,
	)
}

func (z *Replication) Exec(c app_control.Control) error {
	l := c.Log().With(esl.String("srcMemberEmail", z.SrcMemberEmail), esl.String("dstMemberEmail", z.DstMemberEmail))
	svm := sv_member.NewCached(z.Peer.Client())
	srcMember, err := svm.ResolveByEmail(z.SrcMemberEmail)
	if err != nil {
		l.Debug("Unable to resolve src member", esl.Error(err))
		return err
	}
	l.Debug("src member resolved", esl.Any("srcMember", srcMember))

	dstMember, err := svm.ResolveByEmail(z.DstMemberEmail)
	if err != nil {
		l.Debug("Unable to resolve dst member", esl.Error(err))
		return err
	}
	l.Debug("dst member resolved", esl.Any("dstMember", dstMember))

	ctxSrc := z.Peer.Client().AsMemberId(srcMember.TeamMemberId)
	ctxDst := z.Peer.Client().AsMemberId(dstMember.TeamMemberId)
	mirror := uc_file_mirror.New(ctxSrc, ctxDst)

	l.Debug("Try mirroring", esl.String("srcPath", z.SrcPath.Path()), esl.String("dstPath", z.DstPath.Path()))
	err = mirror.Mirror(z.SrcPath, z.DstPath)
	l.Debug("Operation finished", esl.Error(err))
	return err
}

func (z *Replication) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Replication{}, func(r rc_recipe.Recipe) {
		m := r.(*Replication)
		m.DstMemberEmail = "taro@example.com"
		m.DstPath = mo_path.NewDropboxPath("/")
		m.SrcMemberEmail = "hanako@example.net"
		m.SrcPath = mo_path.NewDropboxPath("/Member Files/Taro")
	})
}
