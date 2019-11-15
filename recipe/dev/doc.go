package dev

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/ui/app_lang"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/util/ut_doc"
	"go.uber.org/zap"
)

type DocVO struct {
	Test           bool
	Badge          bool
	MarkdownReadme bool
	Lang           string
	Filename       string
	CommandPath    string
}

type Doc struct {
}

func (z *Doc) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{}
}

func (z *Doc) Console() {
}

func (z *Doc) Hidden() {
}

func (z *Doc) Requirement() app_vo.ValueObject {
	return &DocVO{
		Test:        false,
		Badge:       true,
		Filename:    "README.md",
		CommandPath: "doc/generated/",
	}
}

func (z *Doc) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*DocVO)
	l := k.Log()
	ctl := k.Control()

	if vo.Lang != "" {
		wc := ctl.(app_control_launcher.WithMessageContainer)
		mc, err := app_msg_container_impl.New(app_lang.Select(vo.Lang), ctl)
		if err != nil {
			return err
		}
		ctl = wc.With(mc)
	}

	rme := ut_doc.NewReadme(ctl, vo.Filename, vo.Badge, vo.Test, vo.MarkdownReadme, vo.CommandPath)
	cmd := ut_doc.NewCommand(ctl, vo.CommandPath, vo.Test)
	if err := rme.Generate(); err != nil {
		l.Error("Failed to generate README", zap.Error(err))
		return err
	}
	if err := cmd.GenerateAll(); err != nil {
		l.Error("Failed to generate command manuals", zap.Error(err))
		return err
	}

	return nil
}

func (z *Doc) Test(c app_control.Control) error {
	return z.Exec(app_kitchen.NewKitchen(c, &DocVO{
		Test:     true,
		Badge:    false,
		Filename: "",
	}))
}
