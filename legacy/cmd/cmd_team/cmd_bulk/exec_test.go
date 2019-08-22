package cmd_bulk

import (
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/sequence"
	"github.com/watermint/toolbox/domain/sequence/sq_group"
	"github.com/watermint/toolbox/domain/sequence/sq_sharedfolder"
	"github.com/watermint/toolbox/domain/sequence/sq_test"
	"github.com/watermint/toolbox/domain/service"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/infra/api/api_test"
	"github.com/watermint/toolbox/infra/api/api_util"
	app2 "github.com/watermint/toolbox/legacy/app"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestCmdTeamBulkExec_Exec(t *testing.T) {
	sq_test.DoTestTeamTask(func(biz service.Business) {
		name := fmt.Sprintf("toolbox-test-%x", time.Now().Unix())
		l := biz.Log()

		seqFilePath := filepath.Join(app2.Root().JobsPath(), "test_sequence.json")
		l.Info("Sequence file path", zap.String("path", seqFilePath))

		seqFile, err := os.Create(seqFilePath)
		if err != nil {
			t.Error("Unable to create seq file", err)
			return
		}

		l.Debug("Prepare group")
		group, err := biz.Group().Create(name, sv_group.CompanyManaged())
		if err != nil {
			t.Error("unable to create group", err)
			return
		}

		defer func() {
			l.Info("Clean up group")
			err = biz.Group().Remove(group.GroupId)
			if err != nil {
				t.Error("unable to clean up", err)
			}
		}()

		folderOwner := biz.Admin()

		l.Info("Prepare shared folder")
		path := api_test.ToolboxTestSuiteFolder.ChildPath(name)
		sf, err := biz.SharedFolderAsMember(folderOwner.TeamMemberId).Create(path)
		if err != nil {
			t.Error("unable to create shared folder", err)
			return
		}

		defer func() {
			err = biz.SharedFolderAsMember(folderOwner.TeamMemberId).Remove(sf)
			if err != nil {
				if strings.HasPrefix(api_util.ErrorSummary(err), "internal_error") {
					l.Warn("Internal error. Ignored")
				} else {
					t.Error("unable to clean up", err)
				}
			}
		}()

		l.Info("Determine coworker")
		var coworker *mo_member.Member = nil
		members, err := biz.Member().List()
		if err != nil {
			t.Error("unable to list members", err)
			return
		}
		for _, m := range members {
			if m.EmailVerified && m.TeamMemberId != folderOwner.TeamMemberId {
				coworker = m
				break
			}
		}
		if coworker == nil {
			t.Error("No appropriate coworker account found")
			return
		}

		marshalTask := func(name string, peer, task interface{}) (b []byte, err2 error) {
			t, err2 := json.Marshal(task)
			if err2 != nil {
				return nil, err2
			}
			p, err2 := json.Marshal(peer)
			if err2 != nil {
				return nil, err2
			}
			m := sequence.Metadata{
				TaskName: name,
				TaskData: t,
				Peer:     p,
			}
			n, err2 := json.Marshal(m)
			if err2 != nil {
				return nil, err2
			}
			return n, nil
		}
		writeTask := func(name string, peer, task interface{}) {
			ll := l.With(zap.String("name", name), zap.Any("peer", peer), zap.Any("task", task))
			b, err := marshalTask(name, peer, task)
			if err != nil {
				ll.Error("Unable to write", zap.Error(err))
				t.Error("Unable to write", err)
			}
			if _, err := seqFile.Write(b); err != nil {
				ll.Error("Unable to write", zap.Error(err))
				t.Error("Unable to write", err)
			}
			if _, err := seqFile.Write([]byte("\n")); err != nil {
				ll.Error("Unable to write", zap.Error(err))
				t.Error("Unable to write", err)
			}
		}

		peer := sequence.PeerTeam{
			PeerName: api_test.TestPeerName,
		}
		writeTask("group/add_member", peer, sq_group.AddMember{
			GroupName:   group.GroupName,
			MemberEmail: folderOwner.Email,
		})
		writeTask("group/add_member", peer, sq_group.AddMember{
			GroupName:   group.GroupName,
			MemberEmail: coworker.Email,
		})
		writeTask("shared_folder/add_group", peer, sq_sharedfolder.AddGroup{
			SharedFolderId: sf.SharedFolderId,
			GroupName:      group.GroupName,
			AccessLevel:    mo_sharedfolder_member.AccessTypeViewer,
		})
		writeTask("shared_folder/add_user", peer, sq_sharedfolder.AddUser{
			SharedFolderId: sf.SharedFolderId,
			UserEmail:      coworker.Email,
			AccessLevel:    mo_sharedfolder_member.AccessTypeEditor,
		})
		writeTask("shared_folder/mount", peer, sq_sharedfolder.Mount{
			SharedFolderId: sf.SharedFolderId,
			UserEmail:      coworker.Email,
			MountPoint:     sf.PathLower,
		})

		if err = seqFile.Close(); err != nil {
			t.Error("Unable to close seq file", err)
			return
		}

		cmd2.CmdTest(t, NewCmdTeamBulk(), []string{"exec", "-seq-file", seqFilePath, "-retryable"})

		biz.Purge()
		{
			l.Info("Verify group members")
			groupMembers, err := biz.GroupMember(group.GroupId).List()
			if err != nil {
				t.Error("Unable to fetch members", err)
			} else {
				found := 0
				for _, m := range groupMembers {
					if m.TeamMemberId == folderOwner.TeamMemberId {
						l.Debug("Folder owner found")
						found++
					}
					if m.TeamMemberId == coworker.TeamMemberId {
						l.Debug("Coworker owner found")
						found++
					}
				}
				if found != 2 {
					t.Error("Invalid number of group member found")
				}
			}
		}

		biz.Purge()
		{
			l.Info("Verify shared folder members")
			sfMembers, err := biz.SharedFolderMemberAsMember(sf.SharedFolderId, folderOwner.TeamMemberId).List()
			if err != nil {
				t.Error("Unable to fetch members", err)
			} else {
				foundUser := false
				foundGroup := false
				for _, m := range sfMembers {
					if u, e := m.User(); e {
						if u.TeamMemberId == coworker.TeamMemberId && u.AccessType() == mo_sharedfolder_member.AccessTypeEditor {
							l.Debug("Coworker found")
							foundUser = true
						}
					}
					if g, e := m.Group(); e {
						if g.GroupId == group.GroupId && g.AccessType() == mo_sharedfolder_member.AccessTypeViewer {
							l.Debug("Group found")
							foundGroup = true
						}
					}
				}
				if !foundUser {
					t.Error("User not found")
				}
				if !foundGroup {
					t.Error("Group not found")
				}
			}
		}

		biz.Purge()
		{
			l.Info("Verify mount")
			msf, err := biz.SharedFolderAsMember(coworker.TeamMemberId).Resolve(sf.SharedFolderId)
			if err != nil {
				t.Error("Unable to fetch sf", err)
			}
			if sf.PathLower != msf.PathLower {
				l.Error("Invalid mount point", zap.String("expect", sf.PathLower), zap.String("actual", msf.PathLower))
				t.Error("Invalid mount point")
			}
		}
	})
}
