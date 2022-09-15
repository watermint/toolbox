package update

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_sharedlink"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/ingredient/team/sharedlink"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"time"
)

type Expiry struct {
	rc_recipe.RemarkIrreversible
	Peer         dbx_conn.ConnScopedTeam
	At           mo_time.Time
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
	NoChange     app_msg.Message
	LinkNotFound app_msg.Message
	Updater      *sharedlink.Update
}

func (z *Expiry) Preset() {
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

func (z *Expiry) Exec(c app_control.Control) error {
	l := c.Log()

	if err := z.OperationLog.Open(); err != nil {
		return err
	}
	newExpiry := z.At.Time()
	newExpiryStr := dbx_util.ToApiTimeString(newExpiry)
	updateOpts := uc_team_sharedlink.UpdateOpts{
		Filter: func(target *uc_team_sharedlink.Target) bool {
			return target.Entry.Expires != newExpiryStr
		},
		Opts: func(target *uc_team_sharedlink.Target) (opts []sv_sharedlink.LinkOpt) {
			return []sv_sharedlink.LinkOpt{
				sv_sharedlink.Expires(newExpiry),
			}
		},
		OnMissing: func(url string) {
			z.OperationLog.Skip(z.LinkNotFound, &uc_team_sharedlink.TargetLinks{Url: url})
		},
		OnSkip: func(target *uc_team_sharedlink.Target) {
			l.Debug("Skipped", esl.String("curExpiry", target.Entry.Expires), esl.String("newExpiry", newExpiryStr))
			z.OperationLog.Skip(z.NoChange, &uc_team_sharedlink.TargetLinks{
				Url: target.Entry.Url,
			})
		},
		OnSuccess: func(target *uc_team_sharedlink.Target, updated mo_sharedlink.SharedLink) {
			z.OperationLog.Success(
				&uc_team_sharedlink.TargetLinks{
					Url: target.Entry.Url,
				},
				mo_sharedlink.NewSharedLinkMember(updated, target.Member),
			)
		},
		OnFailure: func(target *uc_team_sharedlink.Target, err error) {
			z.OperationLog.Failure(err, &uc_team_sharedlink.TargetLinks{
				Url: target.Entry.Url,
			})
		},
	}

	return rc_exec.Exec(c, z.Updater, func(r rc_recipe.Recipe) {
		m := r.(*sharedlink.Update)
		m.Opts = updateOpts
		m.File = z.File
		m.Ctx = z.Peer.Client()
	})
}

func (z *Expiry) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("links", "https://www.dropbox.com/scl/fo/fir9vjelf\nhttps://www.dropbox.com/scl/fo/fir9vjelg")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

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
