package commands

import (
	"errors"
	"flag"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/service/detach"
	"os"
)

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

func Detach(args []string) error {
	opts, err := parseDetachOptions(args)
	if err != nil {
		return err
	}

	if opts.User == "" && opts.UserFile == "" {
		return errors.New("Specify user or csv file")
	}

	defer opts.Infra.Shutdown()
	err = opts.Infra.Startup()
	if err != nil {
		seelog.Errorf("Unable to start operation: %s", err)
		return err
	}
	seelog.Tracef("options: %s", util.MarshalObjectToString(opts))

	token, err := opts.Infra.LoadOrAuthBusinessManagement()
	if err != nil || token == "" {
		seelog.Errorf("Unable to acquire token (error: %s)", err)
		return err
	}

	if opts.User != "" {
		detach.DetachUser(token, opts.User, opts.DryRun)
	}
	if opts.UserFile != "" {
		detach.DetachUserByList(token, opts.UserFile, opts.DryRun)
	}
	return nil
}
