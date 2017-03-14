package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/knowledge"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/service/compare"
	"os"
)

func usage() {
	tmpl := `{{.AppName}} {{.AppVersion}} ({{.AppHash}}):

Compare local path and Dropbox path
{{.Command}} LOCALPATH [DROPBOXPATH]
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

type CmpOptions struct {
	Infra       *infra.InfraOpts
	LocalPath   string
	DropboxPath string
}

func parseArgs() (*CmpOptions, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	opts := CmpOptions{}
	opts.Infra = infra.PrepareInfraFlags(f)

	f.SetOutput(os.Stderr)
	f.Parse(os.Args[1:])
	args := f.Args()
	if len(args) < 1 {
		usage()
		f.PrintDefaults()
		return nil, errors.New("Missing LOCALPATH and/or DROPBOXPATH")
	}

	opts.LocalPath = args[0]
	if len(args) == 1 {
		opts.DropboxPath = ""
	} else {
		opts.DropboxPath = args[1]
	}

	return &opts, nil
}

func main() {
	opts, err := parseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		return
	}
	defer opts.Infra.Shutdown()
	err = opts.Infra.Startup()
	if err != nil {
		seelog.Errorf("Unable to start operation: %s", err)
		return
	}
	seelog.Tracef("Options: %s", util.MarshalObjectToString(opts))

	token, err := opts.Infra.LoadOrAuthDropboxFull()
	if err != nil || token == "" {
		seelog.Errorf("Unable to acquire token (error: %s)", err)
		return
	}

	compare.Compare(opts.Infra, token, opts.LocalPath, opts.DropboxPath)
}
