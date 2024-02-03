package uc_insight

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_shutdown"
	"gorm.io/gorm"
)

type IndividualScanner interface {
	// ScanCurrentUser scans current user files and sharing information
	ScanCurrentUser() (err error)
}

type TeamScanner interface {
	// Scan scans all team information
	Scan() (err error)

	Summarize() (err error)

	RetryErrors() (err error)
}

func NewTeamScanner(ctl app_control.Control, client dbx_client.Client, path string) (TeamScanner, error) {
	l := ctl.Log().With(esl.String("path", path))
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
		ctl:              ctl,
		client:           client,
		adb:              adb,
		sdb:              adb,
		disableAutoRetry: false,
		maxRetries:       3,
	}, nil
}

type tsImpl struct {
	ctl    app_control.Control
	client dbx_client.Client
	// adb: API results database
	adb *gorm.DB
	// sdb: summary database
	sdb              *gorm.DB
	disableAutoRetry bool
	maxRetries       int
}
