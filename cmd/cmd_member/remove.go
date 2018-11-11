package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_member"
	"go.uber.org/zap"
)

type CmdMemberRemove struct {
	*cmd.SimpleCommandlet
	optKeepAccount bool
	optWipeData    bool
}

func (CmdMemberRemove) Name() string {
	return "remove"
}
func (CmdMemberRemove) Desc() string {
	return "Remove the member from the team"
}
func (CmdMemberRemove) Usage() string {
	return ""
}
func (z *CmdMemberRemove) FlagConfig(f *flag.FlagSet) {
	descKeepAccount := "Convert account into Dropbox Basic"
	f.BoolVar(&z.optKeepAccount, "keep-account", false, descKeepAccount)

	descWipeData := "Wipe data"
	f.BoolVar(&z.optWipeData, "wipe-data", true, descWipeData)
}

func (z *CmdMemberRemove) Exec(args []string) {
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
	for _, e := range args {
		z.Log().Info("Removing account", zap.String("email", e))
		rm.Remove(apiMgmt, e, z.optWipeData, z.optKeepAccount)
	}
}
