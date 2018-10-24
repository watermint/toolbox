package cmd_member

import (
	"flag"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
	"io"
	"os"
)

type CmdMemberInvite struct {
	*cmdlet.SimpleCommandlet
	optCsv     string
	optSilent  bool
	apiContext *dbx_api.Context
	report     cmdlet.Report
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

func (c *CmdMemberInvite) Exec(ec *infra.ExecContext, args []string) {
	if err := ec.Startup(); err != nil {
		return
	}
	defer ec.Shutdown()
	if c.optCsv == "" {
		seelog.Errorf("Please specify input csv")
		seelog.Flush()
		c.PrintUsage(c)
		return
	}

	apiMgmt, err := ec.LoadOrAuthBusinessManagement()
	if err != nil {
		return
	}

	newMembers, err := c.loadCsv()
	if err != nil {
		return
	}

	type SuccessReport struct {
		Profile *dbx_profile.Profile `json:"profile,omitempty"`
		Role    string               `json:"role,omitempty"`
	}
	type FailureReport struct {
		Email  string `json:"email,omitempty"`
		Reason string `json:"reason,omitempty"`
	}

	type InviteReport struct {
		Result  string         `json:"result"`
		Success *SuccessReport `json:"success,omitempty"`
		Failure *FailureReport `json:"failure,omitempty"`
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
	handleSuccess := func(profile *dbx_profile.Profile, role string) bool {
		c.report.Report(
			InviteReport{
				Result: "success",
				Success: &SuccessReport{
					Profile: profile,
					Role:    role,
				},
			},
		)
		return true
	}

	mi := dbx_team.MembersInvite{
		OnFailure: handleFailure,
		OnSuccess: handleSuccess,
		OnError:   cmdlet.DefaultErrorHandler,
	}
	mi.Invite(apiMgmt, newMembers)
}

func (c *CmdMemberInvite) loadCsv() (newMembers []*dbx_team.NewMember, err error) {
	f, err := os.Open(c.optCsv)
	if err != nil {
		seelog.Warnf("Unable to open file[%s] : error[%s]", c.optCsv, err)
		return nil, err
	}
	csv := util.NewBomAwareCsvReader(f)

	newMembers = make([]*dbx_team.NewMember, 0)

	for {
		cols, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			seelog.Warnf("Unable to read CSV file [%s] : error[%s]", c.optCsv, err)
			return nil, err
		}
		if len(cols) < 1 {
			seelog.Warnf("Skip line: [%v]", cols)
			continue
		}

		newMember := &dbx_team.NewMember{}
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
