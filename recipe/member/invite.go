package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type InviteRow struct {
	Email     string `json:"email"`
	GivenName string `json:"given_name"`
	Surname   string `json:"surname"`
}

func (z *InviteRow) Validate() error {
	if z.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

type InviteVO struct {
}

const (
	reportInvite = "invite"
)

type Invite struct {
	File         fd_file.RowFeed
	Peer         rc_conn.ConnBusinessMgmt
	OperationLog rp_model.TransactionReport
	SilentInvite bool
}

func (z *Invite) Init() {
	z.File.SetModel(&InviteRow{})
	z.OperationLog.Model(&InviteRow{}, &mo_member.Member{})
}

func (z *Invite) Test(c app_control.Control) error {
	return qt_recipe.HumanInteractionRequired()
}

func (z *Invite) Console() {
}

func (z *Invite) msgFromTag(tag string) app_msg.Message {
	return app_msg.M("recipe.member.invite.tag." + tag)
}

func (z *Invite) Exec(k rc_kitchen.Kitchen) error {
	ctx, err := z.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	svm := sv_member.New(ctx)
	err = z.OperationLog.Open()
	if err != nil {
		return err
	}
	defer z.OperationLog.Close()

	return z.File.EachRow(func(row interface{}, rowIndex int) error {
		m := row.(*InviteRow)
		if err = m.Validate(); err != nil {
			if rowIndex > 0 {
				z.OperationLog.Failure(err, m)
			}
			return nil
		}
		opts := make([]sv_member.AddOpt, 0)
		if m.GivenName != "" {
			opts = append(opts, sv_member.AddWithGivenName(m.GivenName))
		}
		if m.Surname != "" {
			opts = append(opts, sv_member.AddWithSurname(m.Surname))
		}
		if z.SilentInvite {
			opts = append(opts, sv_member.AddWithoutSendWelcomeEmail())
		}

		r, err := svm.Add(m.Email, opts...)
		switch {
		case err != nil:
			z.OperationLog.Failure(err, m)
			return nil

		case r.Tag == "success":
			z.OperationLog.Success(m, r)
			return nil

		case r.Tag == "user_already_on_team":
			z.OperationLog.Skip(z.msgFromTag(r.Tag), m)
			return nil

		default:
			// TODO: i18n
			z.OperationLog.Failure(errors.New("failure due to "+r.Tag), m)
			return nil
		}
	})
}
