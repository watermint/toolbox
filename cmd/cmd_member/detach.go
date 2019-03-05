package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_member"
	"go.uber.org/zap"
)

type CmdMemberDetach struct {
	*cmd.SimpleCommandlet
	provision MembersProvision
}

func (CmdMemberDetach) Name() string {
	return "detach"
}
func (CmdMemberDetach) Desc() string {
	return "cmd.member.detach.desc"
}
func (z *CmdMemberDetach) Usage() func(cmd.CommandUsage) {
	return z.provision.Usage()
}
func (z *CmdMemberDetach) FlagConfig(f *flag.FlagSet) {
	z.provision.ec = z.ExecContext
	z.provision.FlagConfig(f)
}

func (z *CmdMemberDetach) Exec(args []string) {
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
			z.ExecContext.Msg("cmd.member.detach.failure").WithData(struct {
				Email  string
				Reason string
			}{
				Email:  email,
				Reason: reason.ErrorTag,
			}).TellFailure()

			z.Log().Debug(
				"Unable to detach user",
				zap.String("email", email),
				zap.String("reason", reason.ErrorTag),
			)
			return true
		},
		OnSuccess: func(email string) bool {
			z.ExecContext.Msg("cmd.member.detach.success").WithData(struct {
				Email string
			}{
				Email: email,
			}).TellSuccess()
			z.Log().Debug("User detached", zap.String("email", email))
			return true
		},
	}
	for _, m := range z.provision.Members {
		z.ExecContext.Msg("cmd.member.detach.progress").WithData(struct {
			Email string
		}{
			Email: m.Email,
		}).TellSuccess()
		z.Log().Debug("Detaching account", zap.String("email", m.Email))
		rm.Remove(apiMgmt, m.Email, false, true)
	}
}
