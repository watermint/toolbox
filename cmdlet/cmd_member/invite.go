package cmd_member

import (
	"errors"
	"flag"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/async"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
	"io"
	"os"
	"time"
)

type CmdMemberInvite struct {
	optCsv       string
	apiContext   *api.ApiContext
	infraContext *infra.InfraContext
}

func NewCmdMemberInvite() *CmdMemberInvite {
	c := CmdMemberInvite{
		infraContext: &infra.InfraContext{},
	}
	return &c
}

func (c *CmdMemberInvite) Name() string {
	return "invite"
}

func (c *CmdMemberInvite) Desc() string {
	return "Invite members"
}

func (c *CmdMemberInvite) UsageTmpl() string {
	return `
Usage: {{.Command}} -csv MEMBER_FILENAME
`
}

func (c *CmdMemberInvite) FlagSet() (f *flag.FlagSet) {
	f = flag.NewFlagSet(c.Name(), flag.ExitOnError)

	descCsv := "CSV file name"
	f.StringVar(&c.optCsv, "csv", "", descCsv)

	c.infraContext.PrepareFlags(f)
	return f
}

func (c *CmdMemberInvite) Exec(cc cmdlet.CommandletContext) error {
	_, err := cmdlet.ParseFlags(cc, c)
	if err != nil {
		return err
	}
	if c.optCsv == "" {
		return &cmdlet.CommandShowUsageError{
			Context:     cc,
			Instruction: "missing `-csv` option.",
		}
	}
	c.infraContext.Startup()
	defer c.infraContext.Shutdown()
	seelog.Debugf("invite:%s", util.MarshalObjectToString(c))
	c.apiContext, err = c.infraContext.LoadOrAuthBusinessManagement()

	if c.optCsv == "" {
		return &cmdlet.CommandError{
			Context:     cc,
			ReasonTag:   "member/invite:missing_csv",
			Description: "Missing CSV file",
		}
	}
	return c.inviteByCsv(c.optCsv)
}

func (c *CmdMemberInvite) inviteByCsv(csvFile string) error {
	f, err := os.Open(csvFile)
	if err != nil {
		seelog.Warnf("Unable to open file[%s] : error[%s]", csvFile, err)
		return err
	}
	csv := util.NewBomAwareCsvReader(f)

	for {
		cols, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			seelog.Warnf("Unable to read CSV file [%s] : error[%s]", csvFile, err)
			return err
		}
		if len(cols) < 1 {
			seelog.Warnf("Skip line: [%v]", cols)
			continue
		}
		var email, givenName, surName string
		email = cols[0]
		if len(cols) >= 2 {
			givenName = cols[1]
		}
		if len(cols) >= 3 {
			surName = cols[2]
		}

		seelog.Infof("Trying invite email[%s] givenName[%s] surName[%s]", email, givenName, surName)

		c.invite(email, givenName, surName)
	}
	return nil
}

func (c *CmdMemberInvite) invite(email, givenName, surname string) error {
	client := team.New(dropbox.Config{
		Token: c.apiContext.Token,
	})

	inv := team.NewMemberAddArg(email)
	inv.MemberGivenName = givenName
	inv.MemberSurname = surname

	arg := team.NewMembersAddArg([]*team.MemberAddArg{inv})
	res, err := client.MembersAdd(arg)
	if err != nil {
		seelog.Warnf("Unable to invite member email[%s] givenName[%s] surName[%s] : error[%s]", email, givenName, surname, err)
		return err
	}
	var added []*team.MemberAddResult
	if res.AsyncJobId != "" {
		added, err = c.waitForAsync(res.AsyncJobId, email, givenName, surname)
		if err != nil {
			seelog.Warnf("Unable to confirm result of invite member email[%s] givenName[%s] surName[%s] : error[%s]", email, givenName, surname, err)
			return err
		}
	} else {
		added = res.Complete
	}

	if len(added) < 1 {
		seelog.Warnf("Unable to invite member email[%s] givenName[%s] surName[%s] : error[%s]", email, givenName, surname, err)
		return errors.New("no one invited")
	}

	for _, m := range added {
		if m.Success != nil {
			seelog.Info("Invited: TeamMemberId[%s] Email[%s] GivenName[%s] SurName[%s]",
				m.Success.Profile.TeamMemberId,
				m.Success.Profile.Email,
				m.Success.Profile.Name.GivenName,
				m.Success.Profile.Name.Surname)
		} else {
			seelog.Warnf("Invitation failed: Email[%s] Reason[%s]", util.MarshalObjectToString(m))
		}
	}
	return nil
}

func (c *CmdMemberInvite) waitForAsync(asyncJobId, email, givenName, surname string) ([]*team.MemberAddResult, error) {
	client := team.New(dropbox.Config{
		Token: c.apiContext.Token,
	})
	for {
		time.Sleep(5 * time.Second)
		res, err := client.MembersAddJobStatusGet(async.NewPollArg(asyncJobId))
		if err != nil {
			seelog.Warnf("Unable to check status : error[%s]", err)
			return nil, err
		}
		if res.Tag == "in_progress" {
			seelog.Debugf("Process status `in_progress`: async_job_id[%s]", asyncJobId)
			continue
		}
		if res.Failed != "" {
			seelog.Warnf("Failed to add member email[%s] givenName[%s] surName[%s] : error[%s]",
				email, givenName, surname, res.Failed)
			return nil, err
		}
		if res.Complete != nil {
			return res.Complete, nil
		}
	}
}
