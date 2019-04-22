package api_recipe_vo

import (
	"flag"
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/app/app_recipe/api_recipe_msg"
	"github.com/watermint/toolbox/app/app_recipe/api_recipe_report"
	"github.com/watermint/toolbox/app/app_recipe/api_recpie_ctl"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"go.uber.org/zap"
	"reflect"
)

type Validator struct {
}

func (z *Validator) Error(key string, placeHolders ...api_recipe_msg.PlaceHolder) {
}

func (z *Validator) AssertFileExists(path string) {

}
func (z *Validator) AssertEmailFormat(email string) {

}
func AssertEmailFormat(email string) error {
	return nil
}

type InvalidRowError error

func InvalidRow(key string, placeHolders ...api_recipe_msg.PlaceHolder) InvalidRowError {
	return nil
}

type NoDataRowError error

func NoDataRow() NoDataRowError {
	return nil
}

type ValueObject interface {
	Validate(t *Validator)
}
type Recipe struct {
	Value func() ValueObject
	Exec  Cook
}

type Cook interface {
	Exec(rc RecipeContext) error
}

type RecipeContext interface {
	Value() ValueObject
	Control() api_recpie_ctl.Controller
	UI() api_recpie_ctl.UI
	Log() *zap.Logger
	Report() api_recipe_report.Report
}

type ApiRecipeContext interface {
	RecipeContext
	Context() api_context.Context
}

func WithBusinessFile(exec func(rc ApiRecipeContext) error) Cook {
	panic("implement me")
}

func WithBusinessManagement(exec func(rc ApiRecipeContext) error) Cook {
	panic("implement me")
}

func MakeFlagSet(f *flag.FlagSet, vo interface{}) {
	vot := reflect.TypeOf(vo)
	vov := reflect.ValueOf(vo)
	if vot.Kind() == reflect.Ptr {
		vot = reflect.ValueOf(vo).Elem().Type()
		vov = reflect.ValueOf(vo).Elem()
	}

	if vot.Kind() != reflect.Struct {
		return
	}

	nf := vot.NumField()
	for i := 0; i < nf; i++ {
		vof := vot.Field(i)
		vvf := vov.Field(i)
		kname := strcase.ToKebab(vof.Name)

		switch vof.Type.Kind() {
		case reflect.Bool:
			var dv bool
			f.BoolVar(&dv, kname, vvf.Bool(), vof.Name)
		}
	}
}
