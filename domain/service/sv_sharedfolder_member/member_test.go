package sv_sharedfolder_member

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_test"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"testing"
)

func TestMemberImpl_List(t *testing.T) {
	api_test.DoTestTokenFull(func(ctx api_context.Context) {
		sf := sv_sharedfolder.New(ctx)
		folders, err := sf.List()
		if err != nil {
			t.Error(err)
			return
		}

		for _, folder := range folders {
			sm := New(ctx, folder)
			members, err := sm.List()
			if err != nil {
				t.Error(err)
				return
			}

			for _, m := range members {
				if m.MemberType() == "" {
					t.Error("invalid")
				}
			}
		}
	})
}
