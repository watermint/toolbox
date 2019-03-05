package cmd_file

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_file"
	"github.com/watermint/toolbox/report"
)

type CmdFileMetadata struct {
	*cmd.SimpleCommandlet
	report report.Factory
}

func (z *CmdFileMetadata) Name() string {
	return "metadata"
}

func (z *CmdFileMetadata) Desc() string {
	return "cmd.file.metadata.desc"
}

func (z *CmdFileMetadata) Usage() func(usage cmd.CommandUsage) {
	return func(usage cmd.CommandUsage) {
		z.ExecContext.Msg("cmd.file.metadata.usage").WithData(usage).Tell()
	}
}

func (z *CmdFileMetadata) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdFileMetadata) Exec(args []string) {
	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	ac, err := au.Auth(dbx_auth.DropboxTokenFull)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	for _, p := range args {
		md := dbx_file.Metadata{
			Path:                            p,
			IncludeDeleted:                  true,
			IncludeHasExplicitSharedMembers: true,
			IncludeMediaInfo:                true,
			OnError:                         z.DefaultErrorHandler,
			OnDelete: func(deleted *dbx_file.Deleted) bool {
				z.report.Report(deleted)
				return true
			},
			OnFile: func(file *dbx_file.File) bool {
				z.report.Report(file)
				return true
			},
			OnFolder: func(folder *dbx_file.Folder) bool {
				z.report.Report(folder)
				return true
			},
		}
		md.Get(ac)
	}
}
