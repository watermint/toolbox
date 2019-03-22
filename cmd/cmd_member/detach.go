package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/app/app_io"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/model/dbx_auth"
	"go.uber.org/zap"
	"strings"
)

type CmdMemberDetach struct {
	*cmd.SimpleCommandlet
	optCsv string
}

func (CmdMemberDetach) Name() string {
	return "detach"
}
func (CmdMemberDetach) Desc() string {
	return "cmd.member.detach.desc"
}
func (z *CmdMemberDetach) Usage() func(cmd.CommandUsage) {
	return nil
}
func (z *CmdMemberDetach) FlagConfig(f *flag.FlagSet) {
	descCsv := z.ExecContext.Msg("cmd.member.detach.flag.csv").T()
	f.StringVar(&z.optCsv, "csv", "", descCsv)
}

func (z *CmdMemberDetach) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, dbx_auth.DropboxTokenBusinessManagement)
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
			err = svm.Remove(member, sv_member.Downgrade())
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
