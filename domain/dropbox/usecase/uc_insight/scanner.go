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
		adb:    adb,
		sdb:    adb,
		opts:   so,
	}, nil
}

type tsImpl struct {
	ctl    app_control.Control
	client dbx_client.Client
	// adb: API results database
	adb *gorm.DB
	// sdb: summary database
	sdb  *gorm.DB
	opts ScanOpts
}
