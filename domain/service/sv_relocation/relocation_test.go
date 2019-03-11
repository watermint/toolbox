package sv_relocation

import (
	"fmt"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_test"
	"testing"
	"time"
)

func TestCopy(t *testing.T) {
	api_test.DoTestTokenFull(func(ctx api_context.Context) {
		r := New(ctx)
		src := api_test.ToolboxTestSuiteFolder.ChildPath("F0.jpg")
		name := fmt.Sprintf("copy-%x.jpg", time.Now().Unix())
		dest := api_test.ToolboxTestSuiteFolder.ChildPath(name)

		entry, err := r.Copy(src, dest)
		if err != nil {
			t.Error(err)
		}
		if entry.Name() != name {
			t.Error("invalid")
		}
	})
}
