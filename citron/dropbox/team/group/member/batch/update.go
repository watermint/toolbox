package bulk

import (
	"os"

	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

type Update struct {
	rc_recipe.RemarkIrreversible
	Peer         dbx_conn.ConnScopedTeam
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
}

func (z *Update) Preset() {
	z.File.SetModel(&MemberRecord{})
	z.Peer.SetScopes(
		dbx_auth.ScopeGroupsRead,
		dbx_auth.ScopeGroupsWrite,
	)
	z.OperationLog.SetModel(&MemberRecord{}, nil)
}

func (z *Update) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	svg := sv_group.NewCached(z.Peer.Client())

	l := c.Log()
	expectedGroupAndMembers := make(map[string][]string)
	fileLoadErr := z.File.EachRow(func(m interface{}, rowIndex int) error {
		r := m.(*MemberRecord)
		ll := l.With(esl.Any("record", r))
		_, err := svg.ResolveByName(r.GroupName)
		if err != nil {
			ll.Debug("Unable to retrieve group", esl.Error(err))
			return err
		}
		if members, ok := expectedGroupAndMembers[r.GroupName]; ok {
			members = append(members, r.MemberEmail)
			expectedGroupAndMembers[r.GroupName] = members
		} else {
			members = make([]string, 0)
			members = append(members, r.MemberEmail)
			expectedGroupAndMembers[r.GroupName] = members
		}
		return nil
	})
	if fileLoadErr != nil {
		l.Debug("There were an error in the data file", esl.Error(fileLoadErr))
		return fileLoadErr
	}

	currentGroupAndMembers := make(map[string][]string)
	retrieveMembers := func(name string) error {
		ll := l.With(esl.String("groupName", name))
		g, err := svg.ResolveByName(name)
		if err != nil {
			ll.Debug("Unable to resolve the group", esl.Error(err))
			return err
		}

		sgm := sv_group_member.New(z.Peer.Client(), g)
		members, err := sgm.List()
		if err != nil {
			ll.Debug("Unable to retrieve group members", esl.Error(err))
			return err
		}
		currentGroupAndMembers[name] = mo_group_member.GroupMemberEmails(members)
		return nil
	}

	queueIdScanMember := "member_scan"
	queueIdMemberAdd := "member_add"
	queueIdMemberDelete := "member_delete"

	var lastErr error
	c.Sequence().DoThen(func(s eq_sequence.Stage) {
		s.Define(queueIdScanMember, retrieveMembers)
		q := s.Get(queueIdScanMember)
		for n := range expectedGroupAndMembers {
			q.Enqueue(n)
		}

	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	})).Do(func(s eq_sequence.Stage) {
		if lastErr != nil {
			l.Debug("Abort due to an error in prior sequence", esl.Error(lastErr))
			return
		}

		s.Define(queueIdMemberAdd, memberAdd, svg, c, z.Peer.Client(), z.OperationLog)
		s.Define(queueIdMemberDelete, memberDelete, svg, c, z.Peer.Client(), z.OperationLog)
		qa := s.Get(queueIdMemberAdd)
		qd := s.Get(queueIdMemberDelete)

		for groupName, expectedMembers := range expectedGroupAndMembers {
			currentMembers := currentGroupAndMembers[groupName]

			// Create maps for efficient lookup
			expectedMap := make(map[string]bool)
			currentMap := make(map[string]bool)
			for _, m := range expectedMembers {
				expectedMap[m] = true
			}
			for _, m := range currentMembers {
				currentMap[m] = true
			}

			// Find members to add (in expected but not in current)
			var membersToAdd []string
			for _, m := range expectedMembers {
				if !currentMap[m] {
					membersToAdd = append(membersToAdd, m)
				}
			}
			l.Debug("Add members", esl.String("group", groupName), esl.Strings("members", membersToAdd))
			for _, m := range membersToAdd {
				qa.Enqueue(&MemberRecord{
					GroupName:   groupName,
					MemberEmail: m,
				})
			}

			// Find members to delete (in current but not in expected)
			var membersToDelete []string
			for _, m := range currentMembers {
				if !expectedMap[m] {
					membersToDelete = append(membersToDelete, m)
				}
			}
			l.Debug("Delete members", esl.String("group", groupName), esl.Strings("members", membersToDelete))
			for _, m := range membersToDelete {
				qd.Enqueue(&MemberRecord{
					GroupName:   groupName,
					MemberEmail: m,
				})
			}
		}
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	return lastErr
}

func (z *Update) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("update", "Sales,taro@example.com\nSales,hanako@example.com\n")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()
	return rc_exec.ExecMock(c, &Update{}, func(r rc_recipe.Recipe) {
		m := r.(*Update)
		m.File.SetFilePath(f)
	})
}
