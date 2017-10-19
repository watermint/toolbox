package main

import (
	"github.com/watermint/toolbox/infra/knowledge"
	"os"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/service/file"
	"flag"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra/util"
)

func usage() {
	tmpl := `{{.AppName}} {{.AppVersion}} ({{.AppHash}}):

Move files/folders to destination
{{.Command}} move [OPTION]... SRC DEST
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
	fg := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	mc := file.MoveContext{}
	mc.Infra = infra.PrepareInfraFlags(fg)
	mc.SrcPath = "/マーケティング/"
	mc.DestPath = "新マーケティング"

	fg.SetOutput(os.Stderr)
	fg.Parse(os.Args[1:])

	defer mc.Infra.Shutdown()
	err := mc.Infra.Startup()
	if err != nil {
		seelog.Errorf("Unable to start operation: %s", err)
		return
	}
	seelog.Tracef("Options: %s", util.MarshalObjectToString(mc))

	token, err := mc.Infra.LoadOrAuthDropboxFull()
	if err != nil || token == "" {
		seelog.Errorf("Unable to acquire token (error: %s)", err)
		return
	}
	mc.Move(token)
}