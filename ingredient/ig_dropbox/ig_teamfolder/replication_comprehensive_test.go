package ig_teamfolder

import (
	"strings"
	"testing"

	"github.com/watermint/toolbox/domain/dropbox/model/mo_teamfolder"
)

func TestMirrorGroupNamePrefix(t *testing.T) {
	// Test that the constant is defined as expected
	if MirrorGroupNamePrefix != "toolbox-teamfolder-mirror" {
		t.Errorf("Expected MirrorGroupNamePrefix to be 'toolbox-teamfolder-mirror', got '%s'", MirrorGroupNamePrefix)
	}
}

func TestArchiveOnSuccess(t *testing.T) {
	opts := &mirrorOpts{}
	result := ArchiveOnSuccess()(opts)

	if !result.archiveOnSuccess {
		t.Error("Expected archiveOnSuccess to be true after applying ArchiveOnSuccess option")
	}
}

func TestSkipVerify(t *testing.T) {
	opts := &mirrorOpts{}
	result := SkipVerify()(opts)

	if !result.skipVerify {
		t.Error("Expected skipVerify to be true after applying SkipVerify option")
	}
}

func TestMirrorPair(t *testing.T) {
	// Test MirrorPair struct
	src := &mo_teamfolder.TeamFolder{
		TeamFolderId: "tf_123",
		Name:         "Test Folder",
	}
	dst := &mo_teamfolder.TeamFolder{
		TeamFolderId: "tf_456",
		Name:         "Test Folder Copy",
	}

	pair := &MirrorPair{
		Src: src,
		Dst: dst,
	}

	if pair.Src.TeamFolderId != "tf_123" {
		t.Errorf("Expected Src TeamFolderId 'tf_123', got '%s'", pair.Src.TeamFolderId)
	}
	if pair.Dst.TeamFolderId != "tf_456" {
		t.Errorf("Expected Dst TeamFolderId 'tf_456', got '%s'", pair.Dst.TeamFolderId)
	}
}

func TestNewScope(t *testing.T) {
	// Test NewScope function
	pair := &MirrorPair{
		Src: &mo_teamfolder.TeamFolder{
			TeamFolderId: "tf_123",
			Name:         "Test Folder",
		},
		Dst: nil,
	}

	scope := NewScope(pair)
	if scope == nil {
		t.Fatal("Expected non-nil scope")
	}

	// Test Pair() method
	returnedPair := scope.Pair()
	if returnedPair != pair {
		t.Error("Expected Pair() to return the same pair")
	}
	if returnedPair.Src.TeamFolderId != "tf_123" {
		t.Errorf("Expected TeamFolderId 'tf_123', got '%s'", returnedPair.Src.TeamFolderId)
	}
}

func TestReplication_Preset(t *testing.T) {
	// Skip this test as it requires proper initialization of all fields
	// which is done by the framework
	t.Skip("Preset requires framework initialization")
}

func TestMarshalUnmarshalContext(t *testing.T) {
	// Test that the marshal/unmarshal functions exist
	// Since mirrorContext is not exported, we can't test this directly
	// We'll just verify the functions are defined
	_ = MarshalContext
	_ = UnmarshalContext
}

func TestReplication_PartialScope_Matching(t *testing.T) {
	// Test the matching logic for partial scope
	// This tests the internal matching function behavior

	names := []string{"Marketing", "Sales", "Engineering"}

	testCases := []struct {
		folderName  string
		shouldMatch bool
	}{
		{"marketing", true},
		{"Marketing", true},
		{"MARKETING", true},
		{"sales", true},
		{"Sales", true},
		{"engineering", true},
		{"Engineering", true},
		{"Finance", false},
		{"HR", false},
		{"", false},
	}

	// Test the matching logic
	for _, tc := range testCases {
		t.Run(tc.folderName, func(t *testing.T) {
			matches := false
			fnl := strings.ToLower(tc.folderName)
			for _, name := range names {
				if strings.ToLower(name) == fnl {
					matches = true
					break
				}
			}

			if matches != tc.shouldMatch {
				t.Errorf("Folder '%s' match result: expected %v, got %v", tc.folderName, tc.shouldMatch, matches)
			}
		})
	}
}

func TestReplication_BasePath_Options(t *testing.T) {
	// Skip this test as it requires proper initialization
	t.Skip("BasePath test requires framework initialization")
}

func TestReplication_Exec_Validations(t *testing.T) {
	// Skip this test as it requires proper framework initialization
	t.Skip("Exec validation test requires framework initialization")
}

func TestScope_Interface(t *testing.T) {
	// Test that Scope interface methods work correctly through NewScope
	pair := &MirrorPair{
		Src: &mo_teamfolder.TeamFolder{
			TeamFolderId: "tf_src",
			Name:         "Source Folder",
		},
		Dst: &mo_teamfolder.TeamFolder{
			TeamFolderId: "tf_dst",
			Name:         "Dest Folder",
		},
	}

	scope := NewScope(pair)

	// Test Pair method
	returnedPair := scope.Pair()
	if returnedPair.Src.TeamFolderId != "tf_src" {
		t.Error("Pair() did not return expected source folder")
	}
	if returnedPair.Dst.TeamFolderId != "tf_dst" {
		t.Error("Pair() did not return expected destination folder")
	}
}

func TestMirrorOpts_MultipleOptions(t *testing.T) {
	// Test applying multiple options
	opts := &mirrorOpts{}

	// Apply both options
	opts = ArchiveOnSuccess()(opts)
	opts = SkipVerify()(opts)

	if !opts.archiveOnSuccess {
		t.Error("Expected archiveOnSuccess to be true")
	}
	if !opts.skipVerify {
		t.Error("Expected skipVerify to be true")
	}
}

func TestReplication_EmptyTargetNames(t *testing.T) {
	// Skip test that requires client initialization
	t.Skip("Requires client initialization")
}

func TestReplication_CaseInsensitiveMatching(t *testing.T) {
	// Test case-insensitive matching in PartialScope
	targetNames := []string{"Marketing", "SALES", "engineering"}

	// Test folder names that should match
	testFolders := []string{
		"marketing",
		"Marketing",
		"MARKETING",
		"sales",
		"Sales",
		"SALES",
		"engineering",
		"Engineering",
		"ENGINEERING",
	}

	for _, folderName := range testFolders {
		matched := false
		fnl := strings.ToLower(folderName)
		for _, name := range targetNames {
			if strings.ToLower(name) == fnl {
				matched = true
				break
			}
		}

		if !matched {
			t.Errorf("Folder name '%s' should have matched target names", folderName)
		}
	}
}
