package update

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_sharedlink"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"math"
	"os"
	"time"
)

type Expiry struct {
	rc_recipe.RemarkIrreversible
	Peer                       dbx_conn.ConnScopedTeam
	Days                       mo_int.RangeInt
	At                         mo_time.TimeOptional
	File                       fd_file.RowFeed
	OperationLog               rp_model.TransactionReport
	LinkNotFound               app_msg.Message
	NoChange                   app_msg.Message
	ErrorPleaseSpecifyDaysOrAt app_msg.Message
	ErrorInvalidDateTime       app_msg.Message
}

func (z *Expiry) Preset() {
	z.Days.SetRange(0, math.MaxInt32, 0)
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

func (z *Expiry) updateExpiry(target *uc_team_sharedlink.Target, c app_control.Control, sel uc_team_sharedlink.Selector, newExpiry time.Time) error {
	l := c.Log().With(esl.String("member", target.Member.Email), esl.String("url", target.Entry.Url))
	mc := z.Peer.Context().AsMemberId(target.Member.TeamMemberId)

	defer func() {
		_ = sel.Processed(target.Entry.Url)
	}()

	newExpiry8601 := dbx_util.ToApiTimeString(newExpiry)

	if target.Entry.Expires == newExpiry8601 {
		l.Debug("Skipped", esl.String("curExpiry", target.Entry.Expires), esl.String("newExpiry", newExpiry8601))
		z.OperationLog.Skip(z.NoChange, &uc_team_sharedlink.TargetLinks{
			Url: target.Entry.Url,
		})
		return nil
	}

	opts := make([]sv_sharedlink.LinkOpt, 0)
	opts = append(opts, sv_sharedlink.Expires(newExpiry))

	updated, err := sv_sharedlink.New(mc).Update(target.Entry.SharedLink(), opts...)
	if err != nil {
		l.Debug("Unable to update visibility of the link", esl.Error(err))
		z.OperationLog.Failure(err, &uc_team_sharedlink.TargetLinks{
			Url: target.Entry.Url,
		})
		return err
	}

	l.Debug("Updated to new visibility",
		esl.String("curExpiry", target.Entry.Expires),
		esl.String("newExpiry", newExpiry8601),
		esl.Any("updated", updated))
	z.OperationLog.Success(
		&uc_team_sharedlink.TargetLinks{
			Url: target.Entry.Url,
		},
		mo_sharedlink.NewSharedLinkMember(updated, target.Member),
	)
	return nil
}

func (z *Expiry) Exec(c app_control.Control) error {
	ui := c.UI()
	l := c.Log()
	var newExpiry time.Time
	if z.Days.Value() > 0 && z.At.Ok() {
		l.Debug("Both Days/At specified", esl.Int("evo.Days", z.Days.Value()), esl.String("evo.At", z.At.Value()))
		ui.Error(z.ErrorPleaseSpecifyDaysOrAt)
		return errors.New("please specify one of `-days` or `-at`")
	}

	switch {
	case z.Days.Value() > 0:
		newExpiry = dbx_util.RebaseTime(time.Now().Add(time.Duration(z.Days.Value()*24) * time.Hour))
		l.Debug("New expiry", esl.Int("evo.Days", z.Days.Value()), esl.String("newExpiry", newExpiry.String()))

	default:
		if !z.At.Ok() {
			l.Debug("Invalid date/time format for at option", esl.String("evo.At", z.At.Value()))
			ui.Error(z.ErrorInvalidDateTime.With("Time", z.At.Value()))
			return errors.New("invalid date/time format for `at`")
		}
		newExpiry = z.At.Time()
	}

	l = l.With(esl.String("newExpiry", newExpiry.String()))

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

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("update_link", z.updateExpiry, c, sel, newExpiry)
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

func (z *Expiry) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("links", "https://www.dropbox.com/scl/fo/fir9vjelf\nhttps://www.dropbox.com/scl/fo/fir9vjelg")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	// should fail
	{
		err := rc_exec.ExecMock(c, &Expiry{}, func(r rc_recipe.Recipe) {
			rc := r.(*Expiry)
			rc.Days.SetValue(1)
			rc.File.SetFilePath(f)
			rc.At = mo_time.NewOptional(time.Now().Add(1 * 1000 * time.Millisecond))
		})
		if err == nil {
			return errors.New("days and at should not be accepted same time")
		}
	}

	{
		err := rc_exec.ExecMock(c, &Expiry{}, func(r rc_recipe.Recipe) {
			m := r.(*Expiry)
			m.File.SetFilePath(f)
			m.Days.SetValue(7)
		})
		if e, _ := qt_errors.ErrorsForTest(c.Log(), err); e != nil {
			return e
		}
	}

	{
		err := rc_exec.ExecMock(c, &Expiry{}, func(r rc_recipe.Recipe) {
			m := r.(*Expiry)
			m.File.SetFilePath(f)
			m.At = mo_time.NewOptional(time.Now().Add(1 * 1000 * time.Millisecond))
		})
		if e, _ := qt_errors.ErrorsForTest(c.Log(), err); e != nil {
			return e
		}
	}

	return nil
}
