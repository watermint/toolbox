package dev

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_messages"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"go.uber.org/zap"
)

type PreflightVO struct {
	Test bool
}

type Preflight struct {
}

func (z *Preflight) Hidden() {
}

func (z *Preflight) Console() {
}

func (z *Preflight) Requirement() app_vo.ValueObject {
	return &PreflightVO{
		Test: false,
	}
}

func (z *Preflight) Test(c app_control.Control) error {
	return z.Exec(app_kitchen.NewKitchen(c, &PreflightVO{Test: true}))
}

func (z *Preflight) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{}
}

func (z *Preflight) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*PreflightVO)
	l := k.Log()
	{
		l.Info("Generating English documents")
		r := Doc{}
		rv := &DocVO{
			Test:           vo.Test,
			Badge:          true,
			MarkdownReadme: true,
			Lang:           "",
			Filename:       "README.md",
			CommandPath:    "doc/generated/",
		}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), rv))
		if err != nil {
			l.Error("Failed to generate documents", zap.Error(err))
			return err
		}
	}
	{
		l.Info("Generating Japanese documents")
		r := Doc{}
		rv := &DocVO{
			Test:           vo.Test,
			Badge:          true,
			MarkdownReadme: true,
			Lang:           "ja",
			Filename:       "README_ja.md",
			CommandPath:    "doc/generated_ja/",
		}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), rv))
		if err != nil {
			l.Error("Failed to generate documents", zap.Error(err))
			return err
		}
	}

	l.Info("Verify message resources")
	return qt_messages.VerifyMessages(k.Control())
}