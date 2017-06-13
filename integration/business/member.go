package business

import (
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
)

type TeamMemberLoader interface {
	LoadMember(*team.TeamMemberInfo) error
	Finished()
}

type TeamMemberQueueLoader struct {
	Queue chan *team.TeamMemberInfo
}

func (m *TeamMemberQueueLoader) LoadMember(member *team.TeamMemberInfo) error {
	m.Queue <- member
	return nil
}

func (m *TeamMemberQueueLoader) Finished() {
	m.Queue <- nil
}

// Load and enqueue *team.TeamMemberInfo.
// enqueue `nil` at the end of load.
func LoadTeamMembersQueue(token string, queue chan *team.TeamMemberInfo) error {
	ql := &TeamMemberQueueLoader{
		Queue: queue,
	}
	return LoadTeamMembers(token, ql)
}

func LoadTeamMembers(token string, loader TeamMemberLoader) error {
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
		err = loader.LoadMember(m)
		if err != nil {
			seelog.Tracef("Discontinue load : member[%s] error[%s]", m.Profile.Email, err)
			return err
		}
	}

	if !result.HasMore {
		loader.Finished()
		return nil
	}

	cursor := result.Cursor
	for {
		cont := team.NewMembersListContinueArg(cursor)
		result, err = client.MembersListContinue(cont)
		if err != nil {
			seelog.Errorf("Could not load team member (continue): cursor[%s] error[%s]", cursor, err)
			loader.Finished()
			return err
		}
		seelog.Tracef("Load with cursor: %d member(s)", len(result.Members))
		seelog.Tracef("Has More: %t", result.HasMore)
		for _, m := range result.Members {
			err = loader.LoadMember(m)
			if err != nil {
				seelog.Tracef("Discontinue load : member[%s] error[%s]", m.Profile.Email, err)
				return err
			}
		}
		if !result.HasMore {
			loader.Finished()
			return nil
		}
		cursor = result.Cursor
	}
}
