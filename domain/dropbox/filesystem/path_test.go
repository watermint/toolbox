package filesystem

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context_impl"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestNewPath(t *testing.T) {
	{
		root := NewPath("", mo_path.NewDropboxPath(""))
		if !root.IsRoot() {
			t.Error(es_json.ToJsonString(root))
		}
		if root.Path() != "/" {
			t.Error(es_json.ToJsonString(root))
		}

		qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
			fs := NewFileSystem(dbx_context_impl.NewMock(ctl))

			root2, err := fs.Path(root.AsData())
			if err != nil {
				t.Error(err)
				return
			}
			if !root2.IsRoot() {
				t.Error(es_json.ToJsonString(root2))
			}
			if root2.Path() != "/" {
				t.Error(es_json.ToJsonString(root2))
			}
		})
	}

	{
		root := NewPath("", mo_path.NewDropboxPath(""))
		if !root.IsRoot() {
			t.Error(es_json.ToJsonString(root))
		}
		if root.Path() != "/" {
			t.Error(es_json.ToJsonString(root))
		}
	}

	{
		root := NewPath("1234", mo_path.NewDropboxPath("/"))
		if !root.IsRoot() {
			t.Error(es_json.ToJsonString(root))
		}
		if root.Path() != "/" {
			t.Error(es_json.ToJsonString(root))
		}
	}
}
