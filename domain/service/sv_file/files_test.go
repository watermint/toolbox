package sv_file

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_test"
	"testing"
)

func TestFilesImpl_List(t *testing.T) {
	api_test.DoTestTokenFull(func(ctx api_context.Context) {
		ls := newFilesTest(ctx)
		folder := api_test.ToolboxTestSuiteFolder.ChildPath("list_folder")
		entries, err := ls.List(folder)
		if err != nil {
			t.Error(err)
			return
		}
		if len(entries) < 1 {
			t.Error("invalid")
		}
	})
}
