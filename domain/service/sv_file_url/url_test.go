package sv_file_url

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"strings"
	"testing"
)

const (
	DummyImageUrl = "https://dummyimage.com/64x64/888/222.png"
)

func TestUrlImpl_SaveWithTestSuite(t *testing.T) {
	qt_api.DoTestTokenFull(func(ctx api_context.Context) {
		svc := New(ctx)
		path := qt_api.ToolboxTestSuiteFolder.ChildPath("save_url", "f0.png")
		entry, err := svc.Save(path, DummyImageUrl)
		if err != nil {
			t.Error(err.Error())
			return
		}
		if entry.Tag() != "file" {
			t.Error("invalid")
		}

		// file name might be auto renamed into like `f0 (1).png`, if duplicated file there
		if !strings.HasPrefix(entry.Name(), "f0") {
			t.Error("invalid")
		}
		ctx.Log().Debug("entry", zap.Any("entry", entry))
	})
}

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
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Save(qt_recipe.NewTestDropboxFolderPath(), "https://www.dropbox.com")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
