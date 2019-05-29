package member

import (
	"github.com/watermint/toolbox/app86/app_flow"
	"github.com/watermint/toolbox/app86/app_msg"
	"github.com/watermint/toolbox/app86/app_recipe"
	"github.com/watermint/toolbox/app86/app_recipe_util"
	"github.com/watermint/toolbox/app86/app_report"
	"github.com/watermint/toolbox/app86/app_validate"
	"github.com/watermint/toolbox/app86/app_vo"
	"github.com/watermint/toolbox/domain/service/sv_member"
)

type InviteRow struct {
	Email     string
	GivenName string
	Surname   string
	Groups    string
}

type InviteVO struct {
	InviteList app_flow.RowDataFile
}

func (z *InviteVO) Validate(t app_vo.Validator) {
}

func InviteRowValidate(cols []string) error {
	m := InviteRowFromCols(cols)
	return m.Validate()
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

func (z *Invite) Exec(k app_recipe.Kitchen) error {
	return app_recipe_util.WithBusinessManagement(k, func(ak app_recipe_util.ApiKitchen) error {
		var vo interface{} = ak.Value()
		mvo := vo.(*InviteVO)
		svm := sv_member.New(ak.Context())
		rep, err := ak.Report("invite")
		if err != nil {
			return err
		}

		return mvo.InviteList.EachRow(InviteRowValidate, func(cols []string) error {
			m := InviteRowFromCols(cols)
			opts := make([]sv_member.AddOpt, 0)
			if m.GivenName != "" {
				opts = append(opts, sv_member.AddWithGivenName(m.GivenName))
			}
			if m.Surname != "" {
				opts = append(opts, sv_member.AddWithSurname(m.Surname))
			}

			r, err := svm.Add(m.Email, opts...)
			switch {
			case app_flow.IsErrorPrefix("user_already_on_team", err):
				rep.Transaction(app_report.Skip("user already on team"), m, r)
				return nil

			case err != nil:
				rep.Transaction(app_report.Failure(""), m, r)
				return nil

			default:
				rep.Transaction(app_report.Success(), m, r)
				return nil
			}
		})
	})
}
