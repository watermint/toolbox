package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
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
	return "Remove the member from the team"
}
func (z *CmdMemberRemove) Usage() string {
	return z.provision.Usage()
}
func (z *CmdMemberRemove) FlagConfig(f *flag.FlagSet) {
	z.provision.FlagConfig(f)

	descKeepAccount := "Convert account into Dropbox Basic"
	f.BoolVar(&z.optKeepAccount, "keep-account", false, descKeepAccount)

	descWipeData := "Wipe data"
	f.BoolVar(&z.optWipeData, "wipe-data", true, descWipeData)
}

func (z *CmdMemberRemove) Exec(args []string) {
	z.provision.Logger = z.Log()
	err := z.provision.Load(args)
	if err != nil {
		z.PrintUsage(z.ExecContext, z)
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
				"Unable to remove user",
				zap.String("email", email),
				zap.String("reason", reason.ErrorTag),
			)
			return true
		},
		OnSuccess: func(email string) bool {
			z.Log().Info("User removed", zap.String("email", email))
			return true
		},
	}
	for _, m := range z.provision.Members {
		z.Log().Info("Removing account", zap.String("email", m.Email))
		rm.Remove(apiMgmt, m.Email, z.optWipeData, z.optKeepAccount)
	}
}
