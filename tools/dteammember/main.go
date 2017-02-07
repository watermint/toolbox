package main

import (
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/knowledge"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/service/detach"
	"os"
)

func usage() {
	tmpl := `{{.AppName}} {{.AppVersion}} ({{.AppHash}}):

Detach member(s) from the team

{{.Command}} detach -user user@example.com
{{.Command}} detach -csv user-list.csv
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

type DetachOptions struct {
	Infra    *infra.InfraOpts
	User     string
	UserFile string
	DryRun   bool
}

func parseDetachOptions(args []string) (*DetachOptions, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	opts := &DetachOptions{}

	opts.Infra = infra.PrepareInfraFlags(f)

	descUser := "Specify target user by email address"
	f.StringVar(&opts.User, "user", "", descUser)

	descUserFile := "Specify CSV file path of target user email address"
	f.StringVar(&opts.UserFile, "csv", "", descUserFile)

	descDryRun := "Dry run"
	f.BoolVar(&opts.DryRun, "dry-run", true, descDryRun)

	f.SetOutput(os.Stderr)
	f.Parse(args)

	return opts, nil
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	if os.Args[1] != "detach" {
		usage()
		return
	}

	opts, err := parseDetachOptions(os.Args[2:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		usage()
		return
	}

	if opts.User == "" && opts.UserFile == "" {
		fmt.Fprintln(os.Stderr, "Specify user or csv file")
		usage()
		return
	}

	err = opts.Infra.Startup()
	if err != nil {
		seelog.Errorf("Unable to start operation: %s", err)
		return
	}

	defer opts.Infra.Shutdown()

	seelog.Tracef("options: %s", util.MarshalObjectToString(opts))

	token, err := opts.Infra.LoadOrAuthBusinessManagement()
	if err != nil || token == "" {
		seelog.Errorf("Unable to acquire token (error: %s)", err)
		return
	}

	if opts.User != "" {
		detach.DetachUser(token, opts.User, opts.DryRun)
	}
	if opts.UserFile != "" {
		detach.DetachUserByList(token, opts.UserFile, opts.DryRun)
	}
}
