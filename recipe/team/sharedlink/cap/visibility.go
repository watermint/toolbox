package cap

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_sharedlink"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/ingredient/ig_team/ig_sharedlink"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Visibility struct {
	rc_recipe.RemarkIrreversible
	Peer          dbx_conn.ConnScopedTeam
	NewVisibility mo_string.SelectString
	File          fd_file.RowFeed
	OperationLog  rp_model.TransactionReport
	LinkNotFound  app_msg.Message
	NoChange      app_msg.Message
	Updater       *ig_sharedlink.Update
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
		"team_only",
	)
}

const (
	levelPublic = iota
	levelTeamOnly
	levelPassword
	levelTeamAndPassword
	levelSharedFolderOnly
	levelNoOne
	levelUnknown
)

func visibilityLevel(l esl.Logger, visibility string) int {
	switch visibility {
	case "public":
		return levelPublic
	case "team_only":
		return levelTeamOnly
	case "password":
		return levelPassword
	case "team_and_password":
		return levelTeamAndPassword
	case "shared_folder_only":
		return levelSharedFolderOnly
	case "no_one":
		return levelNoOne
	default:
		l.Warn("Unknown visibility value", esl.String("visibility", visibility))
		return levelUnknown
	}
}

func (z *Visibility) Exec(c app_control.Control) error {
	l := c.Log()

	if err := z.OperationLog.Open(); err != nil {
		return err
	}
	newLevel := visibilityLevel(l, z.NewVisibility.Value())
	updateOpts := uc_team_sharedlink.UpdateOpts{
		Filter: func(target *uc_team_sharedlink.Target) bool {
			targetLevel := visibilityLevel(l, target.Entry.Visibility)
			return targetLevel < newLevel
		},
		Opts: func(target *uc_team_sharedlink.Target) (opts []sv_sharedlink.LinkOpt) {
			if z.NewVisibility.Value() == "team_only" {
				return []sv_sharedlink.LinkOpt{
					sv_sharedlink.TeamOnly(),
				}
			} else {
				return []sv_sharedlink.LinkOpt{}
			}
		},
		OnMissing: func(url string) {
			z.OperationLog.Skip(z.LinkNotFound, &uc_team_sharedlink.TargetLinks{Url: url})
		},
		OnSkip: func(target *uc_team_sharedlink.Target) {
			l.Debug("Skipped", esl.String("curVisibility", target.Entry.Visibility), esl.String("newExpiry", z.NewVisibility.Value()))
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
		m := r.(*ig_sharedlink.Update)
		m.Opts = updateOpts
		m.File = z.File
		m.Ctx = z.Peer.Client()
	})
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
