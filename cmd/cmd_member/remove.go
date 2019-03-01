package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_member"
	"go.uber.org/zap"
)

type CmdMemberRemove struct {
	*cmd.SimpleCommandlet
	optKeepAccount bool
	optWipeData    bool
	provision      MembersProvision
}

func (CmdMemberRemove) Name() string {
	return "remove"
}
func (CmdMemberRemove) Desc() string {
	return "cmd.member.remove.desc"
}
func (z *CmdMemberRemove) Usage() func(cmd.CommandUsage) {
	return z.provision.Usage()
}
func (z *CmdMemberRemove) FlagConfig(f *flag.FlagSet) {
	z.provision.ec = z.ExecContext
	z.provision.FlagConfig(f)

	descKeepAccount := z.ExecContext.Msg("cmd.member.remove.flag.keep_account").Text()
	f.BoolVar(&z.optKeepAccount, "keep-account", false, descKeepAccount)

	descWipeData := z.ExecContext.Msg("cmd.member.remove.flag.wipe_data").Text()
	f.BoolVar(&z.optWipeData, "wipe-data", true, descWipeData)
}

func (z *CmdMemberRemove) Exec(args []string) {
	z.provision.Logger = z.Log()
	err := z.provision.Load(args)
	if err != nil {
		z.PrintUsage(z.ExecContext, z)
		return
	}

	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiMgmt, err := au.Auth(dbx_auth.DropboxTokenBusinessManagement)
	if err != nil {
		return
	}

	rm := dbx_member.MemberRemove{
		OnError: z.DefaultErrorHandler,
		OnFailure: func(email string, reason dbx_api.ApiError) bool {
			z.ExecContext.Msg("cmd.member.remove.failure").WithData(struct {
				Email  string
				Reason string
			}{
				Email:  email,
				Reason: reason.ErrorTag,
			}).TellError()
			z.Log().Warn(
				"Unable to remove user",
				zap.String("email", email),
				zap.String("reason", reason.ErrorTag),
			)
			return true
		},
		OnSuccess: func(email string) bool {
			z.ExecContext.Msg("cmd.member.remove.success").WithData(struct {
				Email string
			}{
				Email: email,
			})
			z.Log().Info("User removed", zap.String("email", email))
			return true
		},
	}
	for _, m := range z.provision.Members {
		z.ExecContext.Msg("cmd.member.remove.progress").WithData(struct {
			Email string
		}{
			Email: m.Email,
		}).Tell()

		z.Log().Info("Removing account", zap.String("email", m.Email))
		rm.Remove(apiMgmt, m.Email, z.optWipeData, z.optKeepAccount)
	}
}
