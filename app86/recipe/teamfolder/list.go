package teamfolder

import (
	"github.com/watermint/toolbox/app86/app_msg"
	"github.com/watermint/toolbox/app86/app_recipe"
	"github.com/watermint/toolbox/app86/app_vo"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
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

func (z *List) Exec(rc app_recipe.Kitchen) error {
	return app_recipe.WithBusinessFile(func(rc app_recipe.ApiKitchen) error {
		// TypeAssertionError will be handled by infra
		var vo interface{} = rc.Value()
		fvo := vo.(*ListVO)

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
