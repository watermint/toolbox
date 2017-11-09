package commands

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/knowledge"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/service/sharedlink"
	"os"
	"time"
)

func usage() {
	tmpl := `{{.AppName}} {{.AppVersion}} ({{.AppHash}}):

Expire shared links at +7 days if expiration not set

{{.Command}} link-expire -days 7
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

type ExpireFlags struct {
	Infra      *infra.InfraOpts
	Days       int
	Overwrite  bool
	TargetUser string
}

func parseLinkExpireFlags(args []string) (*ExpireFlags, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	pf := &ExpireFlags{}
	pf.Infra = infra.PrepareInfraFlags(f)

	descDays := "Specify expire date in days"
	f.IntVar(&pf.Days, "days", 0, descDays)

	descOverwrite := "Overwrite expiration if existing expiration exceeds specified duration"
	f.BoolVar(&pf.Overwrite, "overwrite", false, descOverwrite)

	descTargetUser := "Specify target user by email for test purpose"
	f.StringVar(&pf.TargetUser, "target-user", "", descTargetUser)

	f.SetOutput(os.Stderr)
	f.Parse(args)

	return pf, nil
}

func LinkExpire(args []string) error {
	opts, err := parseLinkExpireFlags(args)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		usage()
		return err
	}
	if opts.Days < 1 {
		fmt.Fprintln(os.Stderr, "Expiration days must be grater equal 1")
		return errors.New("expiration days must be grater equal 1")
	}

	defer opts.Infra.Shutdown()
	err = opts.Infra.Startup()
	if err != nil {
		seelog.Errorf("Unable to start operation: %s", err)
		return err
	}
	seelog.Tracef("options: %s", util.MarshalObjectToString(opts))

	token, err := opts.Infra.LoadOrAuthBusinessFile()
	if err != nil || token == "" {
		seelog.Errorf("Unable to acquire token (error: %s)", err)
		return err
	}

	sharedlink.UpdateSharedLinkForTeam(token, sharedlink.UpdateSharedLinkExpireContext{
		TargetUser: opts.TargetUser,
		Expiration: time.Duration(opts.Days) * time.Hour * 24,
		Overwrite:  opts.Overwrite,
	})
	return nil
}
