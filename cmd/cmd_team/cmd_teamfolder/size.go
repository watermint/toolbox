package cmd_teamfolder

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_size"
	"github.com/watermint/toolbox/report"
)

type CmdTeamTeamFolderSize struct {
	*cmd.SimpleCommandlet
	optDepth     int
	optCachePath string
	report       report.Factory
}

func (CmdTeamTeamFolderSize) Name() string {
	return "size"
}

func (CmdTeamTeamFolderSize) Desc() string {
	return "cmd.team.teamfolder.size.desc"
}

func (CmdTeamTeamFolderSize) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamTeamFolderSize) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descOptDepth := z.ExecContext.Msg("cmd.team.teamfolder.size.flag.depth").Text()
	f.IntVar(&z.optDepth, "depth", 2, descOptDepth)

	descUseCached := z.ExecContext.Msg("cmd.team.teamfolder.size.flag.cache").Text()
	f.StringVar(&z.optCachePath, "cache", "", descUseCached)
}

func (z *CmdTeamTeamFolderSize) Exec(args []string) {
	z.report.Init(z.ExecContext)
	defer z.report.Close()

	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiFile, err := au.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	nsz := &dbx_size.NamespaceSizes{}
	nsz.Init(z.ExecContext)
	nsz.OptIncludeTeamFolder = true
	nsz.OptIncludeSharedFolder = false
	nsz.OptIncludeAppFolder = false
	nsz.OptIncludeMemberFolder = false
	nsz.Load(apiFile)

	z.Log().Info("Reporting result")
	for _, sz := range nsz.Sizes {
		z.report.Report(sz)
	}
}
