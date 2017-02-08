package commands

import (
	"errors"
	"flag"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/service/members"
	"os"
)

type ListOptions struct {
	Infra     *infra.InfraOpts
	OutputCsv string
	Status    string
}

func parseListOptions(args []string) (*ListOptions, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	opts := &ListOptions{}
	opts.Infra = infra.PrepareInfraFlags(f)

	descCsv := "Output CSV path"
	f.StringVar(&opts.OutputCsv, "csv", "", descCsv)

	descStatus := "Status (all|active|invited|suspended|removed)"
	f.StringVar(&opts.Status, "status", "all", descStatus)

	f.SetOutput(os.Stderr)
	f.Parse(args)

	switch opts.Status {
	case "all":
	case "active":
	case "invited":
	case "suspended":
	case "removed":
	default:
		return nil, errors.New("Invalid status: " + opts.Status)
	}

	return opts, nil
}

func List(args []string) error {
	opts, err := parseListOptions(args)
	if err != nil {
		return err
	}

	err = opts.Infra.Startup()
	if err != nil {
		seelog.Errorf("Unable to start operation: %s", err)
		return err
	}
	defer opts.Infra.Shutdown()

	seelog.Tracef("options: %s", util.MarshalObjectToString(opts))

	token, err := opts.Infra.LoadOrAuthBusinessInfo()
	if err != nil || token == "" {
		seelog.Errorf("Unable to acquire token (error: %s)", err)
		return err
	}

	if opts.OutputCsv == "" {
		return members.ListMembers(token, os.Stdout, opts.Status)
	} else {
		f, err := os.Open(opts.OutputCsv)
		if err != nil {
			seelog.Errorf("Unable to write file[%s] erorr[%s]", opts.OutputCsv, err)
			return err
		}
		defer f.Close()
		return members.ListMembers(token, f, opts.Status)
	}
}
