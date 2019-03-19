package sv_sharedfolder_mount

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_test"
	"testing"
)

func TestMountImpl_List(t *testing.T) {
	api_test.DoTestTokenFull(func(ctx api_context.Context) {
		svc := New(ctx)
		mounts, err := svc.List()
		if err != nil {
			t.Error(err)
			return
		}

		for _, m := range mounts {
			if m.SharedFolderId == "" || m.Name == "" {
				t.Error("invalid")
			}
		}
	})
}
