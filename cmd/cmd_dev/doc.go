package cmd_dev

import (
	"flag"
	"github.com/watermint/toolbox/app/app_doc"
	"github.com/watermint/toolbox/cmd"
)

type CmdDevDoc struct {
	*cmd.SimpleCommandlet
}

func (z *CmdDevDoc) Name() string {
	return "doc"
}

func (z *CmdDevDoc) Desc() string {
	return "cmd.dev.doc.desc"
}

func (z *CmdDevDoc) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdDevDoc) FlagConfig(f *flag.FlagSet) {
}

func (z *CmdDevDoc) Exec(args []string) {
	var root cmd.Commandlet
	root = z
	for {
		if root.Parent() != nil {
			root = root.Parent()
		} else {
			break
		}
	}

	d := app_doc.CmdDoc{ExecContext: z.ExecContext}
	d.Init()
	d.Parse(root)
	d.Markdown()
}
