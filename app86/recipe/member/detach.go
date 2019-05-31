package member

import (
	"github.com/watermint/toolbox/app86/app_flow"
	"github.com/watermint/toolbox/app86/app_recipe"
	"github.com/watermint/toolbox/app86/app_recipe_util"
	"github.com/watermint/toolbox/app86/app_report"
	"github.com/watermint/toolbox/app86/app_validate"
	"github.com/watermint/toolbox/app86/app_vo"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/service/sv_member"
)

type DetachRow struct {
	Email string
}

func (z *DetachRow) Validate() (err error) {
	if err = app_validate.AssertEmailFormat(z.Email); err != nil {
		return err
	}
	return nil
}

func DetachRowFromCols(cols []string) (row *DetachRow) {
	row = &DetachRow{}
	if len(cols) > 0 {
		row.Email = cols[0]
	}
	return
}

type DetachVO struct {
	File app_flow.RowDataFile
}

func (*DetachVO) Validate(t app_vo.Validator) {
}

type Detach struct {
}

func (z *Detach) Requirement() app_vo.ValueObject {
	return &DetachVO{}
}

func (*Detach) Exec(k app_recipe.Kitchen) error {
	return app_recipe_util.WithBusinessManagement(k, func(ak app_recipe_util.ApiKitchen) error {
		var vo interface{} = ak.Value()
		mvo := vo.(*DetachVO)
		svm := sv_member.New(ak.Context())
		rep, err := ak.Report(
			"detach",
			app_report.TransactionHeader(&DetachRow{}, nil),
		)
		if err != nil {
			return err
		}
		defer rep.Close()

		return mvo.File.EachRow(k.Control(), func(cols []string, rowIndex int) error {
			m := DetachRowFromCols(cols)
			if err = m.Validate(); err != nil {
				if rowIndex > 0 {
					rep.Failure(app_report.MsgInvalidData, m, nil)
				}
				return nil
			}
			mem, err := svm.ResolveByEmail(m.Email)
			if err != nil {
				rep.Failure(api_util.MsgFromError(err), m, nil)
				return nil
			}
			err = svm.Remove(mem, sv_member.Downgrade())
			if err != nil {
				rep.Failure(api_util.MsgFromError(err), m, nil)
			} else {
				rep.Success(m, nil)
			}
			return nil
		})
	})
}
