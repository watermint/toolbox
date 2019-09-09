package cmd_dev

import (
	"flag"
	"github.com/watermint/toolbox/legacy/app/app_doc"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdDevDoc struct {
	*cmd2.SimpleCommandlet
}

func (z *CmdDevDoc) Name() string {
	return "legacydoc"
}

func (z *CmdDevDoc) Desc() string {
	return "cmd.dev.doc.desc"
}

func (z *CmdDevDoc) Usage() func(cmd2.CommandUsage) {
	return nil
}

func (z *CmdDevDoc) FlagConfig(f *flag.FlagSet) {
}

func (z *CmdDevDoc) Exec(args []string) {
	var root cmd2.Commandlet
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
	d.ParseLegacy(root)
	d.Markdown()
}
