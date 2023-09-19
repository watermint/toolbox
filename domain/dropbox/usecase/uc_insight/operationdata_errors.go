package uc_insight

import (
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
)

type NamespaceEntryError struct {
	NamespaceId string `path:"shared_folder_id" gorm:"primaryKey"`
	FolderId    string `path:"folder_id" gorm:"primaryKey"`

	Error string `path:"error_summary"`

	Updated uint64 `gorm:"autoUpdateTime"`
}

func (z tsImpl) RetryErrors() error {
	l := z.ctl.Log()
	admin, err := sv_profile.NewTeam(z.client).Admin()
	if err != nil {
		l.Debug("Unable to retrieve admin profile", esl.Error(err))
		return err
	}
	team, err := sv_team.New(z.client).Info()
	if err != nil {
		l.Debug("Unable to retrieve team info", esl.Error(err))
		return err
	}

	rows, err := z.db.Model(&NamespaceEntryError{}).Rows()
	if err != nil {
		l.Debug("cannot retrieve model", esl.Error(err))
		return err
	}
	defer func() {
		_ = rows.Close()
	}()

	var lastErr error
	z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
		z.defineScanQueues(s, admin, team)
		qNamespaceEntry := s.Get(teamScanQueueNamespaceEntry)

		for rows.Next() {
			namespaceEntryError := &NamespaceEntryError{}
			if err := z.db.ScanRows(rows, namespaceEntryError); err != nil {
				l.Debug("cannot scan row", esl.Error(err))
				return
			}
			qNamespaceEntry.Enqueue(&NamespaceEntryParam{
				NamespaceId: namespaceEntryError.NamespaceId,
				FolderId:    namespaceEntryError.FolderId,
				IsRetry:     true,
			})
		}

	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	db, err := z.db.DB()
	if err != nil {
		return err
	}
	_ = db.Close()

	return lastErr
}
