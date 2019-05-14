package app_recipe

import (
	"github.com/watermint/toolbox/app86/app_control"
	"github.com/watermint/toolbox/app86/app_msg"
	"github.com/watermint/toolbox/app86/app_report"
	"github.com/watermint/toolbox/app86/app_ui"
	"github.com/watermint/toolbox/app86/app_vo"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"go.uber.org/zap"
)

func AssertEmailFormat(email string) error {
	return nil
}

type InvalidRowError error

func InvalidRow(key string, placeHolders ...app_msg.PlaceHolder) InvalidRowError {
	return nil
}

type NoDataRowError error

func NoDataRow() NoDataRowError {
	return nil
}

type Recipe interface {
	Requirement() app_vo.ValueObject
	Exec(rc RecipeContext) error
}

type Cook interface {
	Exec(rc RecipeContext) error
}

type RecipeContext interface {
	Value() app_vo.ValueObject
	Control() app_control.Control
	UI() app_ui.UI
	Log() *zap.Logger
	Report() app_report.Report
}

type ApiRecipeContext interface {
	RecipeContext
	Context() api_context.Context
}

func WithBusinessFile(exec func(rc ApiRecipeContext) error) error {
	panic("implement me")
}

func WithBusinessManagement(exec func(rc ApiRecipeContext) error) error {
	panic("implement me")
}
