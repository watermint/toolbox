package update

import (
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type ProfileRow struct {
	Email     string `json:"email"`
	GivenName string `json:"given_name"`
	Surname   string `json:"surname"`
}

type Profile struct {
	File         fd_file.RowFeed
	Peer         rc_conn.ConnBusinessMgmt
	OperationLog rp_model.TransactionReport
}

func (z *Profile) Preset() {
	z.OperationLog.SetModel(&ProfileRow{}, &mo_member.Member{})
	z.File.SetModel(&ProfileRow{})
}

func (z *Profile) Test(c app_control.Control) error {
	return qt_endtoend.HumanInteractionRequired()
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
			msg := app_msg.M("recipe.member.update.profile.err.member_not_found", app_msg.P{
				"Email": m.Email,
			})
			z.OperationLog.Skip(msg, m)
			return nil
		}

		if m.GivenName != "" {
			member.GivenName = m.GivenName
		}
		if m.Surname != "" {
			member.Surname = m.Surname
		}

		ui.InfoK("recipe.member.update.profile.progress", app_msg.P{
			"Email": m.Email,
		})
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
