package uc_insight

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_team"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
)

func (z tsImpl) saveIfExternalGroup(member mo_sharedfolder_member.Member) {
	g, err := NewGroupFromMember(member)
	if err != nil {
		// not a group
		return
	}
	if mo_sharedfolder_member.IsSameTeam(member.SameTeam()) {
		// not an external group
		return
	}
	z.db.Create(g)
}

func (z tsImpl) dispatchMember(member *mo_member.Member, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	if member.Status == "removed" {
		return nil
	}
	qMount := stage.Get(teamScanQueueMount)
	qMount.Enqueue(&MountParam{
		TeamMemberId: member.TeamMemberId,
	})
	qReceivedFile := stage.Get(teamScanQueueReceivedFile)
	qReceivedFile.Enqueue(&ReceivedFileParam{
		TeamMemberId: member.TeamMemberId,
	})
	qSharedLink := stage.Get(teamScanQueueSharedLink)
	qSharedLink.Enqueue(&SharedLinkParam{
		TeamMemberId: member.TeamMemberId,
	})

	return nil
}

func (z tsImpl) defineScanQueues(s eq_sequence.Stage, admin *mo_profile.Profile, team *mo_team.Info) {
	s.Define(teamScanQueueFileMember, z.scanFileMember, s, admin)
	s.Define(teamScanQueueGroup, z.scanGroup, s, admin)
	s.Define(teamScanQueueGroupMember, z.scanGroupMember, s, admin)
	s.Define(teamScanQueueMember, z.scanMembers, s, admin)
	s.Define(teamScanQueueMount, z.scanMount, s, admin, team, z.opts.BaseNamespace)
	s.Define(teamScanQueueNamespace, z.scanNamespaces, s, admin)
	s.Define(teamScanQueueNamespaceDetail, z.scanNamespaceDetail, s, admin, team)
	s.Define(teamScanQueueNamespaceEntry, z.scanNamespaceEntry, s, admin)
	s.Define(teamScanQueueNamespaceMember, z.scanNamespaceMember, s, admin)
	s.Define(teamScanQueueReceivedFile, z.scanReceivedFile, s, admin, z.opts.BaseNamespace)
	s.Define(teamScanQueueSharedLink, z.scanSharedLink, s, admin, z.opts.BaseNamespace)
	s.Define(teamScanQueueTeamFolder, z.scanTeamFolder, s, admin)
}

func (z tsImpl) Scan() (err error) {
	l := z.ctl.Log()
	admin, err := sv_profile.NewTeam(z.client).Admin()
	if err != nil {
		l.Debug("Unable to retrieve admin profile", esl.Error(err))
		return err
	}
	team, err := sv_team.New(z.client).Info()
	if err != nil {
		l.Debug("Unable to retrieve team info", esl.Error(err))
		return err
	}

	z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
		z.defineScanQueues(s, admin, team)

		qMember := s.Get(teamScanQueueMember)
		qMember.Enqueue(&MemberParam{})
		qNamespace := s.Get(teamScanQueueNamespace)
		qNamespace.Enqueue(&NamespaceParam{
			ScanMemberFolders: z.opts.ScanMemberFolders,
		})
		qGroup := s.Get(teamScanQueueGroup)
		qGroup.Enqueue(&GroupParam{})
		qTeamFolder := s.Get(teamScanQueueTeamFolder)
		qTeamFolder.Enqueue(&TeamFolderParam{})
	})

	for i := 0; i < z.opts.MaxRetries; i++ {
		ll := l.With(esl.Int("retry", i+1))
		numErrors, err := z.hasErrors()
		if err != nil {
			ll.Debug("Unable to check errors", esl.Error(err))
			return err
		}
		if numErrors > 0 {
			ll.Debug("Checking errors", esl.Int64("errorRecords", numErrors))
			if err := z.RetryErrors(); err != nil {
				ll.Debug("Error on retry", esl.Error(err))
			}
		}
	}

	numErrs, err := z.hasErrors()
	if err != nil {
		l.Debug("Unable to check errors", esl.Error(err))
		return err
	}
	if numErrs > 0 {
		l.Debug("There are errors", esl.Int64("errorRecords", numErrs))

		return errors.New(fmt.Sprintf("There are %d errors", numErrs))
	}
	return nil
}
