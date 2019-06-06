package group

import (
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/experimental/app_recipe"
	"github.com/watermint/toolbox/experimental/app_recipe_util"
	"github.com/watermint/toolbox/experimental/app_vo"
)

type List struct {
}

func (*List) Requirement() app_vo.ValueObject {
	return &app_vo.EmptyValueObject{}
}

func (*List) Exec(k app_recipe.Kitchen) error {
	return app_recipe_util.WithBusinessInfo(k, func(ak app_recipe_util.ApiKitchen) error {
		groups, err := sv_group.New(ak.Context()).List()
		if err != nil {
			return err
		}
		rep, err := ak.Report("group", &mo_group.Group{})
		if err != nil {
			return err
		}
		defer rep.Close()
		for _, m := range groups {
			rep.Row(m)
		}
		return nil
	})
}
