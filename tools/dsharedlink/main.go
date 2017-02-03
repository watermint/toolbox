package main

import (
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/knowledge"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/integration/auth"
	"github.com/watermint/toolbox/service/dsharedlink"
	"os"
	"path/filepath"
	"time"
)

var (
	AppKey    string = ""
	AppSecret string = ""
)

func usage() {
	tmpl := `{{.AppName}} {{.AppVersion}} ({{.AppHash}}):

Expire shared links at +7 days if expiration not set

{{.Command}} expire -team -days 7
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

type ListFlags struct {
}

type ExpireFlags struct {
	Team       bool
	Infra      *infra.InfraOpts
	Days       int
	Overwrite  bool
	TargetUser string
}

func parseExpireFlags(args []string) (*ExpireFlags, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	pf := &ExpireFlags{}
	pf.Infra = infra.PrepareInfraFlags(f)

	descDays := "Specify expire date in days"
	f.IntVar(&pf.Days, "days", 0, descDays)

	descTeam := "Apply for Team (Dropbox Business)"
	f.BoolVar(&pf.Team, "team", false, descTeam)

	descOverwrite := "Overwrite expiration if existing expiration exceeds specified duration"
	f.BoolVar(&pf.Overwrite, "overwrite", false, descOverwrite)

	descTargetUser := "Specify target user by email for test purpose"
	f.StringVar(&pf.TargetUser, "target-user", "", descTargetUser)

	f.SetOutput(os.Stderr)
	f.Parse(args)

	return pf, nil
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	if os.Args[1] != "expire" {
		usage()
		return
	}

	ef, err := parseExpireFlags(os.Args[2:])

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		usage()
		return
	}
	if !ef.Team {
		fmt.Fprintln(os.Stderr, "Operation for personal account not yet supported")
		return
	}
	if ef.Days < 1 {
		fmt.Fprintln(os.Stderr, "Expiration days must be grater equal 1")
		return
	}

	err = infra.InfraStartup(ef.Infra)
	if err != nil {
		seelog.Errorf("Unable to start operation: %s", err)
		return
	}

	defer infra.InfraShutdown()

	seelog.Tracef("options: %s", util.MarshalObjectToString(ef))

	a := auth.DropboxAuthenticator{
		AuthFile:  filepath.Join(ef.Infra.WorkPath, knowledge.AppName+".secret"),
		AppKey:    AppKey,
		AppSecret: AppSecret,
	}

	token, err := a.LoadOrAuth(ef.Team)
	if err != nil || token == "" {
		seelog.Errorf("Unable to acquire token (error: %s)", err)
		return
	}
	if ef.Infra.CleanupToken {
		defer auth.RevokeToken(token)
	}

	dsharedlink.UpdateSharedLinkForTeam(token, dsharedlink.UpdateSharedLinkExpireContext{
		TargetUser: ef.TargetUser,
		Expiration: time.Duration(ef.Days) * time.Hour * 24,
		Overwrite:  ef.Overwrite,
	})
}
