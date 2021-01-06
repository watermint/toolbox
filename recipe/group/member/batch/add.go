package bulk

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Add struct {
	rc_recipe.RemarkIrreversible
	Peer         dbx_conn.ConnScopedTeam
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
}

func (z *Add) Preset() {
	z.File.SetModel(&MemberRecord{})
	z.Peer.SetScopes(
		dbx_auth.ScopeGroupsRead,
		dbx_auth.ScopeGroupsWrite,
	)
	z.OperationLog.SetModel(&MemberRecord{}, nil)
}

func (z *Add) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	svg := sv_group.NewCached(z.Peer.Context())

	queueIdAdd := "add"

	var lastErr error

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define(queueIdAdd, memberAdd, svg, c, z.Peer.Context(), z.OperationLog)
		q := s.Get(queueIdAdd)

		lastErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			q.Enqueue(m)
			return nil
		})
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	return lastErr
}

func (z *Add) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("add", "Sales,taro@example.com\nSales,hanako@example.com\n")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()
	return rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.File.SetFilePath(f)
	})
}
