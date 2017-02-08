package members

import (
	"encoding/csv"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/integration/business"
	"io"
	"strconv"
)

func ListMembers(token string, output io.Writer, status string) error {
	members, err := business.TeamMembers(token)
	if err != nil {
		return err
	}

	seelog.Infof("%d members loaded", len(members))

	w := csv.NewWriter(output)
	defer w.Flush()

	head := make([]string, 0)
	head = append(head, "Account ID")
	head = append(head, "Team Member ID")
	head = append(head, "Email")
	head = append(head, "Email verified?")
	head = append(head, "Status")
	head = append(head, "External ID")
	head = append(head, "Given Name")
	head = append(head, "Surname")

	w.Write(head)

	for _, m := range members {
		if status != "all" && status != m.Profile.Status.Tag {
			seelog.Debugf("Skip: status[%s] profile[%s]", status, m.Profile.Status.Tag)
			continue
		}
		line := make([]string, 0)
		line = append(line, m.Profile.AccountId)
		line = append(line, m.Profile.TeamMemberId)
		line = append(line, m.Profile.Email)
		line = append(line, strconv.FormatBool(m.Profile.EmailVerified))
		line = append(line, m.Profile.Status.Tag)
		line = append(line, m.Profile.ExternalId)
		line = append(line, m.Profile.Name.GivenName)
		line = append(line, m.Profile.Name.Surname)

		w.Write(line)
	}

	return nil
}
