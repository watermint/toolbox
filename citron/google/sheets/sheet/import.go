package sheet

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/sheets/model/bo_sheet"
	"github.com/watermint/toolbox/domain/google/sheets/service/sv_sheet"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_griddata"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Import struct {
	Peer     goog_conn.ConnGoogleSheets
	Data     da_griddata.GridDataInput
	Range    string
	Id       string
	InputRaw bool
	Updated  rp_model.RowReport
}

func (z *Import) Preset() {
	z.Peer.SetScopes(goog_auth.ScopeSheetsReadWrite)
	z.Updated.SetModel(&bo_sheet.ValueUpdate{})
}

func (z *Import) Exec(c app_control.Control) error {
	data := make([][]interface{}, 0)
	err := z.Data.EachRow(func(col []interface{}, rowIndex int) error {
		data = append(data, col)
		return nil
	})
	if err != nil {
		return err
	}

	if err := z.Updated.Open(); err != nil {
		return err
	}
	uv, err := sv_sheet.New(z.Peer.Client()).Import(z.Id, z.Range, data, z.InputRaw)
	if err != nil {
		return err
	}
	z.Updated.Row(&uv)
	return nil
}

func (z *Import) Test(c app_control.Control) error {
	csv, err := qt_file.MakeTestCsv("import")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(csv)
	}()

	return rc_exec.ExecMock(c, &Import{}, func(r rc_recipe.Recipe) {
		m := r.(*Import)
		m.Id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		m.Range = "Sheet1"
		m.Data.SetFilePath(csv)
	})
}
