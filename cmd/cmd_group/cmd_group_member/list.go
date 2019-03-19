package cmd_group_member

import (
	"encoding/json"
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"go.uber.org/zap"
)

type CmdGroupMemberList struct {
	*cmd.SimpleCommandlet

	apiContext *dbx_api.DbxContext
	report     app_report.Factory
}

func (z *CmdGroupMemberList) Name() string {
	return "list"
}

func (z *CmdGroupMemberList) Desc() string {
	return "cmd.group.member.list.desc"
}

func (z *CmdGroupMemberList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdGroupMemberList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdGroupMemberList) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, dbx_auth.DropboxTokenBusinessInfo)
	if err != nil {
		return
	}
	gsv := sv_group.New(ctx)
	groups, err := gsv.List()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	for _, group := range groups {
		msv := sv_group_member.New(ctx, group)
		members, err := msv.List()
		if err != nil {
			ctx.ErrorMsg(err).TellError()
			return
		}
		for _, m := range members {
			raw := struct {
				Group  json.RawMessage `json:"group"`
				Member json.RawMessage `json:"member"`
			}{
				Group:  group.Raw,
				Member: m.Raw,
			}
			r, err := json.Marshal(raw)
			if err != nil {
				z.Log().Warn("unable to marshal raw JSON", zap.Error(err))
				r = json.RawMessage("{}")
			}

			type Report struct {
				Raw                 json.RawMessage `json:"-"`
				GroupId             string          `json:"group_id"`
				GroupName           string          `json:"group_name"`
				GroupManagementType string          `json:"group_management_type"`
				AccessType          string          `json:"access_type"`
				AccountId           string          `json:"account_id"`
				TeamMemberId        string          `json:"team_member_id"`
				Email               string          `json:"email"`
				Status              string          `json:"status"`
				Surname             string          `json:"surname"`
				GivenName           string          `json:"given_name"`
			}
			row := Report{
				Raw:                 r,
				GroupId:             group.GroupId,
				GroupName:           group.GroupName,
				GroupManagementType: group.GroupManagementType,
				AccessType:          m.AccessType,
				AccountId:           m.AccountId,
				TeamMemberId:        m.TeamMemberId,
				Email:               m.Email,
				Status:              m.Status,
				Surname:             m.Surname,
				GivenName:           m.GivenName,
			}

			z.report.Report(row)
		}
	}
}
