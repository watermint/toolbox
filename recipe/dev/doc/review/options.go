package review

import (
	"fmt"

	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_options"
)

type Options struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkIrreversible
	MissingOptions rp_model.RowReport
	AllCovered     app_msg.Message
	MissingCount   app_msg.Message
	GenerateHint   app_msg.Message
}

func (z *Options) Preset() {
	z.MissingOptions.SetModel(
		&MissingOption{},
		rp_model.HiddenColumns(),
	)
}

func (z *Options) Exec(c app_control.Control) error {
	l := c.Log()
	ui := c.UI()

	missingOptions, err := qtr_options.VerifySelectStringOptions(c)
	if err != nil {
		return err
	}

	if z.MissingOptions == nil {
		l.Debug("MissingOptions is nil, likely in test mode")
		return nil
	}

	if err := z.MissingOptions.Open(); err != nil {
		return err
	}

	totalOptions := 0
	recipeFieldCount := make(map[string]int)

	for _, opt := range missingOptions {
		totalOptions++
		recipeField := fmt.Sprintf("%s.%s", opt.RecipeName, opt.FieldName)
		recipeFieldCount[recipeField]++

		z.MissingOptions.Row(&MissingOption{
			Recipe:     opt.RecipeName,
			Field:      opt.FieldName,
			Option:     opt.Option,
			MessageKey: opt.Key,
		})
	}

	if totalOptions == 0 {
		ui.Success(z.AllCovered)
		l.Info("All SelectString options have descriptions")
	} else {
		ui.Error(z.MissingCount.With("Count", totalOptions))

		l.Warn("Missing SelectString option descriptions",
			esl.Int("count", totalOptions),
			esl.Int("unique_fields", len(recipeFieldCount)))

		ui.Info(z.GenerateHint)
	}

	return nil
}

func (z *Options) Test(c app_control.Control) error {
	return z.Exec(c)
}

type MissingOption struct {
	Recipe     string `json:"recipe"`
	Field      string `json:"field"`
	Option     string `json:"option"`
	MessageKey string `json:"message_key"`
}
