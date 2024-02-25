package update

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_sharedlink"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
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

type Visibility struct {
	rc_recipe.RemarkIrreversible
	Peer           dbx_conn.ConnScopedTeam
	File           fd_file.RowFeed
	OperationLog   rp_model.TransactionReport
	NewVisibility  mo_string.SelectString
	LinkNotFound   app_msg.Message
	NoChange       app_msg.Message
	NoLinkToUpdate app_msg.Message
}

func (z *Visibility) Preset() {
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
	z.NewVisibility.SetOptions(
		"team_only",
		"public",
		"team_only",
	)
}

func (z *Visibility) updateVisibility(target *uc_team_sharedlink.Target, c app_control.Control, sel uc_team_sharedlink.Selector) error {
	l := c.Log().With(esl.String("member", target.Member.Email), esl.String("url", target.Entry.Url))
	mc := z.Peer.Client().AsMemberId(target.Member.TeamMemberId)

	defer func() {
		_ = sel.Processed(target.Entry.Url)
	}()

	newVis := z.NewVisibility.Value()
	curVis := target.Entry.Visibility
	if curVis == newVis {
		l.Debug("Skipped", esl.String("curVis", curVis), esl.String("newVis", newVis))
		z.OperationLog.Skip(z.NoChange, &uc_team_sharedlink.TargetLinks{
			Url: target.Entry.Url,
		})
		return nil
	}

	opts := make([]sv_sharedlink.LinkOpt, 0)
	switch newVis {
	case "team_only":
		opts = append(opts, sv_sharedlink.TeamOnly())
	case "public":
		opts = append(opts, sv_sharedlink.Public())
	}

	if len(opts) < 1 {
		l.Debug("Skipped", esl.String("curVis", curVis), esl.String("newVis", newVis))
		z.OperationLog.Skip(z.NoChange, &uc_team_sharedlink.TargetLinks{
			Url: target.Entry.Url,
		})
		return nil
	}

	updated, err := sv_sharedlink.New(mc).Update(target.Entry.SharedLink(), opts...)
	if err != nil {
		l.Debug("Unable to update visibility of the link", esl.Error(err))
		z.OperationLog.Failure(err, &uc_team_sharedlink.TargetLinks{
			Url: target.Entry.Url,
		})
		return err
	}

	l.Debug("Updated to new visibility", esl.String("curVis", curVis), esl.String("newVis", newVis), esl.Any("updated", updated))
	z.OperationLog.Success(
		&uc_team_sharedlink.TargetLinks{
			Url: target.Entry.Url,
		},
		mo_sharedlink.NewSharedLinkMember(updated, target.Member),
	)
	return nil
}

func (z *Visibility) Exec(c app_control.Control) error {
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
		c.UI().Info(z.NoLinkToUpdate)
		return nil
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("update_link", z.updateVisibility, c, sel)
		var onSharedLink uc_team_sharedlink.OnSharedLinkMember = func(member *mo_member.Member, entry *mo_sharedlink.SharedLinkMember) {
			l := c.Log().With(esl.Any("member", member), esl.Any("entry", entry))
			if shouldProcess, selErr := sel.IsTarget(entry.Url); selErr != nil {
				l.Warn("Abort delete because of KVS error", esl.Error(selErr))
				return
			} else if shouldProcess {
				qml := s.Get("update_link")
				qml.Enqueue(&uc_team_sharedlink.Target{
					Member: member,
					Entry:  entry,
				})
			}
		}

		s.Define("scan_member", uc_team_sharedlink.RetrieveMemberLinks, c, z.Peer.Client(), onSharedLink)
		qsm := s.Get("scan_member")

		dErr := sv_member.New(z.Peer.Client()).ListEach(func(member *mo_member.Member) bool {
			qsm.Enqueue(member)
			return true
		})
		if dErr != nil {
			l.Debug("Unable to enqueue the member", esl.Error(dErr))
		}
	})

	return sel.Done()
}

func (z *Visibility) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("links", "https://www.dropbox.com/scl/fo/fir9vjelf\nhttps://www.dropbox.com/scl/fo/fir9vjelg")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()
	return rc_exec.ExecMock(c, &Visibility{}, func(r rc_recipe.Recipe) {
		m := r.(*Visibility)
		m.File.SetFilePath(f)
	})
}
