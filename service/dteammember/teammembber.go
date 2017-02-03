package dteammember

import (
	"encoding/csv"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
	"io"
	"os"
)

func DetachUser(token, userEmail string, dryRun bool) error {
	seelog.Infof("Detach user[%s]", userEmail)

	config := dropbox.Config{
		Token: token,
	}
	client := team.New(config)

	sel := &team.UserSelectorArg{
		Email: userEmail,
	}
	sel.Tag = team.UserSelectorArgEmail

	getMembers := make([]*team.UserSelectorArg, 0)
	getMembers = append(getMembers, sel)
	get := team.NewMembersGetInfoArgs(getMembers)
	members, err := client.MembersGetInfo(get)
	if err != nil {
		seelog.Warnf("Unable to find user information[%s] error[%s]", userEmail, err)
		return err
	}
	member := members[0]
	seelog.Infof("User AccountId[%s] TeamMemberId[%s] Name[%s] Email[%s]", member.MemberInfo.Profile.AccountId, member.MemberInfo.Profile.TeamMemberId, member.MemberInfo.Profile.Name.DisplayName, member.MemberInfo.Profile.Email)

	if dryRun {
		seelog.Info("Skip detach operation (Dry run)")
		return nil
	}

	arg := team.NewMembersRemoveArg(sel)
	arg.WipeData = false
	arg.KeepAccount = true

	_, err = client.MembersRemove(arg)
	if err != nil {
		seelog.Warnf("Unable to detach user[%s] error[%s]", userEmail, err)
		return err
	}
	return nil
}

func DetachUserByList(token, userEmailFile string, dryRun bool) error {
	userFile, err := os.Open(userEmailFile)
	if err != nil {
		seelog.Warnf("Unable to load file [%s] error[%s]", userEmailFile, err)
		return err
	}
	defer userFile.Close()
	reader := csv.NewReader(userFile)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			seelog.Warnf("Unabble to read record. error[%s]", err)
			return err
		}
		DetachUser(token, line[0], dryRun)
	}
}
