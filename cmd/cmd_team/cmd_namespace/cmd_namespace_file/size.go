package cmd_namespace_file

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_size"
	"github.com/watermint/toolbox/report"
)

type CmdTeamNamespaceFileSize struct {
	*cmd.SimpleCommandlet
	report report.Factory
	nsz    dbx_size.NamespaceSizes
}

func (CmdTeamNamespaceFileSize) Name() string {
	return "size"
}

func (CmdTeamNamespaceFileSize) Desc() string {
	return "Calculate size of namespaces"
}

func (CmdTeamNamespaceFileSize) Usage() string {
	return ""
}

func (z *CmdTeamNamespaceFileSize) FlagConfig(f *flag.FlagSet) {
	z.report.FlagConfig(f)

	descIncludeTeamFolder := "Include team folders"
	f.BoolVar(&z.nsz.OptIncludeTeamFolder, "include-team-folder", true, descIncludeTeamFolder)

	descIncludeSharedFolder := "Include shared folders"
	f.BoolVar(&z.nsz.OptIncludeSharedFolder, "include-shared-folder", true, descIncludeSharedFolder)

	descIncludeAppFolder := "Include app folders"
	f.BoolVar(&z.nsz.OptIncludeAppFolder, "include-app-folder", false, descIncludeAppFolder)

	descIncludeMemberFolder := "Include team member folders"
	f.BoolVar(&z.nsz.OptIncludeMemberFolder, "include-member-folder", false, descIncludeMemberFolder)

	descUseCached := "Use cached information, or create cache if not exist"
	f.StringVar(&z.nsz.OptCachePath, "cache", "", descUseCached)

	descOptDepth := "Depth directories deep"
	f.IntVar(&z.nsz.OptDepth, "depth", 2, descOptDepth)
}

func (z *CmdTeamNamespaceFileSize) Exec(args []string) {
	z.report.Init(z.Log())
	defer z.report.Close()

	apiFile, err := z.ExecContext.LoadOrAuthBusinessFile()
	if err != nil {
		return
	}

	z.nsz.Init(z.Log())
	z.nsz.Load(apiFile)

	z.Log().Info("Reporting result")
	for _, sz := range z.nsz.Sizes {
		z.report.Report(sz)
	}
}
