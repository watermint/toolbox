package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
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
	return "Convert account into Dropbox Basic"
}
func (z *CmdMemberDetach) Usage() string {
	return z.provision.Usage()
}
func (z *CmdMemberDetach) FlagConfig(f *flag.FlagSet) {
	z.provision.FlagConfig(f)
}

func (z *CmdMemberDetach) Exec(args []string) {
	z.provision.Logger = z.Log()
	err := z.provision.Load(args)
	if err != nil {
		z.PrintUsage(z)
		return
	}

	apiMgmt, err := z.ExecContext.LoadOrAuthBusinessManagement()
	if err != nil {
		return
	}

	rm := dbx_member.MemberRemove{
		OnError: z.DefaultErrorHandler,
		OnFailure: func(email string, reason dbx_api.ApiError) bool {
			z.Log().Error(
				"Unable to detach user",
				zap.String("email", email),
				zap.String("reason", reason.ErrorTag),
			)
			return true
		},
		OnSuccess: func(email string) bool {
			z.Log().Info("User detached", zap.String("email", email))
			return true
		},
	}
	for _, m := range z.provision.Members {
		z.Log().Info("Detaching account", zap.String("email", m.Email))
		rm.Remove(apiMgmt, m.Email, false, true)
	}
}
