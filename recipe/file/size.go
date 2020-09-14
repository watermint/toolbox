package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_size"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_file_size"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_file_traverse"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Size struct {
	Peer          dbx_conn.ConnUserFile
	NamespaceSize rp_model.RowReport
	Errors        rp_model.TransactionReport
	Path          mo_path.DropboxPath
	Depth         mo_int.RangeInt
}

func (z *Size) Preset() {
	z.NamespaceSize.SetModel(
		&mo_file_size.NamespaceSize{},
		rp_model.HiddenColumns(
			"namespace_name",
			"namespace_id",
			"namespace_type",
			"owner_team_member_id",
		),
	)
	z.Errors.SetModel(&uc_file_traverse.TraverseEntry{}, nil)
	z.Depth.SetRange(1, 300, 2)
	z.Path = mo_path.NewDropboxPath("/")
}

func (z *Size) Exec(c app_control.Control) error {
	if err := z.Errors.Open(rp_model.NoConsoleOutput()); err != nil {
		return err
	}
	if err := z.NamespaceSize.Open(); err != nil {
		return err
	}

	sum := uc_file_size.NewSum(z.Depth.Value())
	handlerEntries := func(te uc_file_traverse.TraverseEntry, entries []mo_file.Entry) {
		sum.Eval(te.Path, entries)
	}
	handlerError := func(te uc_file_traverse.TraverseEntry, err error) {
		z.Errors.Failure(err, &te)
	}
	traverseQueueId := "scan"
	traverse := uc_file_traverse.NewTraverse(
		z.Peer.Context(),
		c,
		traverseQueueId,
		handlerEntries,
		handlerError,
	)

	var lastErr error
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define(traverseQueueId, traverse.Traverse, s)
		q := s.Get(traverseQueueId)
		q.Enqueue(uc_file_traverse.TraverseEntry{
			Path:      z.Path.Path(),
			Namespace: &mo_namespace.Namespace{},
		})
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	sum.Each(func(path mo_path.DropboxPath, size mo_file_size.Size) {
		z.NamespaceSize.Row(size)
	})

	return lastErr
}

func (z *Size) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Size{}, func(r rc_recipe.Recipe) {
		m := r.(*Size)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath()
	})
}
