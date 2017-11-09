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

type ExtIdOptions struct {
	Infra          *infra.InfraOpts
	OpAssign       bool
	OpShow         bool
	DryRun         bool
	TargetUser     string
	TargetAllUsers bool
}

func parseExtIdOptions(args []string) (*ExtIdOptions, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	opts := &ExtIdOptions{}
	opts.Infra = infra.PrepareInfraFlags(f)

	descPseudo := "Assign new pseudo external ID (like: 'Email email@address')"
	f.BoolVar(&opts.OpAssign, "assign-pseudo-id", false, descPseudo)

	descShow := "List external IDs"
	f.BoolVar(&opts.OpShow, "list", false, descShow)

	descDryRun := "Dry run"
	f.BoolVar(&opts.DryRun, "dry-run", true, descDryRun)

	descUser := "Specify user email address to apply change"
	f.StringVar(&opts.TargetUser, "user", "", descUser)

	descAllUsers := "Apply changes for all team members"
	f.BoolVar(&opts.TargetAllUsers, "all-users", false, descAllUsers)

	f.SetOutput(os.Stderr)
	f.Parse(args)

	return opts, nil
}

func ExtId(arg []string) error {
	opts, err := parseExtIdOptions(arg)
	if err != nil {
		return err
	}

	defer opts.Infra.Shutdown()
	err = opts.Infra.Startup()
	if err != nil {
		seelog.Warnf("Unable to start operation : error[%s]", err)
		return err
	}
	seelog.Tracef("options: %s", util.MarshalObjectToString(opts))

	token, err := opts.Infra.LoadOrAuthBusinessManagement()
	if err != nil || token == "" {
		seelog.Warnf("Unable to acquire token : error[%s]", err)
		return err
	}

	if opts.OpShow {
		return extIdShow(token, opts)
	}
	if opts.OpAssign {
		return extIdAssign(token, opts)
	}

	return errors.New("Specify operation mode (-list or -assign-pseudo-id)")
}

func extIdShow(token string, opts *ExtIdOptions) error {
	if opts.TargetUser != "" {
		return members.ShowExtIdByEmail(token, opts.TargetUser)
	}
	if opts.TargetAllUsers {
		return members.ShowExtIdForTeam(token)
	}
	return errors.New("Specify target user(s)")
}

func extIdAssign(token string, opts *ExtIdOptions) error {
	if opts.TargetUser != "" {
		return members.AssignPseudoExtIdByEmail(token, opts.TargetUser, opts.DryRun)
	}
	if opts.TargetAllUsers {
		return members.AssignPseudoExtIdForTeam(token, opts.DryRun)
	}
	return errors.New("Specify target user(s)")
}
