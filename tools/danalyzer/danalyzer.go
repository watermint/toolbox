package main

import (
	"github.com/watermint/toolbox/infra/knowledge"
	"os"
	"github.com/watermint/toolbox/infra"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/service/tree"
)

func usage() {
	tmpl := `{{.AppName}} {{.AppVersion}} ({{.AppHash}}):

Analyze file tree:
{{.Command}} report [OPTIONS]...

Create pseudo file/folder tree from DATA
{{.Command}} mockup [OPTIONS]...

`

	data := struct {
		AppName    string
		AppVersion string
		AppHash    string
		Command    string
	}{
		AppName:    knowledge.AppName,
		AppVersion: knowledge.AppVersion,
		AppHash:    knowledge.AppHash,
		Command:    os.Args[0],
	}
	infra.ShowUsage(tmpl, data)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	switch os.Args[1] {
	case "report":
		err := tree.ExecReport(os.Args[2:])
		if err != nil {
			seelog.Error(err)
		}

	case "mockup":
		err := tree.ExecMockup(os.Args[2:])
		if err != nil {
			seelog.Error(err)
		}

	default:
		usage()
	}
}