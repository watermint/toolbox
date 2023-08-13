package uc_insight

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_team"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"gorm.io/gorm"
	"reflect"
)

type IndividualScanner interface {
	// ScanCurrentUser scans current user files and sharing information
	ScanCurrentUser() (err error)
}

type TeamScanner interface {
	// Scan scans all team information
	Scan() (err error)

	Summarize() (err error)

	RetryErrors() (err error)
}

func newDatabase(ctl app_control.Control, path string) (*gorm.DB, error) {
	l := ctl.Log().With(esl.String("path", path))
	db, err := ctl.NewOrm(path)
	if err != nil {
		return nil, err
	}

	tables := []interface{}{
		&FileMember{},
		&GroupMember{},
		&Group{},
		&Member{},
		&Mount{},
		&NamespaceDetail{},
		&NamespaceEntryError{},
		&NamespaceEntry{},
		&NamespaceMember{},
		&Namespace{},
		&ReceivedFile{},
		&SharedLink{},
		&SummaryEntry{},
		&SummaryFolderAndNamespace{},
		&SummaryFolderImmediateCount{},
		&SummaryFolderPath{},
		&SummaryFolderRecursive{},
		&SummaryNamespace{},
		&TeamFolder{},
	}

	for _, t := range tables {
		tableName := reflect.ValueOf(t).Elem().Type().Name()
		l.Debug("Migrating", esl.String("table", tableName))
		if err = db.AutoMigrate(t); err != nil {
			l.Debug("Unable to migrate", esl.Error(err), esl.String("table", tableName))
			return nil, err
		}
	}

	return db, nil
}
func NewTeamScanner(ctl app_control.Control, client dbx_client.Client, path string) (TeamScanner, error) {
	db, err := newDatabase(ctl, path)
	if err != nil {
		return nil, err
	}
	return &tsImpl{
		ctl:    ctl,
		client: client,
		db:     db,
	}, nil
}

const (
	teamScanQueueFileMember      = "scan_file_member"
	teamScanQueueGroup           = "scan_group"
	teamScanQueueGroupMember     = "scan_group_member"
	teamScanQueueMember          = "scan_member"
	teamScanQueueMount           = "scan_mount"
	teamScanQueueNamespace       = "scan_team_namespace"
	teamScanQueueNamespaceDetail = "scan_namespace"
	teamScanQueueNamespaceEntry  = "scan_folder"
	teamScanQueueNamespaceMember = "scan_namespace_member"
	teamScanQueueReceivedFile    = "scan_received_file"
	teamScanQueueSharedLink      = "scan_shared_link"
	teamScanQueueTeamFolder      = "scan_team_folder"
	teamSummarizeEntry           = "resolve_entry"
	teamSummarizeFolderImmediate = "resolve_folder_immediate"
	teamSummarizeFolderPath      = "resolve_folder_path"
	teamSummarizeFolderRecursive = "resolve_folder_recursive"
	teamSummarizeNamespace       = "resolve_namespace"
)

type NamespaceEntryParam struct {
	NamespaceId string `json:"namespaceId" path:"namespace_id"`
	FolderId    string `json:"folderId" path:"folder_id"`
	IsRetry     bool   `json:"isRetry" path:"is_retry"`
}

type FileMemberParam struct {
	NamespaceId string `json:"namespaceId" path:"namespace_id"`
	FileId      string `json:"fileId" path:"file_id"`
}

type tsImpl struct {
	ctl    app_control.Control
	client dbx_client.Client
	db     *gorm.DB
}

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
	qMount.Enqueue(member.TeamMemberId)
	qReceivedFile := stage.Get(teamScanQueueReceivedFile)
	qReceivedFile.Enqueue(member.TeamMemberId)
	qSharedLink := stage.Get(teamScanQueueSharedLink)
	qSharedLink.Enqueue(member.TeamMemberId)

	return nil
}

func (z tsImpl) defineScanQueues(s eq_sequence.Stage, admin *mo_profile.Profile, team *mo_team.Info) {
	s.Define(teamScanQueueFileMember, z.scanFileMember, s, admin)
	s.Define(teamScanQueueGroup, z.scanGroup, s, admin)
	s.Define(teamScanQueueGroupMember, z.scanGroupMember, s, admin)
	s.Define(teamScanQueueMember, z.scanMembers, s, admin)
	s.Define(teamScanQueueMount, z.scanMount, s, admin, team)
	s.Define(teamScanQueueNamespace, z.scanNamespaces, s, admin)
	s.Define(teamScanQueueNamespaceDetail, z.scanNamespaceDetail, s, admin, team)
	s.Define(teamScanQueueNamespaceEntry, z.scanNamespaceEntry, s, admin)
	s.Define(teamScanQueueNamespaceMember, z.scanNamespaceMember, s, admin)
	s.Define(teamScanQueueReceivedFile, z.scanReceivedFile, s, admin)
	s.Define(teamScanQueueSharedLink, z.scanSharedLink, s, admin)
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

	var lastErr error
	z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
		z.defineScanQueues(s, admin, team)

		qMember := s.Get(teamScanQueueMember)
		qMember.Enqueue("")
		qNamespace := s.Get(teamScanQueueNamespace)
		qNamespace.Enqueue("")
		qGroup := s.Get(teamScanQueueGroup)
		qGroup.Enqueue("")
		qTeamFolder := s.Get(teamScanQueueTeamFolder)
		qTeamFolder.Enqueue("")

	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	db, err := z.db.DB()
	if err != nil {
		return err
	}
	_ = db.Close()

	return lastErr
}

func (z tsImpl) defineSummarizeQueues(s eq_sequence.Stage) {
	s.Define(teamSummarizeFolderImmediate, z.summarizeFolderImmediateCount)
	s.Define(teamSummarizeFolderPath, z.summarizeFolderPaths)
	s.Define(teamSummarizeFolderRecursive, z.summarizeFolderRecursive)
	s.Define(teamSummarizeNamespace, z.summarizeNamespace)
	s.Define(teamSummarizeEntry, z.summarizeEntry)
}

func (z tsImpl) Summarize() error {
	l := z.ctl.Log()
	var lastErr error

	summaryTables := []interface{}{
		&SummaryEntry{},
		&SummaryFolderAndNamespace{},
		&SummaryFolderImmediateCount{},
		&SummaryFolderPath{},
		&SummaryFolderRecursive{},
		&SummaryNamespace{},
	}
	for _, st := range summaryTables {
		z.db.Delete(st)
	}

	l.Debug("Stage 1")
	{
		namespaceModel := &Namespace{}
		namespaceRows, err := z.db.Model(namespaceModel).Distinct("namespace_id").Select("namespace_id").Rows()
		if err != nil {
			l.Debug("Unable to get namespace rows", esl.Error(err))
			return err
		}
		defer func() {
			_ = namespaceRows.Close()
		}()

		z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
			z.defineSummarizeQueues(s)

			qNamespace := s.Get(teamSummarizeNamespace)
			for namespaceRows.Next() {
				if err := z.db.ScanRows(namespaceRows, namespaceModel); err != nil {
					l.Debug("Unable to scan namespace row", esl.Error(err))
					lastErr = err
					return
				}
				qNamespace.Enqueue(namespaceModel.NamespaceId)
			}
		}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
			lastErr = err
		}))

		if lastErr != nil {
			l.Debug("Failure on processing namespace rows", esl.Error(lastErr))
			return lastErr
		}
	}

	l.Debug("Stage 2")
	{
		folderEntry := &NamespaceEntry{}
		folderRows, err := z.db.Model(folderEntry).Distinct("file_id").Where("entry_type = ?", "folder").Rows()
		if err != nil {
			l.Debug("Unable to get folder rows", esl.Error(err))
			return err
		}
		defer func() {
			_ = folderRows.Close()
		}()

		z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
			z.defineSummarizeQueues(s)

			qFolderPath := s.Get(teamSummarizeFolderPath)
			qFolderImmediate := s.Get(teamSummarizeFolderImmediate)
			for folderRows.Next() {
				if err := z.db.ScanRows(folderRows, folderEntry); err != nil {
					l.Debug("cannot scan row", esl.Error(err))
					return
				}
				qFolderPath.Enqueue(folderEntry.FileId)
				qFolderImmediate.Enqueue(folderEntry.FileId)
			}

		}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
			lastErr = err
		}))

		if lastErr != nil {
			l.Debug("Failure on processing folder rows", esl.Error(lastErr))
			return lastErr
		}
	}

	l.Debug("Stage 3")
	{
		folderEntry := &NamespaceEntry{}
		folderRows, err := z.db.Model(folderEntry).Distinct("file_id").Where("entry_type = ?", "folder").Rows()
		if err != nil {
			l.Debug("Unable to get folder rows", esl.Error(err))
			return err
		}
		defer func() {
			_ = folderRows.Close()
		}()

		z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
			z.defineSummarizeQueues(s)

			qFolderRecursive := s.Get(teamSummarizeFolderRecursive)
			for folderRows.Next() {
				if err := z.db.ScanRows(folderRows, folderEntry); err != nil {
					l.Debug("cannot scan row", esl.Error(err))
					return
				}
				qFolderRecursive.Enqueue(folderEntry.FileId)
			}

		}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
			lastErr = err
		}))

		if lastErr != nil {
			l.Debug("Failure on processing folder rows", esl.Error(lastErr))
			return lastErr
		}
	}

	db, err := z.db.DB()
	if err != nil {
		l.Debug("Unable to get DB", esl.Error(err))
		return err
	}
	_ = db.Close()

	return nil
}
