package member

import (
	"github.com/watermint/toolbox/atbx/app_recipe"
	"github.com/watermint/toolbox/atbx/app_recipe_util"
	"github.com/watermint/toolbox/atbx/app_vo"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
)

type List struct {
}

func (*List) Requirement() app_vo.ValueObject {
	return &app_vo.EmptyValueObject{}
}

func (*List) Exec(k app_recipe.Kitchen) error {
	return app_recipe_util.WithBusinessManagement(k, func(ak app_recipe_util.ApiKitchen) error {
		members, err := sv_member.New(ak.Context()).List()
		if err != nil {
			return err
		}

		rep, err := ak.Report("member", &mo_member.Member{})
		if err != nil {
			return err
		}
		defer rep.Close()
		for _, m := range members {
			rep.Row(m)
		}
		return nil
	})
}
