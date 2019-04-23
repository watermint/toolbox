package api_recipe_vo

import (
	"flag"
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_recipe/api_recipe_msg"
	"github.com/watermint/toolbox/app/app_recipe/api_recipe_report"
	"github.com/watermint/toolbox/app/app_recipe/api_recpie_ctl"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"go.uber.org/zap"
	"reflect"
	"runtime"
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

type ValueContainer struct {
	Values map[string]interface{}
}

func NewValueContainer(vo interface{}) *ValueContainer {
	vc := &ValueContainer{
		Values: make(map[string]interface{}),
	}
	vc.From(vo)
	return vc
}

func (z *ValueContainer) From(vo interface{}) {
	l := app.Root().Log()
	vot := reflect.TypeOf(vo)
	vov := reflect.ValueOf(vo)
	if vot.Kind() == reflect.Ptr {
		vot = reflect.ValueOf(vo).Elem().Type()
		vov = reflect.ValueOf(vo).Elem()
	}

	if vot.Kind() != reflect.Struct {
		l.Error("ValueObject is not a struct", zap.String("name", vot.Name()), zap.String("pkg", vot.PkgPath()))
		return
	}

	nf := vot.NumField()
	for i := 0; i < nf; i++ {
		vof := vot.Field(i)
		vvf := vov.Field(i)
		kn := vof.Name
		ll := l.With(zap.String("key", kn))

		switch vof.Type.Kind() {
		case reflect.Bool:
			v := vvf.Bool()
			z.Values[kn] = &v
		case reflect.Int:
			v := vvf.Int()
			z.Values[kn] = &v
		case reflect.String:
			v := vvf.String()
			z.Values[kn] = &v
		default:
			ll.Debug("Unsupported type", zap.Any("kind", vof.Type.Kind()))
		}
	}
}

func (z *ValueContainer) Apply(vo interface{}) {
	l := app.Root().Log()
	defer func() {
		if r := recover(); r != nil {
			switch r0 := r.(type) {
			case *runtime.TypeAssertionError:
				l.Debug("Unable to convert type", zap.Error(r0))
			default:
				l.Debug("Unexpected error", zap.Any("r", r))
			}
		}
	}()

	vot := reflect.TypeOf(vo)
	vov := reflect.ValueOf(vo)

	// follow pointer
	if vot.Kind() == reflect.Ptr {
		vot = reflect.ValueOf(vo).Elem().Type()
		vov = reflect.ValueOf(vo).Elem()
	}

	if vot.Kind() != reflect.Struct {
		l.Error("ValueObject is not a struct", zap.String("name", vot.Name()), zap.String("pkg", vot.PkgPath()))
		return
	}

	nf := vot.NumField()
	for i := 0; i < nf; i++ {
		vof := vot.Field(i)
		vvf := vov.Field(i)
		kn := vof.Name
		ll := l.With(zap.String("key", kn))

		switch vof.Type.Kind() {
		case reflect.Bool:
			if v, e := z.Values[kn]; e {
				vvf.SetBool(*v.(*bool))
			} else {
				ll.Debug("Unable to find value")
			}
		case reflect.Int:
			if v, e := z.Values[kn]; e {
				vvf.SetInt(*v.(*int64))
			} else {
				ll.Debug("Unable to find value")
			}
		case reflect.String:
			if v, e := z.Values[kn]; e {
				vvf.SetString(*v.(*string))
			} else {
				ll.Debug("Unable to find value")
			}
		default:
			ll.Debug("Not supported type", zap.Any("kind", vof.Type.Kind()))
		}
	}
}

func (z *ValueContainer) MakeFlagSet(f *flag.FlagSet) {
	for n, d := range z.Values {
		kf := strcase.ToKebab(n)
		desc := n

		switch dv := d.(type) {
		case *bool:
			f.BoolVar(dv, kf, *dv, desc)
		case *int64:
			f.Int64Var(dv, kf, *dv, desc)
		case *string:
			f.StringVar(dv, kf, *dv, desc)
		}
	}
}
