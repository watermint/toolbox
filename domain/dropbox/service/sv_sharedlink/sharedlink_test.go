package sv_sharedlink

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_url"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"strings"
	"testing"
	"time"
)

func TestEndToEndSharedLinkImpl_CreateListRemove(t *testing.T) {
	qt_api.DoTestTokenFull(func(ctx api_context.DropboxApiContext) {
		svc := New(ctx)
		src := qt_api.ToolboxTestSuiteFolder.ChildPath("copy/F0.jpg")
		links, err := svc.ListByPath(src)
		if err != nil {
			t.Error(err)
			return
		}

		// clean up existing link
		for _, link := range links {
			if err := svc.Remove(link); err != nil {
				t.Error(err)
			}
		}

		// Default link
		{
			link, err := svc.Create(src)
			if err != nil {
				t.Error(err)
				return
			}
			if link.LinkVisibility() != "public" {
				t.Error("invalid visibility")
			}
			if strings.ToLower(link.LinkName()) != "f0.jpg" {
				t.Error("invalid name")
			}
			if link.LinkExpires() != "" {
				t.Error("invalid expire")
			}

			// clean up
			err = svc.Remove(link)
			if err != nil {
				t.Error(err)
			}
		}

		// Team only link
		{
			link, err := svc.Create(src, TeamOnly())
			if err != nil {
				t.Error(err)
				return
			}
			if link.LinkVisibility() != "team_only" {
				t.Error("invalid visibility")
			}
			if strings.ToLower(link.LinkName()) != "f0.jpg" {
				t.Error("invalid name")
			}
			if link.LinkExpires() != "" {
				t.Error("invalid expire")
			}

			// clean up
			err = svc.Remove(link)
			if err != nil {
				t.Error(err)
			}
		}

		// Password protected link
		{
			link, err := svc.Create(src, Password("toolbox!"))
			if err != nil {
				t.Error(err)
				return
			}
			if link.LinkVisibility() != "password" {
				t.Error("invalid visibility")
			}
			if strings.ToLower(link.LinkName()) != "f0.jpg" {
				t.Error("invalid name")
			}
			if link.LinkExpires() != "" {
				t.Error("invalid expire")
			}

			// clean up
			err = svc.Remove(link)
			if err != nil {
				t.Error(err)
			}
		}
	})
}

// Mock tests

func TestSharedLinkImpl_Create(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
		sv := New(ctx)
		_, err := sv.Create(qt_recipe.NewTestDropboxFolderPath(), Public())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
		_, err = sv.Create(qt_recipe.NewTestDropboxFolderPath(), TeamOnly())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
		_, err = sv.Create(qt_recipe.NewTestDropboxFolderPath(), Expires(time.Now()))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedLinkImpl_List(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
		sv := New(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedLinkImpl_ListByPath(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
		sv := New(ctx)
		_, err := sv.ListByPath(qt_recipe.NewTestDropboxFolderPath())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedLinkImpl_Remove(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
		sv := New(ctx)
		err := sv.Remove(&mo_sharedlink.Metadata{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedLinkImpl_Resolve(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
		sv := New(ctx)
		_, err := sv.Resolve(mo_url.NewEmptyUrl(), "test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedLinkImpl_Update(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
		sv := New(ctx)
		_, err := sv.Update(&mo_sharedlink.Metadata{}, RemoveExpiration())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
