package app_doc

import (
	"fmt"
	app2 "github.com/watermint/toolbox/legacy/app"
	"github.com/watermint/toolbox/legacy/cmd"
	"sort"
	"strings"
)

type CmdDoc struct {
	ExecContext *app2.ExecContext
	docs        map[string]string
}

func (z *CmdDoc) Init() {
	z.docs = make(map[string]string)
}

func (z *CmdDoc) ParseLegacy(c cmd.Commandlet) {
	z.parseLegacyCmd(c, []string{})
}

func (z *CmdDoc) parseLegacyCmd(c cmd.Commandlet, line []string) {
	if c.IsHidden() {
		return
	}
	w := strings.Join(line, " ")
	switch x := c.(type) {
	case *cmd.CommandletGroup:
		q := len(line)
		sl := make([]string, q+1)
		copy(sl, line)
		for _, y := range x.SubCommands {
			sl[q] = y.Name()
			z.parseLegacyCmd(y, sl)
		}

	default:
		z.docs[w] = z.ExecContext.Msg(x.Desc()).T()
	}
}

func (z *CmdDoc) Markdown() {
	lenCmd := 0
	lenDesc := 0
	for k, v := range z.docs {
		if len(k) > lenCmd {
			lenCmd = len(k)
		}
		if len(v) > lenDesc {
			lenDesc = len(v)
		}
	}
	fmtCmd := "| %-" + fmt.Sprintf("%d", lenCmd+2) + "s | %-" + fmt.Sprintf("%d", lenDesc) + "s |"
	fmtHeader := "|" + strings.Repeat("-", lenCmd+4) + "|" + strings.Repeat("-", lenDesc+2) + "|"

	fmt.Printf(fmtCmd, "command", "description")
	fmt.Println()
	fmt.Println(fmtHeader)
	cmds := make([]string, 0)
	for k := range z.docs {
		cmds = append(cmds, k)
	}
	sort.Strings(cmds)
	for _, k := range cmds {
		if d, e := z.docs[k]; e {
			fmt.Printf(fmtCmd, "`"+k+"`", d)
			fmt.Println()
		}
	}
}
