package sv_sharedfolder

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_test"
	"testing"
)

func TestSharedFolderImpl_List(t *testing.T) {
	api_test.DoTestTokenFull(func(ctx api_context.Context) {
		svc := New(ctx)
		folders, err := svc.List()
		if err != nil {
			t.Error(err)
			return
		}
		for _, f := range folders {
			if f.Name == "" {
				t.Error("invalid")
			}
			if f.SharedFolderId == "" {
				t.Error("invalid")
			}
		}
	})
}
