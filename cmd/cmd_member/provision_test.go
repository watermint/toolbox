package cmd_member

import (
	"bufio"
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/app/app_util"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_member"
	"github.com/watermint/toolbox/model/dbx_profile"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

type ProvisionTest struct {
	Logger *zap.Logger
}

func (z *ProvisionTest) PrepareCsv(numAccounts int, pattern string) (*os.File, error) {
	f, err := ioutil.TempFile("", "provisioning_data")
	if err != nil {
		z.Logger.Error("unable to create temp file", zap.Error(err))
		return nil, err
	}

	memberSyncCsvContent := ""
	for i := 1; i <= numAccounts; i++ {
		line, _ := app_util.CompileTemplate(pattern, struct {
			Index int
		}{
			Index: i,
		})
		memberSyncCsvContent = memberSyncCsvContent + line + "\n"
	}
	if _, err := f.WriteString(memberSyncCsvContent); err != nil {
		z.Logger.Error("unable to add line(s) into the temp file", zap.Error(err))
		return nil, err
	}
	if err := f.Close(); err != nil {
		z.Logger.Error("unable to close file", zap.Error(err))
		return nil, err
	}
	return f, nil
}

func (z *ProvisionTest) ListTestAccounts(t *testing.T) ([]*dbx_profile.Member, error) {
	memberListPath, err := ioutil.TempDir("", "member_list")
	if err != nil {
		t.Error(err)
		return nil, err
	}

	members := make([]*dbx_profile.Member, 0)

	cmd.CmdTest(t, NewCmdMember(), []string{"list", "-report-path", memberListPath})

	memberListJson := filepath.Join(memberListPath, "Member.json")
	memberListFile, err := os.Open(memberListJson)
	if err != nil {
		z.Logger.Info("unable to open Member.json", zap.Error(err))
		return members, nil
	}
	defer memberListFile.Close()

	memberListReader := bufio.NewReader(memberListFile)
	for {
		line, _, err := memberListReader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error(err)
			return nil, err
		}
		member := &dbx_profile.Member{}
		err = json.Unmarshal(line, member)
		if err != nil {
			t.Error(err)
			return nil, err
		}

		if strings.HasSuffix(member.Profile.Email, "@example.com") {
			z.Logger.Debug("Test member found", zap.String("email", member.Profile.Email))
			members = append(members, member)
		}
	}
	return members, nil
}

func (z *ProvisionTest) CleanTestAccounts(t *testing.T) error {
	members, err := z.ListTestAccounts(t)
	if err != nil {
		return err
	}

	for _, m := range members {
		cmd.CmdTest(t, NewCmdMember(), []string{
			"remove",
			"-keep-account=false",
			"-wipe-data=true",
			m.Profile.Email,
		})
	}
	return nil
}

func (z *ProvisionTest) VerifyInviteReport(reportFile *os.File, t *testing.T, v func(*dbx_member.InviteReport, *testing.T)) {
	syncedMemberReader := bufio.NewReader(reportFile)
	for {
		line, _, err := syncedMemberReader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error(err)
			return
		}
		invite := &dbx_member.InviteReport{}
		err = json.Unmarshal(line, invite)
		if err != nil {
			t.Error(err)
			return
		}

		// 2.3. Verify account
		if invite.Result != "success" {
			if invite.Failure.Reason == "team_license_limit" {
				z.Logger.Info("Test failed due to `team_license_limit`")
				continue
			} else {
				t.Error("Invitation failure", invite.Failure)
				return
			}
		}

		if strings.HasSuffix(invite.Success.Profile.Email, "@example.com") {
			v(invite, t)
		}
	}
	if err := reportFile.Close(); err != nil {
		t.Error("unable to close report file", err)
	}
}

func TestCmdMemberInvite(t *testing.T) {
	log, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err)
		return
	}
	pt := ProvisionTest{
		Logger: log,
	}

	// 1. Clean up existing test users
	if err := pt.CleanTestAccounts(t); err != nil {
		t.Error("unable to clean up existing test accounts")
		return
	}

	// 2. Invite by csv
	csvFile, err := pt.PrepareCsv(
		3,
		"toolbox{{.Index}}@example.com,invite,{{.Index}}",
	)
	if err != nil {
		t.Error(err)
		return
	}
	memberInviteResult, err := ioutil.TempDir("", "member_sync")
	if err != nil {
		t.Error(err)
		return
	}

	cmd.CmdTestWithTimeout(t, NewCmdMember(), []string{
		"invite",
		"-report-path",
		memberInviteResult,
		"-csv",
		csvFile.Name(),
	}, time.Duration(300)*time.Second)

	// 3. verify
	// 3.1. verify local result
	invitedMemberJson := filepath.Join(memberInviteResult, "InviteReport.json")
	invitedMemberFile, err := os.Open(invitedMemberJson)
	if err != nil {
		t.Error(err)
		return
	}
	pt.VerifyInviteReport(invitedMemberFile, t, func(report *dbx_member.InviteReport, t *testing.T) {
		profile := gjson.ParseBytes(report.Success.Profile.Profile)
		givenName := profile.Get("name.given_name").String()
		surname := profile.Get("name.surname").String()

		log.Info(
			"Test account",
			zap.String("profile", profile.Raw),
		)

		if !strings.HasPrefix(givenName, "invite") {
			t.Errorf("Unexpected test account given_name[%s]", givenName)
		}
		if _, err := strconv.Atoi(surname); err != nil {
			t.Errorf("Unexpected test account surname[%s]", surname)
		}
	})

	// 3.2. verify remote result
	syncedMembers, err := pt.ListTestAccounts(t)
	for _, m := range syncedMembers {
		profile := gjson.ParseBytes(m.Profile.Profile)
		givenName := profile.Get("name.given_name").String()
		surname := profile.Get("name.surname").String()

		log.Info(
			"Test account",
			zap.String("profile", profile.Raw),
		)

		if !strings.HasPrefix(givenName, "invite") {
			t.Errorf("Unexpected test account surname[%s]", surname)
		}
		if _, err := strconv.Atoi(surname); err != nil {
			t.Errorf("Unexpected test account givenname[%s]", givenName)
		}
	}

	// 4. Clean up existing test users
	if err := pt.CleanTestAccounts(t); err != nil {
		t.Error("unable to clean up existing test accounts")
		return
	}
}

func TestCmdMemberSync(t *testing.T) {
	log, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err)
		return
	}
	pt := ProvisionTest{
		Logger: log,
	}

	// 1. Clean up existing test users
	if err := pt.CleanTestAccounts(t); err != nil {
		t.Error("unable to clean up existing test accounts")
		return
	}

	// 2. Provision with sync
	memberSyncCsvFile, err := pt.PrepareCsv(
		3,
		"toolbox{{.Index}}@example.com,toolbox,{{.Index}}",
	)
	if err != nil {
		t.Error(err)
		return
	}

	// 2.2. Sync
	memberSyncResult, err := ioutil.TempDir("", "member_sync")
	if err != nil {
		t.Error(err)
		return
	}

	cmd.CmdTestWithTimeout(t, NewCmdMember(), []string{
		"sync",
		"-report-path",
		memberSyncResult,
		"-csv",
		memberSyncCsvFile.Name(),
	}, time.Duration(300)*time.Second)

	// 2.3. Sync result validation
	syncedMemberJson := filepath.Join(memberSyncResult, "InviteReport.json")
	syncedMemberFile, err := os.Open(syncedMemberJson)
	if err != nil {
		log.Info("Skip validation, because of the report wasn't generated", zap.Error(err))
		return
	}
	pt.VerifyInviteReport(syncedMemberFile, t, func(report *dbx_member.InviteReport, t *testing.T) {
		profile := gjson.ParseBytes(report.Success.Profile.Profile)
		givenName := profile.Get("name.given_name").String()
		surname := profile.Get("name.surname").String()

		log.Info(
			"Test account",
			zap.String("profile", profile.Raw),
		)

		if !strings.HasPrefix(givenName, "toolbox") {
			t.Errorf("Unexpected test account given_name[%s]", givenName)
		}
		if _, err := strconv.Atoi(surname); err != nil {
			t.Errorf("Unexpected test account surname[%s]", surname)
		}
	})

	// 3. Update names

	// 3.1. Prepare test data
	syncUpdateCsvFile, err := pt.PrepareCsv(
		3,
		"toolbox{{.Index}}@example.com,{{.Index}},tbx",
	)
	if err != nil {
		t.Error(err)
		return
	}

	// 3.2. Sync
	memberSyncResult, err = ioutil.TempDir("", "member_sync")
	if err != nil {
		t.Error(err)
		return
	}

	cmd.CmdTestWithTimeout(t, NewCmdMember(), []string{
		"sync",
		"-report-path",
		memberSyncResult,
		"-csv",
		syncUpdateCsvFile.Name(),
	}, time.Duration(300)*time.Second)

	// 3.3. Sync result validation
	syncedMembers, err := pt.ListTestAccounts(t)
	for _, m := range syncedMembers {
		profile := gjson.ParseBytes(m.Profile.Profile)
		givenName := profile.Get("name.given_name").String()
		surname := profile.Get("name.surname").String()

		log.Info(
			"Test account",
			zap.String("profile", profile.Raw),
		)

		if !strings.HasPrefix(surname, "tbx") {
			t.Errorf("Unexpected test account surname[%s]", surname)
		}
		if _, err := strconv.Atoi(givenName); err != nil {
			t.Errorf("Unexpected test account givenname[%s]", givenName)
		}
	}

	// 4. Clean up existing test users
	if err := pt.CleanTestAccounts(t); err != nil {
		t.Error("unable to clean up existing test accounts")
		return
	}
}
