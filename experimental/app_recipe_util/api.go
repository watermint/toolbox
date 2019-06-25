package app_recipe_util

import (
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/experimental/app_control"
	"github.com/watermint/toolbox/experimental/app_kitchen"
	"github.com/watermint/toolbox/experimental/app_report"
	"github.com/watermint/toolbox/experimental/app_ui"
	"github.com/watermint/toolbox/experimental/app_vo"
	"go.uber.org/zap"
)

type ApiKitchen interface {
	app_kitchen.Kitchen
	Context() api_context.Context
}

type apiKitchenImpl struct {
	ctx     api_context.Context
	kitchen app_kitchen.Kitchen
}

func (z *apiKitchenImpl) Report(name string, row interface{}) (r app_report.Report, err error) {
	return z.kitchen.Report(name, row)
}

func (z *apiKitchenImpl) Value() app_vo.ValueObject {
	return z.kitchen.Value()
}

func (z *apiKitchenImpl) Control() app_control.Control {
	return z.kitchen.Control()
}

func (z *apiKitchenImpl) UI() app_ui.UI {
	return z.kitchen.UI()
}

func (z *apiKitchenImpl) Log() *zap.Logger {
	return z.kitchen.Log()
}

func (z *apiKitchenImpl) Context() api_context.Context {
	return z.ctx
}

func WithBusinessFile(kitchen app_kitchen.Kitchen, exec func(k ApiKitchen) error) error {
	return withToken(kitchen, api_auth.DropboxTokenBusinessFile, exec)
}

func WithBusinessManagement(kitchen app_kitchen.Kitchen, exec func(k ApiKitchen) error) error {
	return withToken(kitchen, api_auth.DropboxTokenBusinessManagement, exec)
}

func WithBusinessInfo(kitchen app_kitchen.Kitchen, exec func(k ApiKitchen) error) error {
	return withToken(kitchen, api_auth.DropboxTokenBusinessInfo, exec)
}

func withToken(kitchen app_kitchen.Kitchen, tokenType string, exec func(k ApiKitchen) error) error {
	c := api_auth_impl.NewKc(kitchen)
	ctx, err := c.Auth(tokenType)
	if err != nil {

		kitchen.Control().Abort(app_control.Reason(app_control.FailureAuthenticationFailedOrCancelled))
	}
	ak := &apiKitchenImpl{
		ctx:     ctx,
		kitchen: kitchen,
	}
	return exec(ak)
}
