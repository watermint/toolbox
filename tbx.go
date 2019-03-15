package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/cmd/cmd_root"
	"os"
	"sort"
	"strings"
)

type CmdDoc struct {
	ec   *app.ExecContext
	docs map[string]string
}

func (z *CmdDoc) init() {
	z.docs = make(map[string]string)
}

func (z *CmdDoc) parse(c cmd.Commandlet, line []string) {
	w := strings.Join(line, " ")
	switch x := c.(type) {
	case *cmd.CommandletGroup:
		q := len(line)
		sl := make([]string, q+1)
		copy(sl, line)
		for _, y := range x.SubCommands {
			sl[q] = y.Name()
			z.parse(y, sl)
		}

	default:
		z.docs[w] = z.ec.Msg(x.Desc()).T()
	}
}

func (z *CmdDoc) markdown() {
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

func main() {
	bx := rice.MustFindBox("resources")
	ec, err := app.NewExecContext(bx)
	if err != nil {
		return
	}
	cmds := cmd_root.NewCommands()
	if len(os.Args) > 1 && os.Args[1] == "-markdown" {
		d := CmdDoc{ec: ec}
		d.init()
		d.parse(cmds.RootCommand(), []string{})
		d.markdown()
		return
	}
	cmds.Exec(ec, os.Args)
}
