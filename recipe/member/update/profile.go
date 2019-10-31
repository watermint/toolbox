package update

import (
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_file"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_report"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type ProfileRow struct {
	Email     string
	GivenName string
	Surname   string
}

type ProfileVO struct {
	File app_file.Data
	Peer app_conn.ConnBusinessMgmt
}

type Profile struct {
}

func (z *Profile) Test(c app_control.Control) error {
	return qt_test.HumanInteractionRequired()
}

func (z *Profile) Console() {
}

func (z *Profile) Requirement() app_vo.ValueObject {
	return &ProfileVO{}
}

func (z *Profile) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*ProfileVO)
	ui := k.UI()

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	members, err := sv_member.New(ctx).List()
	if err != nil {
		return err
	}
	emailToMember := mo_member.MapByEmail(members)

	if err := vo.File.Model(k.Control(), &ProfileRow{}); err != nil {
		return err
	}

	rep, err := k.Report(
		"update_profile",
		app_report.TransactionHeader(&ProfileRow{}, &mo_member.Member{}),
	)
	if err != nil {
		return err
	}
	defer rep.Close()

	return vo.File.EachRow(func(row interface{}, rowIndex int) error {
		m := row.(*ProfileRow)
		member, ok := emailToMember[m.Email]
		if !ok {
			msg := app_msg.M("recipe.member.update.profile.err.member_not_found", app_msg.P{
				"Email": m.Email,
			})
			rep.Skip(msg, m, nil)
			return nil
		}

		if m.GivenName != "" {
			member.GivenName = m.GivenName
		}
		if m.Surname != "" {
			member.Surname = m.Surname
		}

		ui.Info("recipe.member.update.profile.progress", app_msg.P{
			"Email": m.Email,
		})
		r, err := sv_member.New(ctx).Update(member)
		switch {
		case err != nil:
			rep.Failure(api_util.MsgFromError(err), m, nil)
			return err

		default:
			rep.Success(m, r)
			return nil
		}
	})
}
