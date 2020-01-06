package file

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model_impl"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type Watch struct {
	Peer      rc_conn.ConnUserFile
	Path      mo_path.DropboxPath
	Recursive bool
}

func (z *Watch) Exec(c app_control.Control) error {
	ctx := z.Peer.Context()
	opts := make([]sv_file.ListOpt, 0)
	if z.Recursive {
		opts = append(opts, sv_file.Recursive())
	}
	w := rp_model_impl.NewJsonWriter("entries", c, true)
	if err := w.Open(c, &mo_file.ConcreteEntry{}); err != nil {
		return err
	}
	defer w.Close()

	return sv_file.NewFiles(ctx).Poll(z.Path, func(entry mo_file.Entry) {
		w.Row(entry.Concrete())
	}, opts...)
}

func (z *Watch) Test(c app_control.Control) error {
	return qt_endtoend.NoTestRequired()
}

func (z *Watch) Preset() {
}
