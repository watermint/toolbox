package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_shutdown"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"reflect"
)

type IndividualScanner interface {
	// ScanCurrentUser scans current user files and sharing information
	ScanCurrentUser() (err error)
}

type TeamScanner interface {
	// Scan scans all team information
	Scan() (err error)

	// RetryErrors retry errors
	RetryErrors() (err error)

	// ReportLastErrors reports last errors
	ReportLastErrors(onErrorRecords func(errCategory string, errMessage string, errTag string, detail string)) (count int, err error)
}

type ScanOpts struct {
	MaxRetries        int
	ScanMemberFolders bool
}

func (z ScanOpts) Apply(opts []ScanOpt) ScanOpts {
	x := z
	for _, o := range opts {
		x = o(x)
	}
	return x
}

type ScanOpt func(opts ScanOpts) ScanOpts

func MaxRetries(maxRetries int) ScanOpt {
	return func(opts ScanOpts) ScanOpts {
		opts.MaxRetries = maxRetries
		return opts
	}
}

func ScanMemberFolders(enabled bool) ScanOpt {
	return func(opts ScanOpts) ScanOpts {
		opts.ScanMemberFolders = enabled
		return opts
	}
}

func DatabaseFromPath(ctl app_control.Control, path string) (db *gorm.DB, err error) {
	l := ctl.Log().With(esl.String("path", path))
	if err := os.MkdirAll(path, 0700); err != nil {
		l.Debug("Unable to create directory", esl.Error(err))
		return nil, err
	}

	db, err = ctl.NewOrm(filepath.Join(path, databaseName))
	if err != nil {
		l.Debug("Unable to open database", esl.Error(err))
		return nil, err
	}
	return db, nil
}

func NewTeamScanner(ctl app_control.Control, client dbx_client.Client, path string, opts ...ScanOpt) (TeamScanner, error) {
	l := ctl.Log().With(esl.String("path", path))
	so := ScanOpts{}.Apply(opts)
	adb, err := newDatabase(ctl, path)
	if err != nil {
		l.Debug("Unable to open database", esl.Error(err))
		return nil, err
	}

	app_shutdown.AddShutdownHook(func() {
		if db, err := adb.DB(); err == nil {
			_ = db.Close()
		}
	})

	return &tsImpl{
		ctl:    ctl,
		client: client,
		db:     adb,
		opts:   so,
	}, nil
}

type tsImpl struct {
	ctl    app_control.Control
	client dbx_client.Client
	db     *gorm.DB
	opts   ScanOpts
}

func (z tsImpl) ReportLastErrors(onErrorRecords func(errCategory string, errMessage string, errTag string, detail string)) (count int, err error) {
	l := z.ctl.Log()
	if onErrorRecords == nil {
		l.Debug("No error handler, skip reporting")
		return 0, nil
	}

	reportTable := func(t interface{}) {
		tableName := reflect.ValueOf(t).Elem().Type().Name()
		ll := l.With(esl.String("table", tableName))

		rows, err := z.db.Model(t).Rows()
		if err != nil {
			ll.Debug("Unable to retrieve model", esl.Error(err))
			return
		}
		defer func() {
			_ = rows.Close()
		}()

		for rows.Next() {
			record := reflect.New(reflect.TypeOf(t).Elem()).Interface()
			if err := z.db.ScanRows(rows, record); err != nil {
				ll.Debug("Unable to scan row", esl.Error(err))
				return
			}
			count++
			apiErrField := reflect.ValueOf(record).Elem().FieldByName("ApiError")
			apiErr := apiErrField.Interface().(ApiError)

			serialized, err := json.Marshal(record)
			if err != nil {
				ll.Debug("Unable to serialize record", esl.Error(err))
				continue
			}

			onErrorRecords(
				tableName,
				apiErr.Error,
				apiErr.ErrorTag,
				string(serialized),
			)
		}
	}

	for _, t := range adbErrorTables {
		reportTable(t)
	}
	return count, nil
}
