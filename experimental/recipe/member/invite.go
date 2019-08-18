package member

import (
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/experimental/app_conn"
	"github.com/watermint/toolbox/experimental/app_file"
	"github.com/watermint/toolbox/experimental/app_kitchen"
	"github.com/watermint/toolbox/experimental/app_msg"
	"github.com/watermint/toolbox/experimental/app_report"
	"github.com/watermint/toolbox/experimental/app_validate"
	"github.com/watermint/toolbox/experimental/app_vo"
)

type InviteRow struct {
	Email     string
	GivenName string
	Surname   string
	Groups    string
}

type InviteVO struct {
	File     app_file.RowDataFile
	PeerName app_conn.ConnBusinessMgmt
}

func (z *InviteVO) Validate(t app_vo.Validator) {
}

func (z *InviteRow) Validate() (err error) {
	if (z.Surname == "" && z.GivenName != "") ||
		(z.Surname != "" && z.GivenName == "") {
		return app_validate.InvalidRow(
			"err.surname_or_givenname_is_empty",
			app_msg.P("GivenName", z.GivenName),
			app_msg.P("Surname", z.Surname),
		)
	}
	if err = app_validate.AssertEmailFormat(z.Email); err != nil {
		return err
	}

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
