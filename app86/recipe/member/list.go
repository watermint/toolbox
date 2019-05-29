package member

import (
	"github.com/watermint/toolbox/app86/app_recipe"
	"github.com/watermint/toolbox/app86/app_recipe_util"
	"github.com/watermint/toolbox/app86/app_vo"
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

		rep, err := ak.Report("member")
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
