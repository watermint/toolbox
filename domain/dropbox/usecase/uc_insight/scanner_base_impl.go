package uc_insight

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"reflect"
)

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

const (
	databaseName = "scan.db"
)

var (
	adbTables = []interface{}{
		&FileMember{},
		&GroupMember{},
		&Group{},
		&Member{},
		&Mount{},
		&NamespaceDetail{},
		&NamespaceEntry{},
		&NamespaceMember{},
		&Namespace{},
		&ReceivedFile{},
		&SharedLink{},
		&TeamFolder{},
	}
	adbErrorTables = []interface{}{
		&FileMemberError{},
		&GroupError{},
		&GroupMemberError{},
		&MemberError{},
		&MountError{},
		&NamespaceDetailError{},
		&NamespaceEntryError{},
		&NamespaceError{},
		&NamespaceMemberError{},
		&ReceivedFileError{},
		&SharedLinkError{},
		&TeamFolderError{},
	}

	sdbTables = []interface{}{
		&SummaryEntry{},
		&SummaryFolderAndNamespace{},
		&SummaryFolderError{},
		&SummaryFolderImmediateCount{},
		&SummaryFolderPath{},
		&SummaryFolderRecursive{},
		&SummaryNamespace{},
		&SummaryTeamFolderEntry{},
	}
)

func newDatabase(ctl app_control.Control, path string) (adb *gorm.DB, err error) {
	l := ctl.Log().With(esl.String("path", path))
	if err := os.MkdirAll(path, 0700); err != nil {
		l.Debug("Unable to create directory", esl.Error(err))
		return nil, err
	}

	adbPath := filepath.Join(path, databaseName)
	adb, err = ctl.NewOrm(adbPath)
	if err != nil {
		l.Debug("Unable to open database", esl.Error(err), esl.String("path", adbPath))
		return nil, err
	}

	for _, t := range adbTables {
		tableName := reflect.ValueOf(t).Elem().Type().Name()
		l.Debug("Migrating API tables", esl.String("table", tableName))
		if err = adb.AutoMigrate(t); err != nil {
			l.Debug("Unable to migrate", esl.Error(err), esl.String("table", tableName))
			return nil, err
		}
	}
	for _, t := range adbErrorTables {
		tableName := reflect.ValueOf(t).Elem().Type().Name()
		l.Debug("Migrating API error tables", esl.String("table", tableName))
		if err = adb.AutoMigrate(t); err != nil {
			l.Debug("Unable to migrate", esl.Error(err), esl.String("table", tableName))
			return nil, err
		}
	}
	for _, t := range sdbTables {
		tableName := reflect.ValueOf(t).Elem().Type().Name()
		l.Debug("Migrating summary tables", esl.String("table", tableName))
		if adb.Migrator().HasTable(t) {
			l.Debug("Try removing existing data", esl.String("table", tableName))
			if err = adb.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(t).Error; err != nil {
				l.Debug("Unable to delete", esl.Error(err), esl.String("table", tableName))
				return nil, err
			}
		}
		if err = adb.AutoMigrate(t); err != nil {
			l.Debug("Unable to migrate", esl.Error(err), esl.String("table", tableName))
			return nil, err
		}
	}

	return adb, nil
}
