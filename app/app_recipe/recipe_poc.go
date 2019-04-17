package app_recipe

import (
	"github.com/watermint/toolbox/app/app_recipe/api_recipe_flow"
	"github.com/watermint/toolbox/app/app_recipe/api_recipe_msg"
	"github.com/watermint/toolbox/app/app_recipe/api_recipe_report"
	"github.com/watermint/toolbox/app/app_recipe/api_recipe_vo"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
)

type TeamFolderListVO struct {
	Recursive    bool
	NonRecursive bool
}

func (z *TeamFolderListVO) Validate(t *api_recipe_vo.ValueObjectValidator) {
	if z.Recursive && z.NonRecursive {
		t.Error("err.inconsistent",
			api_recipe_msg.P("Recursive", z.Recursive),
			api_recipe_msg.P("NonRecursive", z.NonRecursive),
		)
	}
}

type MemberInviteVO struct {
	FilePath string
}

func (z *MemberInviteVO) Validate(t *api_recipe_vo.ValueObjectValidator) {
	t.AssertFileExists(z.FilePath)
}

type MemberInviteRow struct {
	Email     string `json:"email"`
	GivenName string `json:"given_name"`
	Surname   string `json:"surname"`
	Groups    string `json:"groups"`
}

func MemberInviteRowValidate(cols []string) error {
	m := MemberInviteRowFromCols(cols)
	return m.Validate()
}

func (z *MemberInviteRow) Validate() (err error) {
	if (z.Surname == "" && z.GivenName != "") ||
		(z.Surname != "" && z.GivenName == "") {
		return api_recipe_vo.InvalidRow(
			"err.surname_or_givenname_is_empty",
			api_recipe_msg.P("GivenName", z.GivenName),
			api_recipe_msg.P("Surname", z.Surname),
		)
	}
	if err = api_recipe_vo.AssertEmailFormat(z.Email); err != nil {
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

func TeamFolderList() api_recipe_vo.Cook {
	return api_recipe_vo.WithBusinessFile(func(rc api_recipe_vo.ApiRecipeContext) error {
		// TypeAssertionError will be handled by infra
		fvo := rc.Value().(*TeamFolderListVO)

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
			rc.Log().Info("Do Recursively!")
		}
		return nil
	})
}

func MemberInvite() api_recipe_vo.Cook {
	return api_recipe_vo.WithBusinessManagement(func(rc api_recipe_vo.ApiRecipeContext) error {
		mvo := rc.Value().(*MemberInviteVO)
		svm := sv_member.New(rc.Context())

		return api_recipe_flow.EachRow(mvo.FilePath, MemberInviteRowValidate, func(cols []string) error {
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
			case api_recipe_flow.IsErrorPrefix("user_already_on_team", err):
				rc.Report().Result(api_recipe_report.Skip("user already on team"), m, r)
				return nil

			case err != nil:
				rc.Report().Result(api_recipe_report.Failure(err), m, r)
				return nil

			default:
				rc.Report().Result(api_recipe_report.Success(), m, r)
				return nil
			}
		})
	})
}

func poc() (prefix string, recipes map[string]api_recipe_vo.Recipe) {
	recipes = make(map[string]api_recipe_vo.Recipe)

	// Reporting task
	recipes["list"] = api_recipe_vo.Recipe{
		// give default value
		Value: func() api_recipe_vo.ValueObject {
			return &TeamFolderListVO{
				Recursive: false,
			}
		},
		Exec: TeamFolderList(),
	}

	// Transactional task
	recipes["invite"] = api_recipe_vo.Recipe{
		// give default value
		Value: func() api_recipe_vo.ValueObject {
			return &MemberInviteVO{}
		},

		Exec: MemberInvite(),
	}

	return "app.team", recipes
}
