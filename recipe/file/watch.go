package file

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/report/rp_model_deprecated"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type Watch struct {
	Peer      rc_conn.OldConnUserFile
	Path      mo_path.DropboxPath
	Recursive bool
}

func (z *Watch) Exec(k rc_kitchen.Kitchen) error {
	ctx, err := z.Peer.Connect(k.Control())
	if err != nil {
		return err
	}
	opts := make([]sv_file.ListOpt, 0)
	if z.Recursive {
		opts = append(opts, sv_file.Recursive())
	}
	rep, _ := rp_model_deprecated.NewJsonForQuiet("entries", k.Control())

	return sv_file.NewFiles(ctx).Poll(z.Path, func(entry mo_file.Entry) {
		rep.Row(entry.Concrete())
	}, opts...)
}

func (z *Watch) Test(c app_control.Control) error {
	return qt_recipe.NoTestRequired()
}

func (z *Watch) Preset() {
}
