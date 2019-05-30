package teamfolder

import (
	"github.com/watermint/toolbox/app86/app_msg"
	"github.com/watermint/toolbox/app86/app_recipe"
	"github.com/watermint/toolbox/app86/app_recipe_util"
	"github.com/watermint/toolbox/app86/app_vo"
	"github.com/watermint/toolbox/domain/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"go.uber.org/zap"
)

type ListVO struct {
	Recursive    bool
	NonRecursive bool
}

func (z *ListVO) Validate(t app_vo.Validator) {
	if z.Recursive && z.NonRecursive {
		t.Invalid("err.inconsistent",
			app_msg.P("Recursive", z.Recursive),
			app_msg.P("NonRecursive", z.NonRecursive),
		)
	}
}

type List struct {
}

func (z *List) Requirement() app_vo.ValueObject {
	return &ListVO{
		Recursive: false,
	}
}

func (z *List) Exec(k app_recipe.Kitchen) error {
	return app_recipe_util.WithBusinessFile(k, func(rc app_recipe_util.ApiKitchen) error {
		// TypeAssertionError will be handled by infra
		var vo interface{} = rc.Value()
		fvo := vo.(*ListVO)

		folders, err := sv_teamfolder.New(rc.Context()).List()
		if err != nil {
			// ApiError will be reported by infra
			return err
		}

		rep, err := rc.Report("teamfolder", &mo_teamfolder.TeamFolder{})
		if err != nil {
			return err
		}
		defer rep.Close()
		for _, folder := range folders {
			rc.Log().Debug("Folder", zap.Any("folder", folder))
			rep.Row(folder)
		}

		if fvo.Recursive {
			rc.UI().Info("info.do_recursively")
		}
		return nil
	})
}
