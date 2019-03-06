package cmd_file

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_file"
	"github.com/watermint/toolbox/model/dbx_file/copy_ref"
	"go.uber.org/zap"
)

type CmdFileMirror struct {
	*cmd.SimpleCommandlet
	optFromAccount string
	optToAccount   string
	optFromPath    string
	optToPath      string
}

func (CmdFileMirror) Name() string {
	return "mirror"
}

func (CmdFileMirror) Desc() string {
	return "cmd.file.mirror.desc"
}

func (CmdFileMirror) Usage() func(usage cmd.CommandUsage) {
	return nil
}

func (z *CmdFileMirror) FlagConfig(f *flag.FlagSet) {
	descFromAccount := z.ExecContext.Msg("cmd.file.mirror.flag.from_account").Text()
	f.StringVar(&z.optFromAccount, "from-account", "", descFromAccount)

	descToAccount := z.ExecContext.Msg("cmd.file.mirror.flag.to_account").Text()
	f.StringVar(&z.optToAccount, "to-account", "", descToAccount)

	descFromPath := z.ExecContext.Msg("cmd.file.mirror.flag.from_path").Text()
	f.StringVar(&z.optFromPath, "from-path", "", descFromPath)

	descToPath := z.ExecContext.Msg("cmd.file.mirror.flag.to_path").Text()
	f.StringVar(&z.optToPath, "to-path", "", descToPath)
}

func (z *CmdFileMirror) Exec(args []string) {
	if z.optFromAccount == "" ||
		z.optToAccount == "" ||
		z.optFromPath == "" ||
		z.optToPath == "" {

		//TODO msg.json
		z.ExecContext.Msg("cmd.file.mirror.err.not_enough_params").TellError()
		return
	}

	// Ask for FROM account authentication
	z.ExecContext.Msg("cmd.file.mirror.prompt.ask_from_account_auth").Tell()
	auFrom := dbx_auth.NewAuth(z.ExecContext, z.optFromAccount)
	acFrom, err := auFrom.Auth(dbx_auth.DropboxTokenFull)
	if err != nil {
		return
	}

	// Ask for TO account authentication
	z.ExecContext.Msg("cmd.file.mirror.prompt.ask_to_account_auth").Tell()
	auTo := dbx_auth.NewAuth(z.ExecContext, z.optToAccount)
	acTo, err := auTo.Auth(dbx_auth.DropboxTokenFull)
	if err != nil {
		return
	}

	crs := copy_ref.CopyRefSave{
		OnError: z.DefaultErrorHandler,
		OnFile: func(file *dbx_file.File) bool {
			z.ExecContext.Msg("cmd.file.mirror.progress.file.done").WithData(struct {
				FromPath string
				ToPath   string
			}{
				FromPath: z.optFromPath,
				ToPath:   file.PathDisplay,
			}).Tell()
			return true
		},
		OnFolder: func(folder *dbx_file.Folder) bool {
			z.ExecContext.Msg("cmd.file.mirror.progress.folder.done").WithData(struct {
				FromPath string
				ToPath   string
			}{
				FromPath: z.optFromPath,
				ToPath:   folder.PathDisplay,
			}).Tell()
			return true
		},
	}
	crg := copy_ref.CopyRefGet{
		OnError: z.DefaultErrorHandler,
		OnEntry: func(ref copy_ref.CopyRef) bool {
			z.ExecContext.Log().Debug("Trying to copy", zap.String("ref", ref.CopyReference), zap.String("toPath", z.optToPath))
			crs.Save(acTo, ref, z.optToPath)
			return true
		},
	}
	crg.Get(acFrom, z.optFromPath)
}
