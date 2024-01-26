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
	"os"
	"path/filepath"
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

func newDatabase(ctl app_control.Control, path string) (adb, sdb *gorm.DB, err error) {
	l := ctl.Log().With(esl.String("path", path))
	if err := os.MkdirAll(path, 0700); err != nil {
		l.Debug("Unable to create directory", esl.Error(err))
		return nil, nil, err
	}

	adbPath := filepath.Join(path, "api.db")
	adb, err = ctl.NewOrm(adbPath)
	if err != nil {
		l.Debug("Unable to open database", esl.Error(err), esl.String("path", adbPath))
		return nil, nil, err
	}
	sdbPath := filepath.Join(path, "summary.db")
	sdb, err = ctl.NewOrm(sdbPath)
	if err != nil {
		l.Debug("Unable to open database", esl.Error(err), esl.String("path", sdbPath))
		return nil, nil, err
	}

	adbTables := []interface{}{
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
		&TeamFolder{},
	}
	sdbTables := []interface{}{
		&SummaryEntry{},
		&SummaryFolderAndNamespace{},
		&SummaryFolderImmediateCount{},
		&SummaryFolderPath{},
		&SummaryFolderRecursive{},
		&SummaryNamespace{},
		&SummaryTeamFolderEntry{},
	}

	for _, t := range adbTables {
		tableName := reflect.ValueOf(t).Elem().Type().Name()
		l.Debug("Migrating", esl.String("table", tableName))
		if err = adb.AutoMigrate(t); err != nil {
			l.Debug("Unable to migrate", esl.Error(err), esl.String("table", tableName))
			return nil, nil, err
		}
	}
	for _, t := range sdbTables {
		tableName := reflect.ValueOf(t).Elem().Type().Name()
		l.Debug("Migrating", esl.String("table", tableName))
		if err = sdb.AutoMigrate(t); err != nil {
			l.Debug("Unable to migrate", esl.Error(err), esl.String("table", tableName))
			return nil, nil, err
		}
	}

	return adb, sdb, nil
}
func NewTeamScanner(ctl app_control.Control, client dbx_client.Client, path string) (TeamScanner, error) {
	l := ctl.Log().With(esl.String("path", path))
	adb, sdb, err := newDatabase(ctl, path)
	if err != nil {
		l.Debug("Unable to open database", esl.Error(err))
		return nil, err
	}
	return &tsImpl{
		ctl:    ctl,
		client: client,
		adb:    adb,
		sdb:    sdb,
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
	teamSummarizeTeamFolder      = "resolve_team_folder"
	teamSummarizeTeamFolderEntry = "resolve_team_folder_entry"
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
	// adb: API results database
	adb *gorm.DB
	// sdb: summary database
	sdb *gorm.DB
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
	z.adb.Create(g)
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

	db, err := z.adb.DB()
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
	s.Define(teamSummarizeTeamFolder, z.summarizeTeamFolder, s)
	s.Define(teamSummarizeTeamFolderEntry, z.summarizeTeamFolderEntry, s)
}

// Stage1: summarize namespaces
func (z tsImpl) summarizeStage1() error {
	l := z.ctl.Log()
	var lastErr error

	namespaceModel := &Namespace{}
	namespaceRows, err := z.adb.Model(namespaceModel).Distinct("namespace_id").Select("namespace_id").Rows()
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
			namespaceModel = &Namespace{}
			if err := z.adb.ScanRows(namespaceRows, namespaceModel); err != nil {
				l.Debug("Unable to scan namespace row", esl.Error(err))
				lastErr = err
				return
			}
			qNamespace.Enqueue(namespaceModel.NamespaceId)
		}
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))
	_ = namespaceRows.Close()

	if lastErr != nil {
		l.Debug("Failure on processing namespace rows", esl.Error(lastErr))
		return lastErr
	}

	return nil
}

// Stage2: summarize folders
func (z tsImpl) summarizeStage2() error {
	l := z.ctl.Log()
	var lastErr error

	folderEntry := &NamespaceEntry{}
	folderRows, err := z.adb.Model(folderEntry).Distinct("file_id").Where("entry_type = ?", "folder").Rows()
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
			folderEntry = &NamespaceEntry{}
			if err := z.adb.ScanRows(folderRows, folderEntry); err != nil {
				l.Debug("cannot scan row", esl.Error(err))
				lastErr = err
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

	return nil
}

// Stage3: summarize files
func (z tsImpl) summarizeStage3() error {
	l := z.ctl.Log()
	var lastErr error

	folderEntry := &NamespaceEntry{}
	folderRows, err := z.adb.Model(folderEntry).Distinct("file_id").Where("entry_type = ?", "folder").Rows()
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
			folderEntry = &NamespaceEntry{}
			if err := z.adb.ScanRows(folderRows, folderEntry); err != nil {
				l.Debug("cannot scan row", esl.Error(err))
				lastErr = err
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
	return nil
}

// Stage4: summarize entries
func (z tsImpl) summarizeStage4() error {
	l := z.ctl.Log()
	var lastErr error

	folderEntry := &NamespaceEntry{}
	folderRows, err := z.adb.Model(folderEntry).Distinct("file_id").Where("entry_type = ?", "folder").Rows()
	if err != nil {
		l.Debug("Unable to get folder rows", esl.Error(err))
		return err
	}
	defer func() {
		_ = folderRows.Close()
	}()

	z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
		z.defineSummarizeQueues(s)

		qEntry := s.Get(teamSummarizeEntry)
		for folderRows.Next() {
			folderEntry = &NamespaceEntry{}
			if err := z.adb.ScanRows(folderRows, folderEntry); err != nil {
				l.Debug("cannot scan row", esl.Error(err))
				lastErr = err
				return
			}
			qEntry.Enqueue(folderEntry.FileId)
		}

	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	if lastErr != nil {
		l.Debug("Failure on processing folder rows", esl.Error(lastErr))
		return lastErr
	}
	return nil
}

// Stage5: summarize team folder child entries
func (z tsImpl) summarizeStage5() error {
	l := z.ctl.Log()
	var lastErr error

	teamFolder := &TeamFolder{}
	teamFolderRows, err := z.adb.Model(teamFolder).Rows()
	if err != nil {
		l.Debug("Unable to get team folder rows", esl.Error(err))
		return err
	}
	defer func() {
		_ = teamFolderRows.Close()
	}()

	z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
		z.defineSummarizeQueues(s)

		qEntry := s.Get(teamSummarizeTeamFolder)
		for teamFolderRows.Next() {
			teamFolder := &TeamFolder{}
			if err := z.adb.ScanRows(teamFolderRows, teamFolder); err != nil {
				l.Debug("cannot scan row", esl.Error(err))
				lastErr = err
				return
			}
			qEntry.Enqueue(teamFolder)
		}
	})

	if lastErr != nil {
		l.Debug("Failure on processing folder rows", esl.Error(lastErr))
		return lastErr
	}
	return nil
}

func (z tsImpl) Summarize() error {
	l := z.ctl.Log()

	summaryTables := []interface{}{
		&SummaryEntry{},
		&SummaryFolderAndNamespace{},
		&SummaryFolderImmediateCount{},
		&SummaryFolderPath{},
		&SummaryFolderRecursive{},
		&SummaryNamespace{},
		&SummaryTeamFolderEntry{},
	}
	for _, st := range summaryTables {
		z.adb.Delete(st)
	}

	if err := z.summarizeStage1(); err != nil {
		l.Debug("Stage 1 failed", esl.Error(err))
		return err
	}
	if err := z.summarizeStage2(); err != nil {
		l.Debug("Stage 2 failed", esl.Error(err))
		return err
	}
	if err := z.summarizeStage3(); err != nil {
		l.Debug("Stage 3 failed", esl.Error(err))
		return err
	}
	if err := z.summarizeStage4(); err != nil {
		l.Debug("Stage 4 failed", esl.Error(err))
		return err
	}
	if err := z.summarizeStage5(); err != nil {
		l.Debug("Stage 5 failed", esl.Error(err))
		return err
	}

	db, err := z.adb.DB()
	if err != nil {
		l.Debug("Unable to get DB", esl.Error(err))
		return err
	}
	_ = db.Close()

	return nil
}
