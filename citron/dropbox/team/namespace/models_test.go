package namespace

import (
	"testing"
)

func TestMemberNamespaceSummary_Fields(t *testing.T) {
	// Test that the struct can be created and fields set
	summary := MemberNamespaceSummary{
		Email:             "test@example.com",
		TotalNamespaces:   10,
		MountedNamespaces: 8,
		OwnerNamespaces:   5,
		TeamFolders:       2,
		InsideTeamFolders: 3,
		ExternalFolders:   1,
		AppFolders:        4,
	}

	if summary.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got %s", summary.Email)
	}
	if summary.TotalNamespaces != 10 {
		t.Errorf("Expected 10 total namespaces, got %d", summary.TotalNamespaces)
	}
	if summary.MountedNamespaces != 8 {
		t.Errorf("Expected 8 mounted namespaces, got %d", summary.MountedNamespaces)
	}
	if summary.OwnerNamespaces != 5 {
		t.Errorf("Expected 5 owner namespaces, got %d", summary.OwnerNamespaces)
	}
	if summary.TeamFolders != 2 {
		t.Errorf("Expected 2 team folders, got %d", summary.TeamFolders)
	}
	if summary.InsideTeamFolders != 3 {
		t.Errorf("Expected 3 inside team folders, got %d", summary.InsideTeamFolders)
	}
	if summary.ExternalFolders != 1 {
		t.Errorf("Expected 1 external folder, got %d", summary.ExternalFolders)
	}
	if summary.AppFolders != 4 {
		t.Errorf("Expected 4 app folders, got %d", summary.AppFolders)
	}
}

func TestTeamNamespaceSummary_Fields(t *testing.T) {
	summary := TeamNamespaceSummary{
		NamespaceType:  "shared_folder",
		NamespaceCount: 42,
	}

	if summary.NamespaceType != "shared_folder" {
		t.Errorf("Expected namespace type 'shared_folder', got %s", summary.NamespaceType)
	}
	if summary.NamespaceCount != 42 {
		t.Errorf("Expected 42 namespaces, got %d", summary.NamespaceCount)
	}
}

func TestTeamFolderSummary_Fields(t *testing.T) {
	summary := TeamFolderSummary{
		Name:                "Engineering Team Folder",
		NumNamespacesInside: 15,
	}

	if summary.Name != "Engineering Team Folder" {
		t.Errorf("Expected name 'Engineering Team Folder', got %s", summary.Name)
	}
	if summary.NumNamespacesInside != 15 {
		t.Errorf("Expected 15 namespaces inside, got %d", summary.NumNamespacesInside)
	}
}

func TestFolderWithoutParent_Type(t *testing.T) {
	// Test that FolderWithoutParent is an alias for mo_sharedfolder.SharedFolder
	var _ FolderWithoutParent = FolderWithoutParent{
		SharedFolderId: "test_id",
		Name:           "Test Folder",
	}
}
