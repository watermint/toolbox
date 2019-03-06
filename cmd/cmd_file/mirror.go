package cmd_file

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_file/copy_ref"
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

	m := copy_ref.Mirror{
		FromApi:          acFrom,
		FromPath:         z.optFromPath,
		FromAccountAlias: z.optFromAccount,
		ToApi:            acTo,
		ToPath:           z.optToPath,
		ToAccountAlias:   z.optToAccount,
		ExecContext:      z.ExecContext,
	}
	m.Mirror()
}
