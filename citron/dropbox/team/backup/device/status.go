package device

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_activity"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_backup"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_device"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_activity"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_device"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"gorm.io/gorm"
	"path/filepath"
	"time"
)

type Status struct {
	Peer                           dbx_conn.ConnScopedTeam
	StartTime                      mo_time.Time
	EndTime                        mo_time.TimeOptional
	Devices                        rp_model.RowReport
	ProgressScanningDeviceSessions app_msg.Message
	ProgressScanningActivity       app_msg.Message
	ProgressEvaluateStatus         app_msg.Message
}

const (
	StatusCategoryNoStatusUpdate = "no_status_update"
	StatusCategoryBackupComplete = "backup_enabled"
	StatusCategoryBackupDisabled = "backup_disabled"
)

func (z *Status) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeEventsRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSessionsList,
	)
	z.Devices.SetModel(&mo_backup.TeamDeviceBackupStatus{},
		rp_model.HiddenColumns(
			"actor_user_team_member_id",
			"session_info_session_id",
		))
}

func (z *Status) scanDeviceSessions(orm *gorm.DB, c app_control.Control) error {
	l := c.Log()
	c.UI().Progress(z.ProgressScanningDeviceSessions)
	sessions, err := sv_device.New(z.Peer.Client()).List()
	if err != nil {
		l.Debug("Unable to list sessions", esl.Error(err))
		return err
	}
	for _, session := range sessions {
		l.Debug("Session", esl.Any("session", session))
		if d, ok := session.Desktop(); ok {
			l.Debug("Desktop", esl.Any("desktop", d))
			if err := orm.Save(d).Error; err != nil {
				l.Debug("Unable to create desktop", esl.Error(err))
				return err
			}
		}
	}

	return nil
}

func (z *Status) scanActivity(orm *gorm.DB, c app_control.Control) error {
	l := c.Log()
	c.UI().Progress(z.ProgressScanningActivity)

	opts := make([]sv_activity.ListOpt, 0)
	opts = append(opts, sv_activity.StartTime(z.StartTime.Iso8601()))
	if !z.EndTime.IsZero() {
		opts = append(opts, sv_activity.EndTime(z.EndTime.Iso8601()))
	}
	opts = append(opts, sv_activity.Category("devices"))

	handler := func(event *mo_activity.Event) error {
		if event.EventType != "device_sync_backup_status_changed" {
			l.Debug("Skip event", esl.Any("event", event))
			return nil
		}
		bev := &mo_backup.TeamActivityDeviceBackupEvent{}
		raw, err := es_json.Parse(event.Raw)
		if err != nil {
			l.Debug("Unable to unmarshal event", esl.Error(err))
			return err
		}
		if err := raw.Model(bev); err != nil {
			l.Debug("Unable to unmarshal into the model", esl.Error(err))
			return err
		}
		if err := orm.Save(bev).Error; err != nil {
			l.Debug("Unable to create event", esl.Error(err))
			return err
		}
		return nil
	}
	return sv_activity.New(z.Peer.Client()).List(handler, opts...)
}

func (z *Status) evaluateStatus(orm *gorm.DB, c app_control.Control) error {
	l := c.Log()
	c.UI().Progress(z.ProgressEvaluateStatus)

	svm := sv_member.NewCached(z.Peer.Client())

	desktop := &mo_device.Desktop{}
	desktopRows, err := orm.Model(desktop).Rows()
	if err != nil {
		l.Debug("Unable to query desktop", esl.Error(err))
		return err
	}
	defer func() {
		_ = desktopRows.Close()
	}()
	for desktopRows.Next() {
		desktop := &mo_device.Desktop{}
		if err := orm.ScanRows(desktopRows, desktop); err != nil {
			l.Debug("Unable to scan desktop", esl.Error(err))
			return err
		}

		event := &mo_backup.TeamActivityDeviceBackupEvent{}
		if err := orm.Where("session_info_session_id = ?", desktop.Id).Order("timestamp desc").First(event).Error; err != nil {
			l.Debug("Unable to find event", esl.Error(err), esl.Any("desktop", desktop))
			member, err := svm.Resolve(desktop.TeamMemberId)
			if err != nil {
				l.Debug("Unable to resolve member", esl.Error(err))
				z.Devices.Row(&mo_backup.TeamDeviceBackupStatus{
					Timestamp:                "",
					LatestStatus:             StatusCategoryNoStatusUpdate,
					ActorUserTeamMemberId:    desktop.TeamMemberId,
					ActorUserEmail:           "",
					ActorUserDisplayName:     "",
					SessionInfoSessionId:     desktop.Id,
					SessionInfoIpAddress:     desktop.IpAddress,
					SessionInfoHostName:      desktop.HostName,
					SessionInfoUpdated:       desktop.Updated,
					SessionInfoClientType:    desktop.ClientType,
					SessionInfoClientVersion: desktop.ClientVersion,
					SessionInfoPlatform:      desktop.Platform,
				})
			} else {
				z.Devices.Row(&mo_backup.TeamDeviceBackupStatus{
					Timestamp:                "",
					LatestStatus:             StatusCategoryNoStatusUpdate,
					ActorUserTeamMemberId:    desktop.TeamMemberId,
					ActorUserEmail:           member.Email,
					ActorUserDisplayName:     member.DisplayName,
					SessionInfoSessionId:     desktop.Id,
					SessionInfoIpAddress:     desktop.IpAddress,
					SessionInfoHostName:      desktop.HostName,
					SessionInfoUpdated:       desktop.Updated,
					SessionInfoClientType:    desktop.ClientType,
					SessionInfoClientVersion: desktop.ClientVersion,
					SessionInfoPlatform:      desktop.Platform,
				})
			}

		} else {
			l.Debug("Event", esl.Any("event", event), esl.Any("desktop", desktop))
			z.Devices.Row(&mo_backup.TeamDeviceBackupStatus{
				Timestamp:                event.Timestamp,
				LatestStatus:             event.NewValue,
				ActorUserTeamMemberId:    event.ActorUserTeamMemberId,
				ActorUserEmail:           event.ActorUserEmail,
				ActorUserDisplayName:     event.ActorUserDisplayName,
				SessionInfoSessionId:     event.SessionInfoSessionId,
				SessionInfoIpAddress:     event.SessionInfoIpAddress,
				SessionInfoHostName:      event.SessionInfoHostName,
				SessionInfoUpdated:       event.SessionInfoUpdated,
				SessionInfoClientType:    event.SessionInfoClientType,
				SessionInfoClientVersion: event.SessionInfoClientVersion,
				SessionInfoPlatform:      event.SessionInfoPlatform,
			})
		}
	}
	return nil
}

func (z *Status) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.Devices.Open(); err != nil {
		l.Debug("Unable to open report", esl.Error(err))
		return err
	}
	path := filepath.Join(c.Workspace().Job(), "device.db")
	orm, err := c.NewOrm(path)
	if err != nil {
		l.Debug("Unable to open database", esl.Error(err))
		return err
	}
	tables := []interface{}{
		&mo_device.Desktop{},
		&mo_backup.TeamActivityDeviceBackupEvent{},
	}
	for _, t := range tables {
		if err := orm.AutoMigrate(t); err != nil {
			l.Debug("Unable to migrate table", esl.Error(err))
			return err
		}
	}

	if err := z.scanDeviceSessions(orm, c); err != nil {
		l.Debug("Unable to scan device sessions", esl.Error(err))
		return err
	}
	if err := z.scanActivity(orm, c); err != nil {
		l.Debug("Unable to scan activity", esl.Error(err))
		return err
	}
	if err := z.evaluateStatus(orm, c); err != nil {
		l.Debug("Unable to evaluate status", esl.Error(err))
		return err
	}

	return nil
}

func (z *Status) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Status{}, func(r rc_recipe.Recipe) {
		m := r.(*Status)
		m.StartTime = mo_time.New(time.Now().Add(-24 * time.Hour))
	})
}
