package cmd_teamfolder

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_size"
)

type CmdTeamFolderSize struct {
	*cmd.SimpleCommandlet
	optDepth     int
	optCachePath string
	report       app_report.Factory
}

func (CmdTeamFolderSize) Name() string {
	return "size"
}

func (CmdTeamFolderSize) Desc() string {
	return "cmd.teamfolder.size.desc"
}

func (CmdTeamFolderSize) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamFolderSize) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descOptDepth := z.ExecContext.Msg("cmd.teamfolder.size.flag.depth").T()
	f.IntVar(&z.optDepth, "depth", 2, descOptDepth)

	descUseCached := z.ExecContext.Msg("cmd.teamfolder.size.flag.cache").T()
	f.StringVar(&z.optCachePath, "cache", "", descUseCached)
}

func (z *CmdTeamFolderSize) Exec(args []string) {
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
