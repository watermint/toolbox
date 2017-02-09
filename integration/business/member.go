package business

import (
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
)

// Load and enqueue *team.TeamMemberInfo.
// enqueue `nil` at the end of load.
func LoadTeamMembers(token string, queue chan *team.TeamMemberInfo) error {
	client := team.New(dropbox.Config{
		Token: token,
	})

	arg := team.NewMembersListArg()
	result, err := client.MembersList(arg)

	if err != nil {
		return err
	}
	for _, m := range result.Members {
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
			seelog.Debugf("Could not load team member (continue): cursor[%s]", cursor)
			queue <- nil
			return err
		}
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
