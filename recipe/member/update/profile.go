package update

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

type ProfileRow struct {
	Email     string `json:"email"`
	GivenName string `json:"given_name"`
	Surname   string `json:"surname"`
}

type Profile struct {
	File                fd_file.RowFeed
	Peer                dbx_conn.ConnBusinessMgmt
	OperationLog        rp_model.TransactionReport
	ErrorMemberNotFound app_msg.Message
	ProgressUpdating    app_msg.Message
}

func (z *Profile) Preset() {
	z.OperationLog.SetModel(
		&ProfileRow{},
		&mo_member.Member{},
		rp_model.HiddenColumns(
			"result.team_member_id",
			"result.member_folder_id",
			"result.account_id",
			"result.persistent_id",
			"result.familiar_name",
			"result.abbreviated_name",
			"result.external_id",
		),
	)
	z.File.SetModel(&ProfileRow{})
}

func (z *Profile) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Profile{}, func(r rc_recipe.Recipe) {
		f, err := qt_file.MakeTestFile("member-update-profile", "john@example.com,john,smith\nalex@example.com,alex,king\n")
		if err != nil {
			return
		}
		m := r.(*Profile)
		m.File.SetFilePath(f)
	})
	if e, _ := qt_errors.ErrorsForTest(c.Log(), err); e != nil {
		return e
	}
	return qt_errors.ErrorHumanInteractionRequired
}

func (z *Profile) Exec(c app_control.Control) error {
	ui := c.UI()

	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}
	emailToMember := mo_member.MapByEmail(members)

	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	return z.File.EachRow(func(row interface{}, rowIndex int) error {
		m := row.(*ProfileRow)
		member, ok := emailToMember[m.Email]
		if !ok {
			z.OperationLog.Skip(z.ErrorMemberNotFound.With("Email", m.Email), m)
			return nil
		}

		if m.GivenName != "" {
			member.GivenName = m.GivenName
		}
		if m.Surname != "" {
			member.Surname = m.Surname
		}

		ui.Info(z.ProgressUpdating.With("Email", m.Email))
		r, err := sv_member.New(z.Peer.Context()).Update(member)
		switch {
		case err != nil:
			z.OperationLog.Failure(err, m)
			return err

		default:
			z.OperationLog.Success(m, r)
			return nil
		}
	})
}
