package business

import (
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
)

// Load and enqueue *team.TeamMemberInfo.
// enqueue `nil` at the end of load.
func LoadTeamMembers(token string, queue chan *team.TeamMemberInfo) error {
	seelog.Trace("Start loading team members")
	client := team.New(dropbox.Config{
		Token: token,
	})

	arg := team.NewMembersListArg()
	result, err := client.MembersList(arg)

	if err != nil {
		seelog.Errorf("Could not load team member: error[%s]", err)
		return err
	}
	seelog.Tracef("Load: %d member(s)", len(result.Members))
	seelog.Tracef("Has More: %t", result.HasMore)

	for _, m := range result.Members {
		seelog.Tracef("Enqueue: %s", m.Profile.AccountId)
		queue <- m
	}

	if !result.HasMore {
		queue <- nil
		return nil
	}

	cursor := result.Cursor
	for {
		cont := team.NewMembersListContinueArg(cursor)
		result, err = client.MembersListContinue(cont)
		if err != nil {
			seelog.Errorf("Could not load team member (continue): cursor[%s] error[%s]", cursor, err)
			queue <- nil
			return err
		}
		seelog.Tracef("Load with cursor: %d member(s)", len(result.Members))
		seelog.Tracef("Has More: %t", result.HasMore)
		for _, m := range result.Members {
			queue <- m
		}
		if !result.HasMore {
			queue <- nil
			return nil
		}
		cursor = result.Cursor
	}
}
