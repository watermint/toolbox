package member

import (
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_file"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_report"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type InviteRow struct {
	Email     string
	GivenName string
	Surname   string
	Groups    string
}

type InviteVO struct {
	File     app_file.ColDataFile
	PeerName app_conn.ConnBusinessMgmt
}

func (z *InviteRow) Validate() (err error) {
	return nil
}

func InviteRowFromCols(cols []string) (row *InviteRow) {
	row = &InviteRow{}

	switch {
	case len(cols) < 1:
		return row
	case len(cols) < 2:
		row.Email = cols[0]
	case len(cols) < 4:
		row.Email, row.GivenName, row.Surname = cols[0], cols[1], cols[2]
	default:
		row.Email, row.GivenName, row.Surname, row.Groups = cols[0], cols[1], cols[2], cols[3]
	}
	return row
}

type Invite struct {
}

func (z *Invite) Test(c app_control.Control) error {
	return nil
}

func (z *Invite) Console() {
}

func (z *Invite) Requirement() app_vo.ValueObject {
	return &InviteVO{}
}

func (z *Invite) msgFromTag(tag string) app_msg.Message {
	return app_msg.M("recipe.member.invite.tag." + tag)
}

func (z *Invite) Exec(k app_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	mvo := vo.(*InviteVO)

	connMgmt, err := mvo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	svm := sv_member.New(connMgmt)
	rep, err := k.Report(
		"invite",
		app_report.TransactionHeader(&InviteRow{}, &mo_member.Member{}),
	)
	if err != nil {
		return err
	}
	defer rep.Close()

	return mvo.File.EachRow(k.Control(), func(cols []string, rowIndex int) error {
		m := InviteRowFromCols(cols)
		if err = m.Validate(); err != nil {
			if rowIndex > 0 {
				rep.Failure(app_report.MsgInvalidData, m, nil)
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

		r, err := svm.Add(m.Email, opts...)
		switch {
		case err != nil:
			rep.Failure(api_util.MsgFromError(err), m, nil)
			return nil

		case r.Tag == "success":
			rep.Success(m, r)
			return nil

		case r.Tag == "user_already_on_team":
			rep.Skip(z.msgFromTag(r.Tag), m, nil)
			return nil

		default:
			rep.Failure(z.msgFromTag(r.Tag), m, nil)
			return nil
		}
	})
}
