package clear

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"strings"
)

var (
	ErrorUnableToClearExternalId = errors.New("unable to clear external id")
)

type EmailRow struct {
	Email string `json:"email"`
}

type Externalid struct {
	Peer         dbx_conn.ConnBusinessMgmt
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
}

func (z *Externalid) Preset() {
	z.File.SetModel(&EmailRow{})
	z.OperationLog.SetModel(&EmailRow{}, &mo_member.Member{})
}

func (z *Externalid) Exec(c app_control.Control) error {
	l := c.Log()

	if err := z.OperationLog.Open(); err != nil {
		return err
	}
	svm := sv_member.NewCached(z.Peer.Context())

	return z.File.EachRow(func(m interface{}, rowIndex int) error {
		row := m.(*EmailRow)

		l.Debug("Resolving member", esl.String("email", row.Email))
		member, err := svm.ResolveByEmail(row.Email)
		if err != nil {
			z.OperationLog.Failure(err, row)
			return err
		}

		updated, err := svm.Update(member, sv_member.ClearExternalId())
		if err != nil {
			z.OperationLog.Failure(err, row)
			return err
		}
		if updated.ExternalId != "" {
			z.OperationLog.Failure(ErrorUnableToClearExternalId, row)
			return ErrorUnableToClearExternalId
		}

		z.OperationLog.Success(row, updated)
		return nil
	})
}

func (z *Externalid) Test(c app_control.Control) error {
	// replay test
	{
		dummyEmails := make([]string, 0)
		for i := 0; i < 8; i++ {
			dummyEmails = append(dummyEmails, fmt.Sprintf("test%d@example.com", i))
		}
		content := strings.Join(dummyEmails, "\n")
		path, err := qt_file.MakeTestFile("member.csv", content)
		if err != nil {
			return err
		}
		err = rc_exec.ExecReplay(c, &Externalid{}, "recipe-member-clear-externalid.json.gz", func(r rc_recipe.Recipe) {
			m := r.(*Externalid)
			m.File.SetFilePath(path)
		})
		if err != nil {
			return err
		}
	}

	err := rc_exec.ExecMock(c, &Externalid{}, func(r rc_recipe.Recipe) {
		f, err := qt_file.MakeTestFile("member-clear-externalid", "john@example.com\nalex@example.com\n")
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
