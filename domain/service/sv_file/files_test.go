package sv_file

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_test"
	"testing"
)

func TestFilesImpl_List(t *testing.T) {
	api_test.DoTestTokenFull(func(ctx api_context.Context) {
		svc := newFilesTest(ctx)
		folder := api_test.ToolboxTestSuiteFolder.ChildPath("list_folder")
		entries, err := svc.List(folder)
		if err != nil {
			t.Error(err)
			return
		}
		if len(entries) < 1 {
			t.Error("invalid")
		}
		for i, e := range entries {
			if i > 10 {
				break
			}
			f, err := svc.Resolve(e.Path())
			if err != nil {
				t.Error(err)
			}
			if f.Tag() != e.Tag() || f.PathLower() != e.PathLower() {
				t.Error("invalid")
			}
		}
	})
}
