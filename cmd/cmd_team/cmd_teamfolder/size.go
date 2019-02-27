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
	return "Calculate size of team folder"
}

func (CmdTeamTeamFolderSize) Usage() string {
	return ""
}

func (z *CmdTeamTeamFolderSize) FlagConfig(f *flag.FlagSet) {
	z.report.FlagConfig(f)

	descOptDepth := "Depth directories deep"
	f.IntVar(&z.optDepth, "depth", 2, descOptDepth)

	descUseCached := "Use cached information, or create cache if not exist"
	f.StringVar(&z.optCachePath, "cache", "", descUseCached)
}

func (z *CmdTeamTeamFolderSize) Exec(args []string) {
	z.report.Init(z.Log())
	defer z.report.Close()

	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiFile, err := au.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	nsz := &dbx_size.NamespaceSizes{}
	nsz.Init(z.Log())
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
