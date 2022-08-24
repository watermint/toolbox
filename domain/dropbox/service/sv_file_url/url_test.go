package sv_file_url

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
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
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.Save(qtr_endtoend.NewTestDropboxFolderPath(), "https://www.dropbox.com")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
