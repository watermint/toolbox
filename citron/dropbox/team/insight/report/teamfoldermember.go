package report

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_insight"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_insight_reports"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"gorm.io/gorm"
)

type Teamfoldermember struct {
	Database                          mo_path.FileSystemPath
	Entry                             rp_model.RowReport
	ErrDatabaseIsNotReadyForReporting app_msg.Message
}

var (
	ErrDatabaseIsNotReady = errors.New("database is not ready for reporting")
)

func (z *Teamfoldermember) Preset() {
	z.Entry.SetModel(&uc_insight_reports.TeamFolderMember{})
}

func (z *Teamfoldermember) reportTeamFolderMember(teamFolder *uc_insight.TeamFolder, path *uc_insight.SummaryTeamFolderEntry, db *gorm.DB, c app_control.Control) error {
	l := c.Log().With(esl.String("teamFolderId", teamFolder.TeamFolderId)).With(esl.String("path", path.Path))

	memberRows, err := db.Model(&uc_insight.NamespaceMember{}).Where("namespace_id = ?", path.EntryNamespaceId).Rows()
	if err != nil {
		l.Debug("Unable to retrieve model", esl.Error(err))
		return err
	}

	defer func() {
		_ = memberRows.Close()
	}()

	for memberRows.Next() {
		member := &uc_insight.NamespaceMember{}
		if err := db.ScanRows(memberRows, member); err != nil {
			l.Debug("Unable to scan member", esl.Error(err))
			return err
		}
		z.Entry.Row(&uc_insight_reports.TeamFolderMember{
			TeamFolderId:     teamFolder.TeamFolderId,
			TeamFolderName:   teamFolder.Name,
			PathDisplay:      path.PathDisplay,
			AccessType:       member.AccessType,
			IsInherited:      member.IsInherited,
			MemberType:       member.MemberType,
			SameTeam:         member.SameTeam,
			GroupId:          member.GroupId,
			GroupName:        member.GroupName,
			GroupType:        member.GroupType,
			GroupMemberCount: member.GroupMemberCount,
			InviteeEmail:     member.InviteeEmail,
			UserTeamMemberId: member.UserTeamMemberId,
			UserEmail:        member.UserEmail,
			UserDisplayName:  member.UserDisplayName,
			UserAccountId:    member.UserAccountId,
		})
	}

	return nil
}

func (z *Teamfoldermember) reportTeamFolder(teamFolder *uc_insight.TeamFolder, db *gorm.DB, c app_control.Control) error {
	l := c.Log().With(esl.String("teamFolderId", teamFolder.TeamFolderId))

	pathRows, err := db.Model(&uc_insight.SummaryTeamFolderEntry{}).Where("team_folder_id = ?", teamFolder.TeamFolderId).Rows()
	if err != nil {
		l.Debug("Unable to retrieve model", esl.Error(err))
		return err
	}
	defer func() {
		_ = pathRows.Close()
	}()

	for pathRows.Next() {
		path := &uc_insight.SummaryTeamFolderEntry{}
		if err := db.ScanRows(pathRows, path); err != nil {
			l.Debug("Unable to scan path", esl.Error(err))
			return err
		}

		if err := z.reportTeamFolderMember(teamFolder, path, db, c); err != nil {
			l.Debug("Unable to report team folder member", esl.Error(err))
			return err
		}
	}
	return nil
}

func (z *Teamfoldermember) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.Entry.Open(); err != nil {
		return err
	}

	db, err := uc_insight.DatabaseFromPath(c, z.Database.Path())
	if err != nil {
		return err
	}
	if ok, err := uc_insight.HasEntry(db); err != nil {
		return err
	} else if !ok {
		c.UI().Error(z.ErrDatabaseIsNotReadyForReporting)
		return ErrDatabaseIsNotReady
	}

	teamFolderRows, err := db.Model(&uc_insight.TeamFolder{}).Rows()
	if err != nil {
		return err
	}
	defer func() {
		_ = teamFolderRows.Close()
	}()

	for teamFolderRows.Next() {
		teamFolder := &uc_insight.TeamFolder{}
		if err := db.ScanRows(teamFolderRows, teamFolder); err != nil {
			l.Debug("Unable to scan team folder", esl.Error(err))
			return err
		}

		if err := z.reportTeamFolder(teamFolder, db, c); err != nil {
			l.Debug("Unable to report team folder", esl.Error(err))
			return err
		}
	}

	return nil
}

func (z *Teamfoldermember) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
