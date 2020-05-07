package rc_spec

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_value"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewCommonValue() *CommonValues {
	com := &app_opt.CommonOpts{}
	repo := rc_value.NewRepository(com)

	return &CommonValues{
		repo: repo,
	}
}

type CommonValues struct {
	repo rc_recipe.Repository
}

func (z *CommonValues) Value(name string) rc_recipe.Value {
	return z.repo.FieldValue(name)
}

func (z *CommonValues) Debug() map[string]interface{} {
	return z.repo.Debug()
}

func (z *CommonValues) Opts() app_opt.CommonOpts {
	return *z.repo.(*rc_value.RepositoryImpl).Current().(*app_opt.CommonOpts)
}

func (z *CommonValues) Apply() {
	z.repo.Apply()
}

func (z *CommonValues) ValueNames() []string {
	return z.repo.FieldNames()
}

func (z *CommonValues) ValueDesc(name string) app_msg.Message {
	return z.repo.FieldDesc(name)
}

func (z *CommonValues) ValueDefault(name string) interface{} {
	return z.repo.FieldValueText(name)
}

func (z *CommonValues) ValueCustomDefault(name string) app_msg.MessageOptional {
	return z.repo.FieldCustomDefault(name)
}

func (z *CommonValues) SetFlags(f *flag.FlagSet, ui app_ui.UI) {
	z.repo.ApplyFlags(f, ui)
}
