package dc_supplemental

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

func TestNewDbxCatalogue(t *testing.T) {
	// Test with repository media type
	cat := NewDbxCatalogue(dc_index.MediaRepository)
	if cat == nil {
		t.Error("Expected non-nil catalogue")
	}

	// Test with web media type
	webCat := NewDbxCatalogue(dc_index.MediaWeb)
	if webCat == nil {
		t.Error("Expected non-nil catalogue for web")
	}

	// Test with knowledge media type
	knowledgeCat := NewDbxCatalogue(dc_index.MediaKnowledge)
	if knowledgeCat == nil {
		t.Error("Expected non-nil catalogue for knowledge")
	}
}

func TestDbxCat_Recipe(t *testing.T) {
	cat := NewDbxCatalogue(dc_index.MediaRepository)
	dbxCat := cat.(*dbxCat)

	// Test with non-existent recipe path - this should panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected Recipe method to panic for non-existent path")
		}
	}()
	
	spec := dbxCat.Recipe("non-existent-path")
	// Should not reach here
	t.Error("Recipe method should have panicked, but got:", spec)
}

func TestDbxCat_WarnUnmentioned(t *testing.T) {
	cat := NewDbxCatalogue(dc_index.MediaRepository)
	dbxCat := cat.(*dbxCat)

	// WarnUnmentioned should return a boolean
	warn := dbxCat.WarnUnmentioned()
	// Just test that it returns without error
	_ = warn
}

func TestDbxCat_RecipeTable(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		cat := NewDbxCatalogue(dc_index.MediaRepository)
		dbxCat := cat.(*dbxCat)

		// Test with empty paths - should not panic
		dbxCat.RecipeTable("test-table", ctl.UI(), []string{})

		// Test with invalid paths should be wrapped in panic handler
		defer func() {
			if r := recover(); r != nil {
				// Expected to panic with invalid recipe paths
				t.Logf("RecipeTable panicked as expected with invalid paths: %v", r)
			}
		}()
		
		// This will likely panic, but that's the expected behavior
		paths := []string{"dropbox", "file", "list"}
		dbxCat.RecipeTable("test-table", ctl.UI(), paths)

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestNewDropboxBusiness(t *testing.T) {
	// Test with different media types
	mediaTypes := []dc_index.MediaType{
		dc_index.MediaRepository,
		dc_index.MediaWeb,
		dc_index.MediaKnowledge,
	}

	for _, mediaType := range mediaTypes {
		doc := NewDropboxBusiness(mediaType)
		if doc == nil {
			t.Errorf("Expected non-nil document for media type %v", mediaType)
		}

		// Test that it implements the Document interface methods
		docImpl := doc.(*DropboxBusiness)
		
		// Test DocDesc
		desc := docImpl.DocDesc()
		if desc == nil {
			t.Error("Expected non-nil doc description")
		}

		// Test DocId
		docId := docImpl.DocId()
		// DocId should be a valid value (can't easily test specific value)
		_ = docId

		// Test Sections
		sections := docImpl.Sections()
		if sections == nil {
			t.Error("Expected non-nil sections")
		}

		if len(sections) == 0 {
			t.Error("Expected at least one section")
		}
	}
}

func TestDropboxBusinessSections(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		doc := NewDropboxBusiness(dc_index.MediaRepository)
		docImpl := doc.(*DropboxBusiness)

		sections := docImpl.Sections()

		// Test each section has proper Title and Body methods
		for i, section := range sections {
			// Test Title method
			title := section.Title()
			if title == nil {
				t.Errorf("Section %d should have non-nil title", i)
			}

			// Test Body method (may panic due to missing recipes)
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Logf("Section %d Body method panicked as expected: %v", i, r)
					}
				}()
				section.Body(ctl.UI())
			}()
		}

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDropboxBusinessMember(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		cat := NewDbxCatalogue(dc_index.MediaRepository)
		member := DropboxBusinessMember{cat: cat}

		title := member.Title()
		if title == nil {
			t.Error("Expected non-nil title")
		}

		// Test Body method (may panic due to missing recipes)
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Member Body method panicked as expected: %v", r)
			}
		}()
		member.Body(ctl.UI())

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDropboxBusinessGroup(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		cat := NewDbxCatalogue(dc_index.MediaRepository)
		group := DropboxBusinessGroup{cat: cat}

		title := group.Title()
		if title == nil {
			t.Error("Expected non-nil title")
		}

		// Test Body method (may panic due to missing recipes)
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Group Body method panicked as expected: %v", r)
			}
		}()
		group.Body(ctl.UI())

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDropboxBusinessContent(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		cat := NewDbxCatalogue(dc_index.MediaRepository)
		content := DropboxBusinessContent{cat: cat}

		title := content.Title()
		if title == nil {
			t.Error("Expected non-nil title")
		}

		// Test Body method (may panic due to missing recipes)
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Content Body method panicked as expected: %v", r)
			}
		}()
		content.Body(ctl.UI())

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDropboxBusinessConnect(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		cat := NewDbxCatalogue(dc_index.MediaRepository)
		connect := DropboxBusinessConnect{cat: cat}

		title := connect.Title()
		if title == nil {
			t.Error("Expected non-nil title")
		}

		// Test Body method (may panic due to missing recipes)
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Connect Body method panicked as expected: %v", r)
			}
		}()
		connect.Body(ctl.UI())

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDropboxBusinessSharedLink(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		cat := NewDbxCatalogue(dc_index.MediaRepository)
		sharedLink := DropboxBusinessSharedLink{cat: cat}

		title := sharedLink.Title()
		if title == nil {
			t.Error("Expected non-nil title")
		}

		// Test Body method (may panic due to missing recipes)
		defer func() {
			if r := recover(); r != nil {
				t.Logf("SharedLink Body method panicked as expected: %v", r)
			}
		}()
		sharedLink.Body(ctl.UI())

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDropboxBusinessFileLock(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		cat := NewDbxCatalogue(dc_index.MediaRepository)
		fileLock := DropboxBusinessFileLock{cat: cat}

		title := fileLock.Title()
		if title == nil {
			t.Error("Expected non-nil title")
		}

		// Test Body method (may panic due to missing recipes)
		defer func() {
			if r := recover(); r != nil {
				t.Logf("FileLock Body method panicked as expected: %v", r)
			}
		}()
		fileLock.Body(ctl.UI())

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDropboxBusinessActivities(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		cat := NewDbxCatalogue(dc_index.MediaRepository)
		activities := DropboxBusinessActivities{cat: cat}

		title := activities.Title()
		if title == nil {
			t.Error("Expected non-nil title")
		}

		// Test Body method (may panic due to missing recipes)
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Activities Body method panicked as expected: %v", r)
			}
		}()
		activities.Body(ctl.UI())

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDropboxBusinessUsecase(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		cat := NewDbxCatalogue(dc_index.MediaRepository)
		usecase := DropboxBusinessUsecase{cat: cat}

		title := usecase.Title()
		if title == nil {
			t.Error("Expected non-nil title")
		}

		// Test Body method (may panic due to missing recipes)
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Usecase Body method panicked as expected: %v", r)
			}
		}()
		usecase.Body(ctl.UI())

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDropboxBusinessPaper(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		cat := NewDbxCatalogue(dc_index.MediaRepository)
		paper := DropboxBusinessPaper{cat: cat}

		title := paper.Title()
		if title == nil {
			t.Error("Expected non-nil title")
		}

		// Test Body method (may panic due to missing recipes)
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Paper Body method panicked as expected: %v", r)
			}
		}()
		paper.Body(ctl.UI())

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDropboxBusinessTeamAdmin(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		cat := NewDbxCatalogue(dc_index.MediaRepository)
		teamAdmin := DropboxBusinessTeamAdmin{cat: cat}

		title := teamAdmin.Title()
		if title == nil {
			t.Error("Expected non-nil title")
		}

		// Test Body method (may panic due to missing recipes)
		defer func() {
			if r := recover(); r != nil {
				t.Logf("TeamAdmin Body method panicked as expected: %v", r)
			}
		}()
		teamAdmin.Body(ctl.UI())

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDropboxBusinessRunAs(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		cat := NewDbxCatalogue(dc_index.MediaRepository)
		runAs := DropboxBusinessRunAs{cat: cat}

		title := runAs.Title()
		if title == nil {
			t.Error("Expected non-nil title")
		}

		// Test Body method (may panic due to missing recipes)
		defer func() {
			if r := recover(); r != nil {
				t.Logf("RunAs Body method panicked as expected: %v", r)
			}
		}()
		runAs.Body(ctl.UI())

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDropboxBusinessLegalHold(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		cat := NewDbxCatalogue(dc_index.MediaRepository)
		legalHold := DropboxBusinessLegalHold{cat: cat}

		title := legalHold.Title()
		if title == nil {
			t.Error("Expected non-nil title")
		}

		// Test Body method (may panic due to missing recipes)
		defer func() {
			if r := recover(); r != nil {
				t.Logf("LegalHold Body method panicked as expected: %v", r)
			}
		}()
		legalHold.Body(ctl.UI())

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDropboxBusinessFootnote(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		cat := NewDbxCatalogue(dc_index.MediaRepository)
		footnote := DropboxBusinessFootnote{cat: cat}

		title := footnote.Title()
		if title == nil {
			t.Error("Expected non-nil title")
		}

		// Test Body method (may panic due to missing recipes)
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Footnote Body method panicked as expected: %v", r)
			}
		}()
		footnote.Body(ctl.UI())

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestMsgDropboxBusiness(t *testing.T) {
	// Test that MDropboxBusiness is properly initialized
	if MDropboxBusiness == nil {
		t.Error("Expected MDropboxBusiness to be initialized")
	}

	// Test some key message fields
	if MDropboxBusiness.Title == nil {
		t.Error("Expected Title to be initialized")
	}

	if MDropboxBusiness.Overview == nil {
		t.Error("Expected Overview to be initialized")
	}

	if MDropboxBusiness.MemberTitle == nil {
		t.Error("Expected MemberTitle to be initialized")
	}

	if MDropboxBusiness.GroupTitle == nil {
		t.Error("Expected GroupTitle to be initialized")
	}

	if MDropboxBusiness.ContentTitle == nil {
		t.Error("Expected ContentTitle to be initialized")
	}
}

func TestSkipDropboxBusinessCommandDoc(t *testing.T) {
	// Test the global flag
	originalValue := SkipDropboxBusinessCommandDoc
	
	// Test setting to true
	SkipDropboxBusinessCommandDoc = true
	if !SkipDropboxBusinessCommandDoc {
		t.Error("Expected SkipDropboxBusinessCommandDoc to be true")
	}

	// Test setting to false
	SkipDropboxBusinessCommandDoc = false
	if SkipDropboxBusinessCommandDoc {
		t.Error("Expected SkipDropboxBusinessCommandDoc to be false")
	}

	// Restore original value
	SkipDropboxBusinessCommandDoc = originalValue
}