package ig_sharedlink

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_sharedlink"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Update struct {
	Ctx  dbx_client.Client
	File fd_file.RowFeed
	Opts uc_team_sharedlink.UpdateOpts
}

func (z *Update) Preset() {
	z.File.SetModel(&uc_team_sharedlink.TargetLinks{})
}

func (z *Update) Exec(c app_control.Control) error {
	l := c.Log()

	sel, err := uc_team_sharedlink.NewSelector(c, z.Opts.OnMissing)
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

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("update_link", uc_team_sharedlink.Update, c, z.Ctx, sel, z.Opts)
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

		s.Define("scan_member", uc_team_sharedlink.RetrieveMemberLinks, c, z.Ctx, onSharedLink)
		qsm := s.Get("scan_member")

		dErr := sv_member.New(z.Ctx).ListEach(func(member *mo_member.Member) bool {
			qsm.Enqueue(member)
			return true
		})
		if dErr != nil {
			l.Debug("Unable to enqueue the member", esl.Error(dErr))
		}
	})

	return sel.Done()
}

func (z *Update) Test(c app_control.Control) error {
	return qt_errors.ErrorScenarioTest
}
