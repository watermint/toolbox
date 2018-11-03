package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_member"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/report"
	"go.uber.org/zap"
	"io"
	"os"
)

type CmdMemberInvite struct {
	*cmdlet.SimpleCommandlet
	optCsv     string
	optSilent  bool
	apiContext *dbx_api.Context
	report     report.Factory
}

func (c *CmdMemberInvite) Name() string {
	return "invite"
}

func (c *CmdMemberInvite) Desc() string {
	return "Invite members"
}

func (c *CmdMemberInvite) Usage() string {
	return `{{.Command}} -csv MEMBER_FILENAME`
}

func (c *CmdMemberInvite) FlagConfig(f *flag.FlagSet) {
	descCsv := "CSV file name"
	f.StringVar(&c.optCsv, "csv", "", descCsv)

	descSilent := "Silent provisioning"
	f.BoolVar(&c.optSilent, "silent", false, descSilent)

	c.report.FlagConfig(f)
}

func (c *CmdMemberInvite) Exec(args []string) {
	if c.optCsv == "" {
		c.Log().Error("Please specify input csv")
		c.PrintUsage(c)
		return
	}

	apiMgmt, err := c.ExecContext.LoadOrAuthBusinessManagement()
	if err != nil {
		return
	}

	newMembers, err := c.loadCsv()
	if err != nil {
		return
	}

	c.report.Init(c.Log())
	defer c.report.Close()

	type FailureReport struct {
		Email  string `json:"email,omitempty"`
		Reason string `json:"reason,omitempty"`
	}
	type InviteReport struct {
		Result  string              `json:"result"`
		Success *dbx_profile.Member `json:"success,omitempty"`
		Failure *FailureReport      `json:"failure,omitempty"`
	}
	handleFailure := func(email string, reason string) bool {
		c.report.Report(
			InviteReport{
				Result: "failure",
				Failure: &FailureReport{
					Email:  email,
					Reason: reason,
				},
			},
		)
		return true
	}
	handleSuccess := func(member *dbx_profile.Member) bool {
		c.report.Report(
			InviteReport{
				Result:  "success",
				Success: member,
			},
		)
		return true
	}

	mi := dbx_member.MembersInvite{
		OnFailure: handleFailure,
		OnSuccess: handleSuccess,
		OnError:   c.DefaultErrorHandler,
	}
	mi.Invite(apiMgmt, newMembers)
}

func (c *CmdMemberInvite) loadCsv() (newMembers []*dbx_member.NewMember, err error) {
	f, err := os.Open(c.optCsv)
	if err != nil {
		c.Log().Warn(
			"Unable to open file",
			zap.String("file", c.optCsv),
			zap.Error(err),
		)
		return nil, err
	}
	csv := util.NewBomAwareCsvReader(f)

	newMembers = make([]*dbx_member.NewMember, 0)

	for {
		cols, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.Log().Warn(
				"Unable to read CSV file",
				zap.String("file", c.optCsv),
				zap.Error(err),
			)
			return nil, err
		}
		if len(cols) < 1 {
			c.Log().Warn("No column found in the row. Skip")
			continue
		}

		newMember := &dbx_member.NewMember{}
		newMember.MemberEmail = cols[0]
		if len(cols) >= 2 {
			newMember.MemberGivenName = cols[1]
		}
		if len(cols) >= 3 {
			newMember.MemberSurname = cols[2]
		}
		if c.optSilent {
			newMember.SendWelcomeEmail = false
		}

		newMembers = append(newMembers, newMember)
	}
	return newMembers, nil
}
