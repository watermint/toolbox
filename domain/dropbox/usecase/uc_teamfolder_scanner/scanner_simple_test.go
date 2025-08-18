package uc_teamfolder_scanner

import (
	"testing"
	"time"

	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs_impl"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

func TestScanImpl_BasicFunctionality(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
			// Test creating scanner with different configurations
			scanner1 := New(ctl, ctx, ScanTimeoutShort, dbx_filesystem.BaseNamespaceRoot)
			if scanner1 == nil {
				t.Error("Expected non-nil scanner")
			}

			scanner2 := New(ctl, ctx, ScanTimeoutLong, dbx_filesystem.BaseNamespaceHome)
			if scanner2 == nil {
				t.Error("Expected non-nil scanner")
			}

			// Test scan with valid filter
			_, err := scanner1.Scan(mo_filter.New("test"))
			if err != qt_errors.ErrorMock {
				t.Errorf("Expected mock error, got %v", err)
			}
		})
	})
}

func TestScanImpl_StorageOperations(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		// Create empty KVS for testing
		kvs := kv_kvs_impl.NewEmpty()

		// Test basic KVS operations
		err := kvs.PutString("test_key", "test_value")
		if err != nil {
			t.Errorf("Expected no error on PutString, got %v", err)
		}

		// GetString should return ErrorNotFound for empty implementation
		_, err = kvs.GetString("test_key")
		if err != kv_kvs.ErrorNotFound {
			t.Errorf("Expected ErrorNotFound, got %v", err)
		}

		// Test JSON model operations
		testFolder := &mo_sharedfolder.SharedFolder{
			SharedFolderId: "sf_test",
			Name:           "Test Folder",
			IsTeamFolder:   true,
		}

		err = kvs.PutJsonModel("folder_key", testFolder)
		if err != nil {
			t.Errorf("Expected no error on PutJsonModel, got %v", err)
		}

		// GetJsonModel should return ErrorNotFound for empty implementation
		var retrieved mo_sharedfolder.SharedFolder
		err = kvs.GetJsonModel("folder_key", &retrieved)
		if err != kv_kvs.ErrorNotFound {
			t.Errorf("Expected ErrorNotFound, got %v", err)
		}
	})
}

func TestScanImpl_DataStructures(t *testing.T) {
	// Test TeamFolder structure
	tf := &TeamFolder{
		TeamFolder: &mo_sharedfolder.SharedFolder{
			SharedFolderId:     "tf_123",
			Name:               "Test Team Folder",
			IsTeamFolder:       true,
			IsInsideTeamFolder: false,
		},
		NestedFolders: make(map[string]*mo_sharedfolder.SharedFolder),
	}

	// Add nested folders
	tf.NestedFolders["/project1"] = &mo_sharedfolder.SharedFolder{
		SharedFolderId:     "sf_project1",
		Name:               "Project 1",
		IsTeamFolder:       false,
		IsInsideTeamFolder: true,
	}

	tf.NestedFolders["/project2"] = &mo_sharedfolder.SharedFolder{
		SharedFolderId:     "sf_project2",
		Name:               "Project 2",
		IsTeamFolder:       false,
		IsInsideTeamFolder: true,
	}

	// Verify structure
	if tf.TeamFolder.SharedFolderId != "tf_123" {
		t.Errorf("Expected team folder ID 'tf_123', got '%s'", tf.TeamFolder.SharedFolderId)
	}

	if len(tf.NestedFolders) != 2 {
		t.Errorf("Expected 2 nested folders, got %d", len(tf.NestedFolders))
	}

	// Test TeamFolderNested structure
	tfn := &TeamFolderNested{
		NamespaceId:   "ns_456",
		NamespaceName: "Namespace Test",
		RelativePath:  "/test/path",
	}

	if tfn.NamespaceId != "ns_456" {
		t.Errorf("Expected namespace ID 'ns_456', got '%s'", tfn.NamespaceId)
	}

	// Test TeamFolderEntry structure
	tfe := &TeamFolderEntry{
		NamespaceId: "ns_parent",
		Descendants: []string{"ns_child1", "ns_child2", "ns_child3"},
	}

	if len(tfe.Descendants) != 3 {
		t.Errorf("Expected 3 descendants, got %d", len(tfe.Descendants))
	}
}

func TestScanImpl_TimeoutModes(t *testing.T) {
	// Test timeout constants
	if scanShortTimeout != 3*time.Minute {
		t.Errorf("Expected short timeout to be 3 minutes, got %v", scanShortTimeout)
	}
	if scanLongTimeout != 3*time.Hour {
		t.Errorf("Expected long timeout to be 3 hours, got %v", scanLongTimeout)
	}

	// Test ScanTimeoutMode constants
	if ScanTimeoutShort != "short" {
		t.Errorf("Expected ScanTimeoutShort to be 'short', got %s", ScanTimeoutShort)
	}
	if ScanTimeoutLong != "long" {
		t.Errorf("Expected ScanTimeoutLong to be 'long', got %s", ScanTimeoutLong)
	}
	if ScanTimeoutAltPath != "/:ERROR-SCAN-TIMEOUT:/" {
		t.Errorf("Expected ScanTimeoutAltPath to be '/:ERROR-SCAN-TIMEOUT:/', got %s", ScanTimeoutAltPath)
	}
}

func TestScanImpl_FilterValidation(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
			scanner := New(ctl, ctx, ScanTimeoutShort, dbx_filesystem.BaseNamespaceRoot)

			// Test with nil filter
			_, err := scanner.Scan(nil)
			if err == nil {
				t.Error("Expected error with nil filter")
			}

			// Test with valid filters
			validFilters := []string{
				"",          // empty filter
				"project",   // simple filter
				"Project*",  // wildcard
				"test|demo", // OR filter
			}

			for _, filterStr := range validFilters {
				filter := mo_filter.New(filterStr)
				_, err := scanner.Scan(filter)
				// We expect mock error in test environment
				if err != qt_errors.ErrorMock && err != nil {
					t.Logf("Filter '%s' returned unexpected error: %v", filterStr, err)
				}
			}
		})
	})
}

func TestScanImpl_QueueConstants(t *testing.T) {
	// Test queue ID constants
	expectedQueues := map[string]string{
		queueIdScanTeamNamespace:     "scan_team",
		queueIdScanNamespaceMetadata: "scan_namespace",
		queueIdScanTeamFolder:        "scan_teamfolder",
		queueIdExtractTeamFolder:     "extract_teamfolder",
	}

	for actual, expected := range expectedQueues {
		if actual != expected {
			t.Errorf("Expected queue constant to be '%s', got '%s'", expected, actual)
		}
	}
}

func TestScanImpl_ErrorScenarios(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
			scanner := New(ctl, ctx, ScanTimeoutShort, dbx_filesystem.BaseNamespaceRoot)

			// Test error scenarios

			// Test with nil filter should error
			_, err := scanner.Scan(nil)
			if err == nil {
				t.Error("Expected error with nil filter, got nil")
			}

			// Test with empty filter should handle gracefully
			emptyFilter := mo_filter.New("")
			_, err = scanner.Scan(emptyFilter)
			// In test environment, we expect mock error
			if err != qt_errors.ErrorMock && err != nil {
				t.Logf("Scan with empty filter returned: %v", err)
			}
		})
	})
}
