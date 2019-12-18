package update

import (
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recpie/rc_conn"
	"github.com/watermint/toolbox/infra/recpie/rc_kitchen"
	"github.com/watermint/toolbox/infra/recpie/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type ProfileRow struct {
	Email     string `json:"email"`
	GivenName string `json:"given_name"`
	Surname   string `json:"surname"`
}

type ProfileVO struct {
	File fd_file.Feed
	Peer rc_conn.ConnBusinessMgmt
}

const (
	reportProfile = "update_profile"
)

type Profile struct {
}

func (z *Profile) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportProfile, rp_model.TransactionHeader(&ProfileRow{}, &mo_member.Member{})),
	}
}

func (z *Profile) Test(c app_control.Control) error {
	return qt_recipe.HumanInteractionRequired()
}

func (z *Profile) Console() {
}

func (z *Profile) Requirement() rc_vo.ValueObject {
	return &ProfileVO{}
}

func (z *Profile) Exec(k rc_kitchen.Kitchen) error {
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

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportProfile)
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
			rep.Skip(msg, m)
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
			rep.Failure(err, m)
			return err

		default:
			rep.Success(m, r)
			return nil
		}
	})
}
