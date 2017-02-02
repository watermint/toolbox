package business

import (
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
)

func TeamMembers(token string) ([]*team.TeamMemberInfo, error) {
	client := team.New(dropbox.Config{
		Token: token,
	})

	members := make([]*team.TeamMemberInfo, 0)

	arg := team.NewMembersListArg()
	result, err := client.MembersList(arg)

	if err != nil {
		return nil, err
	}
	members = append(members, result.Members...)
	if !result.HasMore {
		return members, nil
	}

	cursor := result.Cursor
	for {
		cont := team.NewMembersListContinueArg(cursor)
		result, err = client.MembersListContinue(cont)
		if err != nil {
			seelog.Debugf("Could not load team member (continue): cursor[%s]", cursor)
			return members, err
		}
		members = append(members, result.Members...)
		if !result.HasMore {
			return members, nil
		}
		cursor = result.Cursor
	}
}
