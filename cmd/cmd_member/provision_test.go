package cmd_member

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/dbx_api/dbx_member"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
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

func TestCmdMemberProvisioning(t *testing.T) {
	log, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err)
		return
	}

	// 1. Clean up existing test users

	// 1.1. List members
	memberListPath, err := ioutil.TempDir("", "member_list")
	if err != nil {
		t.Error(err)
		return
	}

	cmd.CmdTest(t, NewCmdMember(), []string{"list", "-report-path", memberListPath})

	memberListJson := filepath.Join(memberListPath, "Member.json")
	memberListFile, err := os.Open(memberListJson)
	if err != nil {
		log.Warn("Quit when first test failed", zap.Error(err))
		return
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
			return
		}
		member := &dbx_profile.Member{}
		err = json.Unmarshal(line, member)
		if err != nil {
			t.Error(err)
			return
		}

		// 1.2. Remove account if a test account found:
		if strings.HasSuffix(member.Profile.Email, "@example.com") {
			cmd.CmdTest(t, NewCmdMember(), []string{
				"remove",
				"-keep-account=false",
				"-wipe-data=true",
				member.Profile.Email,
			})
		}
	}

	// 2. Provision with sync

	// 2.1. Prepare test data
	memberSyncCsvFile, err := ioutil.TempFile("", "member_sync")
	if err != nil {
		t.Error(err)
		return
	}

	memberSyncCsvContent := ""
	for i := 0; i < 3; i++ {
		memberSyncCsvContent = memberSyncCsvContent + fmt.Sprintf("toolbox%d@example.com,toolbox,%d\n", i, i)
	}
	if _, err := memberSyncCsvFile.WriteString(memberSyncCsvContent); err != nil {
		t.Error(err)
	}
	if err := memberSyncCsvFile.Close(); err != nil {
		t.Error(err)
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
		t.Error(err)
		return
	}

	syncedMemberReader := bufio.NewReader(syncedMemberFile)
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
			t.Error("Invitation failure", invite.Failure)
			return
		}
		if strings.HasSuffix(invite.Success.Profile.Email, "@example.com") {
			profile := gjson.ParseBytes(invite.Success.Profile.Profile)
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
		}
	}
	syncedMemberFile.Close()

	// 3. Update names

	// 3.1. Prepare test data
	syncUpdateCsvFile, err := ioutil.TempFile("", "sync_update")
	if err != nil {
		t.Error(err)
		return
	}

	syncUpdateCsvContent := ""
	for i := 0; i < 3; i++ {
		syncUpdateCsvContent = memberSyncCsvContent + fmt.Sprintf("toolbox%d@example.com,%d,tbx\n", i, i)
	}
	if _, err := syncUpdateCsvFile.WriteString(syncUpdateCsvContent); err != nil {
		t.Error(err)
	}
	if err := syncUpdateCsvFile.Close(); err != nil {
		t.Error(err)
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
	syncedMemberJson = filepath.Join(memberSyncResult, "Member.json")
	syncedMemberFile, err = os.Open(syncedMemberJson)
	if err != nil {
		t.Error(err)
		return
	}
	defer syncedMemberFile.Close()

	syncedMemberReader = bufio.NewReader(syncedMemberFile)
	for {
		line, _, err := syncedMemberReader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error(err)
			return
		}
		member := &dbx_profile.Member{}
		err = json.Unmarshal(line, member)
		if err != nil {
			t.Error(err)
			return
		}

		// 3.4. Verify account
		if strings.HasSuffix(member.Profile.Email, "@example.com") {
			profile := gjson.ParseBytes(member.Profile.Profile)
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
	}
	syncedMemberFile.Close()

}
