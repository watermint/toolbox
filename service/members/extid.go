package members

import (
	"errors"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
	"github.com/watermint/toolbox/integration/business"
	"sync"
)

func AssignPseudoExtId(token string, member *team.TeamMemberInfo, dryRun bool) error {
	newExternalId := fmt.Sprintf("Email %s", member.Profile.Email)

	seelog.Tracef("Trying to assign new External Id: (Old Ext Id: [%s]) -> (New Ext Id: [%s])", member.Profile.ExternalId, newExternalId)

	if dryRun {
		seelog.Infof("DRYRUN: Assign new External Id: [%s] -> [%s]", member.Profile.ExternalId, newExternalId)
		return nil
	}

	config := dropbox.Config{
		Token: token,
	}
	client := team.New(config)
	s := &team.UserSelectorArg{}
	s.Tag = "email" // Workaround
	s.Email = member.Profile.Email

	a := team.NewMembersSetProfileArg(s)
	a.NewExternalId = newExternalId

	m, err := client.MembersSetProfile(a)
	if err != nil {
		seelog.Warnf("Unable to update member external Id : email[%s] error[%s]", member.Profile.Email, err)
		return err
	}
	seelog.Infof("Assigned new External Id: [%s] -> [%s]", member.Profile.ExternalId, m.Profile.ExternalId)

	return nil
}

func AssignPseudoExtIdByEmail(token string, email string, dryRun bool) error {
	seelog.Tracef("Trying to assign pseudo ext id by email[%s]", email)
	config := dropbox.Config{
		Token: token,
	}
	client := team.New(config)
	s := &team.UserSelectorArg{
		Email: email,
	}
	u := make([]*team.UserSelectorArg, 0)
	u = append(u, s)
	a := team.NewMembersGetInfoArgs(u)
	m, err := client.MembersGetInfo(a)
	if err != nil {
		seelog.Warnf("Unable to get member info : email[%s] error[%s]", email, err)
		return err
	}
	if len(m) != 1 {
		seelog.Warnf("Unable to get member info : email[%s] error(no result)", email)
		return errors.New("No member found for email")
	}

	return AssignPseudoExtId(token, m[0].MemberInfo, dryRun)
}

type assignPseudoExtIdForTeamMember struct {
	token  string
	dryRun bool
}

func (a *assignPseudoExtIdForTeamMember) LoadMember(member *team.TeamMemberInfo) error {
	seelog.Tracef("Assign pseudo external id: Email[%s] OrigExtId[%s]", member.Profile.Email, member.Profile.ExternalId)
	return AssignPseudoExtId(a.token, member, a.dryRun)
}

func (a *assignPseudoExtIdForTeamMember) Finished() {
	// NOP
}

func AssignPseudoExtIdForTeam(token string, dryRun bool) error {
	a := &assignPseudoExtIdForTeamMember{
		token:  token,
		dryRun: dryRun,
	}

	err := business.LoadTeamMembers(token, a)
	if err != nil {
		seelog.Warnf("Unable to load members : error[%s]", err)
	}

	return nil
}

func ShowExtIdByEmail(token string, email string) error {
	seelog.Tracef("Trying to show ext id by email[%s]", email)
	config := dropbox.Config{
		Token: token,
	}
	client := team.New(config)
	s := &team.UserSelectorArg{
		Email: email,
	}
	a := team.NewMembersGetInfoArgs([]*team.UserSelectorArg{s})
	m, err := client.MembersGetInfo(a)
	if err != nil {
		seelog.Warnf("Unable to get member info : email[%s] error[%s]", email, err)
		return err
	}
	if len(m) != 1 {
		seelog.Warnf("Unable to get member info : email[%s] error(no result)", email)
		return errors.New("No member found for email")
	}
	seelog.Infof("Email[%s] ExtId[%s]", m[0].MemberInfo.Profile.Email, m[0].MemberInfo.Profile.ExternalId)
	return nil
}

func showExtIdForTeamMember(queue chan *team.TeamMemberInfo, wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	for member := range queue {
		if member == nil {
			seelog.Trace("Reached to the end")
			break
		}

		seelog.Infof("Email[%s] ExtId[%s]", member.Profile.Email, member.Profile.ExternalId)
	}

	return nil
}

func ShowExtIdForTeam(token string) error {
	wg := &sync.WaitGroup{}
	queue := make(chan *team.TeamMemberInfo)

	go showExtIdForTeamMember(queue, wg)

	err := business.LoadTeamMembersQueue(token, queue)
	if err != nil {
		seelog.Warnf("Unable to load members : error[%s]", err)
		return err
	}

	wg.Wait()
	return nil
}
