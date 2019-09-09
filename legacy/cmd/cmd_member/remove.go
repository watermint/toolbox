package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_io"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	"go.uber.org/zap"
	"strings"
)

type CmdMemberRemove struct {
	*cmd2.SimpleCommandlet
	optKeepAccount bool
	optWipeData    bool
	optCsv         string
}

func (CmdMemberRemove) Name() string {
	return "remove"
}
func (CmdMemberRemove) Desc() string {
	return "cmd.member.remove.desc"
}
func (z *CmdMemberRemove) Usage() func(cmd2.CommandUsage) {
	return nil
}
func (z *CmdMemberRemove) FlagConfig(f *flag.FlagSet) {
	descKeepAccount := z.ExecContext.Msg("cmd.member.remove.flag.keep_account").T()
	f.BoolVar(&z.optKeepAccount, "keep-account", false, descKeepAccount)

	descWipeData := z.ExecContext.Msg("cmd.member.remove.flag.wipe_data").T()
	f.BoolVar(&z.optWipeData, "wipe-data", true, descWipeData)

	descCsv := z.ExecContext.Msg("cmd.member.remove.flag.csv").T()
	f.StringVar(&z.optCsv, "csv", "", descCsv)

}

func (z *CmdMemberRemove) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessManagement())
	if err != nil {
		return
	}
	svm := sv_member.New(ctx)

	err = app_io.NewCsvLoader(z.ExecContext, z.optCsv).
		OnRow(func(cols []string) error {
			if len(cols) < 1 {
				return nil
			}
			email := strings.TrimSpace(cols[0])
			if !api_util.RegexEmail.MatchString(email) {
				z.Log().Debug("skip: the data is not looking alike an email address", zap.String("email", email))
				return nil
			}
			member, err := svm.ResolveByEmail(email)
			if err != nil {
				z.Log().Debug("member not found", zap.Error(err))
				return nil
			}
			opts := make([]sv_member.RemoveOpt, 0)
			if z.optKeepAccount {
				opts = append(opts, sv_member.RemoveWipeData())
			}
			if z.optWipeData {
				opts = append(opts, sv_member.RemoveWipeData())
			}
			err = svm.Remove(member, opts...)
			if err != nil {
				z.Log().Debug("unable to detach", zap.Error(err))
				return nil
			}
			return nil
		}).
		Load()

	if err != nil {
		z.Log().Debug("Unable to load", zap.Error(err))
	}
}
