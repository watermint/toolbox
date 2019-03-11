package sv_relocation

import (
	"fmt"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_test"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"testing"
	"time"
)

func TestCopy(t *testing.T) {
	api_test.DoTestTokenFull(func(ctx api_context.Context) {
		r := New(ctx)
		dest := fmt.Sprintf("/2019-03-06/F%x.jpg", time.Now().Unix())
		entry, err := r.Copy(mo_path.NewPath("/2019-03-06/F0.jpg"), mo_path.NewPath(dest))
		if err != nil {
			t.Error(err)
		}
		if entry.PathDisplay() != dest {
			t.Error("invalid")
		}
	})
}
