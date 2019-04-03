package main

import (
	"flag"
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_ui"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/cmd/cmd_root"
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func traverse(t *testing.T, ec *app.ExecContext, x cmd.Commandlet) {
	x.Setup(ec)
	if x.Name() == "" {
		t.Error("name is null", zap.String("class", reflect.TypeOf(x).Name()))
	}
	f := flag.NewFlagSet("traverse", flag.ContinueOnError)
	x.FlagConfig(f)

	if usage := x.Usage(); usage != nil {
		usage(cmd.CommandUsage{Command: "traverse"})
	}

	switch cg := x.(type) {
	case *cmd.CommandletGroup:
		for _, y := range cg.SubCommands {
			traverse(t, ec, y)
		}
	}
}

func TestAllCommands(t *testing.T) {
	bx := rice.MustFindBox("resources")
	ec := app.NewExecContextForTest(app.WithBox(bx))

	cmds := cmd_root.NewCommands()
	root := cmds.RootCommand()

	traverse(t, ec, root)

	for k := range app_ui.Missing() {
		t.Error("Missing key", k)
	}
}
