package uc_insight

import (
	"path/filepath"
	"testing"

	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

func TestScanOpts_Apply_NoOptions(t *testing.T) {
	opts := ScanOpts{
		MaxRetries:        5,
		ScanMemberFolders: true,
		BaseNamespace:     dbx_filesystem.BaseNamespaceRoot,
	}

	result := opts.Apply([]ScanOpt{})

	if result.MaxRetries != 5 {
		t.Errorf("Expected MaxRetries to remain 5, got %d", result.MaxRetries)
	}
	if !result.ScanMemberFolders {
		t.Error("Expected ScanMemberFolders to remain true")
	}
	if result.BaseNamespace != dbx_filesystem.BaseNamespaceRoot {
		t.Errorf("Expected BaseNamespace to remain Root, got %v", result.BaseNamespace)
	}
}

func TestScanOpts_Apply_SingleOption(t *testing.T) {
	opts := ScanOpts{}

	result := opts.Apply([]ScanOpt{MaxRetries(10)})

	if result.MaxRetries != 10 {
		t.Errorf("Expected MaxRetries to be 10, got %d", result.MaxRetries)
	}
}

func TestScanOpts_Apply_MultipleOptions(t *testing.T) {
	opts := ScanOpts{}

	result := opts.Apply([]ScanOpt{
		MaxRetries(3),
		ScanMemberFolders(true),
		BaseNamespace(dbx_filesystem.BaseNamespaceHome),
	})

	if result.MaxRetries != 3 {
		t.Errorf("Expected MaxRetries to be 3, got %d", result.MaxRetries)
	}
	if !result.ScanMemberFolders {
		t.Error("Expected ScanMemberFolders to be true")
	}
	if result.BaseNamespace != dbx_filesystem.BaseNamespaceHome {
		t.Errorf("Expected BaseNamespace to be UserRoot, got %v", result.BaseNamespace)
	}
}

func TestMaxRetries(t *testing.T) {
	opts := ScanOpts{}

	newOpts := MaxRetries(7)(opts)

	if newOpts.MaxRetries != 7 {
		t.Errorf("Expected MaxRetries to be 7, got %d", newOpts.MaxRetries)
	}
}

func TestScanMemberFolders(t *testing.T) {
	opts := ScanOpts{}

	// Test enabling
	newOpts := ScanMemberFolders(true)(opts)
	if !newOpts.ScanMemberFolders {
		t.Error("Expected ScanMemberFolders to be true after enabling")
	}

	// Test disabling
	newOpts = ScanMemberFolders(false)(newOpts)
	if newOpts.ScanMemberFolders {
		t.Error("Expected ScanMemberFolders to be false after disabling")
	}
}

func TestBaseNamespace(t *testing.T) {
	opts := ScanOpts{}

	newOpts := BaseNamespace(dbx_filesystem.BaseNamespaceHome)(opts)

	if newOpts.BaseNamespace != dbx_filesystem.BaseNamespaceHome {
		t.Errorf("Expected BaseNamespace to be TeamRoot, got %v", newOpts.BaseNamespace)
	}
}

func TestDatabaseFromPath(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		tempDir := ctl.Workspace().Job()
		dbPath := filepath.Join(tempDir, "test_db")

		db, err := DatabaseFromPath(ctl, dbPath)

		if err != nil {
			t.Errorf("Expected no error creating database, got %v", err)
		}
		if db == nil {
			t.Error("Expected non-nil database")
		}

		// Clean up
		if sqlDB, err := db.DB(); err == nil {
			_ = sqlDB.Close()
		}
	})
}

func TestDatabaseFromPath_InvalidPath(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		// Use an invalid path (empty or with invalid characters)
		invalidPath := ""

		db, err := DatabaseFromPath(ctl, invalidPath)

		// Should handle gracefully - either error or create in current directory
		if err != nil && db != nil {
			t.Error("If error occurs, database should be nil")
		}
		if err == nil && db == nil {
			t.Error("If no error, database should not be nil")
		}

		// Clean up if db was created
		if db != nil {
			if sqlDB, err := db.DB(); err == nil {
				_ = sqlDB.Close()
			}
		}
	})
}

func TestHasEntryOf_EmptyTable(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Fatalf("Failed to create in-memory database: %v", err)
		}

		// Migrate the table first
		err = db.AutoMigrate(&Namespace{})
		if err != nil {
			t.Fatalf("Failed to migrate table: %v", err)
		}

		has, err := HasEntryOf(db, &Namespace{})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if has {
			t.Error("Expected empty table to return false")
		}
	})
}

func TestHasEntryOf_WithData(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Fatalf("Failed to create in-memory database: %v", err)
		}

		// Migrate the table first
		err = db.AutoMigrate(&Namespace{})
		if err != nil {
			t.Fatalf("Failed to migrate table: %v", err)
		}

		// Add test data
		testNamespace := &Namespace{
			NamespaceId:   "test123",
			Name:          "Test Namespace",
			NamespaceType: "user_folder",
		}
		err = db.Create(testNamespace).Error
		if err != nil {
			t.Fatalf("Failed to create test data: %v", err)
		}

		has, err := HasEntryOf(db, &Namespace{})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !has {
			t.Error("Expected table with data to return true")
		}
	})
}

func TestHasEntry(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Fatalf("Failed to create in-memory database: %v", err)
		}

		// Migrate the table first
		err = db.AutoMigrate(&Namespace{})
		if err != nil {
			t.Fatalf("Failed to migrate table: %v", err)
		}

		// Test with empty table
		has, err := HasEntry(db)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if has {
			t.Error("Expected empty table to return false")
		}

		// Add test data
		testNamespace := &Namespace{
			NamespaceId:   "test456",
			Name:          "Test Namespace 2",
			NamespaceType: "team_folder",
		}
		err = db.Create(testNamespace).Error
		if err != nil {
			t.Fatalf("Failed to create test data: %v", err)
		}

		// Test with data
		has, err = HasEntry(db)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !has {
			t.Error("Expected table with data to return true")
		}
	})
}

func TestHasEntryOf_InvalidTable(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Fatalf("Failed to create in-memory database: %v", err)
		}

		// Define a struct that doesn't exist as a table
		type NonExistentTable struct {
			ID   uint `gorm:"primaryKey"`
			Name string
		}

		has, err := HasEntryOf(db, &NonExistentTable{})

		// Should return error for non-existent table
		if err == nil {
			t.Error("Expected error for non-existent table")
		}
		if has {
			t.Error("Expected false for non-existent table")
		}
	})
}

func TestTsImpl_ReportLastErrors_NoHandler(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Fatalf("Failed to create in-memory database: %v", err)
		}

		ts := &tsImpl{
			ctl: ctl,
			db:  db,
		}

		count, err := ts.ReportLastErrors(nil)

		if err != nil {
			t.Errorf("Expected no error with nil handler, got %v", err)
		}
		if count != 0 {
			t.Errorf("Expected count 0 with nil handler, got %d", count)
		}
	})
}

func TestTsImpl_ReportLastErrors_WithHandler(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Fatalf("Failed to create in-memory database: %v", err)
		}

		// Migrate error tables
		err = db.AutoMigrate(&MemberError{})
		if err != nil {
			t.Fatalf("Failed to migrate table: %v", err)
		}

		// Add test error data
		testError := &MemberError{
			Dummy: "test_dummy",
			ApiError: ApiError{
				Error:    "Test error message",
				ErrorTag: "test_error_tag",
			},
		}
		err = db.Create(testError).Error
		if err != nil {
			t.Fatalf("Failed to create test error: %v", err)
		}

		ts := &tsImpl{
			ctl: ctl,
			db:  db,
		}

		var reportedErrors []ApiErrorReport
		count, err := ts.ReportLastErrors(func(errCategory, errMessage, errTag, detail string) {
			reportedErrors = append(reportedErrors, ApiErrorReport{
				Category: errCategory,
				Message:  errMessage,
				Tag:      errTag,
				Detail:   detail,
			})
		})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if count == 0 {
			t.Error("Expected at least one error to be reported")
		}
		if len(reportedErrors) == 0 {
			t.Error("Expected at least one error in reportedErrors")
		}

		if len(reportedErrors) > 0 {
			report := reportedErrors[0]
			if report.Category != "MemberError" {
				t.Errorf("Expected category 'MemberError', got '%s'", report.Category)
			}
			if report.Message != "Test error message" {
				t.Errorf("Expected message 'Test error message', got '%s'", report.Message)
			}
			if report.Tag != "test_error_tag" {
				t.Errorf("Expected tag 'test_error_tag', got '%s'", report.Tag)
			}
		}
	})
}

func TestDatabaseName(t *testing.T) {
	if databaseName != "scan.db" {
		t.Errorf("Expected database name to be 'scan.db', got '%s'", databaseName)
	}
}

func TestQueueConstants(t *testing.T) {
	expectedConstants := map[string]string{
		teamScanQueueFileMember:      "scan_file_member",
		teamScanQueueGroup:           "scan_group",
		teamScanQueueGroupMember:     "scan_group_member",
		teamScanQueueMember:          "scan_member",
		teamScanQueueMount:           "scan_mount",
		teamScanQueueNamespace:       "scan_team_namespace",
		teamScanQueueNamespaceDetail: "scan_namespace",
		teamScanQueueNamespaceEntry:  "scan_folder",
		teamScanQueueNamespaceMember: "scan_namespace_member",
		teamScanQueueReceivedFile:    "scan_received_file",
		teamScanQueueSharedLink:      "scan_shared_link",
		teamScanQueueTeamFolder:      "scan_team_folder",
		teamSummarizeEntry:           "resolve_entry",
		teamSummarizeFolderImmediate: "resolve_folder_immediate",
		teamSummarizeFolderPath:      "resolve_folder_path",
		teamSummarizeFolderRecursive: "resolve_folder_recursive",
		teamSummarizeNamespace:       "resolve_namespace",
		teamSummarizeTeamFolder:      "resolve_team_folder",
		teamSummarizeTeamFolderEntry: "resolve_team_folder_entry",
	}

	for constant, expected := range expectedConstants {
		if constant != expected {
			t.Errorf("Expected constant '%s' to equal '%s'", constant, expected)
		}
	}
}

func TestAdbTables(t *testing.T) {
	expectedCount := 12
	if len(adbTables) != expectedCount {
		t.Errorf("Expected %d adb tables, got %d", expectedCount, len(adbTables))
	}

	// Verify first few tables exist
	if len(adbTables) > 0 {
		if adbTables[0] == nil {
			t.Error("Expected first adb table to be non-nil")
		}
	}
}

func TestAdbErrorTables(t *testing.T) {
	expectedCount := 12
	if len(adbErrorTables) != expectedCount {
		t.Errorf("Expected %d adb error tables, got %d", expectedCount, len(adbErrorTables))
	}

	// Verify first few error tables exist
	if len(adbErrorTables) > 0 {
		if adbErrorTables[0] == nil {
			t.Error("Expected first adb error table to be non-nil")
		}
	}
}

func TestSdbTables(t *testing.T) {
	expectedCount := 8
	if len(sdbTables) != expectedCount {
		t.Errorf("Expected %d sdb tables, got %d", expectedCount, len(sdbTables))
	}

	// Verify first few summary tables exist
	if len(sdbTables) > 0 {
		if sdbTables[0] == nil {
			t.Error("Expected first sdb table to be non-nil")
		}
	}
}

func TestNewDatabase_Success(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		tempDir := ctl.Workspace().Job()

		db, err := newDatabase(ctl, tempDir)

		if err != nil {
			t.Errorf("Expected no error creating database, got %v", err)
		}
		if db == nil {
			t.Error("Expected non-nil database")
		}

		// Verify some tables were created by checking if they can be queried
		var count int64
		err = db.Model(&Namespace{}).Count(&count).Error
		if err != nil {
			t.Errorf("Expected to be able to query Namespace table, got error: %v", err)
		}

		// Clean up
		if sqlDB, err := db.DB(); err == nil {
			_ = sqlDB.Close()
		}
	})
}

func TestNewDatabase_InvalidPath(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		// Use a path that would cause issues
		invalidPath := "/nonexistent/deeply/nested/path/that/cannot/be/created"

		db, err := newDatabase(ctl, invalidPath)

		// Should return error for invalid path
		if err == nil {
			t.Error("Expected error for invalid path")
		}
		if db != nil {
			t.Error("Expected nil database for invalid path")
		}
	})
}
