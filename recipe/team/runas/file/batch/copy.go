package batch

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_file_relocation"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Copy struct {
	rc_recipe.RemarkIrreversible
	Peer         dbx_conn.ConnScopedTeam
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
}

func (z *Copy) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesContentWrite,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeTeamDataMember,
	)
	z.File.SetModel(&CopyMapping{})
	z.OperationLog.SetModel(&CopyMapping{}, nil)
}

func (z *Copy) copy(cm *CopyMapping, svm sv_member.Member, c app_control.Control) error {
	member, err := svm.ResolveByEmail(cm.MemberEmail)
	if err != nil {
		z.OperationLog.Failure(err, cm)
		return err
	}

	ctm := z.Peer.Client().AsMemberId(member.TeamMemberId)

	srcPath, err := es_filepath.FormatPathWithPredefinedVariables(cm.SrcPath)
	if err != nil {
		srcPath = cm.SrcPath
	}
	dstPath, err := es_filepath.FormatPathWithPredefinedVariables(cm.DstPath)
	if err != nil {
		dstPath = cm.DstPath
	}

	uc := uc_file_relocation.New(ctm)
	if err = uc.Copy(mo_path.NewDropboxPath(srcPath), mo_path.NewDropboxPath(dstPath)); err != nil {
		z.OperationLog.Failure(err, cm)
		return err
	} else {
		z.OperationLog.Success(cm, nil)
		return nil
	}
}

func (z *Copy) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	svm := sv_member.NewCached(z.Peer.Client())
	var lastErr, listErr error

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("copy", z.copy, svm, c)
		q := s.Get("copy")

		listErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			q.Enqueue(m)
			return nil
		})
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	return lang.NewMultiErrorOrNull(lastErr, listErr)
}

func (z *Copy) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
	//f, err := qt_file.MakeTestFile("share", "john@example.com,/project,/backup/project\nemma@example.com,/report,/backup/report")
	//if err != nil {
	//	return err
	//}
	//defer func() {
	//	_ = os.Remove(f)
	//}()
	//
	//return rc_exec.ExecMock(c, &Copy{}, func(r rc_recipe.Recipe) {
	//	m := r.(*Copy)
	//	m.File.SetFilePath(f)
	//})
}
