package api_recipe_vo

import (
	"github.com/watermint/toolbox/app/app_recipe/api_recipe_msg"
	"github.com/watermint/toolbox/app/app_recipe/api_recipe_report"
	"github.com/watermint/toolbox/app/app_recipe/api_recpie_ctl"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"go.uber.org/zap"
)

type ValueObjectValidator struct {
}

func (z *ValueObjectValidator) Error(key string, placeHolders ...api_recipe_msg.PlaceHolder) {
}

func (z *ValueObjectValidator) AssertFileExists(path string) {

}
func (z *ValueObjectValidator) AssertEmailFormat(email string) {

}
func AssertEmailFormat(email string) error {

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
	Validate(t *ValueObjectValidator)
}
type Recipe struct {
	Name  string
	Usage string
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
