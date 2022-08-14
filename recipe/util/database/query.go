package database

import (
	"database/sql"
	"encoding/hex"
	_ "github.com/mattn/go-sqlite3"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_griddata"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"
)

type Query struct {
	File   mo_path.FileSystemPath
	Sql    string
	Result da_griddata.GridDataOutput
}

func (z *Query) Preset() {
}

func (z *Query) formatRow(row interface{}) string {
	switch v := row.(type) {
	case *sql.NullString:
		if v.Valid {
			return v.String
		} else {
			return ""
		}
	case *sql.NullBool:
		if v.Valid {
			return strconv.FormatBool(v.Bool)
		} else {
			return ""
		}
	case *sql.NullByte:
		if v.Valid {
			return hex.EncodeToString([]byte{v.Byte})
		} else {
			return ""
		}
	case *sql.NullFloat64:
		if v.Valid {
			return strconv.FormatFloat(v.Float64, 'E', -1, 64)
		} else {
			return ""
		}
	case *sql.NullInt16:
		if v.Valid {
			return strconv.FormatInt(int64(v.Int16), 10)
		} else {
			return ""
		}
	case *sql.NullInt32:
		if v.Valid {
			return strconv.FormatInt(int64(v.Int32), 10)
		} else {
			return ""
		}
	case *sql.NullInt64:
		if v.Valid {
			return strconv.FormatInt(v.Int64, 10)
		} else {
			return ""
		}
	case *sql.NullTime:
		if v.Valid {
			return v.Time.Format(time.RFC3339)
		} else {
			return ""
		}
	}
	return ""
}

func (z *Query) Exec(c app_control.Control) error {
	db, err := sql.Open("sqlite3", z.File.Path())
	if err != nil {
		return err
	}
	r, err := db.Query(z.Sql)
	if err != nil {
		return err
	}
	columns, err := r.Columns()
	if err != nil {
		return err
	}
	columnTypes, err := r.ColumnTypes()
	if err != nil {
		return err
	}
	numColumns := len(columns)
	header := make([]interface{}, numColumns)
	for i := range columns {
		header[i] = columns[i]
	}
	z.Result.Row(header)

	for r.Next() {
		row := make([]interface{}, numColumns)
		for i := range columnTypes {
			row[i] = reflect.New(columnTypes[i].ScanType()).Interface()
		}
		err = r.Scan(row...)
		if err != nil {
			return err
		}
		formattedRow := make([]interface{}, numColumns)
		for i := range row {
			formattedRow[i] = z.formatRow(row[i])
		}
		z.Result.Row(formattedRow)
	}
	return nil
}

func (z *Query) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("sql", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()
	err = rc_exec.Exec(c, &Exec{}, func(r rc_recipe.Recipe) {
		m := r.(*Exec)
		m.File = mo_path.NewFileSystemPath(filepath.Join(f, "test.db"))
		m.Sql = "CREATE TABLE tbx (command TEXT, desc TEXT)"
	})
	if err != nil {
		return err
	}
	err = rc_exec.Exec(c, &Exec{}, func(r rc_recipe.Recipe) {
		m := r.(*Exec)
		m.File = mo_path.NewFileSystemPath(filepath.Join(f, "test.db"))
		m.Sql = "INSERT INTO tbx (command, desc) VALUES ('database exec', 'execute sql')"
	})
	if err != nil {
		return err
	}
	err = rc_exec.Exec(c, &Exec{}, func(r rc_recipe.Recipe) {
		m := r.(*Exec)
		m.File = mo_path.NewFileSystemPath(filepath.Join(f, "test.db"))
		m.Sql = "INSERT INTO tbx (command, desc) VALUES ('database query', 'query database')"
	})
	if err != nil {
		return err
	}
	err = rc_exec.Exec(c, &Query{}, func(r rc_recipe.Recipe) {
		m := r.(*Query)
		m.File = mo_path.NewFileSystemPath(filepath.Join(f, "test.db"))
		m.Sql = "SELECT * FROM tbx"
	})
	return err
}
