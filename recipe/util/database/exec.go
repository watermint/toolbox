package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"

	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Exec struct {
	rc_recipe.RemarkTransient
	File         mo_path.FileSystemPath
	Sql          string
	AffectedRows app_msg.Message
}

func (z *Exec) Preset() {
}

func (z *Exec) Exec(c app_control.Control) error {
	db, err := sql.Open("sqlite3", z.File.Path())
	if err != nil {
		return err
	}
	r, err := db.Exec(z.Sql)
	if err != nil {
		return err
	}
	rows, err := r.RowsAffected()
	if err != nil {
		return nil
	}
	c.UI().Info(z.AffectedRows.With("Rows", rows))
	return nil
}

func (z *Exec) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("sql", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()
	return rc_exec.Exec(c, &Exec{}, func(r rc_recipe.Recipe) {
		m := r.(*Exec)
		m.File = mo_path.NewFileSystemPath(filepath.Join(f, "test.db"))
		m.Sql = "CREATE TABLE tbx (command TEXT)"
	})
}
