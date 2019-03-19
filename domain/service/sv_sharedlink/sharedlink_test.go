package sv_sharedlink

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_test"
	"strings"
	"testing"
)

func TestSharedLinkImpl_CreateListDelete(t *testing.T) {
	api_test.DoTestTokenFull(func(ctx api_context.Context) {
		svc := New(ctx)
		src := api_test.ToolboxTestSuiteFolder.ChildPath("copy/F0.jpg")
		links, err := svc.ListByPath(src)
		if err != nil {
			t.Error(err)
			return
		}

		// clean up existing link
		for _, link := range links {
			if err := svc.Delete(link); err != nil {
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
			err = svc.Delete(link)
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
			err = svc.Delete(link)
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
			err = svc.Delete(link)
			if err != nil {
				t.Error(err)
			}
		}
	})
}
