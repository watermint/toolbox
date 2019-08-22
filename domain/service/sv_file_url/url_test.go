package sv_file_url

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_test"
	"github.com/watermint/toolbox/infra/api/api_util"
	"go.uber.org/zap"
	"strings"
	"testing"
)

const (
	DummyImageUrl = "https://dummyimage.com/64x64/888/222.png"
)

func TestUrlImpl_Save(t *testing.T) {
	api_test.DoTestTokenFull(func(ctx api_context.Context) {
		svc := New(ctx)
		path := api_test.ToolboxTestSuiteFolder.ChildPath("save_url").ChildPath("f0.png")
		entry, err := svc.Save(path, DummyImageUrl)
		if err != nil {
			t.Error(api_util.UIMsgFromError(err).T())
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
