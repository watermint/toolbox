package delete

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_sharedlink"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

const (
	linkStatusEnqueue = "e"
	linkStatusDeleted = "d"
	linkStatusFailure = "f"
)

type Links struct {
	rc_recipe.RemarkIrreversible
	Peer           dbx_conn.ConnScopedTeam
	File           fd_file.RowFeed
	OperationLog   rp_model.TransactionReport
	LinkNotFound   app_msg.Message
	NoLinkToDelete app_msg.Message
}

func (z *Links) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSharingWrite,
		dbx_auth.ScopeTeamDataMember,
	)
	z.File.SetModel(&uc_team_sharedlink.TargetLinks{})
	z.OperationLog.SetModel(
		&uc_team_sharedlink.TargetLinks{},
		&mo_sharedlink.SharedLinkMember{},
		rp_model.HiddenColumns(
			"result.shared_link_id",
			"result.account_id",
			"result.team_member_id",
			"result.status",
		),
	)
}

func (z *Links) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	var onMissing uc_team_sharedlink.SelectorOnMissing = func(url string) {
		z.OperationLog.Skip(z.LinkNotFound, &uc_team_sharedlink.TargetLinks{Url: url})
	}
	sel, err := uc_team_sharedlink.NewSelector(c, onMissing)
	if err != nil {
		return err
	}

	loadErr := z.File.EachRow(func(m interface{}, rowIndex int) error {
		r := m.(*uc_team_sharedlink.TargetLinks)
		return sel.Register(r.Url)
	})
	if loadErr != nil {
		return loadErr
	}
	if sel.NumTargets() < 1 {
		c.UI().Info(z.NoLinkToDelete)
		return nil
	}

	var onDeleteSuccess uc_team_sharedlink.DeleteOnSuccess = func(t *uc_team_sharedlink.Target) {
		l := c.Log()

		// Mark the link as deleted
		if selErr := sel.Processed(t.Entry.Url); selErr != nil {
			l.Warn("Unable to record status link. Report might not accurate", esl.Error(selErr))
		}

		// Report
		z.OperationLog.Success(&uc_team_sharedlink.TargetLinks{Url: t.Entry.Url}, t.Entry)
	}
	var onDeleteFailure uc_team_sharedlink.DeleteOnFailure = func(t *uc_team_sharedlink.Target, cause error) {
		l := c.Log()

		// Mark the link as proceed
		if selErr := sel.Processed(t.Entry.Url); selErr != nil {
			l.Warn("Unable to record status link. Report might not accurate", esl.Error(selErr))
		}

		// Report
		z.OperationLog.Failure(cause, &uc_team_sharedlink.TargetLinks{Url: t.Entry.Url})
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("delete_link", uc_team_sharedlink.DeleteMemberLinkWithSel, c, z.Peer.Context(), onDeleteSuccess, onDeleteFailure, sel)
		var onSharedLink uc_team_sharedlink.OnSharedLinkMember = func(member *mo_member.Member, entry *mo_sharedlink.SharedLinkMember) {
			l := l.With(esl.Any("member", member), esl.Any("entry", entry))
			if shouldProcess, selErr := sel.IsTarget(entry.Url); selErr != nil {
				l.Warn("Abort delete because of KVS error", esl.Error(selErr))
				return
			} else if shouldProcess {
				qml := s.Get("delete_link")
				qml.Enqueue(&uc_team_sharedlink.Target{
					Member: member,
					Entry:  entry,
				})
			}
		}
		s.Define("scan_member", uc_team_sharedlink.RetrieveMemberLinks, c, z.Peer.Context(), onSharedLink)
		qsm := s.Get("scan_member")

		dErr := sv_member.New(z.Peer.Context()).ListEach(func(member *mo_member.Member) bool {
			qsm.Enqueue(member)
			return true
		})
		if dErr != nil {
			l.Debug("Unable to enqueue the member", esl.Error(dErr))
		}
	})

	return sel.Done()
}

func (z *Links) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("links", "https://www.dropbox.com/scl/fo/fir9vjelf\nhttps://www.dropbox.com/scl/fo/fir9vjelg")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()
	return rc_exec.ExecMock(c, &Links{}, func(r rc_recipe.Recipe) {
		m := r.(*Links)
		m.File.SetFilePath(f)
	})
}
