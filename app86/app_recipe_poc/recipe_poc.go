package app_recipe_poc

import (
	"github.com/watermint/toolbox/app86/app_flow"
	"github.com/watermint/toolbox/app86/app_msg"
	"github.com/watermint/toolbox/app86/app_recipe"
	"github.com/watermint/toolbox/app86/app_report"
	"github.com/watermint/toolbox/app86/app_vo"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
)

type TeamFolderListVO struct {
	Recursive    bool
	NonRecursive bool
}

func (z *TeamFolderListVO) Validate(t app_vo.Validator) {
	if z.Recursive && z.NonRecursive {
		t.Invalid("err.inconsistent",
			app_msg.P("Recursive", z.Recursive),
			app_msg.P("NonRecursive", z.NonRecursive),
		)
	}
}

type MemberInviteVO struct {
	InviteList app_flow.RowDataFile
}

func (z *MemberInviteVO) Validate(t app_vo.Validator) {
}

type MemberInviteRow struct {
	Email     string
	GivenName string
	Surname   string
	Groups    string
}

func MemberInviteRowValidate(cols []string) error {
	m := MemberInviteRowFromCols(cols)
	return m.Validate()
}

func (z *MemberInviteRow) Validate() (err error) {
	if (z.Surname == "" && z.GivenName != "") ||
		(z.Surname != "" && z.GivenName == "") {
		return app_recipe.InvalidRow(
			"err.surname_or_givenname_is_empty",
			app_msg.P("GivenName", z.GivenName),
			app_msg.P("Surname", z.Surname),
		)
	}
	if err = app_recipe.AssertEmailFormat(z.Email); err != nil {
		return err
	}

	return nil
}

func MemberInviteRowFromCols(cols []string) (row *MemberInviteRow) {
	row = &MemberInviteRow{}

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

type TeamFolderList struct {
}

func (z *TeamFolderList) Requirement() app_vo.ValueObject {
	return &TeamFolderListVO{
		Recursive: false,
	}
}

func (z *TeamFolderList) Exec(rc app_recipe.RecipeContext) error {
	return app_recipe.WithBusinessFile(func(rc app_recipe.ApiRecipeContext) error {
		// TypeAssertionError will be handled by infra
		var vo interface{} = rc.Value()
		fvo := vo.(*TeamFolderListVO)

		folders, err := sv_teamfolder.New(rc.Context()).List()
		if err != nil {
			// ApiError will be reported by infra
			return err
		}

		for _, folder := range folders {
			rc.Report().Write(folder)
		}

		if fvo.Recursive {
			rc.UI().Info("info.do_recursively")
		}
		return nil
	})
}

type MemberInvite struct {
}

func (z *MemberInvite) Requirement() app_vo.ValueObject {
	return &MemberInviteVO{}
}

func (z *MemberInvite) Exec(rc app_recipe.RecipeContext) error {
	return app_recipe.WithBusinessManagement(func(rc app_recipe.ApiRecipeContext) error {
		var vo interface{} = rc.Value()
		mvo := vo.(*MemberInviteVO)
		svm := sv_member.New(rc.Context())

		return mvo.InviteList.EachRow(MemberInviteRowValidate, func(cols []string) error {
			m := MemberInviteRowFromCols(cols)
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
				rc.Report().Result(app_report.Skip("user already on team"), m, r)
				return nil

			case err != nil:
				rc.Report().Result(app_report.Failure(err), m, r)
				return nil

			default:
				rc.Report().Result(app_report.Success(), m, r)
				return nil
			}
		})
	})
}
