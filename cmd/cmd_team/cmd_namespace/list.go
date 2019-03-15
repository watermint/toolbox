package cmd_namespace

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_namespace"
)

type CmdTeamNamespaceList struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
}

func (CmdTeamNamespaceList) Name() string {
	return "list"
}

func (CmdTeamNamespaceList) Desc() string {
	return "cmd.team.namespace.list.desc"
}

func (CmdTeamNamespaceList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamNamespaceList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdTeamNamespaceList) Exec(args []string) {
	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiFile, err := au.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	l := dbx_namespace.NamespaceList{
		OnError: z.DefaultErrorHandler,
		OnEntry: func(namespace *dbx_namespace.Namespace) bool {
			z.report.Report(namespace)
			return true
		},
	}
	l.List(apiFile)
}
