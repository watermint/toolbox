package clear

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

var (
	ErrorUnableToClearExternalId = errors.New("unable to clear external id")
)

type GroupNameRow struct {
	Name string `json:"name"`
}

type Externalid struct {
	Peer         dbx_conn.ConnScopedTeam
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
}

func (z *Externalid) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeGroupsRead,
		dbx_auth.ScopeGroupsWrite,
	)
	z.File.SetModel(&GroupNameRow{})
	z.OperationLog.SetModel(&GroupNameRow{}, &mo_group.Group{})
}

func (z *Externalid) Exec(c app_control.Control) error {
	l := c.Log()

	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	svg := sv_group.NewCached(z.Peer.Context())
	svNoCache := sv_group.New(z.Peer.Context())

	return z.File.EachRow(func(m interface{}, rowIndex int) error {
		row := m.(*GroupNameRow)

		l.Debug("Resolving group", esl.String("group", row.Name))
		group, err := svg.ResolveByName(row.Name)
		if err != nil {
			z.OperationLog.Failure(err, row)
			return err
		}

		group.GroupExternalId = ""
		_, err = svg.Update(group)
		if err != nil {
			z.OperationLog.Failure(err, row)
			return err
		}
		updated, err := svNoCache.ResolveByName(row.Name)
		if err != nil {
			z.OperationLog.Failure(err, row)
			return err
		}

		if updated.GroupExternalId != "" {
			z.OperationLog.Failure(ErrorUnableToClearExternalId, row)
			return ErrorUnableToClearExternalId
		}

		z.OperationLog.Success(row, updated)
		return nil
	})
}

func (z *Externalid) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Externalid{}, func(r rc_recipe.Recipe) {
		f, err := qt_file.MakeTestFile("group-clear-externalid", "CorpIT\nSales\n")
		if err != nil {
			return
		}
		m := r.(*Externalid)
		m.File.SetFilePath(f)
	})
	if e, _ := qt_errors.ErrorsForTest(c.Log(), err); e != nil {
		return e
	}
	return nil
}
