package batch

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Share struct {
	Peer             dbx_conn.ConnScopedTeam
	File             fd_file.RowFeed
	OperationLog     rp_model.TransactionReport
	AclUpdatePolicy  mo_string.SelectString
	MemberPolicy     mo_string.SelectString
	SharedLinkPolicy mo_string.SelectString
}

func (z *Share) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
		dbx_auth.ScopeTeamDataMember,
	)
	z.File.SetModel(&MemberFolder{})
	z.AclUpdatePolicy.SetOptions(
		"owner",
		"owner", "editor",
	)
	z.MemberPolicy.SetOptions(
		"anyone",
		"team", "anyone",
	)
	z.SharedLinkPolicy.SetOptions(
		"anyone",
		"anyone", "members",
	)
	z.OperationLog.SetModel(&MemberFolder{}, &mo_sharedfolder.SharedFolder{})
}

func (z *Share) share(mf *MemberFolder, svm sv_member.Member, c app_control.Control) error {
	member, err := svm.ResolveByEmail(mf.MemberEmail)
	if err != nil {
		z.OperationLog.Failure(err, mf)
		return err
	}

	cm := z.Peer.Client().AsMemberId(member.TeamMemberId).WithPath(dbx_client.Namespace(member.Profile().RootNamespaceId))
	sf, err := sv_sharedfolder.New(cm).Create(
		mo_path.NewDropboxPath(mf.Path),
		sv_sharedfolder.AclUpdatePolicy(z.AclUpdatePolicy.Value()),
		sv_sharedfolder.MemberPolicy(z.MemberPolicy.Value()),
		sv_sharedfolder.SharedLinkPolicy(z.SharedLinkPolicy.Value()),
	)
	if err != nil {
		z.OperationLog.Failure(err, mf)
		return err
	}
	z.OperationLog.Success(mf, sf)
	return nil
}

func (z *Share) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}
	svm := sv_member.NewCached(z.Peer.Client())

	var lastErr, listErr error

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("share", z.share, svm, c)
		q := s.Get("share")

		listErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			q.Enqueue(m)
			return nil
		})
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	return es_lang.NewMultiErrorOrNull(lastErr, listErr)
}

func (z *Share) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
}
