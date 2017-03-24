package members

import (
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/integration/business"
	"github.com/watermint/toolbox/service/report"
	"sync"
)

func reportMembers(rows chan report.ReportRow, members chan *team.TeamMemberInfo, wg *sync.WaitGroup, status string) {
	seelog.Tracef("Start reporting members: status[%s]", status)
	wg.Add(1)
	defer wg.Done()

	rows <- report.ReportHeader{
		Headers: []string{
			"Account ID",
			"Team Member ID",
			"Email",
			"Email verified?",
			"Status",
			"External ID",
			"Given Name",
			"Surname",
		},
	}

	for m := range members {
		if m == nil {
			break
		}
		seelog.Tracef("Member: %s", util.MarshalObjectToString(m))

		if status != "all" && status != m.Profile.Status.Tag {
			seelog.Debugf("Skip: status[%s] profile[%s]", status, m.Profile.Status.Tag)
			continue
		}

		rows <- report.ReportData{
			Data: []interface{}{
				m.Profile.AccountId,
				m.Profile.TeamMemberId,
				m.Profile.Email,
				m.Profile.EmailVerified,
				m.Profile.Status.Tag,
				m.Profile.ExternalId,
				m.Profile.Name.GivenName,
				m.Profile.Name.Surname,
			},
		}
	}

	rows <- nil
	seelog.Tracef("Finished report")
}

func ListMembers(token string, rows chan report.ReportRow, status string) error {
	seelog.Tracef("Start listing members for status[%s]", status)
	wg := &sync.WaitGroup{}
	members := make(chan *team.TeamMemberInfo)

	go reportMembers(rows, members, wg, status)

	err := business.LoadTeamMembers(token, members)
	if err != nil {
		return err
	}

	wg.Wait()

	return nil
}
