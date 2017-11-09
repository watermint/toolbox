package commands

import (
	"flag"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/users"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/integration/business"
	"github.com/watermint/toolbox/service/report"
	"os"
	"sync"
)

type SpaceOptions struct {
	Infra  *infra.InfraOpts
	Report *report.SimpleReportOpts
}

func parseSpaceOptions(args []string) (*SpaceOptions, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	opts := new(SpaceOptions)
	opts.Infra = infra.PrepareInfraFlags(f)
	opts.Report = report.PrepareSimpleReportFlags(f)

	f.SetOutput(os.Stderr)
	f.Parse(args)

	return opts, nil
}

func Space(args []string) error {
	opts, err := parseSpaceOptions(args)
	if err != nil {
		return err
	}

	defer opts.Infra.Shutdown()
	err = opts.Infra.Startup()
	if err != nil {
		seelog.Errorf("Unable to start operation: %s", err)
		return err
	}
	seelog.Tracef("options: %s", util.MarshalObjectToString(opts))

	token, err := opts.Infra.LoadOrAuthBusinessFile()
	if err != nil || token == "" {
		seelog.Errorf("Unable to acquire token (error: %s)", err)
		return err
	}
	rows := make(chan report.ReportRow)
	opts.Report.Write(rows)

	seelog.Trace("Start listing members")
	wg := &sync.WaitGroup{}
	members := make(chan *team.TeamMemberInfo)

	go ReportSpace(token, rows, members, wg)

	err = business.LoadTeamMembersQueue(token, members)
	if err != nil {
		return err
	}

	wg.Wait()

	opts.Report.Wait()

	return nil
}

func ReportSpace(token string, rows chan report.ReportRow, members chan *team.TeamMemberInfo, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	rows <- report.ReportHeader{
		Headers: []string{
			"Account ID",
			"Team Member ID",
			"Email",
			"Status",
			"External ID",
			"Given Name",
			"Surname",
			"Used",
		},
	}

	for m := range members {
		if m == nil {
			break
		}

		config := dropbox.Config{
			Token:      token,
			AsMemberID: m.Profile.TeamMemberId,
		}
		client := users.New(config)

		space, err := client.GetSpaceUsage()
		if err != nil {
			seelog.Warnf("Unable to load space information for user[%s] error[%s]", m.Profile.Email, err)
			continue
		}

		rows <- report.ReportData{
			Data: []interface{}{
				m.Profile.AccountId,
				m.Profile.TeamMemberId,
				m.Profile.Email,
				m.Profile.Status.Tag,
				m.Profile.ExternalId,
				m.Profile.Name.GivenName,
				m.Profile.Name.Surname,
				space.Used,
			},
		}
	}

	rows <- nil
	seelog.Tracef("Finished report")
}
