package sv_file_url

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

const (
	DummyImageUrl = "https://dummyimage.com/64x64/888/222.png"
)

func TestPathWithName(t *testing.T) {
	base := mo_path.NewDropboxPath("/Test/Path")

	// regular url
	{
		p := PathWithName(base, DummyImageUrl)

		if p.Path() != "/Test/Path/222.png" {
			t.Error("Invalid path", p)
		}
	}

	// invalid url
	{
		p := PathWithName(base, "/Wrong/Url/222.png")

		if p.Path() != "/Test/Path/222.png" {
			t.Error("invalid path", p)
		}
	}
}

func TestUrlImpl_Save(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx dbx_context.Context) {
		sv := New(ctx)
		_, err := sv.Save(qt_recipe.NewTestDropboxFolderPath(), "https://www.dropbox.com")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
