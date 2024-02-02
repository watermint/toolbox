package uc_insight

import (
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"reflect"
)

func (z tsImpl) hasErrors() (countErrors int64, err error) {
	l := z.ctl.Log()
	for _, t := range adbErrorTables {
		tableName := reflect.ValueOf(t).Elem().Type().Name()

		var records int64
		if err := z.adb.Model(t).Count(&records).Error; err != nil {
			z.ctl.Log().Debug("Unable to count errors", esl.Error(err))
			return countErrors, err
		}
		l.Debug("Count error records", esl.String("table", tableName), esl.Int64("records", records))
		countErrors += records
	}

	return countErrors, nil
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

	handleRetry := func(queueId string, errRecordType interface{}, toParam func(record interface{}) interface{}) error {
		typeName := reflect.TypeOf(errRecordType).Elem().Name()
		l.Debug("Handling retry", esl.String("type", typeName))
		rows, err := z.adb.Model(errRecordType).Rows()
		if err != nil {
			l.Debug("Unable to retrieve model", esl.Error(err))
			return err
		}
		defer func() {
			_ = rows.Close()
		}()

		var lastErr error
		z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
			z.defineScanQueues(s, admin, team)
			q := s.Get(queueId)
			for rows.Next() {
				record := reflect.New(reflect.TypeOf(errRecordType).Elem()).Interface()
				if err := rows.Scan(record); err != nil {
					l.Debug("Unable to scan row", esl.Error(err))
					lastErr = err
					return
				}
				q.Enqueue(toParam(record))
			}
		}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
			lastErr = err
		}))
		return lastErr
	}

	type Retryable struct {
		QueueId    string
		RecordType interface{}
		ToParam    func(record interface{}) interface{}
	}

	defer func() {
		db, err := z.adb.DB()
		if err != nil {
			l.Debug("Unable to retrieve db", esl.Error(err))
			return
		}
		_ = db.Close()
	}()

	retryables := []Retryable{
		{
			teamScanQueueFileMember,
			&FileMemberError{},
			func(record interface{}) interface{} {
				r := record.(*FileMemberError)
				return r.ToParam()
			},
		},
		{
			teamScanQueueGroup,
			&GroupError{},
			func(record interface{}) interface{} {
				r := record.(*GroupError)
				return r.ToParam()
			},
		},
		{
			teamScanQueueGroupMember,
			&GroupMemberError{},
			func(record interface{}) interface{} {
				r := record.(*GroupMemberError)
				return r.ToParam()
			},
		},
		{
			teamScanQueueMember,
			&MemberError{},
			func(record interface{}) interface{} {
				r := record.(*MemberError)
				return r.ToParam()
			},
		},
		{
			teamScanQueueMount,
			&MountError{},
			func(record interface{}) interface{} {
				r := record.(*MountError)
				return r.ToParam()
			},
		},
		{
			teamScanQueueNamespaceDetail,
			&NamespaceDetailError{},
			func(record interface{}) interface{} {
				r := record.(*NamespaceDetailError)
				return r.ToParam()
			},
		},
		{
			teamScanQueueNamespaceEntry,
			&NamespaceEntryError{},
			func(record interface{}) interface{} {
				r := record.(*NamespaceEntryError)
				return r.ToParam()
			},
		},
		{
			teamScanQueueNamespace,
			&NamespaceError{},
			func(record interface{}) interface{} {
				r := record.(*NamespaceError)
				return r.ToParam()
			},
		},
		{
			teamScanQueueNamespaceMember,
			&NamespaceMemberError{},
			func(record interface{}) interface{} {
				r := record.(*NamespaceMemberError)
				return r.ToParam()
			},
		},
		{
			teamScanQueueReceivedFile,
			&ReceivedFileError{},
			func(record interface{}) interface{} {
				r := record.(*ReceivedFileError)
				return r.ToParam()
			},
		},
		{
			teamScanQueueSharedLink,
			&SharedLinkError{},
			func(record interface{}) interface{} {
				r := record.(*SharedLinkError)
				return r.ToParam()
			},
		},
		{
			teamScanQueueTeamFolder,
			&TeamFolderError{},
			func(record interface{}) interface{} {
				r := record.(*TeamFolderError)
				return r.ToParam()
			},
		},
	}

	for _, r := range retryables {
		if lastErr := handleRetry(r.QueueId, r.RecordType, r.ToParam); lastErr != nil {
			return lastErr
		}
	}

	return nil
}
