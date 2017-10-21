package dbsugar

import (
	"database/sql"
	"fmt"
	"github.com/cihub/seelog"
	_ "github.com/mattn/go-sqlite3"
	"github.com/watermint/toolbox/infra/util"
)

type DatabaseSugar struct {
	Db             *sql.DB
	DataSourceName string
}

type RowSugar struct {
	Row *sql.Row
	Err error
}

func (d *DatabaseSugar) Open() (err error) {
	d.Db, err = sql.Open("sqlite3", d.DataSourceName)
	return
}

func (d *DatabaseSugar) compileTemplate(queryTmpl, tableName string) (query string, err error) {
	query, err = util.CompileTemplate(queryTmpl, struct {
		TableName string
	}{
		TableName: tableName,
	})
	return
}

func (d *DatabaseSugar) CreateTable(tableName, ddlTmpl string) error {
	dropTable := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
	_, err := d.Db.Exec(dropTable)
	if err != nil {
		seelog.Warnf("Unable to drop table [%s] : error[%s]", tableName, err)
		return err
	}

	ddl, err := d.compileTemplate(ddlTmpl, tableName)
	if err != nil {
		seelog.Warnf("Unable to compile TableName[%s], DDL[%s] : error[%s]", tableName, ddlTmpl, err)
		return err
	}

	_, err = d.Db.Exec(ddl)
	if err != nil {
		seelog.Warnf("Unable to create table [%s] : error[%s]", tableName, err)
		return err
	}
	return nil
}

func (d *DatabaseSugar) QueryRow(query, tableName string, args ...interface{}) (row *RowSugar) {
	row = &RowSugar{}
	q, err := d.compileTemplate(query, tableName)
	if err != nil {
		seelog.Warnf("Unable to compile query[%s] : error[%s]", query, err)
		row.Err = err
		return
	}

	row.Row = d.Db.QueryRow(q, args...)
	return
}

func (d *DatabaseSugar) Query(query, tableName string, args ...interface{}) (*sql.Rows, error) {
	q, err := d.compileTemplate(query, tableName)
	if err != nil {
		seelog.Warnf("Unable to compile query[%s] : error[%s]", query, err)
		return nil, err
	}
	return d.Db.Query(q, args...)
}

func (r *RowSugar) Scan(dest ...interface{}) error {
	if r.Err != nil {
		return r.Err
	}

	return r.Row.Scan(dest...)
}

func (d *DatabaseSugar) Exec(query, tableName string, args ...interface{}) (sql.Result, error) {
	q, err := d.compileTemplate(query, tableName)
	if err != nil {
		seelog.Warnf("Unable to compile query[%s] : error[%s]", query, err)
		return nil, err
	}
	return d.Db.Exec(q, args...)
}
