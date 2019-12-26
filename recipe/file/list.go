package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type List struct {
	Peer             rc_conn.ConnUserFile
	Path             mo_path.DropboxPath
	Recursive        bool
	IncludeDeleted   bool
	IncludeMediaInfo bool
	FileList         rp_model.RowReport
}

func (z *List) Preset() {
	z.FileList.SetModel(&mo_file.ConcreteEntry{})
}

func (z *List) Exec(k rc_kitchen.Kitchen) error {
	ctx := z.Peer.Context()

	opts := make([]sv_file.ListOpt, 0)
	if z.IncludeDeleted {
		opts = append(opts, sv_file.IncludeDeleted())
	}
	if z.IncludeMediaInfo {
		opts = append(opts, sv_file.IncludeMediaInfo())
	}
	if z.Recursive {
		opts = append(opts, sv_file.Recursive())
	}
	opts = append(opts, sv_file.IncludeHasExplicitSharedMembers())

	if err := z.FileList.Open(); err != nil {
		return err
	}

	err := sv_file.NewFiles(ctx).ListChunked(z.Path, func(entry mo_file.Entry) {
		z.FileList.Row(entry.Concrete())
	}, opts...)
	if err != nil {
		k.Log().Debug("Failed to list files")
		return err
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	err := rc_exec.Exec(c, &List{}, func(r rc_recipe.Recipe) {
		r0 := r.(*List)
		r0.Path = mo_path.NewDropboxPath("")
		r0.Recursive = false
	})
	if err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "file_list", func(cols map[string]string) error {
		if _, ok := cols["id"]; !ok {
			return errors.New("`id` is not found")
		}
		if _, ok := cols["path_display"]; !ok {
			return errors.New("`path_display` is not found")
		}
		return nil
	})
}
