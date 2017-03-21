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
	"github.com/watermint/toolbox/service/report"
	"os"
)

func usage() {
	tmpl := `{{.AppName}} {{.AppVersion}} ({{.AppHash}}):

Usage:

{{.Command}}                   LOCALPATH [DROPBOXPATH]
{{.Command}}    -xlsx XLSXPATH LOCALPATH [DROPBOXPATH]
{{.Command}} -csv-dir CSVDIR   LOCALPATH [DROPBOXPATH]

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

func parseArgs() (*compare.CompareOpts, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	opts := compare.CompareOpts{}
	opts.InfraOpts = infra.PrepareInfraFlags(f)
	opts.ReportOpts = report.PrepareMultiReportFlags(f)

	f.SetOutput(os.Stderr)
	f.Parse(os.Args[1:])
	args := f.Args()
	if len(args) < 1 {
		usage()
		f.PrintDefaults()
		return nil, errors.New("Missing LOCALPATH and/or DROPBOXPATH")
	}

	opts.LocalBasePath = args[0]
	if len(args) == 1 {
		opts.DropboxBasePath = ""
	} else {
		opts.DropboxBasePath = args[1]
	}

	return &opts, nil
}

func main() {
	opts, err := parseArgs()
	if err != nil {
		usage()
		fmt.Fprintln(os.Stderr, "Error: ", err)
		return
	}
	err = opts.ReportOpts.ValidateMultiReportOpts()
	if err != nil {
		usage()
		fmt.Fprintln(os.Stderr, "Error: ", err)
		return
	}
	defer opts.InfraOpts.Shutdown()
	err = opts.InfraOpts.Startup()
	if err != nil {
		seelog.Errorf("Unable to start operation: %s", err)
		return
	}
	seelog.Tracef("Options: %s", util.MarshalObjectToString(opts))

	token, err := opts.InfraOpts.LoadOrAuthDropboxFull()
	if err != nil || token == "" {
		seelog.Errorf("Unable to acquire token (error: %s)", err)
		return
	}
	opts.DropboxToken = token

	compare.Compare(opts)
}
