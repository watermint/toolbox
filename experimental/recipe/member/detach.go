package member

import (
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/experimental/app_conn"
	"github.com/watermint/toolbox/experimental/app_file"
	"github.com/watermint/toolbox/experimental/app_kitchen"
	"github.com/watermint/toolbox/experimental/app_report"
	"github.com/watermint/toolbox/experimental/app_validate"
	"github.com/watermint/toolbox/experimental/app_vo"
)

type DetachRow struct {
	Email string
}

func (z *DetachRow) Validate() (err error) {
	if err = app_validate.AssertEmailFormat(z.Email); err != nil {
		return err
	}
	return nil
}

func DetachRowFromCols(cols []string) (row *DetachRow) {
	row = &DetachRow{}
	if len(cols) > 0 {
		row.Email = cols[0]
	}
	return
}

type DetachVO struct {
	File     app_file.RowDataFile
	PeerName app_conn.ConnBusinessMgmt
}

func (*DetachVO) Validate(t app_vo.Validator) {
}

type Detach struct {
}

func (z *Detach) Requirement() app_vo.ValueObject {
	return &DetachVO{}
}

func (*Detach) Exec(k app_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	mvo := vo.(*DetachVO)

	connMgmt, err := mvo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	svm := sv_member.New(connMgmt)
	rep, err := k.Report(
		"detach",
		app_report.TransactionHeader(&DetachRow{}, nil),
	)
	if err != nil {
		return err
	}
	defer rep.Close()

	return mvo.File.EachRow(k.Control(), func(cols []string, rowIndex int) error {
		m := DetachRowFromCols(cols)
		if err = m.Validate(); err != nil {
			if rowIndex > 0 {
				rep.Failure(app_report.MsgInvalidData, m, nil)
			}
			return nil
		}
		mem, err := svm.ResolveByEmail(m.Email)
		if err != nil {
			rep.Failure(api_util.MsgFromError(err), m, nil)
			return nil
		}
		err = svm.Remove(mem, sv_member.Downgrade())
		if err != nil {
			rep.Failure(api_util.MsgFromError(err), m, nil)
		} else {
			rep.Success(m, nil)
		}
		return nil
	})
}
