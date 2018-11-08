package cmd_namespace_file

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/model/model_file"
	"github.com/watermint/toolbox/report"
)

type CmdTeamNamespaceSize struct {
	*cmdlet.SimpleCommandlet
	report report.Factory
	nsz    model_file.NamespaceSizes
}

func (CmdTeamNamespaceSize) Name() string {
	return "size"
}

func (CmdTeamNamespaceSize) Desc() string {
	return "Calculate size of namespaces"
}

func (CmdTeamNamespaceSize) Usage() string {
	return ""
}

func (z *CmdTeamNamespaceSize) FlagConfig(f *flag.FlagSet) {
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

func (z *CmdTeamNamespaceSize) Exec(args []string) {
	z.report.Init(z.Log())
	defer z.report.Close()

	apiFile, err := z.ExecContext.LoadOrAuthBusinessFile()
	if err != nil {
		return
	}

	nsz := &model_file.NamespaceSizes{}
	nsz.Init(z.Log())
	nsz.Load(apiFile)

	z.Log().Info("Reporting result")
	for _, sz := range nsz.Sizes {
		z.report.Report(sz)
	}
}
